package main

import (
	"fmt"
	"time"
)

func main() {
	limit := time.Date(2018, 2, 18, 12, 00, 00, 00, time.Local)
	now := time.Now()
	diff := limit.Sub(now)
	hoursInt := int(diff.Hours())
	fmt.Printf("卒論提出まで%d日と%d時間%d分", hoursInt/24, hoursInt%24, int(diff.Minutes())%60)
}
