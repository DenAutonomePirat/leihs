package leihs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AddGroup ..
func (l *Leihs) AddGroup() {}

// AddToGroup ...
func (l *Leihs) AddToGroup(u *User, g *Group) (err error) {

	req, err := http.NewRequest("PUT", l.url+"/admin/groups/"+g.ID+"/users/"+u.ID, nil)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("resp.Body %s", string(body))
	return nil
}

// FindGroups ...
func (l *Leihs) FindGroups() (g *[]Group, err error) {
	req, err := http.NewRequest("GET", l.url+"/admin/groups/", nil)
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

	groups := &Groups{}

	err = json.Unmarshal(body, groups)
	if err != nil {
		return nil, err
	}
	g = &groups.Groups
	return g, nil
}

//Groups ...
type Groups struct {
	Groups []Group `json:"groups"`
}

// Group ...
type Group struct {
	Name       string      `json:"name"`
	OrgID      interface{} `json:"org_id"`
	ID         string      `json:"id"`
	CountUsers int         `json:"count_users"`
}
