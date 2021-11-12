package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ApplicationReference struct {
	Name                string `json:"name"`
	Icon_Url            string `json:"icon_url"`
	Summary             string `json:"summary"`
	Description         string `json:"description"`
	Unique_Id           string `json:"unique_id"`
	Avalability_Plan_id int    `json:"availability_plan_id"`
	Setup_Instructions  string `json:"setup_instructions"`
	Extension           string `json:"extension"`
	Application_Type    int    `json:"application_type"`
	Categories          string `json:"categories"`
	Documentation_Link  string `json:"documentation_link"`
}
type Integration struct {
	Name                  string               `json:"name"`
	Creation_Date         string               `json:"creation_date"`
	Summary               string               `json:"summary"`
	Description           string               `json:"description"`
	Unique_Id             string               `json:"unique_id"`
	Service               string               `json:"service"`
	Application           string               `json:"application"`
	Application_Reference ApplicationReference `json:"application_reference"`
	Integration_key       string               `json:"integration_key"`
	Created_By            string               `json:"created_by"`
	Is_Enabled            bool                 `json:"is_enabled"`
	Create_Incident_For   int                  `json:"create_incident_for"`
	Integration_Type      int                  `json:"integration_type"`
	Default_Urgency       int                  `json:"default_urggency"`
}

func (c *Client) CreateIntegration(team string, service_id string, integration *Integration) (*Integration, error) {
	j, err := json.Marshal(integration)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/"+service_id+"/integrations/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Integration
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Client) GetIntegrations(team, service_id string) ([]Integration, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/"+service_id+"/integrations/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i []Integration
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (c *Client) GetIntegrationByID(team, service_id, id string) (*Integration, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/services/"+service_id+"/integrations/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Integration
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
