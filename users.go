package leihs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// AddUser ...
func (l *Leihs) AddUser(u *User) (user *User, err error) {
	if !isEmailValid(u.Email) {
		return nil, errors.New("Invalid email")
	}
	userStr, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", l.url+"/admin/users/", bytes.NewBuffer(userStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("Token", l.token)

	resp, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUser ...
func (l *Leihs) FindUser(term string) (user *User, err error) {

	req, err := http.NewRequest("GET", l.url+"/admin/users?term="+url.QueryEscape(term), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("Token", l.token)

	resp, err := l.client.Do(req)
	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	users := &Users{}

	err = json.Unmarshal(body, users)
	if err != nil {
		return nil, err
	}
	req.URL, _ = url.Parse(l.url + "/admin/users/" + users.Users[0].ID)
	if err != nil {
		return nil, err
	}

	resp, err = l.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUsers ..
func (l *Leihs) FindUsers() (users *[]User, err error) {

	//todo iterate pages,
	req, err := http.NewRequest("GET", l.url+"/admin/users?per-page=1000", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("Token", l.token)

	resp, err := l.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	u := &Users{}

	err = json.Unmarshal(body, u)
	if err != nil {
		return nil, err
	}
	users = &u.Users
	return users, nil
}

//Users ...
type Users struct {
	Users []User `json:"users"`
}

//User ...
type User struct {
	Img32URL                string      `json:"img32_url,omitempty"`
	OrgID                   string      `json:"org_id,omitempty"`
	ID                      string      `json:"id,omitempty"`
	Login                   string      `json:"login,omitempty"`
	Firstname               string      `json:"firstname,omitempty"`
	Lastname                string      `json:"lastname,omitempty"`
	Phone                   string      `json:"phone,omitempty"`
	UniqueID                string      `json:"unique_id,omitempty"`
	Email                   string      `json:"email,omitempty"`
	BadgeID                 string      `json:"badge_id,omitempty"`
	Address                 string      `json:"address,omitempty"`
	City                    string      `json:"city,omitempty"`
	Zip                     string      `json:"zip,omitempty"`
	Country                 string      `json:"country,omitempty"`
	LanguageID              string      `json:"language_id,omitempty"`
	ExtendedInfo            string      `json:"extended_info,omitempty"`
	Settings                string      `json:"settings,omitempty"`
	DelegatorUserID         string      `json:"delegator_user_id,omitempty"`
	ContractsCount          int         `json:"contracts_count,omitempty"`
	AccountEnabled          bool        `json:"account_enabled,omitempty"`
	ImgDigest               interface{} `json:"img_digest,omitempty"`
	SecondaryEmail          interface{} `json:"secondary_email,omitempty"`
	IsAdmin                 bool        `json:"is_admin,omitempty"`
	UpdatedAt               time.Time   `json:"updated_at,omitempty"`
	URL                     interface{} `json:"url,omitempty"`
	PasswordSignInEnabled   bool        `json:"password_sign_in_enabled,omitempty"`
	Img256URL               interface{} `json:"img256_url,omitempty"`
	InventoryPoolRolesCount int         `json:"inventory_pool_roles_count,omitempty"`
	CreatedAt               time.Time   `json:"created_at,omitempty"`
}
