package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	client := Client{Token: "0e2038520fca2fbd7f3d9aace062c4fe911be36b"}
// 	newteam := &Roles{}
// 	newteam.Title = "hello"
// 	newteam.Description = "ddddd"
// 	newteam.Unique_Id = ""
// 	task, err := client.UpdateRoles("57cc5f4f-a785-4e98-8bae-4ef1ca148af5", newteam)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(task)

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
