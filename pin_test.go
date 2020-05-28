package alldebrid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetPin(t *testing.T) {

	getPinEndpoint = func() string {
		return "http://127.0.0.1"
	}

	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		c          *Client
		want       GetPinResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name:       "valid request",
			jsonResp:   `{"status":"success","data":{"pin":"ABCD","check":"0fdfb8b5a548aff3d6cc1c62c32761e532d2c1f7","expires_in":600,"user_url":"https:\/\/alldebrid.com\/pin\/?pin=ABCD","base_url": "https:\/\/alldebrid.com\/pin\/","check_url":"https:\/\/api.alldebrid.com\/v4\/pin\/check?check=0fdfb8b5a548aff3d6cc1c62c32761e532d2c1f7&pin=ABCD&agent=test"}}`,
			statusResp: http.StatusOK,
			c:          cl,
			want: GetPinResponse{
				Status: "success",
				Data: getPinData{
					Pin:       "ABCD",
					Check:     "0fdfb8b5a548aff3d6cc1c62c32761e532d2c1f7",
					ExpiresIn: 600,
					UserURL:   "https://alldebrid.com/pin/?pin=ABCD",
					BaseURL:   "https://alldebrid.com/pin/",
					CheckURL:  "https://api.alldebrid.com/v4/pin/check?check=0fdfb8b5a548aff3d6cc1c62c32761e532d2c1f7&pin=ABCD&agent=test",
				},
			},
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(pinget, getPinEndpoint(), tt.c.ic.appName),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.GetPin()

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

}

func TestClient_CheckPin(t *testing.T) {
	getPinEndpoint = func() string {
		return "http://127.0.0.1"
	}

	type args struct {
		check string
		pin   string
	}
	tests := []struct {
		name       string
		jsonResp   string
		statusResp int
		a          args
		c          *Client
		want       CheckPinResponse
		assertion  assert.ErrorAssertionFunc
	}{
		{
			name: "no pin provided",
			jsonResp: `{
				"status": "error",
				"error": {
				  "code": "PIN_INVALID",
				  "message": "The pin is invalid"
				}
			}`,
			statusResp: http.StatusOK,
			a: args{
				"24c23d09687a4c63ff571570eb7aebb3913a3ffb",
				"",
			},
			c:         cl,
			want:      CheckPinResponse{},
			assertion: assert.Error,
		},
		{
			name: "bad json",
			jsonResp: `{
				"status": "error",
				"error": {
				  "code": "PIN_INVALID",
				  "message": "The pin is invalid"
				},
			}`,
			statusResp: http.StatusOK,
			a: args{
				"24c23d09687a4c63ff571570eb7aebb3913a3ffb",
				"",
			},
			c:         cl,
			want:      CheckPinResponse{},
			assertion: assert.Error,
		},
		{
			name: "no error",
			jsonResp: `{
				"status": "success",
				"data": {	
				  "apikey": "12345678abcdefg",
				  "activated": true,
				  "expires_in": 576
				}
			  }`,
			statusResp: http.StatusOK,
			a: args{
				"24c23d09687a4c63ff571570eb7aebb3913a3ffb",
				"ABCD",
			},
			c: cl,
			want: CheckPinResponse{
				Status: "success",
				Data: checkPinResponseData{
					Apikey:    "12345678abcdefg",
					Activated: true,
					ExpiresIn: 576,
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", fmt.Sprintf(pincheck, getPinEndpoint(), tt.c.ic.appName, tt.a.check, tt.a.pin),
				httpmock.NewStringResponder(tt.statusResp, tt.jsonResp))

			got, err := tt.c.CheckPin(tt.a.check, tt.a.pin)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
