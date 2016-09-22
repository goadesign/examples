//************************************************************************//
// API "adder": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/examples/endpoints/design
// --out=$(GOPATH)/src/github.com/goadesign/examples/endpoints
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// User info extracted from security token (default view)
//
// Identifier: application/vnd.goa-cellar.auth+json; view=default
type Auth struct {
	// User email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User ID
	ID string `form:"id" json:"id" xml:"id"`
	// Token issuer
	Issuer *string `form:"issuer,omitempty" json:"issuer,omitempty" xml:"issuer,omitempty"`
}

// Validate validates the Auth media type instance.
func (mt *Auth) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	return
}
