package main

import (
	
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "  \n&& 11 HHHello     world  ",
		expected: []string{"&&","11","hhhello", "world"},
	},
	
}
for _, c := range cases {
	actual := cleanInput(c.input)


	if len(actual)!=len(c.expected){t.Fatalf("Length Mismatch for input %q: got %d , expected %d",c.input,len(actual),len(c.expected))}
	for i := range actual {
	word := actual[i]
	expectedWord := c.expected[i]
	if word !=expectedWord { t.Errorf("word mismatch at index %d for input %q: got %q, expected %q", i, c.input, word, expectedWord)}
}
		}
	}
