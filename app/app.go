package app

import "fmt"

func App() {
	DBConnect()
	fmt.Println("App")
	Routers()
}
