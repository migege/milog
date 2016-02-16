{{template "header.tpl" .}}
<div class="col-xs-12 col-md-9" id="admin-content">
  <h3>All Posts - {{.LoggedUser.DisplayName}}</h3>
  <table class="table table-striped">
    <thead>
      </th>
        <th>#</th>
        <th>Title</th>
        <th>Status</th>
        <th>Created</th>
        <th>Modified</th>
      </tr>
    </thead>
    <tbody>
    {{range .Posts}}
      <tr>
        <td>{{.PostId}}</td>
        <td><a href="/post/{{.PostSlug}}" title="{{.PostTitle}}">{{.PostTitle}}</a></td>
        {{if eq .PostStatus -1}}
        <td><a title="Edit" href="/admin/post-edit/{{.PostId}}" class="btn btn-xs btn-danger">Deleted</a></td>
        {{else}}
        <td><a title="Edit" href="/admin/post-edit/{{.PostId}}" class="btn btn-xs btn-success">Posted</a></td>
        {{end}}
        <td>{{.PostTime}}</td>
        <td>{{.PostModifiedTime}}</td>
      </tr>
    {{end}}
    </tbody>
  </table>
</div>
{{template "admin-sidebar.tpl" .}}
{{template "footer.tpl" .}}
