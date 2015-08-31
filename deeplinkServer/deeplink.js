
<html>
<head>
<title>Launching App...</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
<meta name="viewport" content="width=device-width, user-scalable=no">
<script>
	var appname = "Test";
	var applink = "chegg-tutors://";  // id = nil
	var fallback_url = "https://zhiyihuang.com/";
	var android_package_name = null;
</script>


<link href="https://deeplinkme.s3.amazonaws.com/assets/launcher-58e6b9e24b2fc4ab7732f82f0b8d8e5b.css" media="all" rel="stylesheet" type="text/css" />

</head>
<body>

		<div class="header show_logo"></div>
		<style>
			.header.show_logo {
				background-image: url(http://is3.mzstatic.com/image/pf/us/r30/Purple7/v4/1d/42/e6/1d42e6c2-fdd1-281e-c974-f63b294a1926/mzl.rswgiuei.png);
			}
		</style>

	<div class="content">

		<div id="message">
			<p>
					Hang on, Test is loading…
			</p>
		</div>

		<div id="buttons" class="buttons" style="display:none;">
			<button type="button" class="app" onclick='window.location = applink'>OPEN THE APP</button>
			<button type="button" class="website" onclick='window.location = fallback_url'>
					GO TO THE WEBSITE
			</button>
		</div>
		<div class="i-deeplink"></div>

	</div>

<iframe id="i" style="display:none" height=0 width=0></iframe>

<script>
document.ontouchmove = function(event){
	event.preventDefault();
}



function show_buttons ()
{
	var message = document.getElementById('message');
	var buttons = document.getElementById('buttons');
	message.style.display = 'none';
	buttons.style.display = 'block';
}

/* Deprecated, but keeping for now */
var MESSAGE_ATTEMPTING = appname + ": Attempting to deeplink via " + applink;
var MESSAGE_CLEANUP_REQUIRED = appname + ": You’ve been deeplinked directly into the app!  Feel free to close this browser window.";
var MESSAGE_FALLING_BACK = appname + ": Loading website";
function set_message (msg)
{
	show_buttons();
}




if (location.hash != "" && location.search.indexOf("__deeplinkme_hash=") == -1) {
	fragment = encodeURIComponent(location.hash.substring(1))
	joiner = location.search != "" ? '&' : '?'
	dehashed_url = location.href.replace(/#.*/, joiner + '__deeplinkme_hash=' + fragment)
	location.replace(dehashed_url)
}
else {


	//// CONFIGURATION
	launch_timeout  = 1500;
	margin_of_error = 500;
	close_delay     = 1000;
	close_interval  = 3000;


	if (history.replaceState) {
		history.replaceState({}, "App Launched!  Welcome Back!",
			"/launch/cleanup"
			+ "?appname="              + encodeURIComponent(appname)
			+ "&applink="              + encodeURIComponent(applink)
			+ "&fallback_url="         + encodeURIComponent(fallback_url)
			+ "&android_package_name=" + encodeURIComponent(android_package_name)
			+ "&app_logo_url="         + encodeURIComponent("http://is3.mzstatic.com/image/pf/us/r30/Purple7/v4/1d/42/e6/1d42e6c2-fdd1-281e-c974-f63b294a1926/mzl.rswgiuei.png")
			+ "&launcher_text="        + encodeURIComponent("")
		)
	}

	var has_safari = navigator.userAgent.indexOf('Safari/')  > -1;

	var i = document.getElementById('i');
	launched_at = Date.now();
	setTimeout(function() {
		arrived_at = Date.now();
		if (arrived_at - launched_at > launch_timeout + margin_of_error) {
			set_message(MESSAGE_CLEANUP_REQUIRED);
			if (history.length > 1) {
				history.back();
			}
			else {
				setTimeout( function(){ close() }, close_delay);
				setInterval(function(){ close() }, close_interval);
			}
		} else {
			set_message(MESSAGE_FALLING_BACK);
			if (has_safari || true)
				location.replace(fallback_url);
		}
	}, launch_timeout);

	// Launch!
	if (has_safari) i.src = applink;
	else location.assign(applink);

} 
</script>
</body>
</html>

