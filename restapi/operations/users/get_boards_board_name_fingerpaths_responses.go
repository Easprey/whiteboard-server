// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Easprey/whiteboard-server/models"
)

// GetBoardsBoardNameFingerpathsOKCode is the HTTP code returned for type GetBoardsBoardNameFingerpathsOK
const GetBoardsBoardNameFingerpathsOKCode int = 200

/*GetBoardsBoardNameFingerpathsOK fingerpaths associated with board

swagger:response getBoardsBoardNameFingerpathsOK
*/
type GetBoardsBoardNameFingerpathsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.FingerPath `json:"body,omitempty"`
}

// NewGetBoardsBoardNameFingerpathsOK creates GetBoardsBoardNameFingerpathsOK with default headers values
func NewGetBoardsBoardNameFingerpathsOK() *GetBoardsBoardNameFingerpathsOK {

	return &GetBoardsBoardNameFingerpathsOK{}
}

// WithPayload adds the payload to the get boards board name fingerpaths o k response
func (o *GetBoardsBoardNameFingerpathsOK) WithPayload(payload []*models.FingerPath) *GetBoardsBoardNameFingerpathsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get boards board name fingerpaths o k response
func (o *GetBoardsBoardNameFingerpathsOK) SetPayload(payload []*models.FingerPath) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBoardsBoardNameFingerpathsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.FingerPath, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetBoardsBoardNameFingerpathsBadRequestCode is the HTTP code returned for type GetBoardsBoardNameFingerpathsBadRequest
const GetBoardsBoardNameFingerpathsBadRequestCode int = 400

/*GetBoardsBoardNameFingerpathsBadRequest no associated board

swagger:response getBoardsBoardNameFingerpathsBadRequest
*/
type GetBoardsBoardNameFingerpathsBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetBoardsBoardNameFingerpathsBadRequest creates GetBoardsBoardNameFingerpathsBadRequest with default headers values
func NewGetBoardsBoardNameFingerpathsBadRequest() *GetBoardsBoardNameFingerpathsBadRequest {

	return &GetBoardsBoardNameFingerpathsBadRequest{}
}

// WithPayload adds the payload to the get boards board name fingerpaths bad request response
func (o *GetBoardsBoardNameFingerpathsBadRequest) WithPayload(payload interface{}) *GetBoardsBoardNameFingerpathsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get boards board name fingerpaths bad request response
func (o *GetBoardsBoardNameFingerpathsBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBoardsBoardNameFingerpathsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}