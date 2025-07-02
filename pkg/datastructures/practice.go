package main

import "fmt"

func main3() {
	hobbies := [3]string{"Cooking", "Eating", "Sleeping"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	hobbies2 := hobbies[:2]
	fmt.Println(hobbies2)

	fmt.Println(cap(hobbies2), len(hobbies2))
	hobbies2 = hobbies2[1:3]
	fmt.Println(hobbies2)

	courseGoals := []string{"Learn Basic Go", "Learn Data Structure"}
	fmt.Println(courseGoals)
	courseGoals[1] = "Learn all the details"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "Learn Data Structure")
	fmt.Println(courseGoals)

	products := []Product{
		{
			"Product 1",
			"First Product",
			22.33,
		},
		{
			"Product 2",
			"Second Product",
			44.13,
		},
	}
	fmt.Println(products)

	products = append(products,
		Product{"Product 3", "Third Product", 32.12})
	fmt.Println(products)

}

type Product struct {
	id    string
	title string
	price float64
}
