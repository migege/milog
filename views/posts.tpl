  {{range .Posts}}
  <div class="post" id="post_{{.PostId}}">
    <div class="post-header">
      <h1><a href="/post/{{.PostSlug}}">{{.PostTitle}}</a></h1>
    </div>
    <div class="post-body">
      {{str2html .PostContent}}
    </div>
    <div class="post-footer">
      <span class="label label-default"><a href="/author/{{.Author.AuthorName}}">{{.Author.DisplayName}}</a> posted @ {{.PostTime}}</span>
      {{if $.Views}}<span class="label label-warning">{{index $.Views .PostId}}&nbsp;views</span>{{end}}
      <a href="/post/{{.PostSlug}}#comments" class="label label-success">{{.CommentCount}}&nbsp;{{if eq .CommentCount 1}}comment{{else}}comments{{end}}</a>
      {{range .Tags}}
      <a href="/tag/{{.TagSlug}}" class="label label-info">{{.TagName}}</a>
      {{end}}
    </div>
  </div>
  {{end}}
  {{template "paginator.tpl" .}}
