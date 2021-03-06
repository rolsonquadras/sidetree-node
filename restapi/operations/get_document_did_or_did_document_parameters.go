// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDocumentDidOrDidDocumentParams creates a new GetDocumentDidOrDidDocumentParams object
// no default values defined in spec.
func NewGetDocumentDidOrDidDocumentParams() GetDocumentDidOrDidDocumentParams {

	return GetDocumentDidOrDidDocumentParams{}
}

// GetDocumentDidOrDidDocumentParams contains all the bound params for the get document did or did document operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDocumentDidOrDidDocument
type GetDocumentDidOrDidDocumentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*DidOrDidDocument can either be 1. Fully qualified DID 2. An encoded DID Document prefixed by the DID method name
	  Required: true
	  In: path
	*/
	DidOrDidDocument string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDocumentDidOrDidDocumentParams() beforehand.
func (o *GetDocumentDidOrDidDocumentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDidOrDidDocument, rhkDidOrDidDocument, _ := route.Params.GetOK("DidOrDidDocument")
	if err := o.bindDidOrDidDocument(rDidOrDidDocument, rhkDidOrDidDocument, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDidOrDidDocument binds and validates parameter DidOrDidDocument from path.
func (o *GetDocumentDidOrDidDocumentParams) bindDidOrDidDocument(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.DidOrDidDocument = raw

	return nil
}
