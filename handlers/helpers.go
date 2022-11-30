package handlers

import (
	"encoding/json"
	"github.com/jithinlal-gelato/go_api/errors"
	"github.com/jithinlal-gelato/go_api/objects"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Response interface {
	Json() []byte
	StatusCode() int
}

func WriteResponse(w http.ResponseWriter, res Response) {
	w.WriteHeader(res.StatusCode())
	_, _ = w.Write(res.Json())
}

func WriteError(w http.ResponseWriter, err error) {
	res, ok := err.(*errors.Error)
	if !ok {
		log.Println(err)
		res = errors.ErrInternal
	}
	WriteResponse(w, res)
}

func IntFromString(w http.ResponseWriter, v string) (int, error) {
	if v == "" {
		return 0, nil
	}
	res, err := strconv.Atoi(v)
	if err != nil {
		log.Println(err)
		WriteError(w, errors.ErrInvalidLimit)
	}
	return res, err
}

func Unmarshal(w http.ResponseWriter, data []byte, v interface{}) error {
	if d := string(data); d == "null" || d == "" {
		WriteError(w, errors.ErrObjectIsRequired)
		return errors.ErrObjectIsRequired
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Println(err)
		WriteError(w, errors.ErrBadRequest)
	}
	return err
}

func checkSlot(slot *objects.TimeSlot) error {
	if slot == nil {
		return errors.ErrEventTimingIsRequired
	}
	if !slot.StartTime.After(time.Time{}) {
		return errors.ErrInvalidTimeFormat
	}
	if !slot.EndTime.After(time.Time{}) {
		return errors.ErrInvalidTimeFormat
	}
	return nil
}
