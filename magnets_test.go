package alldebrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_UploadMagnet(t *testing.T) {

	getMagnetEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		magnets []string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       MagnetsUploadResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name: "no magnet",
			jsonResp: `{
				"status": "error",
				"error": {
				  "code": "MAGNET_NO_URI",
				  "message": "No magnet sent"
				}
			}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnets: []string{""},
			},
			assertion: assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status:"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnets: []string{""},
			},
			assertion: assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status": "success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnets: []string{"magnet:?xt=urn:btih:f95c371d5609d15f6615139be84edbb5b94a79bc"},
			},
			want: MagnetsUploadResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("POST", fmt.Sprintf(magnetupload, getMagnetEndpoint(), tt.c.ic.appName, tt.c.ic.apikey),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.UploadMagnet(tt.args.magnets)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_StatusMagnet(t *testing.T) {

	getMagnetEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		magnetID string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       StatusMagnetResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "no id",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			assertion: assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status:"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			assertion: assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			want: StatusMagnetResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(magnetstatus, getMagnetEndpoint(), tt.c.ic.appName, tt.c.ic.apikey, tt.args.magnetID),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.StatusMagnet(tt.args.magnetID)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_DeleteMagnet(t *testing.T) {

	getMagnetEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		magnetID string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       DeleteMagnetResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "no id",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			assertion: assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status:"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			assertion: assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			want: DeleteMagnetResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(magnetdelete, getMagnetEndpoint(), tt.c.ic.appName, tt.c.ic.apikey, tt.args.magnetID),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.DeleteMagnet(tt.args.magnetID)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_RestartMagnet(t *testing.T) {

	getMagnetEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		magnetID string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       RestartMagnetResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "no id",
			jsonResp:   `{"status":"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			assertion: assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status:"error"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			assertion: assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status":"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnetID: "",
			},
			want: RestartMagnetResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(magnetrestart, getMagnetEndpoint(), tt.c.ic.appName, tt.c.ic.apikey, tt.args.magnetID),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.RestartMagnet(tt.args.magnetID)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_InstantAvailability(t *testing.T) {

	getMagnetEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		magnets []string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		args       args
		want       InstantAvailabilityResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name: "no magnet",
			jsonResp: `{
				"status": "error",
				"error": {
				  "code": "MAGNET_NO_URI",
				  "message": "No magnet sent"
				}
			}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnets: []string{""},
			},
			assertion: assert.Error,
		},
		{
			name:       "bad json",
			jsonResp:   `{"status:"success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnets: []string{""},
			},
			assertion: assert.Error,
		},
		{
			name:       "success",
			jsonResp:   `{"status": "success"}`,
			statusResp: http.StatusOK,
			c:          cl,
			args: args{
				magnets: []string{"magnet:?xt=urn:btih:f95c371d5609d15f6615139be84edbb5b94a79bc"},
			},
			want: InstantAvailabilityResponse{
				Status: "success",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("POST", fmt.Sprintf(magnetinstant, getMagnetEndpoint(), tt.c.ic.appName, tt.c.ic.apikey),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.InstantAvailability(tt.args.magnets)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
