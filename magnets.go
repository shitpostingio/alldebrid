package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

//Error is the error struct
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//MagnetsUploadResponse is the response of the upload call
type MagnetsUploadResponse struct {
	Status string `json:"status"`
	Data   struct {
		Magnets []struct {
			Magnet string `json:"magnet"`
			Hash   string `json:"hash,omitempty"`
			Name   string `json:"name,omitempty"`
			Size   int    `json:"size,omitempty"`
			Ready  bool   `json:"ready,omitempty"`
			ID     int    `json:"id,omitempty"`
			Error  struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"error,omitempty"`
		} `json:"magnets"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

//StatusMagnetResponse is the response of the status call
type StatusMagnetResponse struct {
	Status string `json:"status"`
	Data   struct {
		Magnets []struct {
			ID            int    `json:"id"`
			Filename      string `json:"filename"`
			Size          int    `json:"size"`
			Status        string `json:"status"`
			StatusCode    int    `json:"statusCode"`
			Downloaded    int    `json:"downloaded"`
			Uploaded      int    `json:"uploaded"`
			Seeders       int    `json:"seeders"`
			DownloadSpeed int    `json:"downloadSpeed"`
			UploadSpeed   int    `json:"uploadSpeed"`
			UploadDate    int    `json:"uploadDate"`
			Links         []struct {
				Link     string      `json:"link"`
				Filename string      `json:"filename"`
				Size     int         `json:"size"`
				Files    interface{} `json:"files"`
			} `json:"links"`
		} `json:"magnets"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

//DeleteMagnetResponse is the response of the delete call
type DeleteMagnetResponse struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

//RestartMagnetResponse is the response of the restart call
type RestartMagnetResponse struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

// UploadMagnet sends magnet(s) to AllDebrid
func (c *Client) UploadMagnet(magnets []string) (MagnetsUploadResponse, error) {
	client := &http.Client{}

	ms := url.Values{}
	for _, magnet := range magnets {
		ms.Add("magnets[]", magnet)
	}

	resp, err := client.PostForm(fmt.Sprintf(upload, magnetURL, c.AppName, c.APIKEY), ms)

	if err != nil {
		return MagnetsUploadResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var uploadResponse MagnetsUploadResponse

	err = decoder.Decode(&uploadResponse)
	if err != nil {
		return MagnetsUploadResponse{}, err
	}

	if uploadResponse.Status != "success" {
		return MagnetsUploadResponse{}, errors.New(uploadResponse.Error.Message)
	}

	return uploadResponse, nil
}

//StatusMagnet returns the status of an Alldebrid download
func (c *Client) StatusMagnet(id string) (StatusMagnetResponse, error) {
	resp, err := http.Get(fmt.Sprintf(status, magnetURL, c.AppName, c.APIKEY, id))

	if err != nil {
		return StatusMagnetResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var statusResponse StatusMagnetResponse

	err = decoder.Decode(&statusResponse)
	if err != nil {
		return StatusMagnetResponse{}, err
	}

	if statusResponse.Status != "success" {
		return StatusMagnetResponse{}, errors.New(statusResponse.Error.Message)
	}

	return statusResponse, nil
}

//DeleteMagnet removes a download from alldebrid
func (c *Client) DeleteMagnet(id string) (DeleteMagnetResponse, error) {
	resp, err := http.Get(fmt.Sprintf(delete, magnetURL, c.AppName, c.APIKEY, id))

	if err != nil {
		return DeleteMagnetResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var deleteResponse DeleteMagnetResponse

	err = decoder.Decode(&deleteResponse)
	if err != nil {
		return DeleteMagnetResponse{}, err
	}

	if deleteResponse.Status != "success" {
		return DeleteMagnetResponse{}, errors.New(deleteResponse.Error.Message)
	}

	return deleteResponse, nil
}

//RestartMagnet will restart a failed torrent
func (c *Client) RestartMagnet(id string) (RestartMagnetResponse, error) {
	resp, err := http.Get(fmt.Sprintf(restart, magnetURL, c.AppName, c.APIKEY, id))

	if err != nil {
		return RestartMagnetResponse{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var restartResponse RestartMagnetResponse

	err = decoder.Decode(&restartResponse)
	if err != nil {
		return RestartMagnetResponse{}, err
	}

	if restartResponse.Status != "success" {
		return RestartMagnetResponse{}, errors.New(restartResponse.Error.Message)
	}

	return restartResponse, nil
}
