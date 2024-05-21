package main

import (
	"fmt"
	"time"
)

func main() {
	var time12hour string
	_, _ = fmt.Scanf("%s\n", &time12hour)
	dateTime, _ := time.Parse("15:04:05PM", time12hour)
	fmt.Println(dateTime.Format("15:04:05"))
}
