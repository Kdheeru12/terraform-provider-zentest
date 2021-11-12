package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Targets struct {
	Target_type int    `json:"target_type"`
	Target_id   string `json:"target_id"`
}
type Rules struct {
	Delay     int       `json:"delay"`
	Targets   []Targets `json:"targets"`
	Position  int       `json:"position"`
	Unique_Id string    `json:"unique_id"`
}

type EscalationPolicy struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Summary       string  `json:"summary"`
	Team          string  `json:"team"`
	Unique_Id     string  `json:"unique_id"`
	Repeat_Policy int     `json:"repeat_policy"`
	Move_To_Next  bool    `json:"move_to_next"`
	Global_Ep     bool    `json:"global_ep"`
	Rules         []Rules `json:"rules"`
}

func (c *Client) CreateEscalationPolicy(team string, policy *EscalationPolicy) (*EscalationPolicy, error) {
	j, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/escalation_policies/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s EscalationPolicy
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) GetEscalationPolicy(team string) ([]EscalationPolicy, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/escalation_policies/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s []EscalationPolicy
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Client) GetEscalationPolicyById(team, id string) (*EscalationPolicy, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/escalation_policies/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s EscalationPolicy
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) DeleteEscalationPolicy(team, id string) error {
	req, err := http.NewRequest("DELETE", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/escalation_policies/"+id+"/", nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateEscalationPolicy(team, id string, policy *EscalationPolicy) (*EscalationPolicy, error) {
	j, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/escalation_policies/"+id+"/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s EscalationPolicy
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
