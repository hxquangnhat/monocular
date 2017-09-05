package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Resource resource

swagger:model resource
*/
type Resource struct {

	/* attributes

	Required: true
	*/
	Attributes interface{} `json:"attributes"`

	/* id

	Required: true
	Min Length: 1
	*/
	ID *string `json:"id"`

	/* links
	 */
	Links interface{} `json:"links,omitempty"`

	/* relationships
	 */
	Relationships interface{} `json:"relationships,omitempty"`

	/* type

	Required: true
	Min Length: 1
	*/
	Type *string `json:"type"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateAttributes(formats strfmt.Registry) error {

	return nil
}

func (m *Resource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", string(*m.Type), 1); err != nil {
		return err
	}

	return nil
}