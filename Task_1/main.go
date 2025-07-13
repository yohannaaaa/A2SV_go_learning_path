package main

import "fmt"

func main() {
	name := ""
	fmt.Println("Enter your name please:")
	fmt.Scanln(&name)
	n := 0
	fmt.Println("Enter the amount of subjects you take:")
	fmt.Scan(&n)
	data := map[string]map[string]int{
		name: {},
	} 
	
	for i:= 0; i < n; i++{
		course := ""
		
		fmt.Println("Please enter the subject name:")
		fmt.Scan(&course)
		grade := 0
		fmt.Println("Please enter the grade obtained:")
		fmt.Scan(&grade)
		if grade < 0 || grade > 100{
			fmt.Println("You have entered an invalid grade please try agin")
			i -=1
			continue
		}

		data[name][course] = grade

	}
	fmt.Printf("\nThe grades for the subject you entered are as follows:\n")
	for cou, gra := range data[name]{
		fmt.Printf("In %s you scored %d\n", cou, gra)
	}
	average := getAve(data[name])
	fmt.Printf("\nYour average is %f\n", average)

}

func getAve(grades map[string]int) float64{
	ave := 0
	for _, grade := range grades{
		ave += grade
	}
	return float64(ave) / float64(len(grades))
}