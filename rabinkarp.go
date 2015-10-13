/*
Package rabinkarp implements Rabin-Karp* string search algorithm.

Provides searching multiple patterns with an average O(n+m) time complexity
(where n is the length of text and m is the combined length of pattern strings).

* https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
*/
package rabinkarp

import "io"

const base = 16777619

// Search searches given patterns in txt and returns the matched ones. Returns
// empty string slice if there is no match.
func Search(txt string, patterns []string) []string {
	in := indices(txt, patterns)
	matches := make([]string, len(in))
	i := 0
	for j, p := range patterns {
		if _, ok := in[j]; ok {
			matches[i] = p
			i++
		}
	}

	return matches
}

// SearchReader searches given patterns in r(io.Reader) without buffering the
// whole stream.
func SearchReader(r io.Reader, patterns []string) []string {
	min, max := minLen(patterns), maxLen(patterns)
	chunksize := 4096

	if chunksize < 2*max {
		chunksize = 2 * max
	}

	txt := make([]byte, chunksize)
	eof := false
	round := 0
	matches := make(map[int]int)

	if n, err := io.ReadFull(r, txt); err != nil {
		eof = true
		txt = txt[:n]
	}

	hp := hashPatterns(patterns, min)
	h := hash(string(txt[:min]))

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < min-1; i++ {
		mult = (mult * base)
	}

	var cp uint32 // prev char

	for i := 0; i < len(txt); i++ {
		//rem := len(txt) - i - 1
		if !eof && len(txt)-i-1 < max {
			buf2 := make([]byte, chunksize)
			n, err := io.ReadFull(r, buf2)
			if err != nil {
				eof = true
				buf2 = buf2[:n]
			}

			txt = append(txt[i:], buf2...)
			i = 0
			round++
		}

		if eof && len(txt)-i-1 < min {
			break
		}

		if (round > 0) || (round == 0 && i > 0) {
			h = h - mult*cp
			h = h*base + uint32(txt[i+min-1])
		}

		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= len(txt) && pat == string(txt[i:e]) {
					matches[pi] = (round * max) + i
				}
			}
		}

		cp = uint32(txt[i])
	}

	return matchesFromIndices(matches, patterns)
}

func matchesFromIndices(in map[int]int, patterns []string) []string {
	matches := make([]string, len(in))
	i := 0
	for j, p := range patterns {
		if _, ok := in[j]; ok {
			matches[i] = p
			i++
		}
	}

	return matches
}

// indices returns the indices of the first occurence of each pattern in txt.
func indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), minLen(patterns)
	matches := make(map[int]int)

	if n < m || len(patterns) == 0 {
		return matches
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := hashPatterns(patterns, m)
	h := hash(txt[:m])
	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
		if i > 0 {
			h = h - mult*uint32(txt[i-1])
			h = h*base + uint32(txt[i+m-1])
		}

		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

func hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func hashPatterns(patterns []string, l int) map[uint32][]int {
	m := make(map[uint32][]int)
	for i, t := range patterns {
		h := hash(t[:l])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], i)
		} else {
			m[h] = []int{i}
		}
	}

	return m
}

func minLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}

	m := len(patterns[0])
	for i := range patterns {
		if m > len(patterns[i]) {
			m = len(patterns[i])
		}
	}

	return m
}

func maxLen(patterns []string) int {
	m := 0
	for i := range patterns {
		if m < len(patterns[i]) {
			m = len(patterns[i])
		}
	}

	return m
}
