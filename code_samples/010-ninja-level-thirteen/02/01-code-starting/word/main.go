package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//func Count(s string) int {
//	xs := strings.Fields(s)
//	c := 0
//	for i, _ := range xs {
//		c = c + 1
//}
