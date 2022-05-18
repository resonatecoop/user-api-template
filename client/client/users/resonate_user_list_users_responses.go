// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/resonatecoop/user-api/client/models"
)

// ResonateUserListUsersReader is a Reader for the ResonateUserListUsers structure.
type ResonateUserListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResonateUserListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResonateUserListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResonateUserListUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResonateUserListUsersOK creates a ResonateUserListUsersOK with default headers values
func NewResonateUserListUsersOK() *ResonateUserListUsersOK {
	return &ResonateUserListUsersOK{}
}

/* ResonateUserListUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResonateUserListUsersOK struct {
	Payload *models.UserUserListResponse
}

func (o *ResonateUserListUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] resonateUserListUsersOK  %+v", 200, o.Payload)
}
func (o *ResonateUserListUsersOK) GetPayload() *models.UserUserListResponse {
	return o.Payload
}

func (o *ResonateUserListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUserListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResonateUserListUsersDefault creates a ResonateUserListUsersDefault with default headers values
func NewResonateUserListUsersDefault(code int) *ResonateUserListUsersDefault {
	return &ResonateUserListUsersDefault{
		_statusCode: code,
	}
}

/* ResonateUserListUsersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResonateUserListUsersDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the resonate user list users default response
func (o *ResonateUserListUsersDefault) Code() int {
	return o._statusCode
}

func (o *ResonateUserListUsersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] ResonateUser_ListUsers default  %+v", o._statusCode, o.Payload)
}
func (o *ResonateUserListUsersDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ResonateUserListUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}