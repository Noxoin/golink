package server

import (
	"testing"
)

func TestValidateLinkName(t *testing.T) {
	tests := []struct {
		input string
		res   bool
	}{
		{input: "", res: false},
		{input: "foo", res: true},
		{input: "foo2", res: true},
		{input: "/foo", res: false},
		{input: "foo-bar", res: true},
		{input: "!!!S", res: false},
		{input: "-wer", res: false},
		{input: "3wer", res: false},
	}
	for _, test := range tests {
		r, err := validateLinkName(test.input)
		if err != nil {
			t.Errorf("TestValidateLinkName failed on %v: %v", test.input, err.Error())
		}
		if r != test.res {
			t.Errorf("TestValidateLinkName failed on %v: got: %v, wanted %v", test.input, r, test.res)
		}
	}
}

func TestGetLinkName(t *testing.T) {
	tests := []struct {
		input string
		res   string
		err   bool
	}{
		{input: "", err: true},
		{input: "/foo", res: "foo", err: false},
		{input: "/", err: true},
		{input: "/2invalid", err: true},
		{input: "valid", res: "valid", err: false},
	}
	for _, test := range tests {
		r, err := getLinkName(test.input)
		if (err != nil) != test.err {
			t.Errorf("TestGetLinkName failed on error %v: got: %v, wanted: %v", test.input, test.err, err != nil)
		} else if r != test.res {
			t.Errorf("TestGetLinkName failed on %v: got: %v, wanted: %v", test.input, r, test.res)
		}
	}
}
