// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// GetReposReader is a Reader for the GetRepos structure.
type GetReposReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReposReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReposOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetReposInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReposOK creates a GetReposOK with default headers values
func NewGetReposOK() *GetReposOK {
	return &GetReposOK{}
}

/*GetReposOK handles this case with default header values.

A list of repos
*/
type GetReposOK struct {
	Payload *models.GetReposOKBody
}

func (o *GetReposOK) Error() string {
	return fmt.Sprintf("[GET /repos][%d] getReposOK  %+v", 200, o.Payload)
}

func (o *GetReposOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetReposOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReposInternalServerError creates a GetReposInternalServerError with default headers values
func NewGetReposInternalServerError() *GetReposInternalServerError {
	return &GetReposInternalServerError{}
}

/*GetReposInternalServerError handles this case with default header values.

An unexpected error occured.
*/
type GetReposInternalServerError struct {
	Payload *models.Error
}

func (o *GetReposInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repos][%d] getReposInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReposInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
