<div class="col-xs-6 col-md-3" id="sidebar">
  <ul class="list-unstyled">
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">posts</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
          <li><a href="/admin/posts" title="All posts">All Posts <span class="badge">{{.PostCount}}</span></a></li>
          <li><a href="/admin/post-new" title="New post">New Post</a></li>
        </ul>
      </div>
    </li>
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">options</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
          <li><a href="/admin/options" title="Options">Options</a></li>
        </ul>
      </div>
    </li>
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">miscellaneous</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
          <li><a href="javascript:void(0);" onclick="javascript:logout();" title="Click to log out.">Log out</a>&nbsp;{{.LoggedUser.DisplayName}}</li>
        </ul>
      </div>
    </li>
  </ul>
</div>
<div class="clear"></div>
