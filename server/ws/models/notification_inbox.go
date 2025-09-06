package models

type NotificationInbox struct {
	Type string `json:"type"`
	State *int `json:"state"`
	Data *DataSender `json:"data_sender"`
}

type DataSender struct {
	Name string `json:"name"`
	Profile string `json:"profile"`
}