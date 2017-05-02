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
 package model_manifests


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateModelManifestsUpdateHandlerFunc turns a function with the right signature into a weaviate model manifests update handler
type WeaviateModelManifestsUpdateHandlerFunc func(WeaviateModelManifestsUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateModelManifestsUpdateHandlerFunc) Handle(params WeaviateModelManifestsUpdateParams) middleware.Responder {
	return fn(params)
}

// WeaviateModelManifestsUpdateHandler interface for that can handle valid weaviate model manifests update params
type WeaviateModelManifestsUpdateHandler interface {
	Handle(WeaviateModelManifestsUpdateParams) middleware.Responder
}

// NewWeaviateModelManifestsUpdate creates a new http.Handler for the weaviate model manifests update operation
func NewWeaviateModelManifestsUpdate(ctx *middleware.Context, handler WeaviateModelManifestsUpdateHandler) *WeaviateModelManifestsUpdate {
	return &WeaviateModelManifestsUpdate{Context: ctx, Handler: handler}
}

/*WeaviateModelManifestsUpdate swagger:route PUT /modelManifests/{modelManifestId} modelManifests weaviateModelManifestsUpdate

Updates a particular model manifest.

*/
type WeaviateModelManifestsUpdate struct {
	Context *middleware.Context
	Handler WeaviateModelManifestsUpdateHandler
}

func (o *WeaviateModelManifestsUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateModelManifestsUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
