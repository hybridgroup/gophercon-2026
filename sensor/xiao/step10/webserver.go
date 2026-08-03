package main

import (
	_ "embed"
	"sync"

	"github.com/soypat/lneto/http/httphi"
	"github.com/soypat/lneto/http/httpraw"
)

//go:embed index.html
var page string

// Uses Min - a tiny framework that makes websites pretty.
// See https://mincss.com/
//
//go:embed mincss.min.css
var mincss string

var (
	responseActive         = []byte("system active")
	responseInactive       = []byte("system inactive")
	responseStatusActive   = []byte(`{"status": "active"}`)
	responseStatusInactive = []byte(`{"status": "inactive"}`)
)

// scratchPool hands out work buffers to handlers so that parsing a form or
// building a JSON response does not allocate on every request.
var scratchPool sync.Pool

const scratchSize = 128

func startWebServer() {
	scratchPool.New = func() interface{} { return make([]byte, scratchSize) }

	var http httphi.MuxSlice
	http.Handle("/", root)
	http.Handle("/mincss.min.css", css)
	http.Handle("/on", systemActivate)
	http.Handle("/off", systemDeactivate)
	http.Handle("/status", systemStatus)
	http.Handle("/alarmlevel", alarmlevel)

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

const (
	textplain = "text/plain; charset=UTF-8"
	appjson   = "application/json"
)

func systemActivate(exch *httphi.Exchange) {
	systemActive = true
	exch.Respond(httphi.StatusOK, textplain, responseActive)
}

func systemDeactivate(exch *httphi.Exchange) {
	systemActive = false
	exch.Respond(httphi.StatusOK, textplain, responseInactive)
}

func systemStatus(exch *httphi.Exchange) {
	if systemActive {
		exch.Respond(httphi.StatusOK, appjson, responseStatusActive)
	} else {
		exch.Respond(httphi.StatusOK, appjson, responseStatusInactive)
	}
}

func alarmlevel(exch *httphi.Exchange) {
	scratch := scratchPool.Get().([]byte)
	defer scratchPool.Put(scratch)

	if exch.RequestMethod() == httphi.MethPost {
		var form httpraw.Form
		form.Reset(scratch, 1) // 1 form value max: "level".
		const parseURL, prioritizeURL = false, false
		err := exch.RequestParseForm(&form, parseURL, prioritizeURL)
		if err != nil {
			exch.Respond(httphi.StatusInternalServerError, "", nil)
			return
		}
		if lvl := form.Get("level"); len(lvl) > 0 {
			alarmLevel = uint16(bytesToInt(lvl))
		}
	}

	// form values above point into scratch, so only reuse it once we are done with them.
	json := append(scratch[:0], `{"level": `...)
	n := uintToBytes(scratch[len(json):], uint32(alarmLevel))
	json = append(json[:len(json)+n], '}')
	exch.Respond(httphi.StatusOK, appjson, json)
}
