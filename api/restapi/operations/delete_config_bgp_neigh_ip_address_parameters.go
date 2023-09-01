// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteConfigBgpNeighIPAddressParams creates a new DeleteConfigBgpNeighIPAddressParams object
//
// There are no default values defined in the spec.
func NewDeleteConfigBgpNeighIPAddressParams() DeleteConfigBgpNeighIPAddressParams {

	return DeleteConfigBgpNeighIPAddressParams{}
}

// DeleteConfigBgpNeighIPAddressParams contains all the bound params for the delete config bgp neigh IP address operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteConfigBgpNeighIPAddress
type DeleteConfigBgpNeighIPAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Neighbor IP address
	  Required: true
	  In: path
	*/
	IPAddress string
	/*Remote ASN number
	  In: query
	*/
	RemoteAs *int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConfigBgpNeighIPAddressParams() beforehand.
func (o *DeleteConfigBgpNeighIPAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rIPAddress, rhkIPAddress, _ := route.Params.GetOK("ip_address")
	if err := o.bindIPAddress(rIPAddress, rhkIPAddress, route.Formats); err != nil {
		res = append(res, err)
	}

	qRemoteAs, qhkRemoteAs, _ := qs.GetOK("remoteAs")
	if err := o.bindRemoteAs(qRemoteAs, qhkRemoteAs, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIPAddress binds and validates parameter IPAddress from path.
func (o *DeleteConfigBgpNeighIPAddressParams) bindIPAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.IPAddress = raw

	return nil
}

// bindRemoteAs binds and validates parameter RemoteAs from query.
func (o *DeleteConfigBgpNeighIPAddressParams) bindRemoteAs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("remoteAs", "query", "int32", raw)
	}
	o.RemoteAs = &value

	return nil
}
