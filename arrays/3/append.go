package main

import "fmt"


type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	inferredLen := 0
	for i:= 0; i < len(costs) ; i++ {
		day := costs[i].day
		if day > inferredLen {
			inferredLen = day
		}
	}
	slice := make([] float64 , inferredLen + 1)
	for i := 0 ; i < len(slice) ; i++ {
		
		for n:= 0 ; n< len(costs) ; n++ {
			if i == costs[n].day {
				slice[i] += costs[n].value 
			}
		}
		
	}
	return slice
}

// another way to do the above function

func getCostByDayTeach (costs [] cost) [] float64 {
	costByDay := []float64{}
	for i := 0; i < len(costs) ; i++ {
		cost := costs[i]
		for cost.day >= len(costByDay){
			costByDay = append(costByDay, 0.0)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}




func main () {
	newSlice := make([]int, 5)
 	ap := []int{1, 2 , 4, 5}
	
// this is how we add values to a slice then go takes care of relocating enough memory
	newSlice = append(newSlice , 6) // adding one element to the slice
	newSlice = append(newSlice, 7, 9 ,10) // adding multiple since append is a variadic
	newSlice = append(newSlice, ap...) // adding a slice into it 


	// fmt.Printf("%v", newSlice)

	fmt.Printf("%.1f\n" ,getCostsByDay([]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
	}))

	fmt.Printf("%.1f\n" ,getCostByDayTeach([]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
	}))

	
}