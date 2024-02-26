package maskservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaskingSpam(t *testing.T) {

	var prod producer
	var pres presenter
	ser := &Service{prod, pres}

	var tests = []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "middle link",
			input:  "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			output: "Here's my spammy page: http://******************* see you.",
		},
		{
			name:   "begin link",
			input:  "http://hehe see you.",
			output: "http://**** see you.",
		},
		{
			name:   "end link",
			input:  "Here's my spammy page: http://hehef",
			output: "Here's my spammy page: http://*****",
		},
		{
			name:   "link only",
			input:  "http://oul",
			output: "http://***",
		},
		{
			name:   "space",
			input:  "     ",
			output: "     ",
		},
		{
			name:   "empty",
			input:  "",
			output: "",
		},
		{
			name:   "integers and symbols",
			input:  "3526{!7",
			output: "3526{!7",
		},
		{
			name:   "end link link link",
			input:  "Here's http://hehef http://tera http://te",
			output: "Here's http://***** http://**** http://**",
		},
		{
			name:   "middle empty link",
			input:  "Here's http:// you see",
			output: "Here's http:// you see",
		},
		{
			name:   "only empty link",
			input:  "http://",
			output: "http://",
		},
		{
			name:   "not complete link",
			input:  "ttp://",
			output: "ttp://",
		},
		{
			name:   "not complete link and link",
			input:  "ttp:// http://te",
			output: "ttp:// http://**",
		},
		{
			name:   "one char",
			input:  "h",
			output: "h",
		},
	}

	for _, a := range tests {
		t.Run(a.name, func(t *testing.T) {
			res := ser.MaskingSpam(a.input)
			assert.Equalf(t, a.output, res, "For string with name \"%s\" MaskingSpam(%s) = %s, expected %s", a.name, a.input, res, a.output)
		})
	}
}
