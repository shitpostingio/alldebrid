package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//User is the user info struct
type User struct {
	Status string `json:"status"`
	Data   struct {
		User struct {
			Username             string `json:"username"`
			Email                string `json:"email"`
			IsPremium            bool   `json:"isPremium"`
			IsTrial              bool   `json:"isTrial"`
			PremiumUntil         int    `json:"premiumUntil"`
			Lang                 string `json:"lang"`
			PreferedDomain       string `json:"preferedDomain"`
			FidelityPoints       int    `json:"fidelityPoints"`
			LimitedHostersQuotas struct {
				Filefactory int `json:"filefactory"`
				Gigapeta    int `json:"gigapeta"`
				Videobin    int `json:"videobin"`
				Isra        int `json:"isra"`
				Rapidgator  int `json:"rapidgator"`
				Rapidu      int `json:"rapidu"`
				Brupload    int `json:"brupload"`
				Uploadcloud int `json:"uploadcloud"`
				Userscloud  int `json:"userscloud"`
				Wipfiles    int `json:"wipfiles"`
				Ddl         int `json:"ddl"`
				Flashbit    int `json:"flashbit"`
				Anzfile     int `json:"anzfile"`
			} `json:"limitedHostersQuotas"`
		} `json:"user"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

//GetUserInfo retrieves user infos from alldebrid
func (c *Client) GetUserInfo() (User, error) {
	resp, err := http.Get(fmt.Sprintf(userinfo, user, c.AppName, c.APIKEY))
	if err != nil {
		return User{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var usInfo User

	err = decoder.Decode(&usInfo)
	if err != nil {
		return User{}, err
	}

	if usInfo.Status != "success" {
		return User{}, errors.New(usInfo.Error.Message)
	}

	return usInfo, nil
}
