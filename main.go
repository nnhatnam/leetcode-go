package main

import (
	"fmt"
	"encoding/json"
)

var deep = `{
  "root": {
    "internal": {
      "int": 42,
      "string": "string",
      "float64": 6.022e23
    }
  }
}`

// NestedMapLookup
// m:  a map from strings to other maps or values, of arbitrary depth
// ks: successive keys to reach an internal or leaf node (variadic)
// If an internal node is reached, will return the internal map
//
// Returns: (Exactly one of these will be nil)
// rval: the target node (if found)
// err:  an error created by fmt.Errorf
//
func NestedMapLookup(m map[string]interface{}, ks ...string) (rval interface{}, err error) {
	var ok bool

	if len(ks) == 0 { // degenerate input
		return nil, fmt.Errorf("NestedMapLookup needs at least one key")
	}
	temp, _ := m[ks[0]]
	fmt.Println(" ????" , m , ks[0] , temp)
	if rval, ok = m[ks[0]]; !ok {
		return nil, fmt.Errorf("key not found; remaining keys: %v", ks)
	} else if len(ks) == 1 { // we've reached the final key
		return rval, nil
	} else if m, ok = rval.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("malformed structure at %#v", rval)
	} else { // 1+ more keys
		fmt.Println("yo yo " , m , ks  )
		fmt.Println("rval ", rval)
		return NestedMapLookup(m, ks[1:]...)
	}
}

func main() {
	var m = make(map[string]interface{})
	json.Unmarshal([]byte(deep), &m)
	val, err := NestedMapLookup(m, "root", "internal", "float64")
	fmt.Printf("(%T) %v\n", val, val)

	//val, err = NestedMapLookup(m, "root", "infernal")
	fmt.Printf("%v\n", err)
}

//find how many number < v in array arr
//func binarySearch(v int, arr []int) int{
//	if len(arr) == 0 {
//		return 0
//	}
//
//	result := 0
//	lo := 0
//	hi := len(arr) - 1
//	mid := 0
//	for ;lo <= hi; {
//		mid = (lo + hi) / 2
//		if arr[mid] < v {
//			result = mid + 1
//			lo = mid + 1
//		} else {
//			hi = mid - 1
//		}
//	}
//	return result
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//
//	l1 := len(nums1)
//	l2 := len(nums2)
//	if l1 + l2 == 0 {
//		return 0
//	}
//	if l2 > l1 {
//		nums1, nums2 = nums2, nums1
//	}
//
//	mIndex := (l1 + l2 + 1) / 2 - 1
//
//	lo := 0
//	hi := mIndex
//	r1 := mIndex
//	for {
//		v1 := nums1[r1]
//		leftCount := binarySearch(v1, nums2)
//		if r1 + leftCount == mIndex {
//
//			if (l1 + l2) % 2 == 1 {
//				return float64(nums1[r1])
//			} else {
//				if l1 == l2 {
//					return float64(nums1[r1] + nums2[0]) / 2.0
//				} else {
//					return float64(nums1[r1] + nums1[r1 + 1]) / 2.0
//				}
//			}
//
//		} else if r1 + leftCount > mIndex {
//			hi = r1
//			r1 = (lo + hi) / 2
//
//		} else {
//
//			lo = r1 + 1
//			fmt.Println("lo ? ", lo, hi )
//			r1 = (lo + hi) / 2
//		}
//	}
//
//	return 0.0
//}
//
//func main() {
//	fmt.Println(findMedianSortedArrays([]int{1 , 3} , []int{ 2 }))
//}