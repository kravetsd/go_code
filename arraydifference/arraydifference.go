/*#Есть 2 сортированных последовательности целых чисел.
# Хотим напечатать те элементы из первой последовательности,
# которых нет во второй последовательности

#(если число хоть раз встречается во второй последовательности,
#то удаляем все его экземпляры из первой).

#Нужно решение без накопления в памяти и в один проход.
*/

/*
0, 1, 1, 1, 2, 3, 4, 5, 5, 10, 100, 1000
1, 1, 2, 3, 6, 11, 100

0, 4, 5, 5, 10, 1000
*/

package main

import "fmt"

func main() {
	a := []int{0, 1, 1, 1, 2, 3, 4, 5, 5, 10, 100, 1000}
	b := []int{-1, 1, 1, 2, 3, 6, 11, 100}

	compare(a, b)

}

func compare(a []int, b []int) {
	lastJ := 0
	for i := range a {
		found := false
		for j := lastJ; j < len(b); j++ {
			lastJ = j
			if a[i] == b[j] {
				found = true
				break
			} else {
				if a[i] < b[j] {
					break
				}
			}
		}
		if !found {
			fmt.Println(a[i])
		}
	}
}