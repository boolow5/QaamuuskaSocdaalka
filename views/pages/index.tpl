<h1 class="column-header"><a href="#">{{i18n $.Lang "news section"}}</a></h1>
<div class="latest-news">

  <ul class="list-group">
    {{range $val := .NewsItems}}
      <li class="list-group-item">
        <a target="_blank" href="{{$val.WebsiteUrl}}">{{i18n $.Lang "website %s" $val.WebsiteName}}</a> {{i18n $.Lang "wrote"}} {{timeSince $val.CreatedAt $.Lang}} {{i18n $.Lang "before"}}:<br />
        <h3><a href="{{$val.Link}}">{{$val.Title}}</a></h3>
      </li>
    {{end}}
  </ul>

</div>
