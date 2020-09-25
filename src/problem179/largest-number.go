package problem179

import (
	"sort"
	"strconv"
	"strings"
)

// ByLength custom type for string comparision
type ByLength []string

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return (a[i] + a[j]) > (a[j] + a[i]) }

func largestNumber(nums []int) string {
	strnum := make([]string, len(nums))
	for i := range nums {
		strnum[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(ByLength(strnum))
	ans := strings.Join(strnum, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans
}
