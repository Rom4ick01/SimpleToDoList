package event

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Events struct {
	text      string
	err       string
	timeevent string
}

var event []*Events

func NewEvent(t, e, tme string) {
	event = append(event, &Events{t, e, tme})
}

func ShowEvent() {
	fmt.Println("")
	pp.Println(event)
	fmt.Println("")
}
