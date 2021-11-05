// This file is safe to edit. Once it exists it will not be overwritten

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_server

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/openziti/fabric/rest_server/operations"
	"github.com/openziti/fabric/rest_server/operations/router"
	"github.com/openziti/fabric/rest_server/operations/service"
	"github.com/openziti/fabric/rest_server/operations/terminator"
)

//go:generate swagger generate server --target ../../fabric --name ZitiFabric --spec ../specs/swagger.yml --model-package rest_model --server-package rest_server --principal interface{} --exclude-main

func configureFlags(api *operations.ZitiFabricAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ZitiFabricAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

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

	if api.RouterCreateRouterHandler == nil {
		api.RouterCreateRouterHandler = router.CreateRouterHandlerFunc(func(params router.CreateRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.CreateRouter has not yet been implemented")
		})
	}
	if api.ServiceCreateServiceHandler == nil {
		api.ServiceCreateServiceHandler = service.CreateServiceHandlerFunc(func(params service.CreateServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.CreateService has not yet been implemented")
		})
	}
	if api.TerminatorCreateTerminatorHandler == nil {
		api.TerminatorCreateTerminatorHandler = terminator.CreateTerminatorHandlerFunc(func(params terminator.CreateTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.CreateTerminator has not yet been implemented")
		})
	}
	if api.RouterDeleteRouterHandler == nil {
		api.RouterDeleteRouterHandler = router.DeleteRouterHandlerFunc(func(params router.DeleteRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.DeleteRouter has not yet been implemented")
		})
	}
	if api.ServiceDeleteServiceHandler == nil {
		api.ServiceDeleteServiceHandler = service.DeleteServiceHandlerFunc(func(params service.DeleteServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteService has not yet been implemented")
		})
	}
	if api.TerminatorDeleteTerminatorHandler == nil {
		api.TerminatorDeleteTerminatorHandler = terminator.DeleteTerminatorHandlerFunc(func(params terminator.DeleteTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.DeleteTerminator has not yet been implemented")
		})
	}
	if api.RouterDetailRouterHandler == nil {
		api.RouterDetailRouterHandler = router.DetailRouterHandlerFunc(func(params router.DetailRouterParams) middleware.Responder {
			return middleware.NotImplemented("operation router.DetailRouter has not yet been implemented")
		})
	}
	if api.ServiceDetailServiceHandler == nil {
		api.ServiceDetailServiceHandler = service.DetailServiceHandlerFunc(func(params service.DetailServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DetailService has not yet been implemented")
		})
	}
	if api.TerminatorDetailTerminatorHandler == nil {
		api.TerminatorDetailTerminatorHandler = terminator.DetailTerminatorHandlerFunc(func(params terminator.DetailTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.DetailTerminator has not yet been implemented")
		})
	}
	if api.RouterListRouterTerminatorsHandler == nil {
		api.RouterListRouterTerminatorsHandler = router.ListRouterTerminatorsHandlerFunc(func(params router.ListRouterTerminatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation router.ListRouterTerminators has not yet been implemented")
		})
	}
	if api.RouterListRoutersHandler == nil {
		api.RouterListRoutersHandler = router.ListRoutersHandlerFunc(func(params router.ListRoutersParams) middleware.Responder {
			return middleware.NotImplemented("operation router.ListRouters has not yet been implemented")
		})
	}
	if api.ServiceListServiceTerminatorsHandler == nil {
		api.ServiceListServiceTerminatorsHandler = service.ListServiceTerminatorsHandlerFunc(func(params service.ListServiceTerminatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ListServiceTerminators has not yet been implemented")
		})
	}
	if api.ServiceListServicesHandler == nil {
		api.ServiceListServicesHandler = service.ListServicesHandlerFunc(func(params service.ListServicesParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ListServices has not yet been implemented")
		})
	}
	if api.TerminatorListTerminatorsHandler == nil {
		api.TerminatorListTerminatorsHandler = terminator.ListTerminatorsHandlerFunc(func(params terminator.ListTerminatorsParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.ListTerminators has not yet been implemented")
		})
	}
	if api.ServicePatchServiceHandler == nil {
		api.ServicePatchServiceHandler = service.PatchServiceHandlerFunc(func(params service.PatchServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.PatchService has not yet been implemented")
		})
	}
	if api.TerminatorPatchTerminatorHandler == nil {
		api.TerminatorPatchTerminatorHandler = terminator.PatchTerminatorHandlerFunc(func(params terminator.PatchTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.PatchTerminator has not yet been implemented")
		})
	}
	if api.ServiceUpdateServiceHandler == nil {
		api.ServiceUpdateServiceHandler = service.UpdateServiceHandlerFunc(func(params service.UpdateServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.UpdateService has not yet been implemented")
		})
	}
	if api.TerminatorUpdateTerminatorHandler == nil {
		api.TerminatorUpdateTerminatorHandler = terminator.UpdateTerminatorHandlerFunc(func(params terminator.UpdateTerminatorParams) middleware.Responder {
			return middleware.NotImplemented("operation terminator.UpdateTerminator has not yet been implemented")
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
