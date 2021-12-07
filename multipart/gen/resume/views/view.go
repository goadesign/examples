// Code generated by goa v3.5.3, DO NOT EDIT.
//
// resume views
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o
// $(GOPATH)/src/goa.design/examples/multipart

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StoredResumeCollection is the viewed result type that is projected based on
// a view.
type StoredResumeCollection struct {
	// Type to project
	Projected StoredResumeCollectionView
	// View to render
	View string
}

// StoredResumeCollectionView is a type that runs validations on a projected
// type.
type StoredResumeCollectionView []*StoredResumeView

// StoredResumeView is a type that runs validations on a projected type.
type StoredResumeView struct {
	// ID of the resume
	ID *int
	// Time when resume was created
	CreatedAt *string
	// Name in the resume
	Name *string
	// Experience section in the resume
	Experience []*ExperienceView
	// Education section in the resume
	Education []*EducationView
}

// ExperienceView is a type that runs validations on a projected type.
type ExperienceView struct {
	// Name of the company
	Company *string
	// Name of the role in the company
	Role *string
	// Duration (in years) in the company
	Duration *int
}

// EducationView is a type that runs validations on a projected type.
type EducationView struct {
	// Name of the institution
	Institution *string
	// Major name
	Major *string
}

var (
	// StoredResumeCollectionMap is a map indexing the attribute names of
	// StoredResumeCollection by view name.
	StoredResumeCollectionMap = map[string][]string{
		"default": {
			"id",
			"name",
			"experience",
			"education",
			"created_at",
		},
	}
	// StoredResumeMap is a map indexing the attribute names of StoredResume by
	// view name.
	StoredResumeMap = map[string][]string{
		"default": {
			"id",
			"name",
			"experience",
			"education",
			"created_at",
		},
	}
)

// ValidateStoredResumeCollection runs the validations defined on the viewed
// result type StoredResumeCollection.
func ValidateStoredResumeCollection(result StoredResumeCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredResumeCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStoredResumeCollectionView runs the validations defined on
// StoredResumeCollectionView using the "default" view.
func ValidateStoredResumeCollectionView(result StoredResumeCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredResumeView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredResumeView runs the validations defined on StoredResumeView
// using the "default" view.
func ValidateStoredResumeView(result *StoredResumeView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "result"))
	}
	if result.Education == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("education", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	for _, e := range result.Experience {
		if e != nil {
			if err2 := ValidateExperienceView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Education {
		if e != nil {
			if err2 := ValidateEducationView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateExperienceView runs the validations defined on ExperienceView.
func ValidateExperienceView(result *ExperienceView) (err error) {
	if result.Company == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("company", "result"))
	}
	if result.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "result"))
	}
	if result.Duration == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("duration", "result"))
	}
	return
}

// ValidateEducationView runs the validations defined on EducationView.
func ValidateEducationView(result *EducationView) (err error) {
	if result.Institution == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("institution", "result"))
	}
	if result.Major == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("major", "result"))
	}
	return
}
