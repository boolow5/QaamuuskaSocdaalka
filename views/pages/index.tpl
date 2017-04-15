<h1 class="column-header"><a href="#"><i class="fa fa-bolt"></i> {{i18n $.Lang "latest"}}</a></h1>
<div class="latest-news">

  <div class="container-fluid">
    {{range $index, $val := .LatestPosts}}
    <div class="col-sm-6 col-md-4">
      <div class="thumbnail">
        <div class="thumbnail-image" style="background-image:url({{$val.FeaturedImage.Url}})">

        </div>
        <div class="caption">
          <h4>{{shorten_words $val.Title 10}}</h4>
        </div>
        <span>{{$val.Views}}</span>
      </div>
    </div>
    {{end}}

  </div>


  <h1 class="column-header"><a href="#"><i class="fa fa-bars"></i> {{i18n $.Lang "other posts"}}</a></h1>
  <div class="container-fluid">
    {{range $index, $val := .Posts}}
    <div class="col-sm-6 col-md-4 col-lg-3">
      <div class="thumbnail">
        <div class="thumbnail-image" style="background-image:url({{$val.FeaturedImage.Url}})">

        </div>
        <div class="caption">
          <h5>{{shorten_words $val.Title 10}}</h5>
        </div>
        <span>{{$val.Views}}</span>
      </div>
    </div>
    {{end}}

  </div>



</div>
