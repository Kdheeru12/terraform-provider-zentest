package client

import "net/http"

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

type EscalationPolicy struct {
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
	Uniqie_Id     string `json:"unique_id"`
	Repeat_Policy int    `json:"repeat_policy"`
	Move_To_Next  bool   `json:"move_to_next"`
	Global_Ep     bool   `json:"global_ep"`
}
type Team struct {
	Uniqie_Id string `json:"unique_id"`
	Name      string `json:"name"`
}
type User struct {
	Username   string `json:"username"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type OnCall struct {
	EscalationPolicy EscalationPolicy `json:"escalation_policy"`
	Team             Team             `json:"team"`
	Users            []User           `json:"users"`
}
