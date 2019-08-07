package main

import (
	"fmt"
)

/*
	[1 2 3 4]
	output all of the possible order
		1 2 3 4
		1 2 4 3
		1 3 2 4
		1 3 4 2
		...........
 */

 func exportAllOrders(array []int, ret* [][]int) {
 	if len(array) < 1 {
 		return
	}
	if len(array) == 1 {
		*ret = append(*ret, []int{array[0]})
		return
	}
	for i := 0; i < len(array); i++ {
		result := make([][]int, 0)
		tmpArray := make([]int, 0)
		tmpArray = append(tmpArray, array[:i]...)
		if i+1 < len(array) {
			tmpArray = append(tmpArray, array[i+1:]...)
		}
		exportAllOrders(tmpArray, &result)
		for _,order := range result {
			tmpOrder := make([]int, 0)
			tmpOrder = append(tmpOrder, array[i])
			tmpOrder = append(tmpOrder, order...)
			*ret = append(*ret, tmpOrder)
		}
	}
 }

 /*
 [0,1,0,1,2,0,2,1..................]
 sort a array of which element is 0 or 1 or 2
  */
  func swap(array []int, i, j int) {
  	tmp := array[i]
  	array[i] = array[j]
  	array[j] = tmp
  }

  func sort(array []int) {
  	if len(array) < 1 {
  		return
	}
	index0 := -1
	index1 := -1
	index2 := -1
	for i := 0; i < len(array); i++ {
		if array[i] == 2 {index2++}
		if array[i] == 0 {
			swap(array, index0+1, i)
			index0++
			if index1 == -1 && index2 != -1 {
				index2++
			} else if index1 != -1 && index2 == -1 {
				index1++
			} else if index1 != -1 && index2 != -1 {
				swap(array, index1+1, i)
				index1++
				index2++
			}
		}
		if array[i] == 1 {
			if index2 == -1 {
				index1 = i
			} else {
				if index1 == -1 {
					swap(array, index0+1,i)
					index1 = index0+1
					index2++
				} else {
					swap(array, index1+1, i)
					index1++
					index2++
				}
			}
		}
	}
  }

 func main() {
 	orders := make([][]int, 0)
 	exportAllOrders([]int{1,2,3,4}, &orders)
 	fmt.Println(orders)

 	random := []int{1,2,0,2,1,0,1,2,1,0,0,2,1,2,0,2,1,0,0,0,1,2,2,2,0,0}
 	sort(random)
 	fmt.Println(random)
 }

