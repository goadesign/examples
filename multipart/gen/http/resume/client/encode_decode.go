// Code generated by goa v3.7.4, DO NOT EDIT.
//
// resume HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o multipart

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	resume "goa.design/examples/multipart/gen/resume"
	resumeviews "goa.design/examples/multipart/gen/resume/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "resume" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListResumePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("resume", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the resume
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("resume", "list", err)
			}
			p := NewListStoredResumeCollectionOK(body)
			view := "default"
			vres := resumeviews.StoredResumeCollection{Projected: p, View: view}
			if err = resumeviews.ValidateStoredResumeCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("resume", "list", err)
			}
			res := resume.NewStoredResumeCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("resume", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "resume" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddResumePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("resume", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the resume add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.([]*resume.Resume)
		if !ok {
			return goahttp.ErrInvalidType("resume", "add", "[]*resume.Resume", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("resume", "add", err)
		}
		return nil
	}
}

// NewResumeAddEncoder returns an encoder to encode the multipart request for
// the "resume" service "add" endpoint.
func NewResumeAddEncoder(encoderFn ResumeAddEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.([]*resume.Resume)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeAddResponse returns a decoder for responses returned by the resume add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("resume", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("resume", "add", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStoredResumeResponseToResumeviewsStoredResumeView builds a value of
// type *resumeviews.StoredResumeView from a value of type
// *StoredResumeResponse.
func unmarshalStoredResumeResponseToResumeviewsStoredResumeView(v *StoredResumeResponse) *resumeviews.StoredResumeView {
	res := &resumeviews.StoredResumeView{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		Name:      v.Name,
	}
	res.Experience = make([]*resumeviews.ExperienceView, len(v.Experience))
	for i, val := range v.Experience {
		res.Experience[i] = unmarshalExperienceResponseToResumeviewsExperienceView(val)
	}
	res.Education = make([]*resumeviews.EducationView, len(v.Education))
	for i, val := range v.Education {
		res.Education[i] = unmarshalEducationResponseToResumeviewsEducationView(val)
	}

	return res
}

// unmarshalExperienceResponseToResumeviewsExperienceView builds a value of
// type *resumeviews.ExperienceView from a value of type *ExperienceResponse.
func unmarshalExperienceResponseToResumeviewsExperienceView(v *ExperienceResponse) *resumeviews.ExperienceView {
	res := &resumeviews.ExperienceView{
		Company:  v.Company,
		Role:     v.Role,
		Duration: v.Duration,
	}

	return res
}

// unmarshalEducationResponseToResumeviewsEducationView builds a value of type
// *resumeviews.EducationView from a value of type *EducationResponse.
func unmarshalEducationResponseToResumeviewsEducationView(v *EducationResponse) *resumeviews.EducationView {
	res := &resumeviews.EducationView{
		Institution: v.Institution,
		Major:       v.Major,
	}

	return res
}

// marshalResumeResumeToResumeRequestBody builds a value of type
// *ResumeRequestBody from a value of type *resume.Resume.
func marshalResumeResumeToResumeRequestBody(v *resume.Resume) *ResumeRequestBody {
	res := &ResumeRequestBody{
		Name: v.Name,
	}
	if v.Experience != nil {
		res.Experience = make([]*ExperienceRequestBody, len(v.Experience))
		for i, val := range v.Experience {
			res.Experience[i] = marshalResumeExperienceToExperienceRequestBody(val)
		}
	}
	if v.Education != nil {
		res.Education = make([]*EducationRequestBody, len(v.Education))
		for i, val := range v.Education {
			res.Education[i] = marshalResumeEducationToEducationRequestBody(val)
		}
	}

	return res
}

// marshalResumeExperienceToExperienceRequestBody builds a value of type
// *ExperienceRequestBody from a value of type *resume.Experience.
func marshalResumeExperienceToExperienceRequestBody(v *resume.Experience) *ExperienceRequestBody {
	if v == nil {
		return nil
	}
	res := &ExperienceRequestBody{
		Company:  v.Company,
		Role:     v.Role,
		Duration: v.Duration,
	}

	return res
}

// marshalResumeEducationToEducationRequestBody builds a value of type
// *EducationRequestBody from a value of type *resume.Education.
func marshalResumeEducationToEducationRequestBody(v *resume.Education) *EducationRequestBody {
	if v == nil {
		return nil
	}
	res := &EducationRequestBody{
		Institution: v.Institution,
		Major:       v.Major,
	}

	return res
}

// marshalResumeRequestBodyToResumeResume builds a value of type *resume.Resume
// from a value of type *ResumeRequestBody.
func marshalResumeRequestBodyToResumeResume(v *ResumeRequestBody) *resume.Resume {
	res := &resume.Resume{
		Name: v.Name,
	}
	if v.Experience != nil {
		res.Experience = make([]*resume.Experience, len(v.Experience))
		for i, val := range v.Experience {
			res.Experience[i] = marshalExperienceRequestBodyToResumeExperience(val)
		}
	}
	if v.Education != nil {
		res.Education = make([]*resume.Education, len(v.Education))
		for i, val := range v.Education {
			res.Education[i] = marshalEducationRequestBodyToResumeEducation(val)
		}
	}

	return res
}

// marshalExperienceRequestBodyToResumeExperience builds a value of type
// *resume.Experience from a value of type *ExperienceRequestBody.
func marshalExperienceRequestBodyToResumeExperience(v *ExperienceRequestBody) *resume.Experience {
	if v == nil {
		return nil
	}
	res := &resume.Experience{
		Company:  v.Company,
		Role:     v.Role,
		Duration: v.Duration,
	}

	return res
}

// marshalEducationRequestBodyToResumeEducation builds a value of type
// *resume.Education from a value of type *EducationRequestBody.
func marshalEducationRequestBodyToResumeEducation(v *EducationRequestBody) *resume.Education {
	if v == nil {
		return nil
	}
	res := &resume.Education{
		Institution: v.Institution,
		Major:       v.Major,
	}

	return res
}
