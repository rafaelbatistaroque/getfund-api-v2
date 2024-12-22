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
	s.Params["Send:from"] = to
	s.Params["Send:to"] = to
	s.Params["Send:subject"] = subject
	s.Params["Send:content"] = content
	s.Params["Send:replyTo"] = replyTo

	s.CallsCount["Send"]++

	return s.ErrorResult["Send"]
}
