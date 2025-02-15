// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"fmt"
	"path"
	"sort"

	"github.com/gohugoio/hugo/compare"

	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/resources/resource"
)

// The TaxonomyList is a list of all taxonomies and their values
// e.g. List['tags'] => TagTaxonomy (from above)
type TaxonomyList map[string]Taxonomy

func (tl TaxonomyList) String() string {
	return fmt.Sprintf("TaxonomyList(%d)", len(tl))
}

// A Taxonomy is a map of keywords to a list of pages.
// For example
//    TagTaxonomy['technology'] = page.WeightedPages
//    TagTaxonomy['go']  =  page.WeightedPages
type Taxonomy map[string]page.WeightedPages

// OrderedTaxonomy is another representation of an Taxonomy using an array rather than a map.
// Important because you can't order a map.
type OrderedTaxonomy []OrderedTaxonomyEntry

// OrderedTaxonomyEntry is similar to an element of a Taxonomy, but with the key embedded (as name)
// e.g:  {Name: Technology, page.WeightedPages: TaxonomyPages}
type OrderedTaxonomyEntry struct {
	Name string
	page.WeightedPages
}

// Get the weighted pages for the given key.
func (i Taxonomy) Get(key string) page.WeightedPages {
	return i[key]
}

// Count the weighted pages for the given key.
func (i Taxonomy) Count(key string) int { return len(i[key]) }

func (i Taxonomy) add(key string, w page.WeightedPage) {
	i[key] = append(i[key], w)
}

// TaxonomyArray returns an ordered taxonomy with a non defined order.
func (i Taxonomy) TaxonomyArray() OrderedTaxonomy {
	ies := make([]OrderedTaxonomyEntry, len(i))
	count := 0
	for k, v := range i {
		ies[count] = OrderedTaxonomyEntry{Name: k, WeightedPages: v}
		count++
	}
	return ies
}

// Alphabetical returns an ordered taxonomy sorted by key name.
func (i Taxonomy) Alphabetical() OrderedTaxonomy {
	name := func(i1, i2 *OrderedTaxonomyEntry) bool {
		return compare.LessStrings(i1.Name, i2.Name)
	}

	ia := i.TaxonomyArray()
	oiBy(name).Sort(ia)
	return ia
}

// ByCount returns an ordered taxonomy sorted by # of pages per key.
// If taxonomies have the same # of pages, sort them alphabetical
func (i Taxonomy) ByCount() OrderedTaxonomy {
	count := func(i1, i2 *OrderedTaxonomyEntry) bool {
		li1 := len(i1.WeightedPages)
		li2 := len(i2.WeightedPages)

		if li1 == li2 {
			return compare.LessStrings(i1.Name, i2.Name)
		}
		return li1 > li2
	}

	ia := i.TaxonomyArray()
	oiBy(count).Sort(ia)
	return ia
}

// Pages returns the Pages for this taxonomy.
func (ie OrderedTaxonomyEntry) Pages() page.Pages {
	return ie.WeightedPages.Pages()
}

// Count returns the count the pages in this taxonomy.
func (ie OrderedTaxonomyEntry) Count() int {
	return len(ie.WeightedPages)
}

// Term returns the name given to this taxonomy.
func (ie OrderedTaxonomyEntry) Term() string {
	return ie.Name
}

// Reverse reverses the order of the entries in this taxonomy.
func (t OrderedTaxonomy) Reverse() OrderedTaxonomy {
	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return t
}

// A type to implement the sort interface for TaxonomyEntries.
type orderedTaxonomySorter struct {
	taxonomy OrderedTaxonomy
	by       oiBy
}

// Closure used in the Sort.Less method.
type oiBy func(i1, i2 *OrderedTaxonomyEntry) bool

func (by oiBy) Sort(taxonomy OrderedTaxonomy) {
	ps := &orderedTaxonomySorter{
		taxonomy: taxonomy,
		by:       by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Stable(ps)
}

// Len is part of sort.Interface.
func (s *orderedTaxonomySorter) Len() int {
	return len(s.taxonomy)
}

// Swap is part of sort.Interface.
func (s *orderedTaxonomySorter) Swap(i, j int) {
	s.taxonomy[i], s.taxonomy[j] = s.taxonomy[j], s.taxonomy[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *orderedTaxonomySorter) Less(i, j int) bool {
	return s.by(&s.taxonomy[i], &s.taxonomy[j])
}

// taxonomyNodeInfo stores additional metadata about a taxonomy.
type taxonomyNodeInfo struct {
	plural string

	// Maps "tags" to "tag".
	singular string

	// The term key as used in the taxonomy map, e.g "tag1".
	// The value is normalized for paths, but may or not be lowercased
	// depending on the disablePathToLower setting.
	termKey string

	// The original, unedited term name. Useful for titles etc.
	term string

	dates resource.Dates

	parent *taxonomyNodeInfo

	// Either of Kind taxonomyTerm (parent) or taxonomy
	owner *page.PageWrapper
}

func (t *taxonomyNodeInfo) UpdateFromPage(p page.Page) {

	// Select the latest dates
	t.dates.UpdateDateAndLastmodIfAfter(p)
}

func (t *taxonomyNodeInfo) TransferValues(p *pageState) {
	t.owner.Page = p
	if p.Lastmod().IsZero() && p.Date().IsZero() {
		p.m.Dates.UpdateDateAndLastmodIfAfter(t.dates)
	}
}

// Maps either plural or plural/term to a taxonomy node.
// TODO(bep) consolidate somehow with s.Taxonomies
type taxonomyNodeInfos struct {
	m      map[string]*taxonomyNodeInfo
	getKey func(string) string
}

// map[string]*taxonomyNodeInfo
func (t taxonomyNodeInfos) key(parts ...string) string {
	return path.Join(parts...)
}

// GetOrAdd will get or create and add a new taxonomy node to the parent identified with plural.
// It will panic if the parent does not exist.
func (t taxonomyNodeInfos) GetOrAdd(plural, term string) *taxonomyNodeInfo {
	parent := t.GetOrCreate(plural, "")
	if parent == nil {
		panic(fmt.Sprintf("no parent found with plural %q", plural))
	}
	child := t.GetOrCreate(plural, term)
	child.parent = parent
	return child
}

func (t taxonomyNodeInfos) GetOrCreate(plural, term string) *taxonomyNodeInfo {
	termKey := t.getKey(term)
	key := t.key(plural, termKey)

	n, found := t.m[key]
	if found {
		return n
	}

	n = &taxonomyNodeInfo{
		plural:  plural,
		termKey: termKey,
		term:    term,
		owner:   &page.PageWrapper{}, // Page will be assigned later.
	}

	t.m[key] = n

	return n
}

func (t taxonomyNodeInfos) Get(sections ...string) *taxonomyNodeInfo {
	key := t.key(sections...)

	n, found := t.m[key]
	if found {
		return n
	}

	return nil
}
