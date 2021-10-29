package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	Token      string
}

func NewClient(token string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		Token:      token,
	}
}

// type EscalationPolicy struct {
// 	Name          string `json:"name"`
// 	Summary       string `json:"summary"`
// 	Description   string `json:"description"`
// 	Uniqie_Id     string `json:"unique_id"`
// 	Repeat_Policy int    `json:"repeat_policy"`
// 	Move_To_Next  bool   `json:"move_to_next"`
// 	Global_Ep     bool   `json:"global_ep"`
// }
// type Team struct {
// 	Uniqie_Id string `json:"unique_id"`
// 	Name      string `json:"name"`
// }
// type User struct {
// 	Username   string `json:"username"`
// 	First_Name string `json:"first_name"`
// 	Last_Name  string `json:"last_name"`
// }

// type OnCall struct {
// 	EscalationPolicy EscalationPolicy `json:"escalation_policy"`
// 	Team             Team             `json:"team"`
// 	Users            []User           `json:"users"`
// }

type Team struct {
	Unique_Id     string `json:"unique_id"`
	Name          string `json:"name"`
	Account       string `json:"account"`
	Creation_Date string `json:"creation_date"`
	Owner         string `json:"owner"`
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "token "+c.Token)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK || resp.StatusCode != http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}
}

func (c *Client) CreateTeam(team *Team) (*Team, error) {
	j, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://www.zenduty.com/api/account/teams/", bytes.NewBuffer(j))
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

func (c *Client) GetTeam(uniqie_id string) (*Team, error) {
	req, err := http.NewRequest("GET", "https://www.zenduty.com/api/account/teams/"+uniqie_id, nil)
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
