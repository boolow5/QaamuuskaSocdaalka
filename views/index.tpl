<!DOCTYPE html>
<!--[if lt IE 7 ]><html class="ie ie6" lang="en"> <![endif]-->
<!--[if IE 7 ]><html class="ie ie7" lang="en"> <![endif]-->
<!--[if IE 8 ]><html class="ie ie8" lang="en"> <![endif]-->
<!--[if (gte IE 9)|!(IE)]><!--><html lang="en"> <!--<![endif]-->
<head>

    <!-- Basic Page Needs
  ================================================== -->
	<meta charset="utf-8">
	<title>{{.Title}} | Qaamuuska Socdaalka!</title>
	<meta name="description" content="Free Responsive Html5 Css3 Templates | zerotheme.com">
	<meta name="author" content="www.zerotheme.com">

    <!-- Mobile Specific Metas
  ================================================== -->
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

    <!-- CSS
  ================================================== -->
  	<link rel="stylesheet" href="/static/css/zerogrid.css">
	<link rel="stylesheet" href="/static/css/style.css">
	<link rel="stylesheet" href="/static/css/responsiveslides.css">

	<!-- Custom Fonts -->
	<link href="/static/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">

	<script src="/static/js/jquery-latest.min.js"></script>
	<script src="/static/js/script.js"></script>
    <script src="/static/js/jquery183.min.js"></script>
    <script src="/static/js/responsiveslides.min.js"></script>
    <script>
		// You can also use "$(window).load(function() {"
		$(function () {
		  // Slideshow
		  $("#slider").responsiveSlides({
			auto: true,
			pager: false,
			nav: true,
			speed: 500,
			namespace: "callbacks",
			before: function () {
			  $('.events').append("<li>before event fired.</li>");
			},
			after: function () {
			  $('.events').append("<li>after event fired.</li>");
			}
		  });
		});
	</script>


	<!--[if lt IE 8]>
       <div style=' clear: both; text-align:center; position: relative;'>
         <a href="http://windows.microsoft.com/en-US/internet-explorer/products/ie/home?ocid=ie6_countdown_bannercode">
           <img src="http://storage.ie6countdown.com/assets/100/images/banners/warning_bar_0000_us.jpg" border="0" height="42" width="820" alt="You are using an outdated browser. For a faster, safer browsing experience, upgrade for free today." />
        </a>
      </div>
    <![endif]-->
    <!--[if lt IE 9]>
		<script src="/static/js/html5.js"></script>
		<script src="/static/js/css3-mediaqueries.js"></script>
	<![endif]-->

</head>
<body>
<div class="wrap-body">

<!--////////////////////////////////////Header-->
<header >
	<div class="wrap-header zerogrid">
		<div class="row">
			<div id="cssmenu">
				<ul>
				   <li class='active'><a href="index.html">Home</a></li>
				   <li><a href="archive.html">Countries</a></li>
				   <li><a href="single.html">Cities</a></li>
				   <li><a href="contact.html">Contact</a></li>
				</ul>
			</div>
			<a href='/' class="logo"><img src="/static/images/logo.png" /></a>
		</div>
	</div>
</header>
<div class="slider">
	<!-- Slideshow -->
	<div class="callbacks_container" >
		<ul class="rslides" id="slider">
			<li>
				<img src="/static/images/slideshow-image1.jpg" alt="">
				<div class="caption">
					<h1>Dalka Moldova Maxaad Ka Taqaan?</h1>
					<span >Lorem ipsum dolor sit amet</span>
				</div>
			</li>
			<li>
				<img src="/static/images/slideshow-image2.jpg" alt="">
				<div class="caption">
					<h1>Dalka Ugu Sharciga Adag?</h1>
					<span >Lorem ipsum dolor sit amet</span>
				</div>
			</li>
			<li>
				<img src="/static/images/slideshow-image3.jpg" alt="">
				<div class="caption">
					<h1>Trump iyo Somalida Yaa Saxan?</h1>
					<span >Lorem ipsum dolor sit amet</span>
				</div>
			</li>
		</ul>
	</div>
	<div class="clear"></div>
</div>


