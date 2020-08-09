// Code generated by goa v2.2.1, DO NOT EDIT.
//
// resume HTTP server types
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o
// $(GOPATH)/src/goa.design/examples/multipart

package server

import (
	resume "goa.design/examples/multipart/gen/resume"
	resumeviews "goa.design/examples/multipart/gen/resume/views"
	goa "goa.design/goa"
)

// StoredResumeResponseCollection is the type of the "resume" service "list"
// endpoint HTTP response body.
type StoredResumeResponseCollection []*StoredResumeResponse

// StoredResumeResponse is used to define fields on response body types.
type StoredResumeResponse struct {
	// ID of the resume
	ID int `form:"id" json:"id" xml:"id"`
	// Name in the resume
	Name string `form:"name" json:"name" xml:"name"`
	// Experience section in the resume
	Experience []*ExperienceResponse `form:"experience" json:"experience" xml:"experience"`
	// Education section in the resume
	Education []*EducationResponse `form:"education" json:"education" xml:"education"`
	// Time when resume was created
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
}

// ExperienceResponse is used to define fields on response body types.
type ExperienceResponse struct {
	// Name of the company
	Company string `form:"company" json:"company" xml:"company"`
	// Name of the role in the company
	Role string `form:"role" json:"role" xml:"role"`
	// Duration (in years) in the company
	Duration int `form:"duration" json:"duration" xml:"duration"`
}

// EducationResponse is used to define fields on response body types.
type EducationResponse struct {
	// Name of the institution
	Institution string `form:"institution" json:"institution" xml:"institution"`
	// Major name
	Major string `form:"major" json:"major" xml:"major"`
}

// ResumeRequestBody is used to define fields on request body types.
type ResumeRequestBody struct {
	// Name in the resume
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Experience section in the resume
	Experience []*ExperienceRequestBody `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
	// Education section in the resume
	Education []*EducationRequestBody `form:"education,omitempty" json:"education,omitempty" xml:"education,omitempty"`
}

// ExperienceRequestBody is used to define fields on request body types.
type ExperienceRequestBody struct {
	// Name of the company
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// Name of the role in the company
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// Duration (in years) in the company
	Duration *int `form:"duration,omitempty" json:"duration,omitempty" xml:"duration,omitempty"`
}

// EducationRequestBody is used to define fields on request body types.
type EducationRequestBody struct {
	// Name of the institution
	Institution *string `form:"institution,omitempty" json:"institution,omitempty" xml:"institution,omitempty"`
	// Major name
	Major *string `form:"major,omitempty" json:"major,omitempty" xml:"major,omitempty"`
}

// NewStoredResumeResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "resume" service.
func NewStoredResumeResponseCollection(res resumeviews.StoredResumeCollectionView) StoredResumeResponseCollection {
	body := make([]*StoredResumeResponse, len(res))
	for i, val := range res {
		body[i] = marshalResumeviewsStoredResumeViewToStoredResumeResponse(val)
	}
	return body
}

// NewAddResume builds a resume service add endpoint payload.
func NewAddResume(body []*ResumeRequestBody) []*resume.Resume {
	v := make([]*resume.Resume, len(body))
	for i, val := range body {
		v[i] = unmarshalResumeRequestBodyToResumeResume(val)
	}
	return v
}

// ValidateResumeRequestBody runs the validations defined on ResumeRequestBody
func ValidateResumeRequestBody(body *ResumeRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	for _, e := range body.Experience {
		if e != nil {
			if err2 := ValidateExperienceRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.Education {
		if e != nil {
			if err2 := ValidateEducationRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateExperienceRequestBody runs the validations defined on
// ExperienceRequestBody
func ValidateExperienceRequestBody(body *ExperienceRequestBody) (err error) {
	if body.Company == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("company", "body"))
	}
	if body.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "body"))
	}
	if body.Duration == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("duration", "body"))
	}
	return
}

// ValidateEducationRequestBody runs the validations defined on
// EducationRequestBody
func ValidateEducationRequestBody(body *EducationRequestBody) (err error) {
	if body.Institution == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("institution", "body"))
	}
	if body.Major == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("major", "body"))
	}
	return
}
