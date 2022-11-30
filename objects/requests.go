package objects

import (
	"encoding/json"
	"net/http"
)

const MaxListLimit = 200

type GetRequest struct {
	Id string `json:"id"`
}

type ListRequest struct {
	Limit int    `json:"limit"`
	After string `json:"after"`
	Name  string `json:"name"`
}

type CreateRequest struct {
	Event *Event `json:"event"`
}

type UpdateDetailsRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CancelRequest struct {
	Id string `json:"id"`
}

type RescheduleRequest struct {
	Id      string    `json:"id"`
	NewSlot *TimeSlot `json:"new_slot"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type EventResponseWrapper struct {
	Event  *Event   `json:"event,omitempty"`
	Events []*Event `json:"events,omitempty"`
	Code   int      `json:"-"`
}

func (e *EventResponseWrapper) Json() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

func (e *EventResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}
	return e.Code
}
