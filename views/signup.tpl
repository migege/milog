{{template "header.tpl" .}}
  <div class="col-xs-2 col-md-4"></div>
  <div class="col-xs-8 col-md-4" id="login">
  {{if .LoggedUser}}
    <div class="logged">
      <p>Already logged as {{.LoggedUser.DisplayName}}. (<a href="javascript:void(0);" onclick="javascript:logout();">Logout</a>)</p>
    </div>
  {{else}}
    <form id="loginform" action="/login" method="post">
      <div class="form-group">
        <label for="user_login">Username</label>
        <input type="text" name="log" id="user_login" class="form-control" value="" size="20" placeholder="Username" />
      </div>
      <div class="form-group">
        <label for="user_pass">Password</label>
        <input type="password" name="pwd" id="user_pass" class="form-control" value="" size="20" placeholder="Password" />
      </div>
      <p class="submit">
        <button type="submit" id="loginsubmit" class="btn btn-primary">Log In</button>
        <input type="hidden" name="ts" id="logints" value="{{.TimeStamp}}" />
        <input type="hidden" name="redirect" id="loginredirect" value="{{.Refer}}" />
      </p>
      <p><span id="loginerror"></span></p>
    </form>
  {{end}}
  </div>
<script type="text/javascript" src="/static/js/hmac-sha256.js"></script>
<script type="text/javascript" src="/static/js/md5.js"></script>
{{template "footer.tpl" .}}