<!--////////////////////////////////////Container-->
<section id="container">
	<div class="wrap-container">
		<section class="content-box box-1">
			<div class="zerogrid">
				<div class="header">
					<h2 class="heading">
						<span>Warbixinada Ugu Shidan</span>
					</h2>
					<p>Halkan waxa la isugu keenay warbixinada loogu akhriska badan yahay. </p>
				</div>
				<div class="row"><!--Start Box-->
					<div class="col-1-3">
						<div class="wrap-col item">
							<div class="zoom-container">
								<img src="/static/images/banner-img1.jpg" />
							</div>
							<div class="item-content">
								<span>XUDUUDAHA UGU FUDUD EE LAGU SAFRO.</span>
								<p>Warbixintan waxa ay is barbar dhigeysaa xuduudaha ay aadka u fududahay inaad ku safarto adigoon la kulmin dhibaato
                  dhanka xadidaada sharciga iyo gumeysiga Soomaalida u khaaska ah.</p>
								<a class="btn" href="single.html">Sii Akhri</a>
							</div>
						</div>
					</div>
					<div class="col-1-3">
						<div class="wrap-col item">
							<div class="zoom-container">
								<img src="/static/images/banner-img2.jpg" />
							</div>
							<div class="item-content">
								<span>DALALKA UGU SOOMAALI NACEYBKA BADAN.</span>
								<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.
								Et vel quod legimus, graeci electram ocurreret at his. Vix at tation facete impetus omnesque ius harum antiopam.</p>
								<a class="btn" href="single.html">More Details</a>
							</div>
						</div>
					</div>
					<div class="col-1-3">
						<div class="wrap-col item">
							<div class="zoom-container">
								<img src="/static/images/banner-img3.jpg" />
							</div>
							<div class="item-content">
								<span>WADAMADA DIHAN BALSE AADKA U HORUMARAY.</span>
								<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.
								Et vel quod legimus, graeci electram ocurreret at his. Vix at tation facete impetus omnesque ius harum antiopam.</p>
								<a class="btn" href="single.html">More Details</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>
		<section class="content-box box-2"><!--Start Box-->
			<div class="zerogrid">
				<div class="row">
					<div class="header">
						<h2 class="heading">
							<span>Soo Dhawoow</span>
						</h2>
					</div>
					<p>Halkan waxa aad ka heleysaa warbixinada xulka ah ee loogu talagalay dadlka doonaya inay u safraan aduunka.</p>
				</div>
			</div>
		</section>
		<section class="content-box box-3"><!--Start Box-->
			<div class="zerogrid">
				<div class="row">
					<div class="header">
						<h2 class="heading">
							<span>HAGAHA SOCAALKA</span>
						</h2>
					</div>
					<div class="post">
						<div class="col-1-2">
							<img src="/static/images/23.jpg"/>
						</div>
						<div class="col-1-2">
							<div class="wrapper">
								<h3>Iska ilaali arimahan marka aad hawada ku safreyso</h3>
								<p>Warbixintan waxa aad ka akhrisan doontaa qaladaadka ugu badan ee inta badan ka dhaca dadka gaar ahaan Soomaalida<br>
                  oo keeni kara inaad xabsi ku gasho ama lagaa reebo duulimaadki aad ku qorneyd.</p>
								<a class="btn" href="#">Read More</a>
							</div>
						</div>
					</div>
					<div class="post">
						<div class="col-1-2 f-right">
							<img src="/static/images/24.jpg" />
						</div>
						<div class="col-1-2">
							<div class="wrapper">
								<h3>Garoomada ugu dhib badan ee Soomaalid ku safraan</h3>
								<p>Halkan waxaad ka aqrisan doontaa liiska garoomada diyaaradaha ee dhibka ugu badan kala kulmaan dadka Soomaalida <br>
                  Garoomadaan ayaa intabadan dadka Soomaalida si gaara u eegta ayagoo inta badan fiirinaya in qofka uu iska caabiyo <br>
                  Weerarka nafsiga ah ee lagu khalqal gelinayo qorka musaafirka ah ee doonaya inuu garoonkaas ka duulo.</p>
								<a class="btn" href="#">Read More</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>
		<section class="content-box box-4"><!--Start Box-->
			<div class="zerogrid">
				<div class="row">
					<div class="header">
						<h2 class="heading">
							<span>Kala Soco Halkan</span>
						</h2>
					</div>
					<div class="col-1-4">
						<div class="wrap-col item">
							<i class="fa fa-bar-chart-o"></i>
							<h3>Qiimeynta Baasaboorada</h3>
							<p>Halkan Kala socodo darajada iyo qiimeynta baasaboorada caalamka iyo heerarka wadamadu uga jiraan liiskan.</p>
						</div>
					</div>
					<div class="col-1-4">
						<div class="wrap-col item">
							<i class='fa fa-road'></i>
							<h3>Xuduudaha Caalamka</h3>
							<p>Sed ut perspiciatis unde om nis natus error sit volup atem accusant dolorem que laudantium. Totam aperiam, eaque ipsa quae ai.</p>
						</div>
					</div>
					<div class="col-1-4">
						<div class="wrap-col item">
							<i class='fa fa-flag-checkered'></i>
							<h3>Winning Culture</h3>
							<p>Sed ut perspiciatis unde om nis natus error sit volup atem accusant dolorem que laudantium. Totam aperiam, eaque ipsa quae ai.</p>
						</div>
					</div>
					<div class="col-1-4">
						<div class="wrap-col item">
							<i class='fa fa-dashboard'></i>
							<h3>Top Performance</h3>
							<p>Sed ut perspiciatis unde om nis natus error sit volup atem accusant dolorem que laudantium. Totam aperiam, eaque ipsa quae ai.</p>
						</div>
					</div>
				</div>
			</div>
		</section>
		<section class="content-box box-5"><!--Start Box-->
			<div class="zerogrid">
				<div class="row">
					<div class="col-1-3">
						<div class="wrap-col item">
							<h3 class="item-header">WELCOME to our site</h3>
							<span>LOREM IPSUM DOLOR SIT AMET, CONSEC TEER ADIPISCING. PRSENT VESTIBULUM.</span>
							<img src="/static/images/logo.png" />
							<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.
								Et vel quod legimus, graeci electram ocurreret at his. Vix at tation facete impetus omnesque ius harum antiopam.</p>
							<a class="btn" href="#">More</a>
						</div>
					</div>
					<div class="col-1-3">
						<div class="wrap-col item">
							<h3 class="item-header">CARE & SUPPORT</h3>
							<span>LOREM IPSUM DOLOR SIT AMET, CONSEC TEER ADIPISCENT VESTIBULUM.</span>
							<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.
							Et vel quod legimus, graeci electram ocurreret at his. Vix at tation facete impetus omnesque ius harum antiopam.</p>
							<ul class="link">
								<li><a href="#"> ASETY KSCABO</a></li>
								<li><a href="#"> NERAFAES</a></li>
								<li><a href="#"> KERTYU ERSVITA ERTYA</a></li>
								<li><a href="#"> SNEMO LASEC VASP</a></li>
								<li><a href="#"> TAIADES GOERTAYSE</a></li>
							</ul>
						</div>
					</div>
					<div class="col-1-3">
						<div class="wrap-col item" style="border-right: none;">
							<h3 class="item-header">LATEST NEWS</h3>
							<span>LOREM IPSUM DOLOR SIT AMET</span>
							<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.</p>
							<hr style="border: 1px dashed #ccc;margin: 17px 0;">
							<span>LOREM IPSUM DOLOR SIT AMET</span>
							<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.</p>
							<hr style="border: 1px dashed #ccc;margin: 17px 0;">
							<span>LOREM IPSUM DOLOR SIT AMET</span>
							<p>His primis omittam intellegat cu, voluptua appetere mea ad, eu harum oporteat vix.</p>
						</div>
					</div>
				</div>
			</div>
		</section>
	</div>
