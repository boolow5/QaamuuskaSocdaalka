<div class="slider">
	<!-- Slideshow -->
	<div class="callbacks_container" >
		<ul class="rslides" id="slider">
			{{range $index, $val := .LatestPosts}}
			<li>
				<a href="#">
	        <div class="img" style="background-image: url({{$val.FeaturedImage.Url}})">
	  				<div class="caption">
	  					<h1>{{shorten_words $val.Title 10}}</h1>
	  					<span >{{shorten_words $val.Content 20 |markdown}}</span>
	  				</div>
	        </div>
				</a>
			</li>
			{{end}}
		</ul>
	</div>

</div>
