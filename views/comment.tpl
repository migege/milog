  <div id="comments" class="comments">
    <h3>{{.Post.CommentCount}}&nbsp;Comment{{if ne .Post.CommentCount 1}}s{{end}}</h3>
    {{range .Comments}}
    <div class="comment" id="comment-{{.CommentId}}">
      <div class="comment-header">
        <span><a href="{{.CommentAuthorUrl}}" rel="nofollow" class="external">{{.CommentAuthor}}</a>&nbsp;@&nbsp;<a href="#comment-{{.CommentId}}">{{.CommentTime}}</a>&nbsp;<a href="javascript:void(0);" onclick="javascript:reply({{.CommentId}});">reply</a></span>
      </div>
      <div class="comment-body">
        <p>{{.CommentContent}}</p>
      </div>
    </div>
    {{else}}
    <div class="comment">
      <div class="comment-body">
        <p>No comments yet.</p>
      </div>
    </div>
    {{end}}
  </div>
  <div id="respond">
    <h3>Leave a Reply</h3>
    {{if .Post.CommentStatus}}
    <form action="/comments-add" method="post" id="commentform" class="row col-md-8">
      {{if not .LoggedUser}}
      <div class="form-group">
        <label for="comment-author" class="control-label">Name</label>
        <input class="form-control" id="comment-author" type="text" name="comment_author" value="{{.CommentAuthor}}" placeholder="Name" />
      </div>
      <div class="form-group">
        <label for="comment-author-mail" class="control-label">Email</label>
        <input class="form-control" id="comment-author-mail" type="email" name="comment_author_mail" value="{{.CommentAuthorMail}}" placeholder="Email" />
      </div>
      <div class="form-group">
        <label for="comment-author-url" class="control-label">URL</label>
        <input class="form-control" id="comment-author-url" type="url" name="comment_author_url" value="{{.CommentAuthorUrl}}" placeholder="Website URL" />
      </div>
      {{end}}
      <div class="form-group comment-form-field comment-textarea">
        <div id="comment-form-comment">
          <textarea class="form-control" rows="4" id="comment" name="comment" title="Enter your comment here..." placeholder="Enter your comment here..."></textarea>
        </div>
      </div>
      <div class="form-group comment-form-field comment-captcha">
        <p>{{create_captcha}}</p>
        <p><input class="form-control" type="text" name="captcha" placeholder="Enter 6 digits above" /></p>
      </div>
      <p class="form-submit" style="display:block">
        <button type="submit" id="comment-submit" class="btn btn-success">Post Comment</button>
        <input type="hidden" name="post_id" id="post_id" value="{{.Post.PostId}}" />
        <input type="hidden" name="comment_parent_id" id="comment_parent_id" value="0" />
      </p>
    </form>
    {{else}}
    <div class="comment">
      <div class="comment-body">
        <p>Comments are currently closed.</p>
      </div>
    </div>
    {{end}}
  </div>
