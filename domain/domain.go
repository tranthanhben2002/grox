package domain

import "time"

type Id string

type TimeStamp struct {
	CreatedTime time.Time `json:"created_time,omitempty" gorethink:"created_time,omitempty"`
	UpdatedTime time.Time `json:"updated_time,omitempty" gorethink:"updated_time,omitempty"`
}