</section>

<!--////////////////////////////////////Footer-->
<footer>
	<div class='embed-container maps'>
		<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3164.289259162295!2d-120.7989351!3d37.5246781!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x8091042b3386acd7%3A0x3b4a4cedc60363dd!2sMain+St%2C+Denair%2C+CA+95316%2C+Hoa+K%E1%BB%B3!5e0!3m2!1svi!2s!4v1434016649434" width="100%" height="250px" frameborder="0" style="border: 0;"></iframe>
	</div>
	<div class="wrap-footer">
		<div class="zerogrid">
			<div class="row">
				<h3>Contact</h3>
				<span>Phone / +80 999 99 9999 </span></br>
				<span>Email / info@domain.com  </span></br>
				<span>Studio / Moonshine St. 14/05 Light City </span></br>
				<span><strong>Copyright 20xx - <a href="http://www.zerotheme.com" rel="nofollow" target="_blank">Html5 Templates</a> Designed by <a href="http://www.zerotheme.com" rel="nofollow" target="_blank">ZEROTHEME</a></strong></span>
			</div>
		</div>
	</div>
</footer>

	<!-- Google Map -->
	<script>
		$('.maps').click(function () {
		$('.maps iframe').css("pointer-events", "auto");
	});

	$( ".maps" ).mouseleave(function() {
	  $('.maps iframe').css("pointer-events", "none");
	});
	</script>

</div>
</body></html>
