// Code generated by go-swagger; DO NOT EDIT.

package feature_flags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFeatureFlagReader is a Reader for the DeleteFeatureFlag structure.
type DeleteFeatureFlagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFeatureFlagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFeatureFlagNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteFeatureFlagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFeatureFlagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFeatureFlagNoContent creates a DeleteFeatureFlagNoContent with default headers values
func NewDeleteFeatureFlagNoContent() *DeleteFeatureFlagNoContent {
	return &DeleteFeatureFlagNoContent{}
}

/*DeleteFeatureFlagNoContent handles this case with default header values.

Action completed successfully.
*/
type DeleteFeatureFlagNoContent struct {
}

func (o *DeleteFeatureFlagNoContent) Error() string {
	return fmt.Sprintf("[DELETE /flags/{projectKey}/{featureFlagKey}][%d] deleteFeatureFlagNoContent ", 204)
}

func (o *DeleteFeatureFlagNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFeatureFlagUnauthorized creates a DeleteFeatureFlagUnauthorized with default headers values
func NewDeleteFeatureFlagUnauthorized() *DeleteFeatureFlagUnauthorized {
	return &DeleteFeatureFlagUnauthorized{}
}

/*DeleteFeatureFlagUnauthorized handles this case with default header values.

Invalid access token.
*/
type DeleteFeatureFlagUnauthorized struct {
}

func (o *DeleteFeatureFlagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /flags/{projectKey}/{featureFlagKey}][%d] deleteFeatureFlagUnauthorized ", 401)
}

func (o *DeleteFeatureFlagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFeatureFlagNotFound creates a DeleteFeatureFlagNotFound with default headers values
func NewDeleteFeatureFlagNotFound() *DeleteFeatureFlagNotFound {
	return &DeleteFeatureFlagNotFound{}
}

/*DeleteFeatureFlagNotFound handles this case with default header values.

Invalid resource specifier.
*/
type DeleteFeatureFlagNotFound struct {
}

func (o *DeleteFeatureFlagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /flags/{projectKey}/{featureFlagKey}][%d] deleteFeatureFlagNotFound ", 404)
}

func (o *DeleteFeatureFlagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
