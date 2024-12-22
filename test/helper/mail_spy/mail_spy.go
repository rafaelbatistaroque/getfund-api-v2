package mail_spy

type MailServiceSpy struct {
	Params     map[string]interface{}
	CallsCount map[string]int

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func New() *MailServiceSpy {
	return &MailServiceSpy{
		Params:        make(map[string]interface{}),
		CallsCount:    make(map[string]int),
		SuccessResult: make(map[string]interface{}),
		ErrorResult:   make(map[string]error),
	}
}

func (s *MailServiceSpy) SendMail(from, to, subject, content string, replyTo []string) error {
	s.Params["SendMail:from"] = from
	s.Params["SendMail:to"] = to
	s.Params["SendMail:subject"] = subject
	s.Params["SendMail:content"] = content
	s.Params["SendMail:replyTo"] = replyTo

	s.CallsCount["SendMail"]++

	return s.ErrorResult["SendMail"]
}
