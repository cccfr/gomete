// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cccfr/gomete/restapi/operations"
	"github.com/cccfr/gomete/restapi/operations/audits"
	"github.com/cccfr/gomete/restapi/operations/barcodes"
	"github.com/cccfr/gomete/restapi/operations/denominations"
	"github.com/cccfr/gomete/restapi/operations/image"
	"github.com/cccfr/gomete/restapi/operations/products"
	"github.com/cccfr/gomete/restapi/operations/server"
	"github.com/cccfr/gomete/restapi/operations/users"
)

//go:generate swagger generate server --target ../../gomete --name SpaceMarket --spec ../../../../../../projects/Space-Market-API/spec/openapi.yaml

func configureFlags(api *operations.SpaceMarketAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SpaceMarketAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()

	if api.BarcodesDeleteBarcodesIDHandler == nil {
		api.BarcodesDeleteBarcodesIDHandler = barcodes.DeleteBarcodesIDHandlerFunc(func(params barcodes.DeleteBarcodesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation barcodes.DeleteBarcodesID has not yet been implemented")
		})
	}
	if api.DenominationsDeleteDenominationsIDHandler == nil {
		api.DenominationsDeleteDenominationsIDHandler = denominations.DeleteDenominationsIDHandlerFunc(func(params denominations.DeleteDenominationsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation denominations.DeleteDenominationsID has not yet been implemented")
		})
	}
	if api.ProductsDeleteProductsIDHandler == nil {
		api.ProductsDeleteProductsIDHandler = products.DeleteProductsIDHandlerFunc(func(params products.DeleteProductsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation products.DeleteProductsID has not yet been implemented")
		})
	}
	if api.UsersDeleteUsersIDHandler == nil {
		api.UsersDeleteUsersIDHandler = users.DeleteUsersIDHandlerFunc(func(params users.DeleteUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUsersID has not yet been implemented")
		})
	}
	if api.AuditsGetAuditsHandler == nil {
		api.AuditsGetAuditsHandler = audits.GetAuditsHandlerFunc(func(params audits.GetAuditsParams) middleware.Responder {
			return middleware.NotImplemented("operation audits.GetAudits has not yet been implemented")
		})
	}
	if api.BarcodesGetBarcodesIDHandler == nil {
		api.BarcodesGetBarcodesIDHandler = barcodes.GetBarcodesIDHandlerFunc(func(params barcodes.GetBarcodesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation barcodes.GetBarcodesID has not yet been implemented")
		})
	}
	if api.DenominationsGetDenominationsHandler == nil {
		api.DenominationsGetDenominationsHandler = denominations.GetDenominationsHandlerFunc(func(params denominations.GetDenominationsParams) middleware.Responder {
			return middleware.NotImplemented("operation denominations.GetDenominations has not yet been implemented")
		})
	}
	if api.DenominationsGetDenominationsIDHandler == nil {
		api.DenominationsGetDenominationsIDHandler = denominations.GetDenominationsIDHandlerFunc(func(params denominations.GetDenominationsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation denominations.GetDenominationsID has not yet been implemented")
		})
	}
	if api.ImageGetImagesHandler == nil {
		api.ImageGetImagesHandler = image.GetImagesHandlerFunc(func(params image.GetImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation image.GetImages has not yet been implemented")
		})
	}
	if api.ImageGetImagesIDHandler == nil {
		api.ImageGetImagesIDHandler = image.GetImagesIDHandlerFunc(func(params image.GetImagesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation image.GetImagesID has not yet been implemented")
		})
	}
	if api.ImageGetImagesIDImgHandler == nil {
		api.ImageGetImagesIDImgHandler = image.GetImagesIDImgHandlerFunc(func(params image.GetImagesIDImgParams) middleware.Responder {
			return middleware.NotImplemented("operation image.GetImagesIDImg has not yet been implemented")
		})
	}
	if api.ServerGetInfoHandler == nil {
		api.ServerGetInfoHandler = server.GetInfoHandlerFunc(func(params server.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation server.GetInfo has not yet been implemented")
		})
	}
	if api.ProductsGetProductsHandler == nil {
		api.ProductsGetProductsHandler = products.GetProductsHandlerFunc(func(params products.GetProductsParams) middleware.Responder {
			return middleware.NotImplemented("operation products.GetProducts has not yet been implemented")
		})
	}
	if api.ProductsGetProductsIDHandler == nil {
		api.ProductsGetProductsIDHandler = products.GetProductsIDHandlerFunc(func(params products.GetProductsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation products.GetProductsID has not yet been implemented")
		})
	}
	if api.UsersGetUsersHandler == nil {
		api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})
	}
	if api.UsersGetUsersBarcodeBarcodeHandler == nil {
		api.UsersGetUsersBarcodeBarcodeHandler = users.GetUsersBarcodeBarcodeHandlerFunc(func(params users.GetUsersBarcodeBarcodeParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsersBarcodeBarcode has not yet been implemented")
		})
	}
	if api.UsersGetUsersIDHandler == nil {
		api.UsersGetUsersIDHandler = users.GetUsersIDHandlerFunc(func(params users.GetUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsersID has not yet been implemented")
		})
	}
	if api.UsersGetUsersStatsHandler == nil {
		api.UsersGetUsersStatsHandler = users.GetUsersStatsHandlerFunc(func(params users.GetUsersStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsersStats has not yet been implemented")
		})
	}
	if api.BarcodesPatchBarcodesIDHandler == nil {
		api.BarcodesPatchBarcodesIDHandler = barcodes.PatchBarcodesIDHandlerFunc(func(params barcodes.PatchBarcodesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation barcodes.PatchBarcodesID has not yet been implemented")
		})
	}
	if api.DenominationsPatchDenominationsIDHandler == nil {
		api.DenominationsPatchDenominationsIDHandler = denominations.PatchDenominationsIDHandlerFunc(func(params denominations.PatchDenominationsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation denominations.PatchDenominationsID has not yet been implemented")
		})
	}
	if api.ProductsPatchProductsIDHandler == nil {
		api.ProductsPatchProductsIDHandler = products.PatchProductsIDHandlerFunc(func(params products.PatchProductsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation products.PatchProductsID has not yet been implemented")
		})
	}
	if api.UsersPatchUsersIDHandler == nil {
		api.UsersPatchUsersIDHandler = users.PatchUsersIDHandlerFunc(func(params users.PatchUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PatchUsersID has not yet been implemented")
		})
	}
	if api.BarcodesPostBarcodesHandler == nil {
		api.BarcodesPostBarcodesHandler = barcodes.PostBarcodesHandlerFunc(func(params barcodes.PostBarcodesParams) middleware.Responder {
			return middleware.NotImplemented("operation barcodes.PostBarcodes has not yet been implemented")
		})
	}
	if api.DenominationsPostDenominationsHandler == nil {
		api.DenominationsPostDenominationsHandler = denominations.PostDenominationsHandlerFunc(func(params denominations.PostDenominationsParams) middleware.Responder {
			return middleware.NotImplemented("operation denominations.PostDenominations has not yet been implemented")
		})
	}
	if api.ImagePostImagesHandler == nil {
		api.ImagePostImagesHandler = image.PostImagesHandlerFunc(func(params image.PostImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation image.PostImages has not yet been implemented")
		})
	}
	if api.ProductsPostProductsHandler == nil {
		api.ProductsPostProductsHandler = products.PostProductsHandlerFunc(func(params products.PostProductsParams) middleware.Responder {
			return middleware.NotImplemented("operation products.PostProducts has not yet been implemented")
		})
	}
	if api.UsersPostUsersHandler == nil {
		api.UsersPostUsersHandler = users.PostUsersHandlerFunc(func(params users.PostUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsers has not yet been implemented")
		})
	}
	if api.UsersPostUsersIDBuyHandler == nil {
		api.UsersPostUsersIDBuyHandler = users.PostUsersIDBuyHandlerFunc(func(params users.PostUsersIDBuyParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersIDBuy has not yet been implemented")
		})
	}
	if api.UsersPostUsersIDBuyBarcodeHandler == nil {
		api.UsersPostUsersIDBuyBarcodeHandler = users.PostUsersIDBuyBarcodeHandlerFunc(func(params users.PostUsersIDBuyBarcodeParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersIDBuyBarcode has not yet been implemented")
		})
	}
	if api.UsersPostUsersIDDepositHandler == nil {
		api.UsersPostUsersIDDepositHandler = users.PostUsersIDDepositHandlerFunc(func(params users.PostUsersIDDepositParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersIDDeposit has not yet been implemented")
		})
	}
	if api.UsersPostUsersIDSpendHandler == nil {
		api.UsersPostUsersIDSpendHandler = users.PostUsersIDSpendHandlerFunc(func(params users.PostUsersIDSpendParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersIDSpend has not yet been implemented")
		})
	}
	if api.UsersPostUsersIDTransferHandler == nil {
		api.UsersPostUsersIDTransferHandler = users.PostUsersIDTransferHandlerFunc(func(params users.PostUsersIDTransferParams) middleware.Responder {
			return middleware.NotImplemented("operation users.PostUsersIDTransfer has not yet been implemented")
		})
	}

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
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
