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

// This file is autogenerated.

// Package embedded defines the internal templates that Hugo provides.
package embedded

// EmbeddedTemplates represents all embedded templates.
var EmbeddedTemplates = [][2]string{
	{`_default/robots.txt`, `User-agent: *`},
	{`_default/rss.xml`, `{{- $pages := .Data.Pages -}}
{{- $limit := .Site.Config.Services.RSS.Limit -}}
{{- if ge $limit 1 -}}
{{- $pages = $pages | first $limit -}}
{{- end -}}
{{- printf "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\" ?>" | safeHTML }}
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{ if eq  .Title  .Site.Title }}{{ .Site.Title }}{{ else }}{{ with .Title }}{{.}} on {{ end }}{{ .Site.Title }}{{ end }}</title>
    <link>{{ .Permalink }}</link>
    <description>Recent content {{ if ne  .Title  .Site.Title }}{{ with .Title }}in {{.}} {{ end }}{{ end }}on {{ .Site.Title }}</description>
    <generator>Hugo -- gohugo.io</generator>{{ with .Site.LanguageCode }}
    <language>{{.}}</language>{{end}}{{ with .Site.Author.email }}
    <managingEditor>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</managingEditor>{{end}}{{ with .Site.Author.email }}
    <webMaster>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</webMaster>{{end}}{{ with .Site.Copyright }}
    <copyright>{{.}}</copyright>{{end}}{{ if not .Date.IsZero }}
    <lastBuildDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</lastBuildDate>{{ end }}
    {{ with .OutputFormats.Get "RSS" }}
	{{ printf "<atom:link href=%q rel=\"self\" type=%q />" .Permalink .MediaType | safeHTML }}
    {{ end }}
    {{ range $pages }}
    <item>
      <title>{{ .Title }}</title>
      <link>{{ .Permalink }}</link>
      <pubDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</pubDate>
      {{ with .Site.Author.email }}<author>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</author>{{end}}
      <guid>{{ .Permalink }}</guid>
      <description>{{ .Summary | html }}</description>
    </item>
    {{ end }}
  </channel>
</rss>`},
	{`_default/sitemap.xml`, `{{ printf "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\" ?>" | safeHTML }}
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
  xmlns:xhtml="http://www.w3.org/1999/xhtml">
  {{ range .Data.Pages }}
  <url>
    <loc>{{ .Permalink }}</loc>{{ if not .Lastmod.IsZero }}
    <lastmod>{{ safeHTML ( .Lastmod.Format "2006-01-02T15:04:05-07:00" ) }}</lastmod>{{ end }}{{ with .Sitemap.ChangeFreq }}
    <changefreq>{{ . }}</changefreq>{{ end }}{{ if ge .Sitemap.Priority 0.0 }}
    <priority>{{ .Sitemap.Priority }}</priority>{{ end }}{{ if .IsTranslated }}{{ range .Translations }}
    <xhtml:link
                rel="alternate"
                hreflang="{{ .Language.Lang }}"
                href="{{ .Permalink }}"
                />{{ end }}
    <xhtml:link
                rel="alternate"
                hreflang="{{ .Language.Lang }}"
                href="{{ .Permalink }}"
                />{{ end }}
  </url>
  {{ end }}
</urlset>`},
	{`_default/sitemapindex.xml`, `{{ printf "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\" ?>" | safeHTML }}
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	{{ range . }}
	<sitemap>
	   	<loc>{{ .SitemapAbsURL }}</loc>
		{{ if not .LastChange.IsZero }}
	   	<lastmod>{{ .LastChange.Format "2006-01-02T15:04:05-07:00" | safeHTML }}</lastmod>
		{{ end }}
	</sitemap>
	{{ end }}
</sitemapindex>
`},
	{`disqus.html`, `{{- $pc := .Site.Config.Privacy.Disqus -}}
{{- if not $pc.Disable -}}
{{ if .Site.DisqusShortname }}<div id="disqus_thread"></div>
<script type="application/javascript">
    var disqus_config = function () {
    {{with .Params.disqus_identifier }}this.page.identifier = '{{ . }}';{{end}}
    {{with .Params.disqus_title }}this.page.title = '{{ . }}';{{end}}
    {{with .Params.disqus_url }}this.page.url = '{{ . | html  }}';{{end}}
    };
    (function() {
        if (["localhost", "127.0.0.1"].indexOf(window.location.hostname) != -1) {
            document.getElementById('disqus_thread').innerHTML = 'Disqus comments not available by default when the website is previewed locally.';
            return;
        }
        var d = document, s = d.createElement('script'); s.async = true;
        s.src = '//' + {{ .Site.DisqusShortname }} + '.disqus.com/embed.js';
        s.setAttribute('data-timestamp', +new Date());
        (d.head || d.body).appendChild(s);
    })();
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
<a href="https://disqus.com" class="dsq-brlink">comments powered by <span class="logo-disqus">Disqus</span></a>{{end}}
{{- end -}}`},
	{`google_analytics.html`, `{{- $pc := .Site.Config.Privacy.GoogleAnalytics -}}
{{- if not $pc.Disable -}}
{{ with .Site.GoogleAnalytics }}
<script type="application/javascript">
{{ template "__ga_js_set_doNotTrack" $ }}
if (!doNotTrack) {
	(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');
	{{- if $pc.UseSessionStorage }}
	if (window.sessionStorage) {
		var GA_SESSION_STORAGE_KEY = 'ga:clientId';
		ga('create', '{{ . }}', {
	    'storage': 'none',
	    'clientId': sessionStorage.getItem(GA_SESSION_STORAGE_KEY)
	   });
	   ga(function(tracker) {
	    sessionStorage.setItem(GA_SESSION_STORAGE_KEY, tracker.get('clientId'));
	   });
   }
	{{ else }}
	ga('create', '{{ . }}', 'auto');
	{{ end -}}
	{{ if $pc.AnonymizeIP }}ga('set', 'anonymizeIp', true);{{ end }}
	ga('send', 'pageview');
}
</script>
{{ end }}
{{- end -}}
{{- define "__ga_js_set_doNotTrack" -}}{{/* This is also used in the async version. */}}
{{- $pc := .Site.Config.Privacy.GoogleAnalytics -}}
{{- if not $pc.RespectDoNotTrack -}}
var doNotTrack = false;
{{- else -}}
var dnt = (navigator.doNotTrack || window.doNotTrack || navigator.msDoNotTrack);
var doNotTrack = (dnt == "1" || dnt == "yes");
{{- end -}}
{{- end -}}`},
	{`google_analytics_async.html`, `{{- $pc := .Site.Config.Privacy.GoogleAnalytics -}}
{{- if not $pc.Disable -}}
{{ with .Site.GoogleAnalytics }}
<script type="application/javascript">
{{ template "__ga_js_set_doNotTrack" $ }}
if (!doNotTrack) {
	window.ga=window.ga||function(){(ga.q=ga.q||[]).push(arguments)};ga.l=+new Date;
	{{- if $pc.UseSessionStorage }}
	if (window.sessionStorage) {
		var GA_SESSION_STORAGE_KEY = 'ga:clientId';
		ga('create', '{{ . }}', {
	    'storage': 'none',
	    'clientId': sessionStorage.getItem(GA_SESSION_STORAGE_KEY)
	   });
	   ga(function(tracker) {
	    sessionStorage.setItem(GA_SESSION_STORAGE_KEY, tracker.get('clientId'));
	   });
   }
	{{ else }}
	ga('create', '{{ . }}', 'auto');
	{{ end -}}
	{{ if $pc.AnonymizeIP }}ga('set', 'anonymizeIp', true);{{ end }}
	ga('send', 'pageview');
}
</script>
<script async src='https://www.google-analytics.com/analytics.js'></script>
{{ end }}
{{- end -}}
`},
	{`google_news.html`, `{{ if .IsPage }}{{ with .Params.news_keywords }}
  <meta name="news_keywords" content="{{ range $i, $kw := first 10 . }}{{ if $i }},{{ end }}{{ $kw }}{{ end }}" />
{{ end }}{{ end }}`},
	{`opengraph.html`, `<meta property="og:title" content="{{ .Title }}" />
<meta property="og:description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end }}" />
<meta property="og:type" content="{{ if .IsPage }}article{{ else }}website{{ end }}" />
<meta property="og:url" content="{{ .Permalink }}" />
{{ with $.Param "images" }}{{ range first 6 . }}
<meta property="og:image" content="{{ . | absURL }}" />
{{ end }}{{ end }}

{{- $iso8601 := "2006-01-02T15:04:05-07:00" -}}
{{- if .IsPage }}
{{- if not .PublishDate.IsZero }}<meta property="article:published_time" {{ .PublishDate.Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />
{{ else if not .Date.IsZero }}<meta property="article:published_time" {{ .Date.Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />
{{ end }}
{{- if not .Lastmod.IsZero }}<meta property="article:modified_time" {{ .Lastmod.Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />{{ end }}
{{- else }}
{{- if not .Date.IsZero }}
<meta property="og:updated_time" {{ .Date.Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />
{{- end }}
{{- end }}{{/* .IsPage */}}

{{- with .Params.audio }}<meta property="og:audio" content="{{ . }}" />{{ end }}
{{- with .Params.locale }}<meta property="og:locale" content="{{ . }}" />{{ end }}
{{- with .Site.Params.title }}<meta property="og:site_name" content="{{ . }}" />{{ end }}
{{- with .Params.videos }}
{{- range . }}
<meta property="og:video" content="{{ . | absURL }}" />
{{ end }}{{ end }}

{{- /* If it is part of a series, link to related articles */}}
{{- $permalink := .Permalink }}
{{- $siteSeries := .Site.Taxonomies.series }}{{ with .Params.series }}
{{- range $name := . }}
  {{- $series := index $siteSeries $name }}
  {{- range $page := first 6 $series.Pages }}
    {{- if ne $page.Permalink $permalink }}<meta property="og:see_also" content="{{ $page.Permalink }}" />{{ end }}
  {{- end }}
{{ end }}{{ end }}

{{- if .IsPage }}
{{- range .Site.Authors }}{{ with .Social.facebook }}
<meta property="article:author" content="https://www.facebook.com/{{ . }}" />{{ end }}{{ with .Site.Social.facebook }}
<meta property="article:publisher" content="https://www.facebook.com/{{ . }}" />{{ end }}
<meta property="article:section" content="{{ .Section }}" />
{{- with .Params.tags }}{{ range first 6 . }}
<meta property="article:tag" content="{{ . }}" />{{ end }}{{ end }}
{{- end }}{{ end }}

{{- /* Facebook Page Admin ID for Domain Insights */}}
{{- with .Site.Social.facebook_admin }}<meta property="fb:admins" content="{{ . }}" />{{ end }}
`},
	{`pagination.html`, `{{ $pag := $.Paginator }}
{{ if gt $pag.TotalPages 1 }}
<ul class="pagination">
    {{ with $pag.First }}
    <li class="page-item">
        <a href="{{ .URL }}" class="page-link" aria-label="First"><span aria-hidden="true">&laquo;&laquo;</span></a>
    </li>
    {{ end }}
    <li class="page-item{{ if not $pag.HasPrev }} disabled{{ end }}">
    <a {{ if $pag.HasPrev }}href="{{ $pag.Prev.URL }}"{{ end }} class="page-link" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a>
    </li>
    {{ $ellipsed := false }}
    {{ $shouldEllipse := false }}
    {{ range $pag.Pagers }}
    {{ $right := sub .TotalPages .PageNumber }}
    {{ $showNumber := or (le .PageNumber 3) (eq $right 0) }}
    {{ $showNumber := or $showNumber (and (gt .PageNumber (sub $pag.PageNumber 2)) (lt .PageNumber (add $pag.PageNumber 2)))  }}
    {{ if $showNumber }}
        {{ $ellipsed = false }}
        {{ $shouldEllipse = false }}
    {{ else }}
        {{ $shouldEllipse = not $ellipsed }}
        {{ $ellipsed = true }}
    {{ end }}
    {{ if $showNumber }}
    <li class="page-item{{ if eq . $pag }} active{{ end }}"><a class="page-link" href="{{ .URL }}">{{ .PageNumber }}</a></li>
    {{ else if $shouldEllipse }}
    <li class="page-item disabled"><span aria-hidden="true">&nbsp;&hellip;&nbsp;</span></li>
    {{ end }}
    {{ end }}
    <li class="page-item{{ if not $pag.HasNext }} disabled{{ end }}">
    <a {{ if $pag.HasNext }}href="{{ $pag.Next.URL }}"{{ end }} class="page-link" aria-label="Next"><span aria-hidden="true">&raquo;</span></a>
    </li>
    {{ with $pag.Last }}
    <li class="page-item">
        <a href="{{ .URL }}" class="page-link" aria-label="Last"><span aria-hidden="true">&raquo;&raquo;</span></a>
    </li>
    {{ end }}
</ul>
{{ end }}
`},
	{`schema.html`, `<meta itemprop="name" content="{{ .Title }}">
<meta itemprop="description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end }}">

{{if .IsPage}}{{ $ISO8601 := "2006-01-02T15:04:05-07:00" }}{{ if not .PublishDate.IsZero }}
<meta itemprop="datePublished" content="{{ .PublishDate.Format $ISO8601 | safeHTML }}" />{{ end }}
{{ if not .Lastmod.IsZero }}<meta itemprop="dateModified" content="{{ .Lastmod.Format $ISO8601 | safeHTML }}" />{{ end }}
<meta itemprop="wordCount" content="{{ .WordCount }}">
{{ with .Params.images }}{{ range first 6 . }}
  <meta itemprop="image" content="{{ . | absURL }}">
{{ end }}{{ end }}

<!-- Output all taxonomies as schema.org keywords -->
<meta itemprop="keywords" content="{{ if .IsPage}}{{ range $index, $tag := .Params.tags }}{{ $tag }},{{ end }}{{ else }}{{ range $plural, $terms := .Site.Taxonomies }}{{ range $term, $val := $terms }}{{ printf "%s," $term }}{{ end }}{{ end }}{{ end }}" />
{{ end }}`},
	{`shortcodes/__h_simple_assets.html`, `{{ define "__h_simple_css" }}{{/* These template definitions are global. */}}
{{- if not (.Page.Scratch.Get "__h_simple_css") -}}
{{/* Only include once */}}
{{-  .Page.Scratch.Set "__h_simple_css" true -}}
<style>
.__h_video {
   position: relative;
   padding-bottom: 56.23%;
   height: 0;
   overflow: hidden;
   width: 100%;
   background: #000;
}
.__h_video img {
   width: 100%;
   height: auto;
   color: #000;
}
.__h_video .play {
   height: 72px;
   width: 72px;
   left: 50%;
   top: 50%;
   margin-left: -36px;
   margin-top: -36px;
   position: absolute;
   cursor: pointer;
}
</style>
{{- end -}}
{{- end -}}
{{- define "__h_simple_icon_play" -}}
<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 61 61"><circle cx="30.5" cy="30.5" r="30.5" opacity=".8" fill="#000"></circle><path d="M25.3 19.2c-2.1-1.2-3.8-.2-3.8 2.2v18.1c0 2.4 1.7 3.4 3.8 2.2l16.6-9.1c2.1-1.2 2.1-3.2 0-4.4l-16.6-9z" fill="#fff"></path></svg>
{{- end -}}
`},
	{`shortcodes/figure.html`, `<figure{{ with .Get "class" }} class="{{ . }}"{{ end }}>
    {{- if .Get "link" -}}
        <a href="{{ .Get "link" }}"{{ with .Get "target" }} target="{{ . }}"{{ end }}{{ with .Get "rel" }} rel="{{ . }}"{{ end }}>
    {{- end }}
    <img src="{{ .Get "src" }}"
         {{- if or (.Get "alt") (.Get "caption") }}
         alt="{{ with .Get "alt" }}{{ . }}{{ else }}{{ .Get "caption" | markdownify| plainify }}{{ end }}"
         {{- end -}}
         {{- with .Get "width" }} width="{{ . }}"{{ end -}}
         {{- with .Get "height" }} height="{{ . }}"{{ end -}}
    /> <!-- Closing img tag -->
    {{- if .Get "link" }}</a>{{ end -}}
    {{- if or (or (.Get "title") (.Get "caption")) (.Get "attr") -}}
        <figcaption>
            {{ with (.Get "title") -}}
                <h4>{{ . }}</h4>
            {{- end -}}
            {{- if or (.Get "caption") (.Get "attr") -}}<p>
                {{- .Get "caption" | markdownify -}}
                {{- with .Get "attrlink" }}
                    <a href="{{ . }}">
                {{- end -}}
                {{- .Get "attr" | markdownify -}}
                {{- if .Get "attrlink" }}</a>{{ end }}</p>
            {{- end }}
        </figcaption>
    {{- end }}
</figure>
`},
	{`shortcodes/gist.html`, `<script type="application/javascript" src="https://gist.github.com/{{ index .Params 0 }}/{{ index .Params 1 }}.js{{if len .Params | eq 3 }}?file={{ index .Params 2 }}{{end}}"></script>
`},
	{`shortcodes/highlight.html`, `{{ if len .Params | eq 2 }}{{ highlight (trim .Inner "\n\r") (.Get 0) (.Get 1) }}{{ else }}{{ highlight (trim .Inner "\n\r") (.Get 0) "" }}{{ end }}`},
	{`shortcodes/instagram.html`, `{{- $pc := .Page.Site.Config.Privacy.Instagram -}}
{{- if not $pc.Disable -}}
{{- if $pc.Simple -}}
{{ template "_internal/shortcodes/instagram_simple.html" . }}
{{- else -}}
{{ $id := .Get 0 }}
{{ $hideCaption := cond (eq (.Get 1) "hidecaption") "1" "0" }}
{{ with getJSON "https://api.instagram.com/oembed/?url=https://instagram.com/p/" $id "/&hidecaption=" $hideCaption  }}{{ .html | safeHTML }}{{ end }}
{{- end -}}
{{- end -}}`},
	{`shortcodes/instagram_simple.html`, `{{- $pc := .Page.Site.Config.Privacy.Instagram -}}
{{- $sc := .Page.Site.Config.Services.Instagram -}}
{{- if not $pc.Disable -}}
{{- $id := .Get 0 -}}
{{- $item := getJSON "https://api.instagram.com/oembed/?url=https://www.instagram.com/p/" $id "/&amp;maxwidth=640&amp;omitscript=true" -}}
{{- $class1 := "__h_instagram" -}}
{{- $class2 := "s_instagram_simple" -}}
{{- $hideCaption := (eq (.Get 1) "hidecaption") -}}
{{ with $item }}
{{- $mediaURL := printf "https://instagram.com/p/%s/" $id | safeURL -}}
{{- if not $sc.DisableInlineCSS -}}
{{ template "__h_simple_instagram_css" $ }}
{{- end -}}
<div class="{{ $class1 }} {{ $class2 }} card" style="max-width: {{ $item.thumbnail_width }}px">
	<div class="card-header">
    <a href="{{ $item.author_url | safeURL }}" class="card-link">{{ $item.author_name }}</a>
  </div>
	<a href="{{ $mediaURL }}" target="_blank"><img class="card-img-top img-fluid" src="{{ $item.thumbnail_url }}" width="{{ $item.thumbnail_width }}"  height="{{ $item.thumbnail_height }}" alt="Instagram Image"></a>
	<div class="card-body">
		{{ if not $hideCaption }}<p class="card-text"><a href="{{ $item.author_url | safeURL }}" class="card-link">{{ $item.author_name }}</a> {{ $item.title}}</p>{{ end }}
		<a href="{{ $item.author_url | safeURL }}" class="card-link">View More on Instagram</a>
	</div>
</div>
{{ end }}
{{- end -}}

{{ define "__h_simple_instagram_css" }}
{{ if not (.Page.Scratch.Get "__h_simple_instagram_css") }}
{{/* Only include once */}}
{{  .Page.Scratch.Set "__h_simple_instagram_css" true }}
<style type="text/css">
   .__h_instagram.card {
      font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif;
      font-size: 14px;
      border: 1px solid rgb(219, 219, 219);
      padding: 0;
	  margin-top: 30px;
   }
   .__h_instagram.card .card-header, .__h_instagram.card .card-body {
      padding: 10px 10px 10px;
   }
   .__h_instagram.card img {
      width: 100%;
    	height: auto;
   }
</style>
{{ end }}
{{ end }}`},
	{`shortcodes/param.html`, `{{- $name := (.Get 0) -}}
{{- with $name -}}
{{- with ($.Page.Param .) }}{{ . }}{{ else }}{{ errorf "Param %q not found: %s" $name $.Position }}{{ end -}}
{{- else }}{{ errorf "Missing param key: %s" $.Position }}{{ end -}}`},
	{`shortcodes/ref.html`, `{{ ref . .Params }}`},
	{`shortcodes/relref.html`, `{{ relref . .Params }}`},
	{`shortcodes/twitter.html`, `{{- $pc := .Page.Site.Config.Privacy.Twitter -}}
{{- if not $pc.Disable -}}
{{- if $pc.Simple -}}
{{ template "_internal/shortcodes/twitter_simple.html" . }}
{{- else -}}
{{- $url := printf "https://api.twitter.com/1/statuses/oembed.json?id=%s&dnt=%t" (index .Params 0) $pc.EnableDNT -}}
{{- $json := getJSON $url -}}
{{ $json.html | safeHTML }}
{{- end -}}
{{- end -}}`},
	{`shortcodes/twitter_simple.html`, `{{- $pc := .Page.Site.Config.Privacy.Twitter -}}
{{- $sc := .Page.Site.Config.Services.Twitter -}}
{{- if not $pc.Disable -}}
{{- $id := .Get 0 -}}
{{- $json := getJSON "https://api.twitter.com/1/statuses/oembed.json?id=" $id "&omit_script=true" -}}
{{- if not $sc.DisableInlineCSS -}}
{{ template "__h_simple_twitter_css" $ }}
{{- end -}}
{{ $json.html | safeHTML }}
{{- end -}}

{{ define "__h_simple_twitter_css" }}
{{ if not (.Page.Scratch.Get "__h_simple_twitter_css") }}
{{/* Only include once */}}
{{  .Page.Scratch.Set "__h_simple_twitter_css" true }}
<style type="text/css">
  .twitter-tweet {
  font: 14px/1.45 -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif;
  border-left: 4px solid #2b7bb9;
  padding-left: 1.5em;
  color: #555;
}
  .twitter-tweet a {
  color: #2b7bb9;
  text-decoration: none;
}
  blockquote.twitter-tweet a:hover,
  blockquote.twitter-tweet a:focus {
  text-decoration: underline;
}
</style>
{{ end }}
{{ end }}`},
	{`shortcodes/vimeo.html`, `{{- $pc := .Page.Site.Config.Privacy.Vimeo -}}
{{- if not $pc.Disable -}}
{{- if $pc.Simple -}}
{{ template "_internal/shortcodes/vimeo_simple.html" . }}
{{- else -}}
{{ if .IsNamedParams }}<div {{ if .Get "class" }}class="{{ .Get "class" }}"{{ else }}style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//player.vimeo.com/video/{{ .Get "id" }}" {{ if not (.Get "class") }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" {{ end }}webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
 </div>{{ else }}
<div {{ if len .Params | eq 2 }}class="{{ .Get 1 }}"{{ else }}style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//player.vimeo.com/video/{{ .Get 0 }}" {{ if len .Params | eq 1 }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" {{ end }}webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
 </div>
{{ end }}
{{- end -}}
{{- end -}}`},
	{`shortcodes/vimeo_simple.html`, `{{ $id := .Get "id" | default (.Get 0) }}
{{- $item := getJSON "https://vimeo.com/api/oembed.json?url=https://vimeo.com/" $id -}}
{{ $class := .Get "class" | default (.Get 1) }}
{{ $hasClass := $class }}
{{ $class := $class | default "__h_video" }}
{{ if not $hasClass }}
{{/* If class is set, assume the user wants to provide his own styles. */}}
{{ template "__h_simple_css" $ }}
{{ end }}
{{ $secondClass := "s_video_simple" }}
<div class="{{ $secondClass }} {{ $class }}">
{{- with $item }}
<a href="{{ .provider_url }}{{ .video_id }}" target="_blank">
{{ $thumb := .thumbnail_url }}
{{ $original := $thumb | replaceRE "(_.*\\.)" "." }}
<img src="{{ $thumb }}" srcset="{{ $thumb }} 1x, {{ $original }} 2x" alt="{{ .title }}">
<div class="play">{{ template "__h_simple_icon_play" $ }}</div></a></div>
{{- end -}}
`},
	{`shortcodes/youtube.html`, `{{- $pc := .Page.Site.Config.Privacy.YouTube -}}
{{- if not $pc.Disable -}}
{{- $ytHost := cond $pc.PrivacyEnhanced  "www.youtube-nocookie.com" "www.youtube.com" -}}
{{- $id := .Get "id" | default (.Get 0) -}}
{{- $class := .Get "class" | default (.Get 1) }}
<div {{ with $class }}class="{{ . }}"{{ else }}style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="//{{ $ytHost }}/embed/{{ $id }}{{ with .Get "autoplay" }}{{ if eq . "true" }}?autoplay=1{{ end }}{{ end }}" {{ if not $class }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" {{ end }}allowfullscreen title="YouTube Video"></iframe>
</div>
{{ end -}}
`},
	{`twitter_cards.html`, `{{- with $.Params.images -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ index . 0 | absURL }}"/>
{{ else -}}
{{- $images := $.Resources.ByType "image" -}}
{{- $featured := $images.GetMatch "*feature*" -}}
{{- $featured := cond (ne $featured nil) $featured ($images.GetMatch "{*cover*,*thumbnail*}") -}}
{{- with $featured -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ $featured.Permalink }}"/>
{{- else -}}
{{- with $.Site.Params.images -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ index . 0 | absURL }}"/>
{{ else -}}
<meta name="twitter:card" content="summary"/>
{{- end -}}
{{- end -}}
{{- end }}
<meta name="twitter:title" content="{{ .Title }}"/>
<meta name="twitter:description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end -}}"/>
{{ with .Site.Social.twitter -}}
<meta name="twitter:site" content="@{{ . }}"/>
{{ end -}}
{{ range .Site.Authors }}
{{ with .twitter -}}
<meta name="twitter:creator" content="@{{ . }}"/>
{{ end -}}
{{ end -}}`},
}
