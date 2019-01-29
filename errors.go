package dexcom

import (
	"fmt"
)

type Error struct {
	Fault `json:"fault"`
}

type Fault struct {
	Fault string `json:"faultString"`
	Detail map[string]string `json:"detail"`
}

func (e Error) Error() string {
	return fmt.Sprintf("dexcom: %s", e.Fault.Fault)
}