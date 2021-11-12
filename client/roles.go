package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Roles struct {
	Team          string `json:"team"`
	Unique_Id     string `json:"unique_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Creation_Date string `json:"creation_date"`
	Rank          int    `json:"rank"`
}

func (c *Client) CreateRole(team string, role *Roles) (*Roles, error) {
	j, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/roles/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var r Roles
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) GetRoles(team string) ([]Roles, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/roles/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var r []Roles
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) UpdateRoles(team string, role *Roles) (*Roles, error) {
	j, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/roles/"+role.Unique_Id+"/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var r Roles
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) DeleteRole(team string, role string) error {
	req, err := http.NewRequest("DELETE", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/roles/"+role+"/", nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
