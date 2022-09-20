package models

type RequestPayload struct {
	Action   string      `json:"action"`
	Auth     AuthPayload `json:"auth,omitempty"`
	Log      LogEntry    `json:"log,omitempty"`
	Mail     MailPayload `json:"mail,omitempty"`
	QueueMsg MqPayload   `json:"queuemsg,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogEntry struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type MqPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
