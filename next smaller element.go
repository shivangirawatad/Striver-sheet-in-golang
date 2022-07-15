package main
import "fmt"

func nextsmallerelement(arr []int, n int) []int
{
    res := make([][]int, n)

    for i := 0; i < n; i++ {
        next := -1;
        for j := i + 1; j < n; j++ {
            if arr[i] > arr[j] {
                next = arr[j];
            }
        }
        res = append(res, next)
    }
    return res
}
 
func main() {
	arr := []int{4, 8, 5, 2, 25}
    n := len(arr)
    res := nextsmallerelement(arr, n);
}