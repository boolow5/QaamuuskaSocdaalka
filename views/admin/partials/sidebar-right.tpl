<div class="col-sm-3 sidebar min-height">
  <h1 class="column-header"><a href="#"><i class="fa fa-pencil-square-o fa-fw"></i> {{i18n $.Lang "drafts"}}</a></h1>
  {{if .Drafts}}
    <ul>
    {{range $index, $val := .Drafts}}
      {{if lessthan $index 10 }}
      <li>{{$val.Title}}</li>
      {{end}}
    {{end}}
    </ul>
  {{else}}
  <p class="text-center">{{i18n $.Lang "no items"}}</p>
  {{end}}

  <h1 class="column-header"><a href="#"><i class="fa fa-newspaper-o fa-fw"></i> {{i18n $.Lang "posts"}}</a></h1>
  {{if .Posts}}
    <ul>
    {{range $index, $val := .Posts}}
      {{if lessthan $index 10 }}
      <li>{{$val.Title}}</li>
      {{end}}
    {{end}}
    </ul>
  {{else}}
  <p class="text-center">{{i18n $.Lang "no items"}}</p>
  {{end}}

  <h1 class="column-header"><a href="#"><i class="fa fa-th-list fa-fw"></i> {{i18n $.Lang "categories"}}</a></h1>
  {{if .Categories}}
    <ul>
    {{range $index, $val := .Categories}}
      {{if lessthan $index 10 }}
      <li>{{$val.Name}}</li>
      {{end}}
    {{end}}
    </ul>
  {{else}}
  <p class="text-center">{{i18n $.Lang "no items"}}</p>
  {{end}}

  <h1 class="column-header"><a href="#"><i class="fa fa-picture-o fa-fw"></i> {{i18n $.Lang "images"}}</a></h1>
  {{if .Images}}
    <ul>
    {{range $index, $val := .Images}}
      {{if lessthan $index 10 }}
      <li>{{$val.Title}}</li>
      {{end}}
    {{end}}
    </ul>
  {{else}}
  <p class="text-center">{{i18n $.Lang "no items"}}</p>
  {{end}}

  <h1 class="column-header"><a href="#"><i class="fa fa-users fa-fw"></i> {{i18n $.Lang "users"}}</a></h1>
  {{if .Users}}
    <ul>
    {{range $val := .Users}}
      <li>{{$val.Username}}</li>
    {{end}}
    </ul>
  {{else}}
  <p class="text-center">{{i18n $.Lang "no items"}}</p>
  {{end}}

</div>
