{{template "header.tpl" .}}
  <div class="col-xs-12 col-md-9" id="admin-content">
    <form action="/admin/option-edit" method="post">
      {{range .Options}}
      <div class="form-group">
        <label for="opt_{{.OptionName}}">{{if .OptionDesc}}{{.OptionDesc}}{{else}}{{.OptionName}}{{end}}</label>
        <input class="form-control" id="opt_{{.OptionName}}" name="{{.OptionName}}" value="{{.OptionValue}}" />
      </div>
      {{else}}
      {{end}}
      <button type="submit" class="btn btn-primary">Save</button>
    </form>
  </div>
{{template "admin-sidebar.tpl" .}}
{{template "footer.tpl" .}}
