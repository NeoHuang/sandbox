package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adjust/backend/core/util"
)

var (
	iosDeeplink  = "lumosity://challenges?challenge_id=41224"
	iosStorelink = "https://itunes.apple.com/app/id577232024?mt=8"
)

func androidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.UserAgent())

	//	redirect := "intent://?adjust_tracker=sc6zi7&adjust_creative=testCreative&adjust_reftag=cVIMvbbqcmV46#Intent;scheme=example;package=com.adjust.example;S.browser_fallback_url=google.com;end" //fmt.Sprintf("intent://%s#Intent;scheme=%s;package=%s;end", "", "example", "com.adjust.example")
	redirect := "http://play.google.com/store/apps/details?id=com.google.android.apps.maps"
	fmt.Println(redirect)
	http.Redirect(w, r, redirect, 302)
}

func iosDirectRedirect(w http.ResponseWriter, r *http.Request) {
	redirect := "https://itunes.apple.com/app/id920877281?mt=8"
	//	redirect := "chegg-tutors://tutors?adjust_reftag=cPwLfyC"
	http.Redirect(w, r, redirect, 302)
}

func twitterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.UserAgent())
	res := clickAnchorScript("example://", "http://play.google.com/store/apps/details?id=com.adjust.example", 100)
	fmt.Fprintf(w, res)
}

func iosRedirectHandler(w http.ResponseWriter, r *http.Request) {
	res := `<html>
        <head>
        <title>adjust deeplinking...</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
        <meta name="viewport" content="width=640, user-scalable=yes" />
				<meta name="apple-itunes-app" content="app-id=920877281" />
        </head>
        <body>
        <script>
				start = Date.now()

                  setTimeout(function () {
										document.getElementById("deeplink-iframe").src = "";
										console.log("lalalalal")
										window.location.replace("http://www.google.com");  
                }, 2000);
                
        </script>

                <iframe id="deeplink-iframe" style="display:none" height=0 width=0 src="chegg-tutors://tutors?adjust_reftag=cPwLfyC"></iframe>

        </body>
        </html>`
	fmt.Fprintf(w, res)
}

func iosHandler(w http.ResponseWriter, r *http.Request) {
	res := "hello world"
	redirect := r.FormValue("redirect")
	fmt.Println(redirect)
	if redirect == "" {
		res = util.IosDeeplinkScript("chegg-tutors://tutors?adjust_reftag=cPwLfyC", "itms-apps://itunes.apple.com/app/id920877281?mt=8", "920877281", 1000)
	} else {
		res = util.IosDeeplinkScript("chegg-tutors://tutors?adjust_reftag=cPwLfyC", redirect, "920877281", 1000)

	}
	fmt.Fprintf(w, res)

}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "")

}

func iframeDeeplinkingScript() string {
	return `<script>
	runner = {
		  timer: null,
			  deeplink: "chegg-tutors://tutors?adjust_reftag=cPwLfyC",
			  fallback: "https://itunes.apple.com/app/id920877281?mt=8",
			  clearTimers : function() {
				    clearTimeout(this.timer)
			  },
			  openLinkWithIframe: function(e) {
				    var i = document.createElement("iframe");
				    i.style.width = "1px", i.style.height = "1px", i.border = "none", i.addEventListener("load", this.clearTimers), i.src = e, document.body.appendChild(i)
			  },
			  run: function() {
				    fallback = this.fallback
						this.timer = setTimeout(function() {
								window.location.replace(fallback)
						}, 100000)
						this.openLinkWithIframe("chegg-tutors://tutors?adjust_reftag=cPwLfyC")
				}
	}`
}

func ios9Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.UserAgent())
	res := delayedLocationScript(iosDeeplink, "http://google.com", 1000)
	fmt.Fprintf(w, res)

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, iosStorelink, 302)
}

