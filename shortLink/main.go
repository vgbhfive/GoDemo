package main

import "fmt"

func main() {
	fmt.Println("Server Run Successful ...")
	a := App{}
	a.Initialize(getEnv())
	a.Run(":9000")

}
