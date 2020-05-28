package alldebrid

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//User is the user info struct
type User struct {
	Status string         `json:"status"`
	Data   userData       `json:"data,omitempty"`
	Error  alldebridError `json:"error,omitempty"`
}

type userData struct {
	User userDataUser `json:"user"`
}

type userDataUser struct {
	Username             string     `json:"username"`
	Email                string     `json:"email"`
	IsPremium            bool       `json:"isPremium"`
	IsTrial              bool       `json:"isTrial"`
	PremiumUntil         int        `json:"premiumUntil"`
	Lang                 string     `json:"lang"`
	PreferedDomain       string     `json:"preferedDomain"`
	FidelityPoints       int        `json:"fidelityPoints"`
	LimitedHostersQuotas userQuotas `json:"limitedHostersQuotas"`
}

type userQuotas struct {
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
	Keep2Share  int `json:"keep2share"`
}

//GetUserInfo retrieves user infos from alldebrid
func (c *Client) GetUserInfo() (User, error) {
	resp, err := http.Get(fmt.Sprintf(userinfo, getUserEndpoint(), c.ic.appName, c.ic.apikey))
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
