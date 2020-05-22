package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	upload = "%s/upload?agent=%s&apikey=%s"
	status = "%s/status?agent=%s&apikey=%s&id=%s"
	delete = "%s/delete?agent=%s&apikey=%s&id=%s"
)

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
	} `json:"data"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type StatusMagnetResponse struct {
	Status string `json:"status"`
	Data   struct {
		Magnets []struct {
			ID            int           `json:"id"`
			Filename      string        `json:"filename"`
			Size          int           `json:"size"`
			Status        string        `json:"status"`
			StatusCode    int           `json:"statusCode"`
			Downloaded    int           `json:"downloaded"`
			Uploaded      int           `json:"uploaded"`
			Seeders       int           `json:"seeders"`
			DownloadSpeed int           `json:"downloadSpeed"`
			UploadSpeed   int           `json:"uploadSpeed"`
			UploadDate    int           `json:"uploadDate"`
			Links         []interface{} `json:"links"`
		} `json:"magnets"`
	} `json:"data"`
}

type DeleteMagnetResponse struct {
	Status string `json:"status"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data"`
}

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
	if uploadResponse.Status != "success" {
		return MagnetsUploadResponse{}, errors.New(uploadResponse.Error.Message)
	}

	return uploadResponse, err

}
