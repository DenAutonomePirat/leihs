package leihs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// AddUser ...
func (l *Leihs) AddUser(u *User) (user *User, err error) {
	if !isEmailValid(u.Email) {
		return nil, errors.New("Invalid email")
	}
	//make payload
	userStr, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	//prep request
	req, err := http.NewRequest("POST", l.url+"admin/users/", bytes.NewBuffer(userStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("Token", l.token)

	//do
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

	user = &User{}

	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindUser ...
func (l *Leihs) FindUser(term string) (user *User, err error) {

	req, err := http.NewRequest("GET", l.url+"admin/users?term="+term, nil)
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

	result := &Users{}

	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	for i, v := range result.Users {
		if strings.Compare(v.Email, term) == 0 {
			return &result.Users[i], nil
		}
	}

	return nil, errors.New("User not found")

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
	SystemAdminProtected  bool      `json:"system_admin_protected,omitempty"`
	Address               string    `json:"address,omitempty"`
	Email                 string    `json:"email,omitempty"`
	PoolProtected         bool      `json:"pool_protected,omitempty"`
	LastSignInAt          string    `json:"last_sign_in_at,omitempty"`
	Img32URL              string    `json:"img32_url,omitempty"`
	AccountEnabled        bool      `json:"account_enabled,omitempty"`
	Lastname              string    `json:"lastname,omitempty"`
	Phone                 string    `json:"phone,omitempty"`
	ImgDigest             string    `json:"img_digest,omitempty"`
	OrgID                 string    `json:"org_id,omitempty"`
	ExtendedInfo          string    `json:"extended_info,omitempty"`
	SecondaryEmail        string    `json:"secondary_email,omitempty"`
	City                  string    `json:"city,omitempty"`
	Settings              string    `json:"settings,omitempty"`
	IsAdmin               bool      `json:"is_admin,omitempty"`
	Organization          string    `json:"organization,omitempty"`
	Login                 string    `json:"login,omitempty"`
	Searchable            string    `json:"searchable,omitempty"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
	Firstname             string    `json:"firstname,omitempty"`
	Zip                   string    `json:"zip,omitempty"`
	ID                    string    `json:"id,omitempty"`
	URL                   string    `json:"url,omitempty"`
	PasswordSignInEnabled bool      `json:"password_sign_in_enabled,omitempty"`
	AccountDisabledAt     string    `json:"account_disabled_at,omitempty"`
	IsSystemAdmin         bool      `json:"is_system_admin,omitempty"`
	BadgeID               string    `json:"badge_id,omitempty"`
	LanguageLocale        string    `json:"language_locale,omitempty"`
	Img256URL             string    `json:"img256_url,omitempty"`
	Country               string    `json:"country,omitempty"`
	DelegatorUserID       string    `json:"delegator_user_id,omitempty"`
	CreatedAt             time.Time `json:"created_at,omitempty"`
	AdminProtected        bool      `json:"admin_protected,omitempty"`
}
