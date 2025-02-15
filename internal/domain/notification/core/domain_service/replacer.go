package replacer

import "strings"

type Replaceable struct {
	Tag   string
	Value string
}

func Build(template *string, options ...Replaceable) {
	for _, opt := range options {
		*template = strings.ReplaceAll(*template, opt.Tag, opt.Value)
	}
}
