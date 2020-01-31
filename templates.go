package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-playground/statics/static"

	"github.com/amattn/deeperror"
)

type AssetPath string
type ContentKey string
type ContentErrorKey string

type AssetManager struct {
	Path          string
	TemplateCache sync.Map
	EmbeddedFiles *static.Files
}

func NewAssetManager(path string) (*AssetManager, error) {
	log.Println(1600691750)
	config := &static.Config{
		UseStaticFiles: true,
		AbsPkgPath:     "assets/",
	}

	embeds, err := newStaticAssets(config)
	if err != nil {
		derr := deeperror.New(292884003, currentFunction()+" Failure: newStaticAssets", err)
		derr.AddDebugField("path", path)
		return nil, derr
	}

	log.Println(1600691751)

	assetManager := new(AssetManager)
	assetManager.Path = path
	assetManager.TemplateCache = sync.Map{}
	assetManager.EmbeddedFiles = embeds
	return assetManager, nil
}

func (assetManager *AssetManager) LoadAssetString(path AssetPath) (string, error) {
	file_bytes, err := assetManager.EmbeddedFiles.ReadFile(string(path))
	if err != nil {
		derr := deeperror.New(3302808785, currentFunction()+" Failure:", err)
		derr.AddDebugField("path", path)
		return "", derr
	}

	return string(file_bytes), nil

}

func ProcessContentTemplate(assetManager *AssetManager, templatePath AssetPath, data ContentData) (*bytes.Buffer, error) {

	if templatePath != BaseHtml {
		contentBuf, err := loadAndExecuteTemplate(assetManager, templatePath, data)
		if err != nil {
			derr := deeperror.New(3520389020, "", err)
			derr.AddDebugField("assetManager", assetManager)
			derr.AddDebugField("templatePath", templatePath)
			derr.AddDebugField("data", data)
			return nil, derr
		}

		contentString := contentBuf.String()
		contentHtml := template.HTML(contentString)
		data.HTMLContent = contentHtml

		if len(data.HTMLHeader) == 0 {
			// load default header
			headerString, err := loadAsset(assetManager, BaseHeaderHtml)
			if err != nil {
				//
			} else {
				data.HTMLHeader = template.HTML(headerString)
			}
		}

		if len(data.HTMLFooter) == 0 {
			// load default footer
			footerString, err := loadAsset(assetManager, BaseFooterHtml)
			if err != nil {
				//
			} else {
				data.HTMLFooter = template.HTML(footerString)
			}

		}
	}

	return ProcessBase(assetManager, data)
}

func ProcessBase(assetManager *AssetManager, data ContentData) (*bytes.Buffer, error) {
	return loadAndExecuteTemplate(assetManager, BaseHtml, data)
}

func ProcessTemplate(assetManager *AssetManager, templatePath AssetPath, data ContentData) (*bytes.Buffer, error) {
	return loadAndExecuteTemplate(assetManager, templatePath, data)
}

func loadAndExecuteTemplate(assetManager *AssetManager, assetName AssetPath, data interface{}) (*bytes.Buffer, error) {
	tmpl, err := loadTemplate(assetManager, assetName)
	if err != nil {
		derr := deeperror.New(3776256257, "load template error", err)
		derr.AddDebugField("assetManager", assetManager)
		derr.AddDebugField("assetName", assetName)
		return nil, derr
	}

	buff := new(bytes.Buffer)

	err = tmpl.Execute(buff, data)
	if err != nil {
		derr := deeperror.New(3049642997, "template execute error", err)
		derr.AddDebugField("assetBox", assetManager)
		derr.AddDebugField("assetName", assetName)
		derr.AddDebugField("data", data)
		return nil, derr
	}

	return buff, nil
}

func loadTemplate(assetManager *AssetManager, assetName AssetPath) (*template.Template, error) {
	// first check cache
	maybeTmpl, exists := assetManager.TemplateCache.Load(assetName)
	if exists {
		tmpl, ok := maybeTmpl.(*template.Template)
		if ok {
			return tmpl, nil
		} else {
			// tmpl is something else
			derr := deeperror.New(1003944010, "unexpected type tmpl is ", nil)
			derr.AddDebugField("assetManager", assetManager)
			derr.AddDebugField("assetName", assetName)
			derr.AddDebugField("type", fmt.Sprintf("%T", tmpl))
			return nil, derr
		}
	}

	// add some helper functions
	helperFuncMap := template.FuncMap{
		"shortTime":     shortTime,
		"shortTimeHour": shortTimeHour,
	}

	// so our desired template doesn't exist...  let's make it.
	raw, err := loadAsset(assetManager, assetName)
	if err != nil {
		derr := deeperror.New(2316963673, "", err)
		derr.AddDebugField("assetBox", assetManager)
		derr.AddDebugField("assetName", assetName)
		return nil, derr
	}

	tmpl, err := template.New(string(assetName)).Funcs(helperFuncMap).Parse(raw)
	if err != nil {
		derr := deeperror.New(2316963674, "Failure Parsing Template", err)
		derr.AddDebugField("assetManager", assetManager)
		derr.AddDebugField("assetName", assetName)
		return nil, derr
	}

	assetManager.TemplateCache.Store(assetName, tmpl)
	return tmpl, err

}

