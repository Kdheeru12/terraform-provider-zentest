package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
}

type members struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         User   `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}
type Team struct {
	Unique_Id     string    `json:"unique_id"`
	Name          string    `json:"name"`
	Account       string    `json:"account"`
	Creation_Date string    `json:"creation_date"`
	Owner         string    `json:"owner"`
	Roles         []Roles   `json:"roles"`
	Members       []members `json:"members"`
}

func (c *Client) CreateTeam(team *Team) (*Team, error) {
	j, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ddd")
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var t Team
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *Client) UpdateTeam(id string, team *Team) (*Team, error) {
	j, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ddd")
	req, err := http.NewRequest("PUT", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+id+"/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var t Team
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *Client) GetTeamById(uniqie_id string) (*Team, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+uniqie_id, nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var t Team
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *Client) GetTeams() ([]Team, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var t []Team
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) DeleteTeam(id string) error {
	req, err := http.NewRequest("DELETE", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+id+"/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
