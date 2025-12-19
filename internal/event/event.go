package event

import (
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
	pp.Println(event)
}
