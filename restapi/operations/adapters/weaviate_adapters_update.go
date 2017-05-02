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

// WeaviateAdaptersUpdateHandlerFunc turns a function with the right signature into a weaviate adapters update handler
type WeaviateAdaptersUpdateHandlerFunc func(WeaviateAdaptersUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateAdaptersUpdateHandlerFunc) Handle(params WeaviateAdaptersUpdateParams) middleware.Responder {
	return fn(params)
}

// WeaviateAdaptersUpdateHandler interface for that can handle valid weaviate adapters update params
type WeaviateAdaptersUpdateHandler interface {
	Handle(WeaviateAdaptersUpdateParams) middleware.Responder
}

// NewWeaviateAdaptersUpdate creates a new http.Handler for the weaviate adapters update operation
func NewWeaviateAdaptersUpdate(ctx *middleware.Context, handler WeaviateAdaptersUpdateHandler) *WeaviateAdaptersUpdate {
	return &WeaviateAdaptersUpdate{Context: ctx, Handler: handler}
}

/*WeaviateAdaptersUpdate swagger:route PUT /adapters/{adapterId} adapters weaviateAdaptersUpdate

Updates an adapter.

*/
type WeaviateAdaptersUpdate struct {
	Context *middleware.Context
	Handler WeaviateAdaptersUpdateHandler
}

func (o *WeaviateAdaptersUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateAdaptersUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
