package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/docker", dockerGet)
	server.GET("/bocker", redirect)

	server.Run(":8080")
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
	<a href="http://localhost:8080/docker">
            <button style="color: blue">Get back to home page</button></a
        >
</body>
</html>`))
}
