/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package adapters


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateAdaptersListHandlerFunc turns a function with the right signature into a weaviate adapters list handler
type WeaviateAdaptersListHandlerFunc func(WeaviateAdaptersListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateAdaptersListHandlerFunc) Handle(params WeaviateAdaptersListParams) middleware.Responder {
	return fn(params)
}

// WeaviateAdaptersListHandler interface for that can handle valid weaviate adapters list params
type WeaviateAdaptersListHandler interface {
	Handle(WeaviateAdaptersListParams) middleware.Responder
}

// NewWeaviateAdaptersList creates a new http.Handler for the weaviate adapters list operation
func NewWeaviateAdaptersList(ctx *middleware.Context, handler WeaviateAdaptersListHandler) *WeaviateAdaptersList {
	return &WeaviateAdaptersList{Context: ctx, Handler: handler}
}

/*WeaviateAdaptersList swagger:route GET /adapters adapters weaviateAdaptersList

Lists all adapters.

*/
type WeaviateAdaptersList struct {
	Context *middleware.Context
	Handler WeaviateAdaptersListHandler
}

func (o *WeaviateAdaptersList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateAdaptersListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
