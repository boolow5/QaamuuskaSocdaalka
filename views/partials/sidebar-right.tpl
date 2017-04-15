<div class="col-sm-3 sidebar min-height">
  <h1 class="column-header"><a href="#"><i class="fa fa-pie-chart"></i> {{i18n $.Lang "categories"}}</a></h1>
  <ul class="list-group">
    {{range $index, $val := .Categories}}
    <li class="list-group-item">
      <span class="badge">{{$val.PostsCount}}</span>
      {{i18n $.Lang $val.Name}}
    </li>
    {{end}}
  </ul>

  <h1 class="column-header"><a href="#"><i class="fa fa-line-chart"></i> {{i18n $.Lang "most popular"}}</a></h1>

  {{range $index, $val := .MostPopularPosts}}
    <div class="thumbnail">
      <div class="thumbnail-image" style="background-image:url({{$val.FeaturedImage.Url}})">

      </div>
      <div class="caption">
        <h4>{{shorten_words $val.Title 12}}</h4>
      </div>
      <span> <i class="fa fa-eye"></i> {{$val.Views}}</span>
    </div>
  {{end}}

</div>
