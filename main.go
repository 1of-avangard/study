package main

import (
	"fmt"
	simpleconnection "study/feature/postgres/simple_connection"
	"study/feature1"
	"study/feature2"
)

func main() {
	fmt.Println("daroy git")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}
