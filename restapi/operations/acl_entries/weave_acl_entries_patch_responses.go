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
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package acl_entries




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

/*WeaveACLEntriesPatchOK Successful response

swagger:response weaveAclEntriesPatchOK
*/
type WeaveACLEntriesPatchOK struct {

	// In: body
	Payload *models.ACLEntry `json:"body,omitempty"`
}

// NewWeaveACLEntriesPatchOK creates WeaveACLEntriesPatchOK with default headers values
func NewWeaveACLEntriesPatchOK() *WeaveACLEntriesPatchOK {
	return &WeaveACLEntriesPatchOK{}
}

// WithPayload adds the payload to the weave Acl entries patch o k response
func (o *WeaveACLEntriesPatchOK) WithPayload(payload *models.ACLEntry) *WeaveACLEntriesPatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave Acl entries patch o k response
func (o *WeaveACLEntriesPatchOK) SetPayload(payload *models.ACLEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveACLEntriesPatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}