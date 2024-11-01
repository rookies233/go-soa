// package main

// // fullPermutation 全排列
// func fullPermutation(nums []int) [][]int {
// 	var res [][]int
// 	var dfs func(int)
// 	dfs = func(x int) {
// 		if x == len(nums)-1 {
// 			tmp := make([]int, len(nums))
// 			copy(tmp, nums)
// 			res = append(res, tmp)
// 			return
// 		}
// 		for i := x; i < len(nums); i++ {
// 			nums[x], nums[i] = nums[i], nums[x]
// 			dfs(x + 1)
// 			nums[x], nums[i] = nums[i], nums[x]
// 		}
// 	}
// 	dfs(0)
// 	return res
// }

package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Request struct {
	XMLName xml.Name `xml:"Request"`
	Nums    []int    `xml:"Nums>Num"`
}

type Response struct {
	XMLName      xml.Name `xml:"Response"`
	Permutations [][]int  `xml:"Permutations>Permutation"`
}

func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(path []int, used []bool)
	used := make([]bool, len(nums))

	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			backtrack(append(path, nums[i]), used)
			used[i] = false
		}
	}

	backtrack([]int{}, used)
	return result
}

func soapHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	permutations := permute(req.Nums)

	resp := Response{Permutations: permutations}
	w.Header().Set("Content-Type", "text/xml")
	if err := xml.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/soap", soapHandler)
	fmt.Println("SOAP service is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
