package leihs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// AddAuthenticationSystem ...
func (l *Leihs) AddAuthenticationSystem(a *AuthenticationSystem) (err error) {
	payload, err := json.Marshal(a)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", l.url+"/admin/system/authentication-systems/", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth("Token", l.token)

	resp, err := l.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}

// FindAuthenticationSystems ...
func (l *Leihs) FindAuthenticationSystems() (a *[]AuthenticationSystem, err error) {

	req, err := http.NewRequest("GET", l.url+"/admin/system/authentication-systems/", nil)
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

	as := &AuthenticationSystems{}

	err = json.Unmarshal(body, as)
	if err != nil {
		return nil, err
	}

	return &as.AuthenticationSystems, nil
}

// AuthenticationSystemByID ...
func (l *Leihs) AuthenticationSystemByID(id string) (a *AuthenticationSystem, err error) {

	req, err := http.NewRequest("GET", l.url+"/admin/system/authentication-systems/"+id, nil)
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

	err = json.Unmarshal(body, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// AuthenticationSystemByName ...
func (l *Leihs) AuthenticationSystemByName(name string) (*AuthenticationSystem, error) {
	as, err := l.FindAuthenticationSystems()
	if err != nil {
		return nil, err
	}
	for _, v := range *as {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, errors.New("Authentication sysytem not found")
}

//AuthenticationSystems ...
type AuthenticationSystems struct {
	AuthenticationSystems []AuthenticationSystem `json:"authentication-systems"`
}

// AuthenticationSystem ...
type AuthenticationSystem struct {
	Description           string    `json:"description,omitempty" yaml:"description"`
	ExternalSignInURL     string    `json:"external_sign_in_url,omitempty" yaml:"external_sign_in_url"`
	InternalPrivateKey    string    `json:"internal_private_key,omitempty" yaml:"internal_private_key"`
	SendOrgID             bool      `json:"send_org_id,omitempty" yaml:"send_org_id"`
	GroupsCount           int       `json:"groups_count,omitempty" yaml:"groups_count"`
	ShortcutSignInEnabled bool      `json:"shortcut_sign_in_enabled,omitempty" yaml:"shortcut_sign_in_enabled"`
	UsersCount            int       `json:"users_count,omitempty" yaml:"users_count"`
	SendEmail             bool      `json:"send_email,omitempty" yaml:"send_email"`
	ExternalPublicKey     string    `json:"external_public_key,omitempty" yaml:"external_public_key"`
	Name                  string    `json:"name,omitempty" yaml:"name"`
	SendLogin             bool      `json:"send_login,omitempty" yaml:"send_login"`
	Type                  string    `json:"type,omitempty" yaml:"type"`
	SignUpEmailMatch      string    `json:"sign_up_email_match,omitempty" yaml:"sign_up_email_match"`
	InternalPublicKey     string    `json:"internal_public_key,omitempty" yaml:"internal_public_key"`
	UpdatedAt             time.Time `json:"updated_at,omitempty" yaml:"-"`
	Priority              int       `json:"priority,omitempty" yaml:"priority"`
	ID                    string    `json:"id,omitempty" yaml:"id"`
	ExternalSignOutURL    string    `json:"external_sign_out_url,omitempty" yaml:"external_sign_out_url"`
	Enabled               bool      `json:"enabled,omitempty" yaml:"enabled"`
	CreatedAt             time.Time `json:"created_at,omitempty" yaml:"-"`
}
