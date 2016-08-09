{{template "header.tpl" .}}
<div class="col-xs-12 col-md-9">
  <h1>{{.Author.DisplayName}}</h1>
  <form>
    <div class="form-group">
      <label class="control-label">Mail</label>
      <p class="form-control-static"><a href="mailto:{{.Author.AuthorMail}}" title="Write to me">{{.Author.AuthorMail}}</a></p>
    </div>
    <div class="form-group">
      <label class="control-label">Website</label>
      <p class="form-control-static"><a href="{{.Author.AuthorUrl}}" rel="nofollow">{{.Author.AuthorUrl}}</a></p>
    </div>
  </form>
  <h3>Posts</h3>
  <div class="posts">
  {{template "posts.tpl" .}}
  </div>
</div>
{{template "sidebar.tpl" .}}
{{template "footer.tpl" .}}
