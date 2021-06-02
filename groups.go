package leihs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// AddGroup adsa group

func (l *Leihs) AddGroup(g *Group) (err error) {
	groupStr, err := json.Marshal(g)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", l.url+"admin/groups/", bytes.NewBuffer(groupStr))
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth("Token", l.token)

	resp, err := l.client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return
}

// GroupByName ...
func (l *Leihs) GroupByName(name string) (*Group, error) {
	g, err := l.FindGroups()
	if err != nil {
		return nil, err
	}
	for _, v := range *g {
		if v.Name == name {
			return &v, nil
		}
	}
	return nil, errors.New("Group not found")
}

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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
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
	Name        string      `json:"name"`
	OrgID       interface{} `json:"org_id"`
	ID          string      `json:"id"`
	CountUsers  int         `json:"count_users"`
	Description string      `json:"description"`
}
