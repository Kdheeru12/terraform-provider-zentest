package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Services struct {
	Name                   string `json:"name"`
	Creation_Date          string `json:"creation_date"`
	Summary                string `json:"summary"`
	Description            string `json:"description"`
	Unique_Id              string `json:"unique_id"`
	Auto_Resolve_Timeout   int    `json:"auto_resolve_timeout"`
	Created_By             string `json:"created_by"`
	Team_Priority          string `json:"team_priority"`
	Task_Template          string `json:"task_template"`
	Acknowledgment_Timeout int    `json:"acknowledge_timeout"`
	Status                 int    `json:"status"`
	Escalation_Policy      string `json:"escalation_policy"`
	Team                   string `json:"team"`
	Sla                    string `json:"sla"`
	Collation_Time         int    `json:"collation_time"`
	Collation              int    `json:"collation"`
	Under_Maintenance      bool   `json:"under_maintenance"`
}

func (c *Client) CreateService(team string, service *Services) (*Services, error) {
	j, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Services
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Client) GetServices(team string) ([]Services, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i []Services
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (c *Client) GetServicesById(team, id string) (*Services, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Services
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Client) UpdateService(team, id string, service *Services) (*Services, error) {
	j, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/"+id+"/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Services
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Client) DeleteService(team string, id string) error {
	req, err := http.NewRequest("DELETE", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/"+id+"/", nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
