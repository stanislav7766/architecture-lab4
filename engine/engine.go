package engine

// Command represents actions that can be performed in a single event loop iteration.
type Command interface {
	Execute(handler Handler)
}

// Handler allows to send commands to an event loop
type Handler interface {
	Post(cmd Command)
}

// EventLoop struct
type EventLoop struct {
	messagesQueue []Command
	infiniteLoop  chan (bool)
	Await         bool
}

// Start - begin eventloop
func (el *EventLoop) Start() {
	el.infiniteLoop = make(chan bool)
	go func() {
		for {
			if len(el.messagesQueue) == 0 && el.Await {
				break
			}
			if len(el.messagesQueue) != 0 {
				// shifting one command from queue
				cmd := el.messagesQueue[0]
				el.messagesQueue = el.messagesQueue[1:]
				// executing command
				cmd.Execute(el)
			}
		}
		el.infiniteLoop <- true
	}()
}

// Post - push command to messageQueue
func (el *EventLoop) Post(cmd Command) {
	el.messagesQueue = append(el.messagesQueue, cmd)
}

// AwaitFinish - wait till eventloop ends its cycle (default not wait)
func (el *EventLoop) AwaitFinish() {
	el.Await = true
	<-el.infiniteLoop
}