{{template "header.tpl" .}}
  <div class="col-xs-12 col-md-9" id="admin-content">
    <form id="post-form" action="{{if .Post}}/admin/post-edit{{else}}/admin/post-new{{end}}" method="post">
      <div class="form-group post-form-header post-form-field">
        <h2><input class="form-control" type="text" name="post-title" placeholder="Enter title here..." value="{{.Post.PostTitle}}"></h2>
      </div>
      <div class="form-group post-form-text post-form-field">
        <p>
          <textarea class="form-control" id="post-text" name="post-content-md" oninput="this.editor.update()" rows="10" cols="60" placeholder="Type **Markdown** here.">{{if .Post}}{{.Post.PostContentMd}}{{else}}Type **Markdown** here.{{end}}</textarea>
        </p>
        <p><div id="post-preview" class="post-body"></div></p>
      </div>
      <div class="form-group">
        <label for="post-tags">Tags</label>
        <select id="post-tags" name="post-tags" class="form-control post-tags" multiple="multiple">
        {{range .Post.Tags}}
          <option selected="selected">{{.TagName}}</option>
        {{else}}
          {{range .AllTags}}
          <option>{{.TagName}}</option>
          {{end}}
        {{end}}
        </select>
      </div>
      <div class="checkbox">
        <label>
          <input type="checkbox" name="comment-status" value="0" {{if .Post}}{{if eq .Post.CommentStatus 0}}checked{{end}}{{end}} />&nbsp;Disable&nbsp;comments
        </label>
      </div>
      <p class="submit">
        <button type="submit" class="btn btn-primary">Post</button>
        <input type="hidden" id="post-id" name="post-id" value="{{.Post.PostId}}" />
        <input type="hidden" id="post-content" name="post-content" value="" />
      </p>
    </form>
  </div>
<script type="text/javascript" src="/static/js/markdown.min.js"></script>
<script type="text/javascript">
function Editor(input,preview,output){
  this.update=function(){
    preview.innerHTML=markdown.toHTML(input.value)
    output.value=preview.innerHTML
  }
  input.editor=this
  this.update()
}
new Editor(document.getElementById("post-text"),document.getElementById("post-preview"),document.getElementById('post-content'))
</script>
{{template "admin-sidebar.tpl" .}}
{{template "footer.tpl" .}}
