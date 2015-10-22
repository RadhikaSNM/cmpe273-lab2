/*
Radhika SNM
009426196
*/
package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    
)


type User struct {
    Name string `json:"name"`
}


 type UserGreeting struct {
    Greeting string `json:"greeting"`
}




func getHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func postHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params){
    
    u:=User{}

    //decode the sent json
    err1 :=json.NewDecoder(req.Body).Decode(&u)
    if err1!=nil{
        fmt.Print("Error occcured in decoding")
    }

    name:=u.Name;
    fmt.Print("Name is :"+name)
    greet:="Hello,"+name+"!"
        
    //creating the greeting obj
    gr:=UserGreeting{greet}
    //marshalling into a json

    grJson, err := json.Marshal(gr)
    if err!=nil{
        fmt.Print("Error occcured in marshalling")
    }

    //sending it in the response
    rw.Header().Set("Content-Type","application/json")
    rw.WriteHeader(http.StatusOK)
    fmt.Fprintf(rw, "%s", grJson)

}


func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", getHello)
    mux.POST("/hello",postHello)

    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }

    server.ListenAndServe()
}