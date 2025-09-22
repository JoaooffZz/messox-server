package models

type Notification struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

type EventUpdateContactRequest struct {
	State int `json:"state"`
	InfoStates map[int]string `json:"info-states"`
}

type EventNewContactRequest struct {
	Data DataContact `json:"data-contact"`
}
type DataContact struct {
	Name string `json:"name"`
	ProfileBytes []byte `json:"profile-bytes"`
}
