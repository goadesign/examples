// Code generated by goa v3.21.1, DO NOT EDIT.
//
// resume HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/multipart/design

package client

import (
	"encoding/json"
	"fmt"

	resume "goa.design/examples/multipart/gen/resume"
)

// BuildAddPayload builds the payload for the resume add endpoint from CLI
// flags.
func BuildAddPayload(resumeAddBody string) ([]*resume.Resume, error) {
	var err error
	var body []*ResumeRequestBody
	{
		err = json.Unmarshal([]byte(resumeAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      {\n         \"education\": [\n            {\n               \"institution\": \"Nihil voluptate dolorem ut ipsam fuga.\",\n               \"major\": \"Ex sequi.\"\n            },\n            {\n               \"institution\": \"Nihil voluptate dolorem ut ipsam fuga.\",\n               \"major\": \"Ex sequi.\"\n            }\n         ],\n         \"experience\": [\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            },\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            },\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            },\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            }\n         ],\n         \"name\": \"In qui dolorum adipisci itaque.\"\n      },\n      {\n         \"education\": [\n            {\n               \"institution\": \"Nihil voluptate dolorem ut ipsam fuga.\",\n               \"major\": \"Ex sequi.\"\n            },\n            {\n               \"institution\": \"Nihil voluptate dolorem ut ipsam fuga.\",\n               \"major\": \"Ex sequi.\"\n            }\n         ],\n         \"experience\": [\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            },\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            },\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            },\n            {\n               \"company\": \"Saepe similique libero sunt.\",\n               \"duration\": 4538661945567210561,\n               \"role\": \"Cum maiores quo at ducimus sit.\"\n            }\n         ],\n         \"name\": \"In qui dolorum adipisci itaque.\"\n      }\n   ]'")
		}
	}
	v := make([]*resume.Resume, len(body))
	for i, val := range body {
		v[i] = marshalResumeRequestBodyToResumeResume(val)
	}
	return v, nil
}
