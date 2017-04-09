<div class="col-sm-3 sidebar min-height">
  <h1 class="column-header"><a href="#"> {{i18n $.Lang "drafts"}}</a></h1>
  {{if .Drats}}
  {{else}}
  <p class="text-center">No items</p>
  {{end}}

  <h1 class="column-header"><a href="#"> {{i18n $.Lang "posts"}}</a></h1>
  {{if .Posts}}
  {{else}}
  <p class="text-center">No items</p>
  {{end}}

  <h1 class="column-header"><a href="#"> {{i18n $.Lang "categories"}}</a></h1>
  {{if .Categories}}
  {{else}}
  <p class="text-center">No items</p>
  {{end}}

</div>
