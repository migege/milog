<div class="col-xs-6 col-md-3" id="sidebar">
  <ul class="list-unstyled">
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">search</span>
      </div>
      <div class="widget-body panel-body">
        <form class="form-inline" id="searchform">
          <div class="form-group">
            <input type="text" class="form-control" id="searchterm" placeholder="Type here..." />
          </div>
          <button class="btn btn-primary" type="submit">Go</button>
        </form>
        <script type="text/javascript">
        function gosearch(){
          location.href="/search/"+$('#searchterm').val().replace(/ /g,'+')
        }
        $('#searchform').submit(function(){
          gosearch()
          return false
        })
        </script>
      </div>
    </li>
    {{if .MostPopular}}
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">most popular</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
        {{range .MostPopular}}
          <li><span class="label label-default">{{.Views}}</span>&nbsp;<a href="/post/{{.Post.PostSlug}}" title="{{.Post.PostTitle}}">{{.Post.PostTitle}}</a></li>
        {{end}}
        </ul>
      </div>
    </li>
    {{end}}
    {{if .LatestComments}}
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">latest comments</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
        {{range .LatestComments}}
          <li><a href="/post/{{.Post.PostSlug}}#comment-{{.CommentId}}" title="{{.Post.PostTitle}}">{{substr .CommentContent 0 30}}</a>&nbsp;by&nbsp;{{.CommentAuthor}}</li>
        {{end}}
        </ul>
      </div>
    </li>
    {{end}}
    {{if .TagCloud}}
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">tag cloud</span>
      </div>
      <div class="widget-body panel-body">
        {{range .TagCloud}}
        <a href="/tag/{{.TagSlug}}" title="{{.Counts}} posts" style="font-size:{{add .Counts 9}}px">{{.TagName}}</a>
        {{end}}
      </div>
    </li>
    {{end}}
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">feed</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
          <li><a href="/feed" title="RSS Feed">RSS</a></li>
        </ul>
      </div>
    </li>
    {{if .Links}}
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">links</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
        {{range .Links}}
          <li><a rel="nofollow" href="{{.LinkUrl}}" title="{{.LinkDesc}}" class="external">{{.LinkText}}</a></li>
        {{end}}
        </ul>
      </div>
    </li>
    {{end}}
    <li class="widget panel panel-default">
      <div class="widget-header panel-heading">
        <span class="text-uppercase">miscellaneous</span>
      </div>
      <div class="widget-body panel-body">
        <ul class="list-unstyled">
          {{if .LoggedUser}}
          <li><a href="/admin" title="Already logged">{{.LoggedUser.DisplayName}}</a></li>
          {{else}}
          <!--li><a href="/signup" title="Sign up">Sign up</a></li-->
          <li><a href="/login" title="Log In">Log In</a></li>
          {{end}}
        </ul>
      </div>
    </li>
  </ul>
</div>
<div class="clear"></div>
