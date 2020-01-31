package main

import (
	"context"
	"log"
	"net/http"

	"github.com/alexedwards/scs"

	"github.com/amattn/deeperror"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	SignUpEndpoint       = "/signup"
	SignUpActionEndpoint = "/signup/action"
	LoginEndpoint        = "/login"
	LoginActionEndpoint  = "/login/action"
	LogoutEndpoint       = "/logout"
	//LogoutActionEndpoint = "/logout/action"

	TestEndpoint    = "/test"
	Test500Endpoint = TestEndpoint + "/500"

	// only if logged in as super admin
	APIEndpoint         = "/api"
	APIUpdateEndpoint   = APIEndpoint + "/update"
	APIRegisterEndpoint = APIEndpoint + "/register"

	// only if logged in
	DashboardEndpoint = "/dashboard"

	// only if logged in as super admin
	AdminEndpoint = "/admin"

	// must be last
	RootEndpoint = "/"
)

func GetRouter(am *AssetManager, sessionKey string) *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(system404Handler)

	compressor := middleware.NewCompressor(5, "gzip")

	router.Use(
		middleware.Recoverer,
		middleware.Logger,
		compressor.Handler(),
		middleware.RedirectSlashes,
		middleware.RequestID,
		AddRequestHelpersToContext(am, sessionKey),
	)

	// just for test purposes:

	router.Get(Test500Endpoint, test500Handler)

	//static assets

	// we want to route not just /img, but anything under it as well /img/file.jpg
	//router.Get("/img/*", staticAssetsHandler)
	//router.Get("/css/*", staticAssetsHandler)
	//router.Get("/js/*", staticAssetsHandler)
	//router.Get("/fonts/*", staticAssetsHandler)

	// must be last
	router.Get(RootEndpoint, rootHandler)

	return router
}

//  #####
// #     #  ####  #    # ##### ###### #    # #####
// #       #    # ##   #   #   #       #  #    #
// #       #    # # #  #   #   #####    ##     #
// #       #    # #  # #   #   #        ##     #
// #     # #    # #   ##   #   #       #  #    #
//  #####   ####  #    #   #   ###### #    #   #
//
// #     #
// #     # ###### #      #####  ###### #####   ####
// #     # #      #      #    # #      #    # #
// ####### #####  #      #    # #####  #    #  ####
// #     # #      #      #####  #      #####       #
// #     # #      #      #      #      #   #  #    #
// #     # ###### ###### #      ###### #    #  ####
//

const RequestHelpersContextKey = "RequestHelpersContextKey"

type RequestHelpers struct {
	assetManager   *AssetManager
	sessionManager *scs.Manager
}

func AddRequestHelpersToContext(am *AssetManager, sessionKey string) func(next http.Handler) http.Handler {
	rh := RequestHelpers{
		am,
		scs.NewCookieManager(sessionKey),
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), RequestHelpersContextKey, rh)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AssetManagerFromContext(ctx context.Context) *AssetManager {
	maybe := ctx.Value(RequestHelpersContextKey)

	unpackedThing, isExpectedType := maybe.(RequestHelpers)
	if isExpectedType {
		return unpackedThing.assetManager
	} else {
		return nil
	}
}

func SessionManagerFromContext(ctx context.Context) *scs.Manager {
	maybe := ctx.Value(RequestHelpersContextKey)

	unpackedThing, isExpectedType := maybe.(RequestHelpers)
	if isExpectedType {
		return unpackedThing.sessionManager
	} else {
		return nil
	}
}

// #     #
// #     # ##### # #       ####
// #     #   #   # #      #
// #     #   #   # #       ####
// #     #   #   # #           #
// #     #   #   # #      #    #
//  #####    #   # ######  ####
//

func LogAllRoutes(mux *chi.Mux) {
	log.Println("LogAllRoutes:")
	walker := func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Println(method, route)
		return nil
	}

	err := chi.Walk(mux, walker)
	if err != nil {
		derr := deeperror.New(3361616010, "", err)
		derr.AddDebugField("mux", mux)
		log.Println(derr.Num, derr)
	}
}
