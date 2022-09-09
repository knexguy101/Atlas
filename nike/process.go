package nike

import "context"

func (t *Task) Run() {
	t.Start()
	t.ctx, t.cancel = context.WithCancel(context.Background())
	t.state = INIT
	for t.running {
		switch t.state {
		case INIT:
			t.state = t.init()
			break
		case LOAD:
			t.state = t.load()
			break
		case LOGIN:
			t.state = t.login()
			break
		case CLEAR:
			t.state = t.clear()
			break
		case ATC:
			t.state = t.cart()
			break
		case ADDRESS:
			t.state = t.address()
			break
		case PAYMENT:
			t.state = t.payment()
			break
		case SUBMIT:
			t.state = t.submit()
			break
		case PENDING:
			break
		case FINISHED:
			t.Stop()
			break
		}
	}
	t.closePageAndBrowser()
	t.Stop()
	t.setStatus("Stopped")
}
