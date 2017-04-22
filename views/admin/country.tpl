<h1 class="column-header"><a href="#"> {{i18n $.Lang "admin forms"}}</a></h1>
{{if .LoggedIn}}
<div class="admin-forms">


  <div class="country-section">
    <h1 class="column-header"><a href="#"> {{i18n $.Lang "new country"}}</a></h1>

    {{if .Country}}
    <form id="category-form" method="post" action="/bol-admin/update/country">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="name" value='{{.Country.Name}}' class="form-control" placeholder='{{i18n $.Lang "name"}}'>
        <select class="form-control" name="capital_city">
          <option value="0">{{i18n $.Lang "select capital city"}}</option>
          {{range $val := .Cities}}
          <option value="{{$val.Id}}">{{$val.Name}}</option>
          {{end}}
        </select>
        <input type="text" value="{{.Country.Citizen}}" name="citizen" class="form-control" placeholder='{{i18n $.Lang "citizen"}}'>
        <input type="text" value="{{.Country.OfficialLanguages}}" name="official_languages" class="form-control" placeholder='{{i18n $.Lang "official languages"}}'>
        <input type="text" value="{{.Country.NorthernBorder}}" name="northern_border" class="form-control" placeholder='{{i18n $.Lang "northern border"}}'>
        <input type="text" value="{{.Country.EasternBorder}}" name="eastern_border" class="form-control" placeholder='{{i18n $.Lang "eastern border"}}'>
        <input type="text" value="{{.Country.SouthernBorder}}" name="southern_border" class="form-control" placeholder='{{i18n $.Lang "southern border"}}'>
        <input type="text" value="{{.Country.WesternBorder}}" name="western_border" class="form-control" placeholder='{{i18n $.Lang "western border"}}'>
        <input type="text" value="{{.Country.Location}}" name="location" class="form-control" placeholder='{{i18n $.Lang "location"}}'>
        <input type="text" value="{{.Country.Population}}" name="population" class="form-control" placeholder='{{i18n $.Lang "population"}}'>
        <input type="text" value="{{.Country.Area}}" name="area" class="form-control" placeholder='{{i18n $.Lang "area"}}'>
        <input type="text" value="{{.Country.AverageCostOfLiving}}" name="cost_of_living" class="form-control" placeholder='{{i18n $.Lang "average cost of living"}}'>
        <input type="text" value="{{.Country.AverageVisaCost}}" name="average_visa_cost" class="form-control" placeholder='{{i18n $.Lang "average_visa_cost"}}'>
        <input type="text" value="{{.Country.NaturalizationPeriodLength}}" name="naturalization_period_length" class="form-control" placeholder='{{i18n $.Lang "naturalization period length"}}'>
        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "update"}}</button>
          <input class="btn btn-warning" type="reset" value='{{i18n $.Lang "clear"}}'/>
        </div>
      </div>
    </form>
    {{else}}
    <form id="category-form" method="post" action="/bol-admin/add/country">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="name" class="form-control" placeholder='{{i18n $.Lang "name"}}'>
        <select class="form-control" name="capital_city">
          <option value="0">{{i18n $.Lang "select city"}}</option>
          {{range $val := .Cities}}
          <option value="{{$val.Id}}">{{i18n $.Lang $val.Name}}</option>
          {{end}}
        </select>
        <input type="text" name="citizen" class="form-control" placeholder='{{i18n $.Lang "citizen"}}'>
        <input type="text" name="official_languages" class="form-control" placeholder='{{i18n $.Lang "official languages"}}'>
        <input type="text" name="northern_border" class="form-control" placeholder='{{i18n $.Lang "northern border"}}'>
        <input type="text" name="eastern_border" class="form-control" placeholder='{{i18n $.Lang "eastern border"}}'>
        <input type="text" name="southern_border" class="form-control" placeholder='{{i18n $.Lang "southern border"}}'>
        <input type="text" name="western_border" class="form-control" placeholder='{{i18n $.Lang "western border"}}'>
        <input type="text" name="location" class="form-control" placeholder='{{i18n $.Lang "location"}}'>
        <input type="text" name="population" class="form-control" placeholder='{{i18n $.Lang "population"}}'>
        <input type="text" name="area" class="form-control" placeholder='{{i18n $.Lang "area"}}'>
        <input type="text" name="cost_of_living" class="form-control" placeholder='{{i18n $.Lang "average cost of living"}}'>
        <input type="text" name="average_visa_cost" class="form-control" placeholder='{{i18n $.Lang "average_visa_cost"}}'>
        <input type="text" name="naturalization_period_length" class="form-control" placeholder='{{i18n $.Lang "naturalization period length"}}'>

        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "save"}}</button>
          <input class="btn btn-warning" type="reset" value='{{i18n $.Lang "clear"}}'/>
        </div>
      </div>
    </form>
    {{end}}

  </div>



</div>
{{else}}
<p>
  {{i18n $.Lang "login required notice"}}
</p>
{{end}}
