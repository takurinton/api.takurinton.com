package model

import "time"

type RequestConditionEventsType struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	APIAppID string `json:"api_app_id"`
	Event    struct {
		Type        string `json:"type"`
		Channel     string `json:"channel"`
		User        string `json:"user"`
		Text        string `json:"text"`
		Ts          string `json:"ts"`
		EventTs     string `json:"event_ts"`
		ChannelType string `json:"channel_type"`
	} `json:"event"`
	Type        string   `json:"type"`
	AuthedTeams []string `json:"authed_teams"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
}

type Condition struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Mental    string    `json:"mental"`
	Physical  string    `json:"physical"`
	CreatedAt time.Time `json:"created_at"`
}
