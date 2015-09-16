# Rabin-Karp multiple pattern search

A golang implementation of [Rabin-Karp](https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm) algorithm.

Provides searching multiple patterns with an average O(n+m) time complexity (where n is the length of text and m is the combined length of pattern strings).


## Installation
```
$ go get github.com/glkz/rabinkarp
```

## Usage

```go
import "github.com/glkz/rabinkarp"

func main() {
  rabinkarp.Search("the text you want to search in", []string{"the", "keywords"})
  // returns []string{"the"}
}
```

See [![GoDoc](https://godoc.org/github.com/glkz/rabinkarp?status.svg)](http://godoc.org/github.com/glkz/rabinkarp)
