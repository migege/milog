{{template "header.tpl" .}}
<div class="col-xs-12 col-md-9">
  <div class="post" id="post_{{.Post.PostId}}">
    <div class="post-header">
      <h2>{{.Post.PostTitle}}</h2>
    </div>
    <div class="post-body">
      {{str2html .Post.PostContent}}
    </div>
    <div class="post-footer">
      <span class="label label-default"><a href="/author/{{.Author.AuthorName}}">{{.Author.DisplayName}}</a> posted @ {{.Post.PostTime}}</span>
      {{range .Post.Tags}}
      <span class="label label-info"><a href="/tag/{{.TagSlug}}">{{.TagName}}</a></span>
      {{end}}
      {{if .LoggedUser}}{{if eq .LoggedUser.AuthorId .Author.AuthorId}}<span class="label label-primary"><a href="/admin/post-edit/{{.Post.PostId}}">Edit</a></span>{{end}}{{end}}
    </div>
  </div>
{{template "comment.tpl" .}}
</div>
{{template "sidebar.tpl" .}}
{{template "footer.tpl" .}}
