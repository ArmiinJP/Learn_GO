package main

import (
	"fmt"
	"time"
)

type tt struct {
	year int
	month int
	day int
	hour int
	minute int
	secound int
}

func AddGigasecond(t time.Time) time.Time{
	var total int
	var tmptime tt
	total += t.Second()
	total += t.Minute()*60
	total += t.Hour()*60*60
	total += t.Day()*24*60*60
	total += int(t.Month())*30*24*60*60
	total += t.Year()*365*24*60*60

	total += 1000000000

	tmptime.year = (total / (365*24*60*60)) 
	total %= (365*24*60*60)

	tmptime.month = total / (30*24*60*60)
	total %= (30*24*60*60)

	tmptime.day = total / (24*60*60)
	total %= (24*60*60)

	tmptime.hour = total / (60*60)
	total %= (60*60)

	tmptime.minute = total / (60)
	total %= 60

	tmptime.secound = total
	
	return time.Date(tmptime.year,time.Month(tmptime.month),tmptime.day,tmptime.hour,tmptime.minute, tmptime.secound,0,t.Location())
}

func main(){
	fmt.Println(time.Now())
	fmt.Println(AddGigasecond(time.Date(2015,1,24,23,59,59,0,time.Now().Location())))
	
}