package main

import "fmt"

func main() {

	var teman = []string{"frisky", "Yoseph Bram", "Fiqri", "clara", "Septanus", "Thalia", "Bayu", "Yudha", "Aulia", "Hasanudin"}
	for _, data := range teman {
		fmt.Println("Halo", data, "salam kenal ya...")
	}
}
