// named-return function

package main

import "fmt"

func post(){
	fmt.Println("Beautiful Memories!")
}

func main(){

	for i:=0; i<5; i++{
		post()
	}

}