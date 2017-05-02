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
 package acl_entries




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateACLEntriesUpdateOKCode is the HTTP code returned for type WeaviateACLEntriesUpdateOK
const WeaviateACLEntriesUpdateOKCode int = 200

/*WeaviateACLEntriesUpdateOK Successful updated.

swagger:response weaviateAclEntriesUpdateOK
*/
type WeaviateACLEntriesUpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models.ACLEntry `json:"body,omitempty"`
}

// NewWeaviateACLEntriesUpdateOK creates WeaviateACLEntriesUpdateOK with default headers values
func NewWeaviateACLEntriesUpdateOK() *WeaviateACLEntriesUpdateOK {
	return &WeaviateACLEntriesUpdateOK{}
}

// WithPayload adds the payload to the weaviate Acl entries update o k response
func (o *WeaviateACLEntriesUpdateOK) WithPayload(payload *models.ACLEntry) *WeaviateACLEntriesUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate Acl entries update o k response
func (o *WeaviateACLEntriesUpdateOK) SetPayload(payload *models.ACLEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateACLEntriesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateACLEntriesUpdateNotImplementedCode is the HTTP code returned for type WeaviateACLEntriesUpdateNotImplemented
const WeaviateACLEntriesUpdateNotImplementedCode int = 501

/*WeaviateACLEntriesUpdateNotImplemented Not (yet) implemented.

swagger:response weaviateAclEntriesUpdateNotImplemented
*/
type WeaviateACLEntriesUpdateNotImplemented struct {
}

// NewWeaviateACLEntriesUpdateNotImplemented creates WeaviateACLEntriesUpdateNotImplemented with default headers values
func NewWeaviateACLEntriesUpdateNotImplemented() *WeaviateACLEntriesUpdateNotImplemented {
	return &WeaviateACLEntriesUpdateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateACLEntriesUpdateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
