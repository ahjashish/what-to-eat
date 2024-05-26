package main

import (
	"fmt"

	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to What to Eat! :)")

	homemap := make(map[string][]string)
	homemap["Greek"] = []string{"Souvlaki", "Gyro", "Tyrokafteri", "Melinzanosalata"}
	homemap["Italian"] = []string{"Margerrita Pizza", "Arabiatta Pasta", "Alfredo Pasta"}
	homemap["Arabic"] = []string{"Shawarma", "Hummus", "Baba Ganoush", "Vine Leaves", "Grilled Chicken"}
	homemap["South Indian"] = []string{"Dosa", "Uttapam", "Idly", "Medu Vada"}
	homemap["North Indian"] = []string{"Butter Chicken", "Biryani", "Chaat", "Rajma Rice", "Kadhi Rice"}
	homemap["Burgers"] = []string{"Chicken Burger"}
	homemap["Desserts"] = []string{"Waffles", "Crepe", "Gujiya", "Ras Malai", "Gulab Jamun", "Kunafa", "Carrot Baklava"}

	officemap := make(map[string][]string)
	officemap["Pasta & Co"] = []string{"Arabiatta Pasta"}
	officemap["Shawarma & Co"] = []string{"Shawarma"}
	officemap["Healthy & Co"] = []string{"Fusili salad", "Black rice salad"}
	officemap["Soup & Co"] = []string{"Lentil soup"}
	officemap["Burger & Co"] = []string{"Chicken Burger"}
	officemap["Acai & Co"] = []string{"Acai Bowl"}
	officemap["Fili"] = []string{"Samosa chat", "Bombay masala omelet", "Maple paneer"}
	officemap["Aloo Beirut"] = []string{"Hummus", "Fattoush", "Lentil soup", "Fransisco sandwich"}


	fmt.Println("Are you ordering from home or office?")
	fmt.Println("1. Home")
	fmt.Println("2. Office")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Printf("You selected option %d\n", 1)
		var homekeys []string
		for k := range homemap {
			homekeys = append(homekeys, k)
		}
		fmt.Println("Select from following choices:")
		fmt.Println("1. Choose a random cuisine and dish")
		fmt.Println("2. Choose a cuisine")
		fmt.Printf("You selected option %d\n", 1)
		fmt.Println("Choosing a random cuisine and dish for today from the following options :")
		for ind, key := range homekeys {
			fmt.Printf("%d. %s\n", ind+1, key)
			for i, d := range homemap[key] {
				fmt.Printf("\t%d. %s\n", i+1, d)
			}
		}

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		lk := len(homekeys)
		rcn := r.Intn(lk)
		rc := homekeys[rcn]
		fmt.Printf("Random cuisine selected for you is : %s\n", rc)

		dishes := homemap[rc]
		ld := len(dishes)
		rdn := r.Intn(ld)
		rd := dishes[rdn]
		fmt.Printf("Random dish selected for you is : %s\n", rd)
	} else if choice == 2 {
		var officekeys []string
		for k := range officemap {
			officekeys = append(officekeys, k)
		}
		fmt.Printf("You selected option %d\n", 2)
		fmt.Println("Choosing a random cuisine and dish for today from the following options :")
		for key, dishes := range officemap {
			fmt.Printf("%s\n", key)
			for i, d := range dishes {
				fmt.Printf("\t%d. %s\n", i+1, d)
			}
		}

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		lk := len(officekeys)
		rcn := r.Intn(lk)
		rc := officekeys[rcn]
		fmt.Printf("Random cuisine selected for you is : %s\n", rc)

		dishes := officemap[rc]
		ld := len(dishes)
		rdn := r.Intn(ld)
		rd := dishes[rdn]
		fmt.Printf("Random dish selected for you is : %s\n", rd)
	} else {
		fmt.Println("Invalid choice")
		return
	}
}
