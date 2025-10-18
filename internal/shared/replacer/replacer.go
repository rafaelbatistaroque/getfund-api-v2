package replacer

import "strings"

// Replaceable represents a key-value pair for replacement,
// where Tag is the placeholder and Value is the content to replace it with.
type Replaceable struct {
	Tag   string
	Value string
}

// Build takes a pointer to a template string and a series of Replaceable options.
// It iterates through the options and replaces all occurrences of each Tag with its corresponding Value in the template.
func Build(template *string, options ...Replaceable) {
	for _, opt := range options {
		*template = strings.ReplaceAll(*template, opt.Tag, opt.Value)
	}
}