func main() {
	fmt.Println("server staring...port 1234\n")
	http.HandleFunc("/android", androidHandler)
	http.HandleFunc("/ios", iosHandler)
	http.HandleFunc("/ios9", ios9Handler)
	http.HandleFunc("/iosRedirect", iosRedirectHandler)
	http.HandleFunc("/iosDirectRedirect", iosDirectRedirect)
	http.HandleFunc("/twitter", twitterHandler)
	http.HandleFunc("/302", redirectHandler)
	http.HandleFunc("/", simpleHandler)
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(10 * time.Millisecond)
			request, _ := http.NewRequest("GET", "http://zhiyis-macbook-pro.local:1234/？", nil)
			client := http.Client{}
			client.Do(request)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(10 * time.Millisecond)
			request, _ := http.NewRequest("GET", "http://google.com/", nil)
			client := http.Client{}
			client.Do(request)
		}
	}()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Errorf("ADJUST SERVER SHUTTING DOWN (%s)\n\n", err)
	}
}

func metarefreshScript(deeplink, store string, delay int) string {
	return fmt.Sprintf(`<html>
        <head>
        <title>adjust deeplinking...</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
        <meta name="viewport" content="width=640, user-scalable=yes" />
				<meta name="apple-itunes-app" content="app-id=920877281" />
				<meta http-equiv="refresh" content="0; url=%s" />
        </head>
        <body>
									<script>

														window.onload =function() {
															setTimeout(function(){

							window.location.replace("%s")
															}, %d)

					}
        </script>
        </body>
        </html>`, deeplink, store, delay)
}
func newClickAnchorScript(deeplink, store string, delay int) string {
	return fmt.Sprintf(`<html>
        <head>
        <title>adjust deeplinking...</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
        <meta name="viewport" content="width=640, user-scalable=yes" />
				<meta name="apple-itunes-app" content="app-id=920877281" />
        </head>
        <body>
					<a href="%s" id="deeplink" ></a>
									<script>

														window.onload =function() {
															setTimeout(function(){

							window.location.replace("%s")
															}, 2000)
						setTimeout(function() {
															simulateClick()
						}, %d)
														}
function simulateClick() {
  var event = new MouseEvent('click', {
    'view': window,
    'bubbles': true,
    'cancelable': true
  });
  var cb = document.getElementById('deeplink'); 
  var canceled = !cb.dispatchEvent(event);
  if (canceled) {
    // A handler called preventDefault.
  } else {
    // None of the handlers called preventDefault.
  }
}

        </script>
        </body>
        </html>`, deeplink, store, delay)
}

func clickAnchorScript(deeplink, store string, delay int) string {
	return fmt.Sprintf(`<html>
        <head>
        <title>adjust deeplinking...</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
        <meta name="viewport" content="width=640, user-scalable=yes" />
				<meta name="apple-itunes-app" content="app-id=920877281" />
        </head>
        <body>
					<a href="%s" id="deeplink" ></a>
									<script>

														window.onload =function() {
															setTimeout(function(){

							window.location.replace("%s")
															}, 10000)
						setTimeout(function() {
															eventFire(document.getElementById('deeplink'), "click")
						}, %d)
														}
														function eventFire(el, etype){
						if (el.fireEvent) {
							el.fireEvent('on' + etype);
						} else {
							var evObj = document.createEvent('Events');
							evObj.initEvent(etype, true, false);
							el.dispatchEvent(evObj);
						}

					}
        </script>
        </body>
        </html>`, deeplink, store, delay)
}

func delayedLocationScript(deeplink, store string, delay int) string {
	return fmt.Sprintf(`<html>
        <head>
        <title>adjust deeplinking...</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge" />
        <meta name="viewport" content="width=640, user-scalable=yes" />
				<meta name="apple-itunes-app" content="app-id=920877281" />
        </head>
        <body>
									<script>

														window.onload =function() {
															window.location = "%s"
															setTimeout(function(){

							window.location = "%s"
															}, %d)

					}
        </script>
        </body>
        </html>`, deeplink, store, delay)

}
