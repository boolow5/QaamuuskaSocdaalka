<h1 class="column-header"><a href="#"><i class="fa fa-bolt"></i> {{.Post.Title}}</a></h1>
<div class="latest-news">

  {{i18n $.Lang "written"}} {{.Post.PublishedDate}} | {{i18n $.Lang "views"}} <span>{{.Post.Views}}</span>
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

  {{.Post.Content |markdown}}

</div>
