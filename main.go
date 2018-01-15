package main

import (
	"fmt"
	"time"
)

func main() {
	limit := time.Date(2018, 2, 18, 12, 00, 00, 00, time.Local)
	now := time.Now()
	diff := limit.Sub(now)
	fmt.Printf("卒論提出まで%.0f日と%.0f時間", diff.Hours()/24, diff.Hours()/24/60)
}
