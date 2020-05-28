package alldebrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_UnlockLink(t *testing.T) {

	getLinksEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		link string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       UnlockLinkResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "error",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status":success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			want: UnlockLinkResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(linkunlock, getLinksEndpoint(), tt.c.ic.appName, tt.c.ic.apikey, tt.args.link),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.UnlockLink(tt.args.link)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_StreamingLink(t *testing.T) {

	getLinksEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		stream string
		id     string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       StreamingResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "error",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status":success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			want: StreamingResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(linkstreaming, getLinksEndpoint(), tt.c.ic.appName, tt.c.ic.apikey, tt.args.stream, tt.args.id),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.StreamingLink(tt.args.stream, tt.args.id)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_DelayedLink(t *testing.T) {

	getLinksEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		delayedID string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       DelayedResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "error",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status":success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			want: DelayedResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(linkdelayed, getLinksEndpoint(), tt.c.ic.appName, tt.c.ic.apikey, tt.args.delayedID),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.DelayedLink(tt.args.delayedID)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
