package views

const (
	Setup    = "setup"
	Loading  = "loading"
	Chatting = "chatting"
)

type Screen struct {
	state string
}

func UpdateView() {}

func renderSetup()   {}
func renderLoading() {}
func renderChat()    {}
