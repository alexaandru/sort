package utils

import "sort"

// Pair will hold one sorted list item (for a map of int to int)
type Pair struct {
    Key   int
    Value int
}

// StrPair will hold one sorted list item (for a map of string to int)
type StrPair struct {
    Key   string
    Value int
}

// PairList holds a list of Pair
type PairList []Pair

// StrPairList holds a list of StrPair
type StrPairList []StrPair

// Implement the sort.Interface interface.
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// SortMapByValue turns a map into a PairList, then sorts and returns it.
func SortMapByValue(m map[int]int) (p PairList) {
    p = make(PairList, len(m))
    i := 0
    for k, v := range m {
        p[i] = Pair{k, v}
        i++
    }
    sort.Sort(p)

    return
}

// SortDescMapByValue same as SortMapByValue but with order reversed.
func SortDescMapByValue(m map[int]int) (p PairList) {
    p = SortMapByValue(m)
    sort.Sort(sort.Reverse(p))

    return
}

// Implement the sort.Interface interface.
func (p StrPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p StrPairList) Len() int           { return len(p) }
func (p StrPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// SortStrMapByValue turns a map into a StrPairList, then sorts and returns it.
func SortStrMapByValue(m map[string]int) (p StrPairList) {
    p = make(StrPairList, len(m))
    i := 0
    for k, v := range m {
        p[i] = StrPair{k, v}
        i++
    }
    sort.Sort(p)

    return
}

// SortDescStrMapByValue same as SortStrMapByValue but with order reversed.
func SortDescStrMapByValue(m map[string]int) (p StrPairList) {
    p = SortStrMapByValue(m)
    sort.Sort(sort.Reverse(p))

    return
}
