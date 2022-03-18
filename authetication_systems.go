package leihs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (l *Leihs) AddToAuthenticationSystem(g *Group, as *AuthenticationSystem) (err error) {

	req, err := http.NewRequest("PUT", l.url+"/admin/system/authentication-systems/"+as.ID+"/groups/"+g.ID, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth("Token", l.token)
	resp, err := l.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// AddAuthenticationSystem ...
func (l *Leihs) AddAuthenticationSystem(a *AuthenticationSystem) (err error) {
	payload, err := json.Marshal(a)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", payload)
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

	// read response body
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	resp.Body.Close()

	// print response body
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}

// FindAuthenticationSystems ...
func (l *Leihs) FindAuthenticationSystems() (a *[]AuthenticationSystem, err error) {

	req, err := http.NewRequest("GET", l.url+"admin/system/authentication-systems/", nil)
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
func (l *Leihs) AuthenticationSystemByID(id string) (*AuthenticationSystem, error) {

	req, err := http.NewRequest("GET", l.url+"admin/system/authentication-systems/"+id, nil)
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
	a := &AuthenticationSystem{}

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
	return nil, errors.New("Authentication system not found")
}

//AuthenticationSystems ...
type AuthenticationSystems struct {
	AuthenticationSystems []AuthenticationSystem `json:"authentication-systems"`
}

// AuthenticationSystem ...
type AuthenticationSystem struct {
	Name                  string    `json:"name,omitempty" yaml:"name"`
	Priority              int       `json:"priority,omitempty" yaml:"priority"`
	Enabled               bool      `json:"enabled,omitempty" yaml:"enabled"`
	Description           string    `json:"description,omitempty" yaml:"description"`
	ExternalSignInURL     string    `json:"external_sign_in_url,omitempty" yaml:"external_sign_in_url"`
	ExternalSignOutURL    string    `json:"external_sign_out_url,omitempty" yaml:"external_sign_out_url"`
	SignUpEmailMatch      string    `json:"sign_up_email_match,omitempty" yaml:"sign_up_email_match"`
	ShortcutSignInEnabled bool      `json:"shortcut_sign_in_enabled,omitempty" yaml:"shortcut_sign_in_enabled"`
	SendOrgID             bool      `json:"send_org_id,omitempty" yaml:"send_org_id"`
	SendEmail             bool      `json:"send_email,omitempty" yaml:"send_email"`
	SendLogin             bool      `json:"send_login,omitempty" yaml:"send_login"`
	Type                  string    `json:"type,omitempty" yaml:"type"`
	ID                    string    `json:"id,omitempty" yaml:"id"`
	InternalPrivateKey    string    `json:"internal_private_key,omitempty" yaml:"-"`
	InternalPublicKey     string    `json:"internal_public_key,omitempty" yaml:"-"`
	ExternalPublicKey     string    `json:"external_public_key,omitempty" yaml:"-"`
	CreatedAt             time.Time `json:"created_at,omitempty" yaml:"-"`
	UsersCount            int       `json:"users_count,omitempty" yaml:"-"`
	GroupsCount           int       `json:"groups_count,omitempty" yaml:"-"`
	UpdatedAt             time.Time `json:"updated_at,omitempty" yaml:"-"`
}
