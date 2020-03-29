package main

import "fmt"

func main() {
	exit := false
	country1 := false
	country2 := false
	country3 := false
	pak := 1571
	france := 2045
	korea := 80231

	for exit != true {
		var input int
		fmt.Println("1 – Print Covid19 cases in Pakistan ")
		fmt.Println("2 – Print Covid19 cases in SouthKorea ")
		fmt.Println("3 – Print Covid19 cases in France ")
		fmt.Println("4 – Print my personalized message to Coronavirus ")
		fmt.Println("5 – Exit ")
		fmt.Println(" ")
		fmt.Print("Enter Number :")
		fmt.Scanln(&input)

		if input == 1 {
			fmt.Print("Covid19 cases in Pakistan :")
			fmt.Println(pak)
			fmt.Println(" ")
			country1 = true

		}
		if input == 2 {
			fmt.Print("Covid19 cases in SouthKorea :")
			fmt.Println(korea)
			fmt.Println(" ")
			country2 = true

		}
		if input == 3 {
			fmt.Print("Covid19 cases in France :")
			fmt.Println(france)
			fmt.Println(" ")
			country3 = true
		}
		if input == 4 {
			var msg string
			fmt.Print("Enter First String: ")
			fmt.Scanln(&msg)
			fmt.Println(" ")

		}
		if input == 5 {
			fmt.Println("EXITING")
			exit = true

		}
		if country1 == true && country2 == true && country3 == true {
			fmt.Println("Seen All Details: EXITING")
			exit = true
		}

	}
}
