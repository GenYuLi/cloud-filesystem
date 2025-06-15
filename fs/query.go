// fs/query.go
package fs

type PatternType int

const (
  PatternRegex PatternType = 1 << iota
  PatterenPrefix
  PatterenSuffix
  PatternExact
)
