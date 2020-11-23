# jokeapi-go
Unofficial golang wrapper for Sv443's jokeapi.

Install-
```go get -u github.com/icelain/jokeapi```

Basic Usage Without Parameters-
```go
import "github.com/icelain/jokeapi"

func main(){
  api := new(jokeapi.JokeAPI)
  response := api.Get(jokeapi.Params{})
}
```
