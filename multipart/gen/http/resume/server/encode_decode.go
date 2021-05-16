// Code generated by goa v3.3.1, DO NOT EDIT.
//
// resume HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o
// $(GOPATH)/src/goa.design/examples/multipart

package server

import (
	"context"
	"net/http"

	resume "goa.design/examples/multipart/gen/resume"
	resumeviews "goa.design/examples/multipart/gen/resume/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the resume
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(resumeviews.StoredResumeCollection)
		enc := encoder(ctx, w)
		body := NewStoredResumeResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeAddResponse returns an encoder for responses returned by the resume
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the resume add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload []*resume.Resume
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewResumeAddDecoder returns a decoder to decode the multipart request for
// the "resume" service "add" endpoint.
func NewResumeAddDecoder(mux goahttp.Muxer, resumeAddDecoderFn ResumeAddDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(*[]*resume.Resume)
			if err := resumeAddDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}

// marshalResumeviewsStoredResumeViewToStoredResumeResponse builds a value of
// type *StoredResumeResponse from a value of type
// *resumeviews.StoredResumeView.
func marshalResumeviewsStoredResumeViewToStoredResumeResponse(v *resumeviews.StoredResumeView) *StoredResumeResponse {
	res := &StoredResumeResponse{
		ID:        *v.ID,
		CreatedAt: *v.CreatedAt,
		Name:      *v.Name,
	}
	if v.Experience != nil {
		res.Experience = make([]*ExperienceResponse, len(v.Experience))
		for i, val := range v.Experience {
			res.Experience[i] = marshalResumeviewsExperienceViewToExperienceResponse(val)
		}
	}
	if v.Education != nil {
		res.Education = make([]*EducationResponse, len(v.Education))
		for i, val := range v.Education {
			res.Education[i] = marshalResumeviewsEducationViewToEducationResponse(val)
		}
	}

	return res
}

// marshalResumeviewsExperienceViewToExperienceResponse builds a value of type
// *ExperienceResponse from a value of type *resumeviews.ExperienceView.
func marshalResumeviewsExperienceViewToExperienceResponse(v *resumeviews.ExperienceView) *ExperienceResponse {
	res := &ExperienceResponse{
		Company:  *v.Company,
		Role:     *v.Role,
		Duration: *v.Duration,
	}

	return res
}

// marshalResumeviewsEducationViewToEducationResponse builds a value of type
// *EducationResponse from a value of type *resumeviews.EducationView.
func marshalResumeviewsEducationViewToEducationResponse(v *resumeviews.EducationView) *EducationResponse {
	res := &EducationResponse{
		Institution: *v.Institution,
		Major:       *v.Major,
	}

	return res
}

// unmarshalResumeRequestBodyToResumeResume builds a value of type
// *resume.Resume from a value of type *ResumeRequestBody.
func unmarshalResumeRequestBodyToResumeResume(v *ResumeRequestBody) *resume.Resume {
	res := &resume.Resume{
		Name: *v.Name,
	}
	if v.Experience != nil {
		res.Experience = make([]*resume.Experience, len(v.Experience))
		for i, val := range v.Experience {
			res.Experience[i] = unmarshalExperienceRequestBodyToResumeExperience(val)
		}
	}
	if v.Education != nil {
		res.Education = make([]*resume.Education, len(v.Education))
		for i, val := range v.Education {
			res.Education[i] = unmarshalEducationRequestBodyToResumeEducation(val)
		}
	}

	return res
}

// unmarshalExperienceRequestBodyToResumeExperience builds a value of type
// *resume.Experience from a value of type *ExperienceRequestBody.
func unmarshalExperienceRequestBodyToResumeExperience(v *ExperienceRequestBody) *resume.Experience {
	if v == nil {
		return nil
	}
	res := &resume.Experience{
		Company:  *v.Company,
		Role:     *v.Role,
		Duration: *v.Duration,
	}

	return res
}

// unmarshalEducationRequestBodyToResumeEducation builds a value of type
// *resume.Education from a value of type *EducationRequestBody.
func unmarshalEducationRequestBodyToResumeEducation(v *EducationRequestBody) *resume.Education {
	if v == nil {
		return nil
	}
	res := &resume.Education{
		Institution: *v.Institution,
		Major:       *v.Major,
	}

	return res
}
