// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBoardsFingerpathsByBoardNameGetParams creates a new BoardsFingerpathsByBoardNameGetParams object
// no default values defined in spec.
func NewBoardsFingerpathsByBoardNameGetParams() BoardsFingerpathsByBoardNameGetParams {

	return BoardsFingerpathsByBoardNameGetParams{}
}

// BoardsFingerpathsByBoardNameGetParams contains all the bound params for the boards fingerpaths by board name get operation
// typically these are obtained from a http.Request
//
// swagger:parameters BoardsFingerpathsByBoardNameGet
type BoardsFingerpathsByBoardNameGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the name of the board
	  Required: true
	  In: path
	*/
	BoardName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewBoardsFingerpathsByBoardNameGetParams() beforehand.
func (o *BoardsFingerpathsByBoardNameGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBoardName, rhkBoardName, _ := route.Params.GetOK("boardName")
	if err := o.bindBoardName(rBoardName, rhkBoardName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBoardName binds and validates parameter BoardName from path.
func (o *BoardsFingerpathsByBoardNameGetParams) bindBoardName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.BoardName = raw

	return nil
}
