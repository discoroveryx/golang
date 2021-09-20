package main
 
import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gin-gonic/gin"
)


func CORSMiddlware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    }
}


func main() {
    router := gin.New()
    router.Use(CORSMiddlware())
    router.GET("/books/", book_list)
    router.Run(":80")
}


var response_json = `
[
    {
        "id":1,
        "name":"Книга 1 название",
        "title":"Книга 1 описание"
    },
    {
        "id":2,
        "name":"Книга 2 название",
        "title":"Книга 2 описание"
    },
    {
        "id":3,
        "name":"Книга 3 название",
        "title":"Книга 3 описание"
    }
]
`

type book struct {
    Id      int     `json:"id"`
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

    // headers := httpHeaders{}
    // c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

    c.JSON(http.StatusOK, response_result)
}
