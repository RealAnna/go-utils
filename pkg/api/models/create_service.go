// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateService create service
// swagger:model CreateService
type CreateService struct {

	// deployment strategies
	DeploymentStrategies map[string]string `json:"deploymentStrategies,omitempty"`

	// helm chart
	HelmChart string `json:"helmChart,omitempty"`

	// service name
	// Required: true
	ServiceName *string `json:"serviceName"`
}

// Validate validates this create service
func (m *CreateService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateService) validateServiceName(formats strfmt.Registry) error {

	if err := validate.Required("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateService) UnmarshalBinary(b []byte) error {
	var res CreateService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}