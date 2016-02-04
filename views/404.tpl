{{template "header.tpl" .}}
<div class="col-xs-12 col-md-9" id="posts">
  <div>
    <p>{{.Content}}</p>
  </div>
  {{range .Posts}}
  <div class="post" id="post_{{.PostId}}">
    <div class="post-header">
      <a href="/post/{{.PostSlug}}"><h2>{{.PostTitle}}</h2></a>
    </div>
    <div class="post-body">
      {{str2html .PostContent}}
    </div>
    <div class="post-footer">
      <span class="label label-default"><a href="/author/{{.Author.AuthorId}}">{{.Author.DisplayName}}</a> posted @ {{.PostTime}}</span>
      {{range .Tags}}
      <span class="label label-info"><a href="/tag/{{.TagSlug}}">{{.TagName}}</a></span>
      {{end}}
    </div>
  </div>
  {{end}}
</div>
{{template "sidebar.tpl" .}}
{{template "footer.tpl" .}}
