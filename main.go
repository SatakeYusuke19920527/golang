package main

import "fmt"

func receiver(name string){
	i := int(0)
	for{
		fmt.Println(name, "receiver", i)
		if i == 10 {
			break
		}
		i++
	}
}

func main() {
		fmt.Println("Hello!");
		go receiver("name1")
		go receiver("name2")

}