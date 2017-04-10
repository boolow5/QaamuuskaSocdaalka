<nav class="navbar blue-color">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand brand-and-logo" href="#"><img src="static/favicon/favicon-96x96.png" alt="logo"> Qaamuuska</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li class="active"><a href="/bol-admin"><i class="fa fa-home fa-fw"></i> {{i18n $.Lang "admin home"}} <span class="sr-only">(current)</span></a></li>
      </ul>

      <ul class="nav navbar-nav navbar-right">
        <li>
          {{if .LoggedIn }}
            <a href="#"><i class="fa fa-user fa-fw"></i> {{i18n .Lang "welcome"}}, {{.CurrentUser|title}}</a>
          {{end}}
        </li>
        <li>
          {{if .LoggedIn }}
            <a href="#" id="logout"><i class="fa fa-sign-out fa-fw"></i> Logout</a>
          {{else}}
            <form id="login-form" class="form nav-login-form" method="post" action="/bol-admin/login">
              {{ .xsrfdata }}

                <input type="text" class="" name="username" placeholder='{{i18n .Lang "username"}}' id="login-username" autofocus>
                <input type="password" class="" name="password" placeholder='{{i18n .Lang "password"}}' id="login-email">

              <button class="btn btn-primary btn-small" id="nav-login-user-btn">{{i18n .Lang "login_btn"}}</button>
              {{i18n .Lang "or"}}
              <a class="btn btn-primary btn-small" href="/register">{{i18n .Lang "register_btn"}}</a>
            </form>
          {{end}}
        </li>
      </ul>

      <ul class="nav navbar-nav navbar-right">
        <li role="presentation" class='dropdown'>
          <a class="dropdown-toggle" data-toggle="dropdown" href="#" role="button" aria-haspopup="true" aria-expanded="false">
            <i class="fa fa-language fa-fw" aria-hidden="true"></i> {{i18n $.Lang "languages"}} <span class="caret"></span>
          </a>
          <ul class="dropdown-menu">
            {{range .RestLangs}}
                <li><a href="javascript::" data-lang="{{.Lang}}" class="lang-changed"> {{i18n $.Lang .Name}}</a></li>
            {{end}}
          </ul>
        </li>

      </ul>

    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>
