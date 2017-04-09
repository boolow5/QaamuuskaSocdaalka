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
        <li class="active"><a href="/"><i class="fa fa-home fa-fw"></i> {{i18n $.Lang "home"}} <span class="sr-only">(current)</span></a></li>
        <li><a href="/countries"><i class="fa fa-globe fa-fw"></i> {{i18n $.Lang "countries"}}</a></li>
        <li><a href="/cities"><i class="fa fa-map-marker fa-fw"></i> {{i18n $.Lang "cities"}}</a></li>
        <li><a href="/contactus"><i class="fa fa-phone fa-fw"></i> {{i18n $.Lang "contact us"}}</a></li>
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
