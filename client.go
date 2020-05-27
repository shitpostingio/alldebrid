package alldebrid

import "errors"

// Client is the base AllDebrid client
type Client struct {
	ic *innerClient
}

type innerClient struct {
	apikey  string
	appName string
}

// New returns a new alldebrid client and error if one ore more vars are missing
func New(key, appname string) (Client, error) {
	if key == "" {
		return Client{}, errors.New("Missing APIKEY")
	}

	if appname == "" {
		return Client{}, errors.New("Missing appname")
	}

	return Client{
		ic: &innerClient{
			appName: appname,
			apikey:  key,
		},
	}, nil
}
