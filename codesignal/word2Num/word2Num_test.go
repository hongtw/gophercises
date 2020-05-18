package word2num

import "testing"

type sample struct {
	input string
	ans   int
}

var tests = []sample{
	{"Nine hundred eighty seven million six hundred fifty-four thousand three hundred and twenty one", 987654321},
	{"Two hundred and thirty-seven", 237},
	{"One thousand, nine hundred and seventy-nine", 1979},
	{"One hundred Twenty three thousand, Four hundred and fifty-six", 123456},
}

func Test_word2num(t *testing.T) {
	for i, test := range tests {
		if got := word2Num(test.input); got != test.ans {
			t.Errorf("[test-%v] got %v, need %v \n", i, got, test.ans)
		}
	}
}
