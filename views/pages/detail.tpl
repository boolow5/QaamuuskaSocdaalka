<h1 class="column-header"><a href="#"><i class="fa fa-newspaper-o"></i> {{.Post.Title}}</a></h1>
<div class="latest-news">

  {{if .Post.FeaturedImage}}

  {{i18n $.Lang "written"}} {{format .Post.PublishedDate "03:04PM 02-01-2006"}} | {{i18n $.Lang "views"}} <span>{{.Post.Views}}</span>
  <hr>
  <div class="full-image-box">
    <div class="image" style="background-image:url({{.Post.FeaturedImage.Url}})">


      <div class="caption">
        <h4>{{.Post.FeaturedImage.Title}}</h4>
        <span>{{.Post.FeaturedImage.Description}}</span>
      </div>
    </div>
  </div>
  <hr>
  {{end}}

  {{.Post.Content |markdown}}

</div>
