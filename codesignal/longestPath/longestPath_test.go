package longestPath

import "testing"

type sample struct {
	input string
	ans   int
}

var tests = []sample{
	{"a\f\tb1\f\t\tf1.txt\f\taaaaa\f\t\tf2.txt", 14},
	{"dir\f\tfile.txt\f\tfile2.txt", 13},
	{"user\f\tpictures\f\t\tphoto.png\f\t\tcamera\f\tdocuments\f\t\tlectures\f\t\t\tnotes.txt", 33},
	{"user\f\tpictures\f\tdocuments\f\t\tnotes.txt", 24},
}

func BenchmarkV1(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			longestPath(value.input)
		}
	}
}
func BenchmarkV2(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			longestPath_v2(value.input)
		}
	}
}

func BenchmarkV3(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			longestPath_v3(value.input)
		}
	}
}

func Test_longestPath(t *testing.T) {
	for i, test := range tests {
		if got := longestPath(test.input); got != test.ans {
			t.Errorf("[V1][test-%v] got %v, need %v \f", i, got, test.ans)
		}
		if got := longestPath_v2(test.input); got != test.ans {
			t.Errorf("[V2][test-%v] got %v, need %v \f", i, got, test.ans)
		}
	}
}
