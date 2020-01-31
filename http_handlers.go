package main

import (
	"net/http"

	"github.com/amattn/deeperror"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	assetManager := AssetManagerFromContext(r.Context())

	data := MakeContentData(r)
	data.Title = "Web App Template Default Title"
	data.PageSubtitle = "Hello!"

	buf, contentErr := ProcessContentTemplate(assetManager, IndexHtml, data)
	WriteContent(w, r, 0, 3392455707, buf, contentErr)
}

func test500Handler(w http.ResponseWriter, r *http.Request) {
	derr := deeperror.New(2752399866, "This is a synthetic 500 response", nil)
	Default500Handler(w, r, derr, 962680333, 4112454076, 1269746600)
}
