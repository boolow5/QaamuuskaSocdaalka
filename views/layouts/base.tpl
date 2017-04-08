<!DOCTYPE html>
<html {{if eq .Lang "ar-SA"}}lang="ar" DIR="RTL"{{else}}lang="en"{{end}}>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>{{i18n $.Lang "page title"}}</title>

    {{.favicons}}

    {{.css}}
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
    <div class="top-menu">
      {{if eq .Lang "ar-SA"}}
        {{.navbarAr}}
      {{else}}
        {{.navbar}}
      {{end}}
    </div>

    {{.featured}}


    <div class="messages">
      <div id="error" class="alert alert-danger alert-dismissible hidden" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <strong>Error!</strong> <span class="message"></span>
      </div>
      <div id="warning" class="alert alert-warning alert-dismissible hidden" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <strong>Warning!</strong> <span class="message"></span>
      </div>
      <div id="success" class="alert alert-success alert-dismissible hidden" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <strong>Success!</strong> <span class="message"></span>
      </div>
    </div>

    <div class="container  main-wrapper">
      <div class="row">
        <div class="col-sm-8 main-content min-height">
          {{.LayoutContent}}
        </div>
        {{.sidebarRight}}

      </div>
    </div>

    {{.footer}}

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/bootstrap/js/bootstrap.min.js"></script>
    <script src="static/js/jquery.cookie.js"></script>
    <script src="static/js/main.js"></script>
  </body>
</html>
