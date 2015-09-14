package rabinkarp

import (
	crand "crypto/rand"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

var SearchTests = []struct {
	txt      string
	patterns []string
	expected []string
}{
	{"hello world", []string{"hell", "yeah", "lo"}, []string{"hell", "lo"}},
	{"lorem ipsum dolor", []string{"lorem", "lor", "dolor"}, []string{"lorem", "lor", "dolor"}},
	{"hel", []string{"hell", "yeah"}, []string{}},
	{"", []string{"hell", "yeah"}, []string{}},
	{"", []string{}, []string{}},
	{"", []string{""}, []string{""}},
	{"ε-δ def: Let ƒ:D→R, ∀ε > 0, ∃ a δ > 0 such ∀x∈D that satisfy 0 < |x - c| < δ, the inequality |ƒ(x) - L| < ε holds.",
		[]string{"ƒ:D→R", "√3", "ε-δ", "ƒ(x)", "f(x)", "∀x"}, []string{"ƒ:D→R", "ε-δ", "ƒ(x)", "∀x"}},
}

func TestSearch(t *testing.T) {
	for _, ct := range SearchTests {
		found := Search(ct.txt, ct.patterns)
		if !eq(found, ct.expected) {
			t.Errorf("Search(%s, %#v) = %#v, want %#v",
				ct.txt, ct.patterns, found, ct.expected)
		}
	}
}

func eq(f, s []string) bool {
	fx := make([]string, len(f))
	sx := make([]string, len(s))
	copy(fx, f)
	copy(sx, s)

	sort.Strings(fx)
	sort.Strings(sx)
	return reflect.DeepEqual(fx, sx)
}

func buildRandStr(n int) string {
	var bytes = make([]byte, n)
	crand.Read(bytes)
	return string(bytes)
}

func buildRandStrSlice(n, min, max int) []string {
	r := make([]string, n)
	if n >= 2 {
		r[0] = buildRandStr(min)
		r[1] = buildRandStr(max)
	}
	for i := 2; i < n; i++ {
		r[i] = buildRandStr(rand.Intn(max-min) + min)
	}

	return r
}

var benchInputTxt = buildRandStr(1 << 20)

func benchmarkSearch(b *testing.B, patterns []string) {
	for i := 0; i < b.N; i++ {
		Search(benchInputTxt, patterns)
	}
}

func BenchmarkSearch_m16_n10(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(10, 16, 32)) }
func BenchmarkSearch_m16_n50(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(50, 16, 32)) }
func BenchmarkSearch_m16_n100(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(100, 16, 32)) }
func BenchmarkSearch_m16_n200(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(200, 16, 32)) }

func BenchmarkSearch_m8_n10(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(10, 8, 32)) }
func BenchmarkSearch_m8_n50(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(50, 8, 32)) }
func BenchmarkSearch_m8_n100(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(100, 8, 32)) }
func BenchmarkSearch_m8_n200(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(200, 8, 32)) }

func BenchmarkSearch_m4_n10(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(10, 4, 32)) }
func BenchmarkSearch_m4_n50(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(50, 4, 32)) }
func BenchmarkSearch_m4_n100(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(100, 4, 32)) }
func BenchmarkSearch_m4_n200(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(200, 4, 32)) }

func BenchmarkSearch_m2_n10(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(10, 2, 32)) }
func BenchmarkSearch_m2_n50(b *testing.B)  { benchmarkSearch(b, buildRandStrSlice(50, 2, 32)) }
func BenchmarkSearch_m2_n100(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(100, 2, 32)) }
func BenchmarkSearch_m2_n200(b *testing.B) { benchmarkSearch(b, buildRandStrSlice(200, 2, 32)) }
