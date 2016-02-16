{{template "header.tpl" .}}
<div class="col-xs-12 col-md-9" id="posts">
  <div class="alert alert-warning">
    <p>{{.Content}}</p>
  </div>
  {{template "posts.tpl" .}}
</div>
{{template "sidebar.tpl" .}}
{{template "footer.tpl" .}}
