package main

import (
	"time"
	"fmt"
	"math/rand"
)

func init(){
	rand.Seed(time.Now().UTC().UnixNano())
}

func main()  {
	loc,_ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc).Format(time.RFC850)
	fmt.Println(t)
	fmt.Println(randNum(5))
	fmt.Println(randNum(5))
	fmt.Println(randNum(5))
	fmt.Println(randNum(5))
	fmt.Println(randNum(89))
	fmt.Println(randNum(5))
	fmt.Println(randNum(84))

	loc,_ =  time.LoadLocation("America/Juneau")
	t1 := time.Unix(1529679600,0).In(loc)
	t2 := time.Now().In(loc)
	fmt.Println(t1,t2)
}

func randNum(val int)int{
	return rand.Intn(val)
}