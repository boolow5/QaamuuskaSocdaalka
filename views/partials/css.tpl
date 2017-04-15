<!-- Bootstrap -->
<link href="static/bootstrap/css/bootstrap.min.css" rel="stylesheet">
<link href="static/font-awesome/css/font-awesome.css" rel="stylesheet">
<link href="static/css/main.css" rel="stylesheet">

<!-- jQuery (necessary for Bootstrap's JavaScript plugins and responsiveSlides) -->
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>

<!-- Sliders -->
<link href="static/slider/responsiveslides.css" rel="stylesheet">
<script src="static/slider/responsiveslides.js"></script>
<script>
  // $(".rslides").responsiveSlides({
  //   auto: true,             // Boolean: Animate automatically, true or false
  //   speed: 100,            // Integer: Speed of the transition, in milliseconds
  //   timeout: 1000,          // Integer: Time between slide transitions, in milliseconds
  //   pager: false,           // Boolean: Show pager, true or false
  //   nav: true,             // Boolean: Show navigation, true or false
  //   random: false,          // Boolean: Randomize the order of the slides, true or false
  //   pause: false,           // Boolean: Pause on hover, true or false
  //   pauseControls: true,    // Boolean: Pause when hovering controls, true or false
  //   prevText: "Previous",   // String: Text for the "previous" button
  //   nextText: "Next",       // String: Text for the "next" button
  //   maxwidth: "",           // Integer: Max-width of the slideshow, in pixels
  //   navContainer: "",       // Selector: Where controls should be appended to, default is after the 'ul'
  //   manualControls: "",     // Selector: Declare custom pager navigation
  //   namespace: "rslides",   // String: Change the default namespace used
  //   before: function(){console.log("before event");},   // Function: Before callback
  //   after: function(){console.log("after event");}     // Function: After callback
  // });
  // You can also use "$(window).load(function() {"
    $(function () {
      // Slideshow
      $("#slider").responsiveSlides({
      auto: true,             // Boolean: Animate automatically, true or false
      pager: false,           // Boolean: Show pager, true or false
      nav: true,             // Boolean: Show navigation, true or false
      speed: 500,            // Integer: Speed of the transition, in milliseconds
      timeout: 3000,          // Integer: Time between slide transitions, in milliseconds
      pause: true,           // Boolean: Pause on hover, true or false
      prevText: " << ",   // String: Text for the "previous" button
      nextText: " >> ",       // String: Text for the "next" button
      namespace: "callbacks",   // String: Change the default namespace used
      before: function () {
        // console.log("before");
      },
      after: function () {
        // console.log("after");
      }
      });
    });
</script>
<style>
  .rslides {
    position: relative;
    list-style: none;
    overflow: hidden;
    width: 100%;
    padding: 0;
    margin: 0;
    border-top: .2em solid #020A31;
    border-right: .2em solid #020A31;
    border-left: .2em solid #020A31;
    border-top-left-radius: 1em;
    border-top-right-radius: 1em;
  }

  .rslides li {
    -webkit-backface-visibility: hidden;
    position: absolute;
    display: none;
    width: 100%;
    left: 0;
    top: 0;
  }

  .rslides li:first-child {
    position: relative;
    display: block;
    float: left;
  }

  .rslides .img {
    background-size: cover;
    background-position: center;
    display: block;
    min-height: 20em;
    max-height: 30em;
    float: left;
    width: 100%;
    border: 0;
    /* text */
    color: #FFF;
    box-shadow: 0 0 2px #000;
  }
  .rslides .img h1 {
    text-transform: uppercase;
  }
  .rslides a:hover {
    text-shadow: 0 0 2px #000;
  }
  .rslides li {
    max-height: 30em;
  }


  a.next {
    float: right;
    //border-top-left-radius: 1em;
    border-bottom-right-radius: 1em;
  }

  a.prev {
    float: left;
    border-bottom-left-radius: 1em;
    //border-top-right-radius: 1em;
  }

  a.next, a.prev {
    color: #FFF;
    font-weight: bolder;
    background-color: #020A31;
    padding-left: 1em;
    padding-right: 1em;
  }
  a.next:hover, a.prev:hover {
    text-decoration: none;

  }

</style>
