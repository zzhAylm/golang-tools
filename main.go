package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	//router := gin.Default()
	//router.GET("/albums", getAlbums)
	//router.GET("/albums/:id", getAlbumByID)
	//router.POST("/albums", postAlbums)
	//
	//router.Run("localhost:8080")

	dial, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	fmt.Println("连接成功，请输入要发送的数据")
	defer dial.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取控制台输入的数据失败")
			return
		}
		if readString == "exit\n" {
			fmt.Println("停止输入，断开连接")
			return
		}
		_, err = dial.Write([]byte(readString))
		if err != nil {
			fmt.Println("发送给服务端失败，断开连接")
			return
		}
		fmt.Printf("客户端发送成功，message is %s", readString)
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
