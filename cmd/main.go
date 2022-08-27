package main

import (
	"fmt"

	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to What to Eat! :)")

	m := make(map[string][]string)
	m["Greek"] = []string{"Souvlaki", "Gyro", "Tyrokafteri", "Melinzanosalata"}
	m["Italian"] = []string{"Margerrita Pizza", "Arabiatta Pasta", "Alfredo Pasta"}
	m["Arabic"] = []string{"Shawarma", "Hummus", "Baba Ganoush", "Vine Leaves", "Grilled Chicken"}
	m["South Indian"] = []string{"Dosa", "Uttapam", "Idly", "Medu Vada"}
	m["North Indian"] = []string{"Butter Chicken", "Biryani", "Chaat", "Rajma Rice", "Kadhi Rice"}
	m["Burgers"] = []string{"Chicken Burger", "Veg Burger"}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println("Choosing a random cuisine and dish for today from the following options :")
	for ind, key := range keys {
		fmt.Printf("%d. %s\n", ind+1, key)
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	lk := len(keys)
	rcn := r.Intn(lk)
	rc := keys[rcn]
	fmt.Printf("Random cuisine selected for you is : %s\n", rc)

	dishes := m[rc]
	ld := len(dishes)
	rdn := r.Intn(ld)
	rd := dishes[rdn]
	fmt.Printf("Random dish selected for you is : %s\n", rd)
}
