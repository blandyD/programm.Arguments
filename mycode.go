package main

import ("fmt"
		"os"
	)

		func main(){
			var argument []string
			argument = os.Args[1:]

			for _,contenus := range argument{
				fmt.Println(contenus)
			}
		}