// load the asset from whatever asset package you are using
func loadAsset(assetManager *AssetManager, assetName AssetPath) (string, error) {

	raw, err := assetManager.LoadAssetString(assetName)
	if err != nil {
		derr := deeperror.New(809174294, "", err)
		derr.AddDebugField("assetManager", assetManager)
		derr.AddDebugField("assetName", assetName)
		return "", derr
	}

	return raw, nil
}

// #     #
// #     # ###### #      #####  ###### #####   ####
// #     # #      #      #    # #      #    # #
// ####### #####  #      #    # #####  #    #  ####
// #     # #      #      #####  #      #####       #
// #     # #      #      #      #      #   #  #    #
// #     # ###### ###### #      ###### #    #  ####
//

const (
	ShortTimeDefaultFormat     = "2006-01-02 3:04pm"
	ShortTimeHourDefaultFormat = "2006-01-02 3pm"
)

func shortTime(t *time.Time) string {

	if t == nil || t.Unix() == 0 {
		return "--"
	}

	userLoc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		userLoc = time.UTC
	}
	userTime := t.In(userLoc)

	return userTime.Format(ShortTimeDefaultFormat)
}

func shortTimeHour(t *time.Time) string {
	if t == nil || t.Unix() == 0 {
		return "--"
	}
	userLoc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		userLoc = time.UTC
	}
	userTime := t.In(userLoc)

	return userTime.Format(ShortTimeHourDefaultFormat)
}

//  #####
// #     #  ####  #    # ##### ###### #    # #####
// #       #    # ##   #   #   #      ##   #   #
// #       #    # # #  #   #   #####  # #  #   #
// #       #    # #  # #   #   #      #  # #   #
// #     # #    # #   ##   #   #      #   ##   #
//  #####   ####  #    #   #   ###### #    #   #
//

type ContentData struct {
	Title           string        // <title> tag in header
	HTMLHeadContent template.HTML // additional content in the header
	PageTitle       template.HTML // default title on page body.  default template typically uses Title if this is empty
	PageSubtitle    template.HTML // default subtitle on page body
	HTMLContent     template.HTML
	HTMLHeader      template.HTML
	HTMLFooter      template.HTML

	CurrentURL *url.URL
	//CurrentUser *AuthLayerUser

	Data map[string]interface{}

	// simple error
	SomethingWentWrong *SomethingWentWrong

	// multiple errors
	Errors map[string]interface{} // misc form validation errors, etc.
}

type SomethingWentWrong struct {
	Tracer    string
	DebugNums []int64
	Message   string
}

func NewSomethingWentWrong(tracer string, debugNums ...int64) *SomethingWentWrong {
	sww := new(SomethingWentWrong)

	sww.DebugNums = debugNums
	sww.Tracer = tracer

	return sww
}

func (sww SomethingWentWrong) String() string {
	return fmt.Sprintf("%s (t:%s d:%v)", sww.Message, sww.Tracer, sww.DebugNums)
}

func MakeContentData(r *http.Request) ContentData {
	//user := UserFromContext(r.Context())

	cd := ContentData{
		CurrentURL: r.URL,
		//CurrentUser: user,
		Data:   map[string]interface{}{},
		Errors: map[string]interface{}{},
	}

	return cd
}

func (cd ContentData) CurrentURLPathAndQuestionMark() string {
	if cd.CurrentURL != nil {
		return cd.CurrentURL.Path + "?" + cd.CurrentURL.Query().Encode()
	}

	return "?"
}

func (cd *ContentData) AddKeyValue(key ContentKey, value interface{}) {
	cd.Data[string(key)] = value
}

func (cd *ContentData) AddErrorKeyValue(key ContentErrorKey, value interface{}) {
	cd.Errors[string(key)] = value
}

func WriteContent(w http.ResponseWriter, r *http.Request, httpStatusCode int, debugNum int64, buf *bytes.Buffer, contentErr error) {
	if contentErr != nil {
		// we can't use Default500Handler here.  Default500Handler calls this method and we will recurse infinitely
		log.Println(2564018910, debugNum)
		log.Println(2564018911, contentErr)
		log.Println(2564018911, r)
		debugString := fmt.Sprintln("Internal Server error:", 2793893260, debugNum)
		http.Error(w, debugString, http.StatusInternalServerError)
		return
	}

	if httpStatusCode != 0 {
		w.WriteHeader(httpStatusCode)
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		derr := deeperror.New(401502009, "FATAL error writing to buffer", err)
		derr.AddDebugField("*http.Request", r)
		log.Println(derr)
	}
}

// use WriteContent if possible... most of the time we already have a buffer...
// no need to convert buf to bytes back to buf.
func WriteContentBytes(w http.ResponseWriter, r *http.Request, httpStatusCode int, debugNum int64, rawBytes []byte, contentErr error) {
	buf := bytes.NewBuffer(rawBytes)
	WriteContent(w, r, httpStatusCode, debugNum, buf, contentErr)
}
