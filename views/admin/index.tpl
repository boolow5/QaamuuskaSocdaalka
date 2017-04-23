<h1 class="column-header"><a href="#"> {{i18n $.Lang "admin forms"}}</a></h1>
{{if .LoggedIn}}
<div class="admin-forms">

  <div class="category-section">
    <h1 class="column-header"><a href="#"> {{i18n $.Lang "new category"}}</a></h1>

    {{if .Category.Name}}
    <form id="category-form" method="post" action="/bol-admin/update/category/{{.Category.Id}}">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="name" value='{{.Category.Name}}' class="form-control" placeholder='{{i18n $.Lang "name"}}'>
        <div class="">
          <input type="submit" class="btn btn-primary" value='{{i18n $.Lang "update"}}' />
          <input class="btn btn-danger delete-item" type="button" data-url="/bol-admin/delete/category/{{.Category.Id}}" value='{{i18n $.Lang "delete"}}'/>
        </div>
      </div>
    </form>

    {{else}}
    <form id="category-form" method="post" action="/bol-admin/add/category">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="name" class="form-control" placeholder='{{i18n $.Lang "name"}}'>
        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "save"}}</button>
          <input class="btn btn-warning" type="reset" value='{{i18n $.Lang "clear"}}'/>
        </div>
      </div>
    </form>
    {{end}}

  </div>

  <div class="image-section">
    <h1 class="column-header"><a href="#"> {{i18n $.Lang "new image"}}</a></h1>

    {{if .Image.Url}}
    <form id="image-form" method="post" enctype="multipart/form-data" action="/bol-admin/update/image/{{.Image.Id}}">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="title" class="form-control" value='{{.Image.Title}}' placeholder='{{i18n $.Lang "title"}}'>
        <input type="text" readonly name="url" class="form-control" value='{{.Image.Url}}' placeholder='{{i18n $.Lang "url"}}'>
        <textarea class="form-control" name="description" rows="8" cols="80"  placeholder='{{i18n $.Lang "description"}}'>{{.Image.Description}}</textarea>
        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "update"}}</button>
          <input class="btn btn-danger delete-item" type="button" data-url="/bol-admin/delete/image/{{.Image.Id}}" value='{{i18n $.Lang "delete"}}'/>
        </div>
      </div>
    </form>
    {{else}}
    <form id="image-form" method="post" enctype="multipart/form-data" action="/bol-admin/add/image">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="title" class="form-control" placeholder='{{i18n $.Lang "title"}}'>
        <input class="form-control" type="file" name="file">
        <textarea class="form-control" name="description" rows="8" cols="80"  placeholder='{{i18n $.Lang "description"}}'></textarea>
        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "save"}}</button>
          <input class="btn btn-warning" type="reset" value='{{i18n $.Lang "clear"}}'/>
        </div>
      </div>
    </form>
    {{end}}

  </div>

  <div class="post-section">
    <h1 class="column-header"><a href="#"> {{i18n $.Lang "new post"}}</a></h1>

    {{if .Post.Title}}
    <form id="post-form" method="post" action="/bol-admin/update/post/{{.Post.Id}}">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="title" value='{{.Post.Title}}' class="form-control" placeholder='{{i18n $.Lang "title"}}'>
        <select class="form-control" name="category">
          {{range $val := .Categories}}
          <option {{if eq $.Post.Category.Id $val.Id}}selected="true"{{end}} value="{{$val.Id}}">{{i18n $.Lang $val.Name}}</option>
          {{end}}
        </select>
        <select class="form-control" name="language">
          {{range $val := .AllLangs}}
          <option {{if eq $.Post.Language $val.Name}}selected="true"{{end}} value="{{$val.Name}}">{{i18n $.Lang $val.Name}}</option>
          {{end}}
        </select>
        <select class="form-control" name="featured_image">
          {{range $val := .Images}}
          <option {{if eq $.Post.FeaturedImage.Id $val.Id}}selected="true"{{end}} value="{{$val.Id}}">{{i18n $.Lang $val.Title}}</option>
          {{end}}
        </select>
        <textarea class="form-control" name="content" rows="8" cols="80"  placeholder="Content">{{.Post.Content}}</textarea>
        <label class="radio-inline"><input type="checkbox" {{if .Post.SaveAsDraft}}checked{{end}}  name="save_as_draft"> {{i18n $.Lang "save as draft"}}</label><br>
        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "update"}}</button>
          <input class="btn btn-danger delete-item" type="button" data-url="/bol-admin/delete/post/{{.Post.Id}}" value='{{i18n $.Lang "delete"}}'/>
        </div>
      </div>
    </form>
    {{else}}
    <form id="post-form" method="post" action="/bol-admin/add/post">
      <div class="form-group">
        {{ .xsrfdata }}
        <input type="text" name="title" class="form-control" placeholder='{{i18n $.Lang "title"}}'>
        <select class="form-control" name="category">
          <option value="0">{{i18n $.Lang "select category"}}</option>
          {{range $val := .Categories}}
          <option value="{{$val.Id}}">{{i18n $.Lang $val.Name}}</option>
          {{end}}
        </select>
        <select class="form-control" name="language">
          <option value="0">{{i18n $.Lang "select language"}}</option>
          {{range $val := .AllLangs}}
          <option value="{{$val.Name}}">{{i18n $.Lang $val.Name}}</option>
          {{end}}
        </select>
        <select class="form-control" name="featured_image">
          <option value="0">{{i18n $.Lang "select image"}}</option>
          {{range $val := .Images}}
          <option value="{{$val.Id}}">{{i18n $.Lang $val.Title}}</option>
          {{end}}
        </select>
        <textarea class="form-control" name="content" rows="8" cols="80"  placeholder="Content"></textarea>
        <label class="radio-inline"><input type="checkbox" checked  name="save_as_draft"> {{i18n $.Lang "save as draft"}}</label><br>
        <div class="">
          <button class="btn btn-primary">{{i18n $.Lang "save"}}</button>
          <input class="btn btn-warning" type="reset" value='{{i18n $.Lang "clear"}}'/>
        </div>
      </div>
    </form>
    {{end}}

  </div>

  <div class="user-section">
    {{if .IsAdmin }}
    <h1 class="column-header"><a href="#"> {{i18n $.Lang "new user"}}</a></h1>

      {{if .User.Role}}
      <form id="user-form" method="post" action="/bol-admin/update/user/{{.User.Id}}">
        <div class="form-group">
          {{ .xsrfdata }}
          <input type="text" name="first_name" value='{{.User.Profile.FirstName}}' class="form-control" placeholder="First Name" required>
          <input type="text" name="middle_name" value='{{.User.Profile.MiddleName}}' class="form-control" placeholder="Middle Name">
          <input type="text" name="last_name" value='{{.User.Profile.LastName}}' class="form-control" placeholder="Last Name">
          <input type="email" name="email" value='{{.User.Profile.Email}}' class="form-control" placeholder="Email">
          <input type="password" name="password" class="form-control" placeholder="New Password" required>
          <label class="radio-inline"><input type="checkbox" name="admin"> {{i18n $.Lang "admin"}}</label><br>
          <div class="">
            <button class="btn btn-primary">{{i18n $.Lang "save"}}</button>
          </div>
        </div>
      </form>
      {{else}}
      <form id="user-form" method="post" action="/bol-admin/add/user">
        <div class="form-group">
          {{ .xsrfdata }}
          <input type="text" name="first_name" class="form-control" placeholder="First Name" required>
          <input type="text" name="middle_name" class="form-control" placeholder="Middle Name">
          <input type="text" name="last_name" class="form-control" placeholder="Last Name">
          <input type="email" name="email" class="form-control" placeholder="Email">
          <input type="text" name="username" class="form-control" placeholder="userame" required>
          <input type="password" name="password" class="form-control" placeholder="Password" required>
          <label class="radio-inline"><input type="checkbox" name="admin"> {{i18n $.Lang "admin"}}</label><br>
          <div class="">
            <button class="btn btn-primary">{{i18n $.Lang "save"}}</button>
            <input class="btn btn-warning" type="reset" value='{{i18n $.Lang "clear"}}'/>
          </div>
        </div>
      </form>
      {{end}}
    {{end}}

  </div>

</div>
{{else}}
<p>
  {{i18n $.Lang "login required notice"}}
</p>
{{end}}
