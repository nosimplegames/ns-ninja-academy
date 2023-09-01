package movement

import "github.com/nosimplegames/ns-framework/hnbEvents"

type DynamicBodyFactory struct {
}

func (factory DynamicBodyFactory) Init(body *DynamicBody) {
	body.EventTarget.AddEventCreator(hnbEvents.EventCreator{
		EventType: "jump-stop",
	})
}

func (factory DynamicBodyFactory) Create() *DynamicBody {
	body := &DynamicBody{}

	factory.Init(body)

	return body
}
