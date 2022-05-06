<div id="disqus_thread"></div>
{{ if .Site.Comments.Disqus.Enable }}
<script>
(function() { // DON'T EDIT BELOW THIS LINE
    var d = document, s = d.createElement('script');
    s.src = '//{{.Site.Comments.Disqus.Site}}.disqus.com/embed.js';
    s.setAttribute('data-timestamp', +new Date());
    (d.head || d.body).appendChild(s);
})();
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
{{ end }}
{{ if .Site.Comments.Giscus.Enable }}
<script src="https://giscus.app/client.js"
        data-repo="{{.Site.Comments.Giscus.Repo}}"
        data-repo-id="{{.Site.Comments.Giscus.RepoID}}"
        data-category="{{.Site.Comments.Giscus.Category}}"
        data-category-id="{{.Site.Comments.Giscus.CategoryID}}"
        data-mapping="{{.Site.Comments.Giscus.Mapping}}"
        data-reactions-enabled="{{.Site.Comments.Giscus.ReactionsEnabled}}"
        data-theme="{{.Site.Comments.Giscus.Theme}}"
        crossorigin="anonymous"
        async>
</script>
<noscript>Please enable JavaScript to view the comments powered by giscus.</noscript>
{{ end }}
