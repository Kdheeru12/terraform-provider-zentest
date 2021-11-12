package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// func main() {
// 	client := Client{Token: "3b44da5b6cc076b459c45a6256b2e0e8b03af91c"}
// 	task, err := client.GetEscalationPolicy("dd518f4d-dbce-4ad2-b5be-ceff597c67f8")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%+v\n", task)

// }
type EmailAccounts struct {
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Role       int    `json:"role"`
}

type Invite struct {
	EmailAccounts []EmailAccounts `json:"email_accounts"`
	Team          string          `json:"team"`
}

type InviteResponse struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         User   `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}

type Member struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         string `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}

type MemberResponse struct {
	Unique_Id    string `json:"unique_id"`
	Team         string `json:"team"`
	User         User   `json:"user"`
	Joining_Date string `json:"joining_date"`
	Role         int    `json:"role"`
}

func (c *Client) CreateInvite(invite *Invite) ([]InviteResponse, error) {
	j, err := json.Marshal(invite)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/invite/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s []InviteResponse
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/dd518f4d-dbce-4ad2-b5be-ceff597c67f8/members/3f16016e-7d53-4153-bda0-1c6415fc5ff0/

func (c *Client) CreateTeamMember(team string, member *Member) (*Member, error) {
	j, err := json.Marshal(member)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/members/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s Member
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) UpdateTeamMember(member *Member) (*Member, error) {
	j, err := json.Marshal(member)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+member.Team+"/members/"+member.Unique_Id+"/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s Member
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) DeleteTeamMember(team string, member string) error {
	req, err := http.NewRequest("DELETE", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/members/"+member+"/", nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetTeamMembers(team string) ([]MemberResponse, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/members/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s []MemberResponse
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Client) GetTeamMembersByID(team, id string) (*MemberResponse, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/members/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s MemberResponse
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
