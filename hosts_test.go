package alldebrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetDomainsOnly(t *testing.T) {
	getHostsEndpoint = func() string {
		return "http://127.0.0.1"
	}

	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		want       Domains
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "bad json",
			jsonResp:   `{"status":"successful",}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "status error",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			assertion:  assert.Error,
		},
		{
			name:       "status successful",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			want: Domains{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(hostsdomains, getHostsEndpoint(), tt.c.ic.appName),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.GetDomainsOnly()

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
