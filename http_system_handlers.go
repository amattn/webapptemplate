package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/amattn/deeperror"
)

// https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/
//
// kubernetes uses liveness probes to know when to restart a Container.
// For example, liveness probes could catch a deadlock, where an application
// is running, but unable to make progress. Restarting a Container in
// such a state can help to make the application more available despite bugs.
//
// kubernetes uses readiness probes to know when a Container is ready to
// start accepting traffic. A Pod is considered ready when all of its
// Containers are ready. One use of this signal is to control which Pods
// are used as backends for Services. When a Pod is not ready, it is removed
// from Service load balancers.

const (
	SystemMountEndpoint = "/sys"

	AliveEndpoint = "/alive"
	ReadyEndpoint = "/ready"
)

// Any data structures, interfaces or variables that are necessary to bootstrap the system
// or help with the alive (did I startup?) or ready (am I ready to handle incoming requests?) methods
type SystemBootstrap struct {
}

func makeSystemHandlerClosure(bootstrap *SystemBootstrap, fn func(http.ResponseWriter, *http.Request, *SystemBootstrap)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, bootstrap)
	}
}

// kubernetes uses liveness probes to know when to restart a Container.
// by convention, this should be routed to /sys/alive
// this function is called periodically, (such as every 10 seconds)
// and should be relatively efficient
func aliveHandler(w http.ResponseWriter, r *http.Request, bootstrap *SystemBootstrap) {
	// for this simple implementation, we have identicaly definitions of ready and alive, but you
	// but with proper backpressure, this could be simplified to
	// a working http response system that can send proper error messages
	// in the case of downstream systems (db, redis, vault, etc.) being down.
	err := IsSystemHealthy(bootstrap)
	if err != nil {
		derr := deeperror.New(9993003620, "", err)
		Default500Handler(w, r, derr, derr.Num)
		log.Println(derr.Num, derr)
		return
	}

	_, _ = fmt.Fprintf(w, "Alive!")
}

// kubernetes uses readiness probes to know when a Container is ready to
// start accepting traffic.
// by convention, this should be routed to /sys/ready
func readyHandler(w http.ResponseWriter, r *http.Request, bootstrap *SystemBootstrap) {
	err := IsSystemHealthy(bootstrap)
	if err != nil {
		derr := deeperror.New(9994051382, "", err)
		Default503Handler(w, r, derr, derr.Num)
		return
	}

	_, _ = fmt.Fprintf(w, "Ready!")
}

// to be alive, you typically just need to make sure the system started up appropriately and isn't hung, frozen, crashed, etc.

// to be healthy or ready, you typically need to be in good working condition
// - you have all appropriate config files, secrets etc.
// - you can talk to, ping your db successfully
// - etc.
func IsSystemHealthy(bootstrap *SystemBootstrap) error {
	if bootstrap == nil {
		derr := deeperror.New(99916734488, "SystemBootstrap pointer is nil", nil)
		return derr
	}

	// fill in any other system checks here...

	return nil
}

// #######
// #       #####  #####   ####  #####
// #       #    # #    # #    # #    #
// #####   #    # #    # #    # #    #
// #       #####  #####  #    # #####
// #       #   #  #   #  #    # #   #
// ####### #    # #    #  ####  #    #
//
// #     #
// #     #   ##   #    # #####  #      ###### #####
// #     #  #  #  ##   # #    # #      #      #    #
// ####### #    # # #  # #    # #      #####  #    #
// #     # ###### #  # # #    # #      #      #####
// #     # #    # #   ## #    # #      #      #   #
// #     # #    # #    # #####  ###### ###### #    #
//

// same as Default404Handler, but with a function signature that matches the handler interface
func system404Handler(w http.ResponseWriter, r *http.Request) {
	Default404Handler(w, r, 1010109090)
}

func Default404Handler(w http.ResponseWriter, r *http.Request, debugNums ...int64) {
	DefaultStatusCodeHandler(w, r, http.StatusNotFound, debugNums...)
}

// 400 = bad request
func Default400Handler(w http.ResponseWriter, r *http.Request, parentErr error, debugNums ...int64) {
	DefaultErrorCodeHandler(w, r, http.StatusBadRequest, parentErr, debugNums...)
}

// 500 = internal service error
func Default500Handler(w http.ResponseWriter, r *http.Request, parentErr error, debugNums ...int64) {
	DefaultErrorCodeHandler(w, r, http.StatusInternalServerError, parentErr, debugNums...)
}

// 503 = service unavailable
func Default503Handler(w http.ResponseWriter, r *http.Request, parentErr error, debugNums ...int64) {
	DefaultErrorCodeHandler(w, r, http.StatusServiceUnavailable, parentErr, debugNums...)
}

// for non-200 status codes
func DefaultStatusCodeHandler(w http.ResponseWriter, r *http.Request, httpStatusCode int, debugNums ...int64) {

	data := MakeContentData(r)
	data.Title = fmt.Sprintf("%d %s", httpStatusCode, http.StatusText(httpStatusCode))
	assetManager := AssetManagerFromContext(r.Context())
	data.SomethingWentWrong = new(SomethingWentWrong)
	data.SomethingWentWrong.DebugNums = debugNums
	data.SomethingWentWrong.Tracer = ObfuscatedRequestId(r.Context())
	data.SomethingWentWrong.Message = ""

	var errorTemplate AssetPath = Error5XXHtml
	if httpStatusCode == http.StatusNotFound {
		errorTemplate = Error404Html
	}

	buf, contentErr := ProcessContentTemplate(assetManager, errorTemplate, data)
	WriteContent(w, r, httpStatusCode, 3499759047, buf, contentErr)
}

// for 5XX status codes, some extra error logging...
func DefaultErrorCodeHandler(w http.ResponseWriter, r *http.Request, httpStatusCode int, parentErr error, debugNums ...int64) {
	debugNum := int64(0)
	if len(debugNums) > 0 {
		debugNum = debugNums[0]
	}

	derr := deeperror.New(debugNum, "defaultErrorCodeHandler", parentErr)
	derr.AddDebugField("http_error_code", httpStatusCode)
	derr.AddDebugField("r", r)
	derr.AddDebugField("r.Context()", r.Context())
	derr.AddDebugField("RequestID", ObfuscatedRequestId(r.Context()))
	for i, dn := range debugNums {
		derr.AddDebugField("debug_num["+strconv.Itoa(i)+"]", dn)
	}

	// TODO send this to a error log instead of standard out pls.
	log.Println(derr.Num, derr)

	DefaultStatusCodeHandler(w, r, httpStatusCode, debugNums...)
}

// TODO 3162424914 not great obfuscation here.  timestamp is pretty clear if you repeatedly generate
func ObfuscatedRequestId(ctx context.Context) string {
	return ""

	// if you use chi router, you can uncomment the following
	//reqId := middleware.GetReqID(ctx)
	//encoded := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(reqId))
	//return encoded
}
