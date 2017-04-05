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
 package commands




import (
	"errors"
	"net/url"

	"github.com/go-openapi/swag"
)

// WeaveCommandsInsertURL generates an URL for the weave commands insert operation
type WeaveCommandsInsertURL struct {
	Alt             *string
	ExecuteAfter    *string
	Fields          *string
	Hl              *string
	Key             *string
	OauthToken      *string
	PrettyPrint     *bool
	QuotaUser       *string
	ResponseAwaitMs *string
	UserIP          *string

	// avoid unkeyed usage
	_ struct{}
}

// Build a url path and query string
func (o *WeaveCommandsInsertURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/commands"

	result.Path = _path

	qs := make(url.Values)

	var alt string
	if o.Alt != nil {
		alt = *o.Alt
	}
	if alt != "" {
		qs.Set("alt", alt)
	}

	var executeAfter string
	if o.ExecuteAfter != nil {
		executeAfter = *o.ExecuteAfter
	}
	if executeAfter != "" {
		qs.Set("executeAfter", executeAfter)
	}

	var fields string
	if o.Fields != nil {
		fields = *o.Fields
	}
	if fields != "" {
		qs.Set("fields", fields)
	}

	var hl string
	if o.Hl != nil {
		hl = *o.Hl
	}
	if hl != "" {
		qs.Set("hl", hl)
	}

	var key string
	if o.Key != nil {
		key = *o.Key
	}
	if key != "" {
		qs.Set("key", key)
	}

	var oauthToken string
	if o.OauthToken != nil {
		oauthToken = *o.OauthToken
	}
	if oauthToken != "" {
		qs.Set("oauth_token", oauthToken)
	}

	var prettyPrint string
	if o.PrettyPrint != nil {
		prettyPrint = swag.FormatBool(*o.PrettyPrint)
	}
	if prettyPrint != "" {
		qs.Set("prettyPrint", prettyPrint)
	}

	var quotaUser string
	if o.QuotaUser != nil {
		quotaUser = *o.QuotaUser
	}
	if quotaUser != "" {
		qs.Set("quotaUser", quotaUser)
	}

	var responseAwaitMs string
	if o.ResponseAwaitMs != nil {
		responseAwaitMs = *o.ResponseAwaitMs
	}
	if responseAwaitMs != "" {
		qs.Set("responseAwaitMs", responseAwaitMs)
	}

	var userIP string
	if o.UserIP != nil {
		userIP = *o.UserIP
	}
	if userIP != "" {
		qs.Set("userIp", userIP)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WeaveCommandsInsertURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WeaveCommandsInsertURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WeaveCommandsInsertURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WeaveCommandsInsertURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WeaveCommandsInsertURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *WeaveCommandsInsertURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}