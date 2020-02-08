package main

import "fmt"

// function to sum elements in slice
func sum(slice []float64) float64 {
	res := 0.0
	for _, v := range slice {
		res += v
	}
	return res
}

func main() {
	evens := []int{2, 4}
	odds := []int{3, 5, 7}
	copy(evens, odds)
	fmt.Printf("Evens: %v\n", evens)

	// another example
	data := []float64{10, 25, 30, 50}
	fmt.Printf("data: %v\n", data)
	// fmt.Printf("Is it gonna rain? %.f%% chance.\n", (data[0]+data[1]+data[2]+data[3])/float64(len(data)))
	fmt.Printf("Is it gonna rain? %.f%% chance.\n", (sum(data))/float64(len(data)))

	newData := []float64{99, 100}
	// copy function also return number of elements copied
	n := copy(data, newData)
	fmt.Printf("data: %v, copied: %d\n", data, n)

	// what if data to be copied is bigger than the data wher it has to be copied.
	freshData := []float64{60, 70, 80, 90, 40, 30}
	// one way to do
	// data = append(data[:0], freshData...)
	// fmt.Printf("data: %v\n", data)
	// fmt.Printf("Is it gonna rain? %.f%% chance.\n", (sum(data))/float64(len(data)))

	// using copy
	// saved := make([]float64, len(freshData))
	saved := append([]float64(nil), data...)
	copy(saved, freshData)
	fmt.Printf("data: %v\n", freshData)
	fmt.Printf("Is it gonna rain? %.f%% chance.\n", (sum(data))/float64(len(data)))

}
