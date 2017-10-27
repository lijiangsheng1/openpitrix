// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// PostAppsReader is a Reader for the PostApps structure.
type PostAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostAppsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAppsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAppsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAppsNoContent creates a PostAppsNoContent with default headers values
func NewPostAppsNoContent() *PostAppsNoContent {
	return &PostAppsNoContent{}
}

/*PostAppsNoContent handles this case with default header values.

App succesfully created.
*/
type PostAppsNoContent struct {
}

func (o *PostAppsNoContent) Error() string {
	return fmt.Sprintf("[POST /apps][%d] postAppsNoContent ", 204)
}

func (o *PostAppsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppsBadRequest creates a PostAppsBadRequest with default headers values
func NewPostAppsBadRequest() *PostAppsBadRequest {
	return &PostAppsBadRequest{}
}

/*PostAppsBadRequest handles this case with default header values.

App couldn't have been created.
*/
type PostAppsBadRequest struct {
}

func (o *PostAppsBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps][%d] postAppsBadRequest ", 400)
}

func (o *PostAppsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAppsInternalServerError creates a PostAppsInternalServerError with default headers values
func NewPostAppsInternalServerError() *PostAppsInternalServerError {
	return &PostAppsInternalServerError{}
}

/*PostAppsInternalServerError handles this case with default header values.

An unexpected error occured.
*/
type PostAppsInternalServerError struct {
	Payload *models.Error
}

func (o *PostAppsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps][%d] postAppsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAppsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
