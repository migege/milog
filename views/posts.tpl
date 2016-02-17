  {{range .Posts}}
  <div class="post" id="post_{{.PostId}}">
    <div class="post-header">
      <a href="/post/{{.PostSlug}}"><h2>{{.PostTitle}}</h2></a>
    </div>
    <div class="post-body">
      {{str2html .PostContent}}
    </div>
    <div class="post-footer">
      <span class="label label-default"><a href="/author/{{.Author.AuthorName}}">{{.Author.DisplayName}}</a> posted @ {{.PostTime}}</span>
      <span class="label label-success">{{index $.Views .PostId}}&nbsp;views&nbsp;&amp;&nbsp;<a href="/post/{{.PostSlug}}#comments">{{.CommentCount}}&nbsp;{{if eq .CommentCount 1}}comment{{else}}comments{{end}}</a></span>
      {{range .Tags}}
      <span class="label label-info"><a href="/tag/{{.TagSlug}}">{{.TagName}}</a></span>
      {{end}}
    </div>
  </div>
  {{end}}
  {{template "paginator.tpl" .}}
