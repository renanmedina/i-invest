package event_store

type PublishableEvent interface {
	Name() string
	ObjectId() string
	ObjectType() string
	EventData() map[string]interface{}
}

type EventHandler interface {
	Handle(event PublishableEvent)
}

type EventPublisher struct {
	handlers map[string][]EventHandler
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		handlers: make(map[string][]EventHandler),
	}
}

func NewEventPublisherWith(handlersSetup map[string][]EventHandler) *EventPublisher {
	return &EventPublisher{
		handlers: handlersSetup,
	}
}

func (p *EventPublisher) Publish(event PublishableEvent) bool {
	eventHandlers, exists := p.handlers[event.Name()]

	if exists {
		for _, handler := range eventHandlers {
			go handler.Handle(event)
		}

		return true
	}

	return false
}

func (p *EventPublisher) Subscribe(eventName string, handler *EventHandler) *EventPublisher {
	eventHandlers, exists := p.handlers[eventName]

	if !exists {
		eventHandlers = make([]EventHandler, 0)
	}

	eventHandlers = append(eventHandlers, *handler)
	p.handlers[eventName] = eventHandlers
	return p
}
