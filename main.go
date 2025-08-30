package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	fmt.Println("Numbers:", cutTheLastOne(numbers))
	numbers = numbers[1:]
	fmt.Println(numbers)
	server := gin.Default()

	server.GET("/docker", dockerGet)
	server.GET("/bocker", redirect)

	server.Run(":8080")
}

func cutTheLastOne(numbers []int) []int {

	numbers = numbers[:len(numbers)-1]

	if len(numbers) == 2 {
		fmt.Println("I can`t cut you anymore. You are too short!")
		fmt.Println("")
		return numbers
	}
	numbers = cutTheLastOne(numbers)
	return numbers
}

func dockerGet(ctx *gin.Context) {
	ctx.Writer.Write([]byte(
		`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Learn Docker</title>
</head>
<body>
	<h1 style="color: blue;">Docker is a crazy tool!</h1>
	 <a href="http://localhost:8080/bocker">
            <button style="color: blue">Get started with Docker</button></a
        >
</body>
</html>`))
}

func redirect(ctx *gin.Context) {
	ctx.Writer.Write([]byte(
		`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Learn Docker</title>
</head>
<body>
	<h1 style="color: blue;">Docker stil is a crazy tool!</h1>
</body>
</html>`))
}
