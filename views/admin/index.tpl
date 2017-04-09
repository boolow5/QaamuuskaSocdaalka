<h1 class="column-header"><a href="#"> {{i18n $.Lang "admin forms"}}</a></h1>
<div class="admin-forms">

  <h1 class="column-header"><a href="#"> {{i18n $.Lang "new category"}}</a></h1>

  <form id="category-form" method="post" action="/bol-admin/add/category">
    <div class="form-group">
      <input type="text" name="name" class="form-control" placeholder="Name">
      <div class="">
        <button class="btn btn-primary">Save</button>
        <input class="btn btn-warning" type="reset" value="Clear"/>
      </div>
    </div>
  </form>

  <h1 class="column-header"><a href="#"> {{i18n $.Lang "new post"}}</a></h1>

  <form id="post-form" method="post" action="/bol-admin/add/post">
    <div class="form-group">
      <input type="text" name="title" class="form-control" placeholder="Title">
      <select class="form-control" name="category">
        <option value="1">Category 1</option>
        <option value="2">Category 2</option>
        <option value="3">Category 3</option>
      </select>
      <textarea class="form-control" name="content" rows="8" cols="80"  placeholder="Content"></textarea>
      <label class="radio-inline"><input type="checkbox" checked  name="save_as_draft"> {{i18n $.Lang "save as draft"}}</label><br>
      <div class="">
        <button class="btn btn-primary">Save</button>
        <input class="btn btn-warning" type="reset" value="Clear"/>
      </div>
    </div>
  </form>

  {{if eq .CurrentUserRole "admin" }}
  <h1 class="column-header"><a href="#"> {{i18n $.Lang "new user"}}</a></h1>

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
        <button class="btn btn-primary">Save</button>
        <input class="btn btn-warning" type="reset" value="Clear"/>
      </div>
    </div>
  </form>
  {{end}}

</div>
