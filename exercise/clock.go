package main 

import "fmt"

// Define the Clock type here.
type Clock struct {
    Hour int
    Min int
}

const timeDay int = 1440

func New(h, m int) Clock {
    var cl Clock
	var time int
    totalMin := h * 60 + m
    switch {
	case totalMin >= timeDay || totalMin == 0 :
		time = totalMin % timeDay 
	case totalMin < 0:
    	if totalMin >= (timeDay * -1) {
            time = timeDay + totalMin
        } else {
        	time = totalMin % timeDay
        	time += timeDay
        }
	case totalMin < timeDay:
		time = totalMin
    }
	cl.Hour = time/60
	cl.Min = time%60
	return cl
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Min + m)
    

}
func (c Clock) Subtract(m int) Clock {
	return c.Add(-1 * m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Min)
}

func main(){
	//a := New(23, 58)
	fmt.Println(-1536%1440)
}