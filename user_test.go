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
			name: "no apikey provided",
			jsonResp: `{
				"status": "error",
				"error": {
				  "code": "AUTH_MISSING_APIKEY",
				  "message": "The auth apikey was not sent"
				}
			  }`,
			statusResp: http.StatusOK,
			c:          noapicl,
			want:       User{},
			assertion:  assert.Error,
		},
		{
			name: "bad json",
			jsonResp: `{
				"status": "error",
				"error": {
				  "code": "AUTH_BAD_APIKEY",
				  "message": "The auth apikey is invalid"
				},
			  }`,
			statusResp: http.StatusOK,
			c:          cl,
			want:       User{},
			assertion:  assert.Error,
		},
		{
			name: "valid request",
			jsonResp: `{
				"status": "success",
				"data": {
				  "user": {
					"username": "User",
					"email": "testing@isfun.com",
					"isPremium": true,
					"isTrial": false,
					"premiumUntil": 1619568038,
					"lang": "en",
					"preferedDomain": "com",
					"fidelityPoints": 100,
					"limitedHostersQuotas": {
					  "filefactory": 3000,
					  "gigapeta": 10000,
					  "videobin": 10000,
					  "isra": 3000,
					  "rapidgator": 50000,
					  "rapidu": 5000,
					  "brupload": 3000,
					  "uploadcloud": 2000,
					  "userscloud": 3000,
					  "wipfiles": 3000,
					  "ddl": 50000,
					  "flashbit": 5000,
					  "anzfile": 3000,
					  "keep2share": 5000
					}
				  }
				}
			  }`,
			statusResp: http.StatusOK,
			c:          cl,
			want: User{
				Status: "success",
				Data: userData{
					userDataUser{
						Username:       "User",
						Email:          "testing@isfun.com",
						IsPremium:      true,
						IsTrial:        false,
						PremiumUntil:   1619568038,
						Lang:           "en",
						PreferedDomain: "com",
						FidelityPoints: 100,
						LimitedHostersQuotas: userQuotas{
							Filefactory: 3000,
							Gigapeta:    10000,
							Videobin:    10000,
							Isra:        3000,
							Rapidgator:  50000,
							Rapidu:      5000,
							Brupload:    3000,
							Uploadcloud: 2000,
							Userscloud:  3000,
							Wipfiles:    3000,
							Ddl:         50000,
							Flashbit:    5000,
							Anzfile:     3000,
							Keep2Share:  5000,
						},
					},
				},
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
