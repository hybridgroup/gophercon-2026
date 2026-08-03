package main

import (
	_ "embed"

	"github.com/soypat/lneto/http/httphi"
)

//go:embed index.html
var page string

// Uses Min - a tiny framework that makes websites pretty.
// See https://mincss.com/
//
//go:embed mincss.min.css
var mincss string

//go:embed tetromino.html
var tetromino string

var (
	responseActive   = []byte("system active")
	responseInactive = []byte("system inactive")
)

func startWebServer() {
	var http httphi.MuxSlice
	http.Handle("/", root)
	http.Handle("/mincss.min.css", css)
	http.Handle("/6", sixlines)
	http.Handle("/on", systemActivate)
	http.Handle("/off", systemDeactivate)

	h, _ := link.Addr()
	host := h.String()
	print("HTTP server listening on http://", host, ":", port, "\n")
	var router httphi.Router
	cfg := httphi.DefaultRouterConfig(4, 2048, http.MaxPathValues())
	err := router.Configure(&http, cfg)
	if err != nil {
		failMessage("router configuration: " + err.Error())
	}
	// ListenAndServe should block indefinetely unless Router shut down
	err = link.ListenAndServe(&router, port)
	failMessage("listen and serve failed: " + err.Error())
}

func root(exch *httphi.Exchange) {
	exch.RespondString(httphi.StatusOK, "text/html", page)
}

func css(exch *httphi.Exchange) {
	exch.RespondString(httphi.StatusOK, "text/css", mincss)
}

// https://fukuno.jig.jp/3267
func sixlines(exch *httphi.Exchange) {
	exch.RespondString(httphi.StatusOK, "text/html", tetromino)
}

const textplain = "text/plain; charset=UTF-8"

func systemActivate(exch *httphi.Exchange) {
	systemActive = true
	exch.Respond(httphi.StatusOK, textplain, responseActive)
}

func systemDeactivate(exch *httphi.Exchange) {
	systemActive = false
	exch.Respond(httphi.StatusOK, textplain, responseInactive)
}
