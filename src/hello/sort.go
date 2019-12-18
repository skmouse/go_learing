package main

import (
	"fmt"
)

func main() {
	number := []int{5, 7, 6, 2, 4}
	BubbleSort(number)
	//SelectSort(number)
	//InsertSortOne(number)
	//ShllSort(number)
	//QuickSort(number, 0, len(number)-1)
	//fmt.Println(number)
}

/**
 * 1.冒泡排序--稳定
 */
func BubbleSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}

	fmt.Println(num)
}

/**
 * 2.选择排序--不稳定
 */
func SelectSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		min := i
		for j := i + 1; j < len(num)-1; j++ {
			if num[j] < num[min] {
				min = j
			}
		}

		tmp := num[i]
		num[i] = num[min]
		num[min] = tmp
	}
	fmt.Println(num)
}

/**
 * 3. 插入排序--稳定
 * 当数据基本有序的时候移动的次数少
 */
func InsertSortOne(num []int) {
	for i := 1; i < len(num); i++ {
		insertToRightPosition(num, i)
	}
}

func insertToRightPosition(num []int, i int) {
	inserted := num[i]
	j := i - 1

	for j >= 0 && num[j] > inserted {
		num[j+1] = num[j]
		j--
	}
	num[j+1] = inserted

	fmt.Println(num)
}

/**
 * 4.希尔排序
 */

func ShllSort(num []int) {
	//先分组
	for gap := len(num) / 2; gap > 0; gap /= 2 {
		//对各个分组进行插入排序
		for i := gap; i < len(num); i++ {
			inserI(num, gap, i)
		}
	}
}

func inserI(num []int, gap, i int) {
	inserted := num[i]
	var j int
	for j = i - gap; j >= 0 && num[j] > inserted; j -= gap {
		num[j+gap] = num[j]
	}
	num[j+gap] = inserted
	fmt.Println(num)
}

/**
 * 快速排序
 */
func QuickSort(num []int, low, high int) {
	if low < high {
		priovt := partion(num, low, high)
		QuickSort(num, low, priovt-1)
		QuickSort(num, priovt+1, high)
	}

}

func partion(num []int, low, high int) int {
	i := low
	j := high
	tmp := num[low]
	for i != j {
		for i < j && num[j] >= tmp {
			j--
		}
		for i < j && num[i] <= tmp {
			i++
		}

		if i < j {
			num[i], num[j] = num[j], num[i]
		}

	}
	num[i], num[low] = num[low], num[i]

	return i
}
