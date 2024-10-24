// Package order provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package ports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /customer/{costumerID}/orders)
	PostCustomerCostumerIDOrders(c *gin.Context, costumerID string)

	// (GET /customer/{costumerID}/orders/{orderID})
	GetCustomerCostumerIDOrdersOrderID(c *gin.Context, costumerID string, orderID string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostCustomerCostumerIDOrders operation middleware
func (siw *ServerInterfaceWrapper) PostCustomerCostumerIDOrders(c *gin.Context) {

	var err error

	// ------------- Path parameter "costumerID" -------------
	var costumerID string

	err = runtime.BindStyledParameterWithOptions("simple", "costumerID", c.Param("costumerID"), &costumerID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter costumerID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCustomerCostumerIDOrders(c, costumerID)
}

// GetCustomerCostumerIDOrdersOrderID operation middleware
func (siw *ServerInterfaceWrapper) GetCustomerCostumerIDOrdersOrderID(c *gin.Context) {

	var err error

	// ------------- Path parameter "costumerID" -------------
	var costumerID string

	err = runtime.BindStyledParameterWithOptions("simple", "costumerID", c.Param("costumerID"), &costumerID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter costumerID: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "orderID" -------------
	var orderID string

	err = runtime.BindStyledParameterWithOptions("simple", "orderID", c.Param("orderID"), &orderID, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter orderID: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCustomerCostumerIDOrdersOrderID(c, costumerID, orderID)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/customer/:costumerID/orders", wrapper.PostCustomerCostumerIDOrders)
	router.GET(options.BaseURL+"/customer/:costumerID/orders/:orderID", wrapper.GetCustomerCostumerIDOrdersOrderID)
}
