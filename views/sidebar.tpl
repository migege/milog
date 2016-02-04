<div class="col-xs-6 col-md-3" id="sidebar">
  <ul class="list-unstyled">
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">miscellaneous</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
          {{if .LoggedUser}}
          <li><a href="/admin" title="Already logged">{{.LoggedUser.DisplayName}}</a></li>
          {{else}}
          <li><a href="/signup" title="Sign Up">Sign Up</a></li>
          <li><a href="/login" title="Log In">Log In</a></li>
          {{end}}
        </ul>
      </div>
    </li>
  </ul>
</div>
<div class="clear"></div>
