package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Restrictions struct {
	Duration       int    `json:"duration"`
	StartDayOfWeek int    `json:"start_day_of_week"`
	StartTimeOfDay string `json:"start_time_of_day"`
	Unique_Id      string `json:"unique_id"`
}
type Users struct {
	User      string `json:"user"`
	Position  int    `json:"position"`
	Unique_Id string `json:"unique_id"`
}

type Overrides struct {
	Name      string `json:"name"`
	User      string `json:"user"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Unique_Id string `json:"unique_id"`
}
type Layers struct {
	ShiftLength       int            `json:"shift_length"`
	Name              string         `json:"name"`
	RotationStartTime string         `json:"rotation_start_time"`
	RotationEndTime   string         `json:"rotation_end_time"`
	UniqueId          string         `json:"unique_id"`
	LastEdited        string         `json:"last_edited"`
	RestrictionType   int            `json:"restriction_type"`
	IsActive          bool           `json:"is_active"`
	Restrictions      []Restrictions `json:"restrictions"`
	Users             []Users        `json:"users"`
}

type Schedules struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Summary     string      `json:"summary"`
	Time_zone   string      `json:"time_zone"`
	Team        string      `json:"team"`
	Unique_Id   string      `json:"unique_id"`
	Layers      []Layers    `json:"layers"`
	Overrides   []Overrides `json:"overrides"`
}

func (c *Client) CreateSchedule(team string, schedule *Schedules) (*Schedules, error) {
	j, err := json.Marshal(schedule)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/schedules/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s Schedules
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) GetSchedules(team string) ([]Schedules, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/schedules/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s []Schedules
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Client) GetScheduleByID(team, id string) (*Schedules, error) {
	req, err := http.NewRequest("GET", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/schedules/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s Schedules
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (c *Client) DeleteScheduleByID(team, id string) error {
	req, err := http.NewRequest("DELETE", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/schedules/"+id+"/", nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateScheduleByID(team, id string, schedule *Schedules) (*Schedules, error) {
	j, err := json.Marshal(schedule)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", "http://zenduty-beanstalk-stage-dev.us-east-1.elasticbeanstalk.com/api/account/teams/"+team+"/schedules/"+id+"/", bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var s Schedules
	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
