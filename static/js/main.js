// change locale and reload page
$(document).on('click', '.lang-changed', function(){
  var $e = $(this);
  var lang = $e.data('lang');
  $.cookie('lang', lang, {path: '/', expires: 365});
  window.location.reload();
});

var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        //var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var xsrftoken   = $('input[name=_xsrf]').attr('value');

        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
        }
        options.headers = headers;
        return ajax(url, options);
    }
});


function getFormData(formId){
  frm = document.getElementById(formId);
  var data = {};
  for (var i=0, ii = frm.length; i< ii; ++i) {
    var input = frm[i];
    if (input.name) {
      data[input.name] = input.value;
    }
  }
  console.log(data);
  return data;
}

function translate(result, ids) {
  returned_word = ""
  if (result["error"]) {
    returned_word = result["error"]
  } else if (result["success"]) {
    returned_word = result["success"]
  }
  $.ajax({
    url: "/translate/"+returned_word+"/-1",
    type: "GET",
    contentType: "application/json",
    success: function(another_result) {
      console.log(another_result);

      if (returned_word) {
        for (var i=0; i<ids.length; i++) {
          $(ids[i]).html(another_result["meaning"]);
        }
      }
    }
  });
}

$("#logout").on("click", function(e){
  e.preventDefault();
  $.ajax({
    url: "/logout",
    success: function(result) {
      console.log(result);
      //window.location.href = "http://localhost:8080";
      if (result["error"]) {
        translate(result, ["#error > .message"])
        ShowMessage("error");
      } else if (result["success"]) {
        translate(result, ["#success > span.message"])
        ShowMessage("#success");
        reloadAfter(5000)
      }
    }
  });
});

// login form
$("#login-form").submit(function(e){
  e.preventDefault();

  console.log("Submiting login");
  console.log(this.action);

  $.ajax({
    url: this.action,
    type: this.method,
    data: JSON.stringify(getFormData("login-form")),
    contentType: "application/json",
    success: function(result) {
      console.log(result);
      //window.location.href = "http://localhost:8080";
      if (result["error"]) {
        ShowMessage("error");
        translate(result, ["#error > .message"])

      } else if (result["success"]) {
        translate(result, ["#success > span.message"])
        ShowMessage("#success");
        $("#login-form").html('');
        console.log("after 3 seconds the page will reload");
        reloadAfter(5000)
        console.log("waited for 3 seconds");
      }
    }
  });
});

function reloadAfter(seconds) {
  if (typeof seconds == "number") {
    window.setTimeout(function(){
      document.location.reload();
    }, seconds)
  }
}

function ShowMessage(id) {
  if (id == "#success") {
    $("#success").removeClass("hidden");
    $("#error").addClass("hidden");
  } else if (id == "#error") {
    $("#error").removeClass("hidden");
    $("#success").addClass("hidden");
  }
}

$(document).ready(function(e) {
  console.log("document is ready");
})
