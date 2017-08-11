package main

import "fmt"

func main()  {
	channel := make(chan int)
	var op1 =1
	op2 := <-channel

	go func() {

		channel <- op1
		fmt.Println(op2)

	}()



}

//Day la van de cua channel ( khong phai channel buffering)
//sender va receiver se bi block cho den khi