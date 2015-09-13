package rabinkarp

import (
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
