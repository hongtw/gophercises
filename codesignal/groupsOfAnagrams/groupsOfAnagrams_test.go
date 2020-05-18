package groupsofanagrams

import (
	"fmt"
	"math/rand"
	"testing"
)

type sample struct {
	v   []string
	ans int
}

var tests = []sample{
	{[]string{"tea", "eat", "eat", "ate", "aet", "eta", "eat", "tae", "tea"}, 1},
	{[]string{"c", "c", "c", "c", "uewuhu", "uwuueh", "ewuuuh", "huuweu", "n", "z"}, 4},
	{[]string{"jsbuvd", "bvujsd", "sjvbud", "ubdjsv", "vdsbju", "duvbsj", "sdbvju", "vdjubs", "bdjvsu", "jdvsub", "vsdubj", "sjudbv", "uvjbsd", "dubsjv", "subdvj", "dvjubs", "jsdbvu", "vjbsud", "sdbjuv", "vdsujb", "jvdbus", "vdusbj", "sjvubd", "dsvjbu", "sjdbvu", "vdjsub", "uvbdjs", "bsjduv", "sjbudv", "ujvsdb", "bvjsdu", "dbvjus", "vujbds", "bsudjv", "djsvbu", "sbvudj", "budsvj", "bvsujd", "dvujsb", "budsvj", "vbsduj", "usvdbj", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "r", "spvfxdy", "xvfdspy", "fdxsypv", "spyvxdf", "pxdysfv", "fydvpxs", "pysxdvf", "sxdyvpf", "otjnxt", "oxttjn", "totnjx", "tnjotx", "totnjx", "ahcqngwmuw", "ngqhumacww", "wqunmahwgc", "qmhgwuncaw", "mnuhqwcawg", "wmguhaqnwc", "cuawmnhwqg", "cwhngmqauw", "hwwauqmcgn", "wgwhanmquc", "mawqugnchw", "mahquncgww", "awhqwgmunc", "nwcmahguwq", "uwhmgcqwan", "hcwagqwnum", "cwqhwmuang", "hcagnumwqw", "hgnmwuawqc", "auqgwwncmh", "gdscjyz", "zjygdsc", "ygcdszj", "gyczsdj", "sdcjzyg", "syzcgjd", "tqr", "qrt"}, 7},
}

func generateRuningSamples() [][]string {
	const n = 1000
	// Guaranteed constraints:
	// 1 ≤ words.length ≤ 10^5,
	// 1 ≤ words[i].length ≤ 10.

	samples := make([][]string, n)
	for i := 0; i < n; i++ {
		wordCnt := rand.Intn(10000) + 1
		wordList := make([]string, wordCnt)
		for j := 0; j < wordCnt; j++ {
			wordLen := rand.Intn(10) + 1
			word := ""
			for k := 0; k < wordLen; k++ {
				word += string('a' + rand.Intn(26))
			}
			wordList[j] = word
		}
		samples[i] = wordList
	}
	for i := 0; i < 3; i++ {
		fmt.Println(samples[i][:5])
	}

	return samples
}

var benchmarkSamples = generateRuningSamples()

func BenchmarkV1(gb *testing.B) {

	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			groupsOfAnagrams(value.v)
		}
		for _, value := range benchmarkSamples {
			groupsOfAnagrams(value)
		}
	}
}
func BenchmarkV2(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			groupsOfAnagramsV2(value.v)
		}
		for _, value := range benchmarkSamples {
			groupsOfAnagramsV2(value)
		}
	}
}

func BenchmarkMaxVersion(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			groupsOfAnagramsMaxVersion(value.v)
		}
		for _, value := range benchmarkSamples {
			groupsOfAnagramsMaxVersion(value)
		}
	}
}

func BenchmarkV3(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			groupsOfAnagramsV3(value.v)
		}
		for _, value := range benchmarkSamples {
			groupsOfAnagramsV3(value)
		}
	}
}

func BenchmarkV4(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range tests {
			groupsOfAnagramsV4(value.v)
		}
		for _, value := range benchmarkSamples {
			groupsOfAnagramsV4(value)
		}
	}
}

func Test_arrayPacking(t *testing.T) {

	for i, test := range tests {
		if got := groupsOfAnagrams(test.v); got != test.ans {
			t.Errorf("[test-%v] v1: got %v, need %v \n", i, got, test.ans)
		}
		if got := groupsOfAnagramsV2(test.v); got != test.ans {
			t.Errorf("[test-%v] v2: got %v, need %v \n", i, got, test.ans)
		}
		if got := groupsOfAnagramsV3(test.v); got != test.ans {
			t.Errorf("[test-%v] v2: got %v, need %v \n", i, got, test.ans)
		}
	}
}
