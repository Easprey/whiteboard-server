// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/swag"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/Easprey/whiteboard-server/handlers"
	"github.com/Easprey/whiteboard-server/restapi/operations"
	"github.com/Easprey/whiteboard-server/restapi/operations/fingerpaths"
)

//go:generate swagger generate server --target ../../whiteboard-server --name Whiteboard --spec ../swagger.yaml

// create a DataBase variable to read in connection string and hold database connection
var db = handlers.DataBase{}

func configureFlags(api *operations.WhiteboardAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "WhiteBoard Server Service Options",
			Options:          &db,
		},
	}
}

func configureAPI(api *operations.WhiteboardAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.FingerpathsFingerPathsGetHandler == nil {
		api.FingerpathsFingerPathsGetHandler = fingerpaths.FingerPathsGetHandlerFunc(handlers.GetFingerPaths)
	}
	if api.FingerpathsFingerPathsPostHandler == nil {
		api.FingerpathsFingerPathsPostHandler = fingerpaths.FingerPathsPostHandlerFunc(handlers.PostFingerPaths)
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
	fmt.Printf("called")
	db.Init()
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
