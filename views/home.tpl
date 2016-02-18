{{template "header.tpl" .}}
<div class="col-xs-12 col-md-9" id="posts">
{{if .Content}}
  <div class="alert alert-info">
    <p>{{.Content}}</p>
  </div>
{{end}}
{{template "posts.tpl" .}}
</div>
{{template "sidebar.tpl" .}}
{{template "footer.tpl" .}}
