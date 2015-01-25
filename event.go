type event struct {
	trigger channel
	subscriptions []channel
}

func (e *event) Subscribe() {

	c := make(channel)
	e.subscriptions = append(e.subscriptions, c)
	return c
}

func (e *event) EventListener() {



}
