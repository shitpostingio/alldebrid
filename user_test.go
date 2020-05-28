package alldebrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetUserInfo(t *testing.T) {

	getUserEndpoint = func() string {
		return "http://127.0.0.1"
	}

	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		want       User
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "error",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          noapicl,
			assertion:  assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status":success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			want:       User{},
			assertion:  assert.Error,
		},
		{
			name:       "valid request",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			want: User{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(userinfo, getUserEndpoint(), tt.c.ic.appName, tt.c.ic.apikey),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.GetUserInfo()

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
