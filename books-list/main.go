package main
 
import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gin-gonic/gin"
)
 
func main() {
    router := gin.Default()
    router.GET("/books/", book_list)
    router.Run(":80")
}


var response_json = `
[
    {
        "name":"Книга 1 название",
        "title":"Книга 1 описание"
    },
    {
        "name":"Книга 2 название",
        "title":"Книга 2 описание"
    }
]
`

type book struct {
    Name    string  `json:"name"`
    Title   string  `json:"title"`
}


type books []book


func book_list(c *gin.Context) {
    response_body := []byte(response_json)

    response_result := books{}

    if err := json.Unmarshal(response_body, &response_result); err != nil {
        panic(err)
    }

    fmt.Println(response_result)

    c.JSON(http.StatusOK, response_result)
}
