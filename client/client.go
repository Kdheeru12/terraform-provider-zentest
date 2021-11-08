package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	client := Client{Token: "3b44da5b6cc076b459c45a6256b2e0e8b03af91c"}
// 	task, err := client.GetTeams()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%+v\n", task)

// }

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
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
type user struct {
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
}

// type OnCall struct {
// 	EscalationPolicy EscalationPolicy `json:"escalation_policy"`
// 	Team             Team             `json:"team"`
// 	Users            []User           `json:"users"`
// }

type members struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         user   `json:"user"`
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

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

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

type Incident struct {
	Service          string `json:"service"`
	EscalationPolicy string `json:"escalation_policy"`
	User             string `json:"user"`
	Title            string `json:"title"`
	Summary          string `json:"summary"`
}

type service_object struct {
	Name                   string `json:"name"`
	Creation_Date          string `json:"creation_date"`
	Summary                string `json:"summary"`
	Description            string `json:"description"`
	Unique_Id              string `json:"unique_id"`
	Auto_Resolve_Timeouts  int    `json:"auto_resolve_timeout"`
	Created_By             string `json:"created_by"`
	Team_Priority          string `json:"team_priority"`
	Task_Template          string `json:"task_template"`
	Acknowledgment_Timeout int    `json:"acknowledge_timeout"`
	Status                 int    `json:"status"`
	EscalationPolicy       string `json:"escalation_policy"`
	Team                   string `json:"team"`
	Sla                    string `json:"sla"`
	Collation_Time         int    `json:"collation_time"`
	Collation              int    `json:"collation"`
}

type escalation_policy_object struct {
	Unique_Id string `json:"unique_id"`
	Name      string `json:"name"`
}

type Incidents struct {
	Summary                  string `json:"summary"`
	Incident_Number          int    `json:"incident_number"`
	Creation_Date            string `json:"creation_date"`
	Status                   int    `json:"status"`
	Unique_Id                string `json:"unique_id"`
	Service_Object           service_object
	Title                    string                   `json:"title"`
	Incident_Key             string                   `json:"incident_key"`
	Service                  string                   `json:"service"`
	Urgency                  int                      `json:"urgency"`
	Merged_With              string                   `json:"merged_with"`
	Assigned_To              string                   `json:"assigned_to"`
	Escalation_Policy        string                   `json:"escalation_policy"`
	Escalation_Policy_Object escalation_policy_object `json:"escalation_policy_object"`
	Assigned_to_name         string                   `json:"assigned_to_name"`
	Resolved_Date            string                   `json:"resolved_date"`
	Acknowledged_Date        string                   `json:"acknowledged_date"`
	Context_Window_start     string                   `json:"context_window_start"`
	Context_Window_end       string                   `json:"context_window_end"`
	Tags                     []string                 `json:"tags"`
	Sla                      string                   `json:"sla"`
	Team_Priority            string                   `json:"team_priority"`
	Team_Priority_Object     string                   `json:"team_priority_object"`
}

type IncidentPagination struct {
	Results  []Incidents `json:"results"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Count    int         `json:"count"`
}

func (c *Client) CreateIncident(incident *Incident) (*Incidents, error) {
	j, err := json.Marshal(incident)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/incidents/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Incidents
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Client) GetIncidents() (*IncidentPagination, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/incidents/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i IncidentPagination
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (c *Client) GetIncidentByNumber(id string) (*Incidents, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/incidents/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var i Incidents
	err = json.Unmarshal(body, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

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

func (c *Client) GetServices(team string) (*[]Services, error) {
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
	return &i, nil
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
