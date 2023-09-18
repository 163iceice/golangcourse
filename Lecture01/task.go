import (
    "sort"
)
type IntSliceSorter []int
func (s IntSliceSorter) Len() int {
    return len(s)
}
func (s IntSliceSorter) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s IntSliceSorter) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func CompareSlices(slice1, slice2 []int) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    copy1 := make([]int, len(slice1))
    copy2 := make([]int, len(slice2))
    copy(copy1, slice1)
    copy(copy2, slice2)
    sort.Sort(IntSliceSorter(copy1))
    sort.Sort(IntSliceSorter(copy2))
    for i := 0; i < len(copy1); i++ {
        if copy1[i] != copy2[i] {
            return false
        }
    }

    return true
}

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := []int{3, 2, 1}
    slice3 := []int{1, 2, 3, 4}

    result1 := CompareSlices(slice1, slice2)
    result2 := CompareSlices(slice1, slice3) 

    println(result1)
    println(result2)
}
