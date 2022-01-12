package main

import "testing"

func TestProperCase(t *testing.T) {
	tc := []struct {
		in   string
		want string
	}{{in: "ABCDE", want: "Abcde"},
		{in: "xxxxx", want: "Xxxxx"},
		{in: "ABCDE ABCDE small", want: "Abcde Abcde Small"},
	}

	for _,test:= range tc{
		got:=properCase(test.in)
		if test.want != got {
			t.Errorf("got %s, want %s", got, test.want)
		}
	}

}
