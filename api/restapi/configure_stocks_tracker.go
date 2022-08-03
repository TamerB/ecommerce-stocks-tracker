// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/TamerB/ecommerce-stocks-tracker/api/handler"
	"github.com/TamerB/ecommerce-stocks-tracker/api/restapi/operations"
	"github.com/TamerB/ecommerce-stocks-tracker/config"
	db "github.com/TamerB/ecommerce-stocks-tracker/db/sqlc"

	_ "github.com/lib/pq" // pq module is necessary for DB transactions
)

//go:generate swagger generate server --target ../../api --name StocksTracker --spec ../../swagger.yaml --principal interface{}

var debugFlags = struct {
	Verbose bool `long:"verbose" short:"v" description:"Show debug log entries"`
}{}

func configureFlags(api *operations.StocksTrackerAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Debug Flags",
			Options:          &debugFlags,
		},
	}
}

func configureAPI(api *operations.StocksTrackerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	config := config.NewConfig()

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	api.StockConsumeProductStockHandler = handler.NewConsumeProductStockHandler(store)
	api.ProductGetProductBySkuHandler = handler.NewGetProductRequestHandler(store)
	api.ProductGetProductStocksBySkuHandler = handler.NewGetProductStocksRequestHandler(store)
	api.HealthGetHealthzHandler = handler.NewHealthzRequestHandler()
	api.HealthGetReadyzHandler = handler.NewReadyzRequestHandler()

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
