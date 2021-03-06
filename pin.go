package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//GetPinResponse contains data about the get pin request
type GetPinResponse struct {
	Status string         `json:"status"`
	Data   getPinData     `json:"data,omitempty"`
	Error  alldebridError `json:"error,omitempty"`
}

type getPinData struct {
	Pin       string `json:"pin"`
	Check     string `json:"check"`
	ExpiresIn int    `json:"expires_in"`
	UserURL   string `json:"user_url"`
	BaseURL   string `json:"base_url"`
	CheckURL  string `json:"check_url"`
}

//CheckPinResponse contains data about the check pin request
type CheckPinResponse struct {
	Status string               `json:"status"`
	Data   checkPinResponseData `json:"data,omitempty"`
	Error  alldebridError       `json:"error,omitempty"`
}

type checkPinResponseData struct {
	Apikey    string `json:"apikey,omitempty"`
	Activated bool   `json:"activated"`
	ExpiresIn int    `json:"expires_in"`
}

//GetPin asks Alldebrid for a new pin
func (c *Client) GetPin() (GetPinResponse, error) {
	resp, err := http.Get(fmt.Sprintf(pinget, getPinEndpoint(), c.ic.appName))
	if err != nil {
		return GetPinResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var pin GetPinResponse

	err = decoder.Decode(&pin)
	if err != nil {
		return GetPinResponse{}, err
	}

	if pin.Status != "success" {
		return GetPinResponse{}, errors.New(pin.Error.Message)
	}

	return pin, nil
}

//CheckPin gives you an apikey after pin validating
func (c *Client) CheckPin(check, authpin string) (CheckPinResponse, error) {
	resp, err := http.Get(fmt.Sprintf(pincheck, getPinEndpoint(), c.ic.appName, check, authpin))
	if err != nil {
		return CheckPinResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var pin CheckPinResponse

	err = decoder.Decode(&pin)
	if err != nil {
		return CheckPinResponse{}, err
	}

	if pin.Status != "success" {
		return CheckPinResponse{}, errors.New(pin.Error.Message)
	}

	return pin, nil
}
