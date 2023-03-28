package main

import (
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name   string
		regexp string
		text   string
		want   bool
	}{
		{name: "match: a, apple", regexp: `a`, text: "apple", want: true},
		{name: "not match: a, dog", regexp: `a`, text: "dog", want: false},

		{name: "match: \\d, 3", regexp: `\d`, text: "3", want: true},
		{name: "not match: \\d, c", regexp: `\d`, text: "c", want: false},

		{name: "match: \\w, foo101", regexp: `\w`, text: "foo101", want: true},
		{name: "not match: \\w, $!?", regexp: `\w`, text: "$!?", want: false},

		{name: "match: [abc], apple", regexp: `[abc]`, text: "apple", want: true},
		{name: "not match: [abc], dog", regexp: `[abc]`, text: "dog", want: false},

		{name: "match: [^abc], cab", regexp: `[^abc]`, text: "cab", want: true},
		{name: "not match: [^abc], dog", regexp: `[^abc]`, text: "dog", want: false},

		{name: "match: \\d apple, 1 apple", regexp: `\d apple`, text: "1 apple", want: true},
		{name: "not match: \\d apple, 1 orange", regexp: `\d apple`, text: "1 orange", want: false},

		{name: "match: \\d\\d\\d apple, 100 apples", regexp: `\d\d\d apple`, text: "100 apples", want: true},
		{name: "not match: \\d\\d\\d apple 1 apples", regexp: `\d\d\d apple`, text: "1 apple", want: false},

		{name: "match: \\d \\w\\w\\ws, 3 dogs", regexp: `\d \w\w\ws`, text: "3 dogs", want: true},
		{name: "not match: \\d \\w\\w\\ws, 4 cats", regexp: `\d \w\w\ws`, text: "4 cats", want: true},
		{name: "not match: \\d \\w\\w\\ws, 1 dog", regexp: `\d \w\w\ws`, text: "1 dog", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.regexp, tt.text); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}