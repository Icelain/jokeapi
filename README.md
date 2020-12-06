# jokeapi-go
Official golang wrapper for Sv443's jokeapi.

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
Response Struct-
```go
type JokesResp struct{
	Error bool
	Category string
	JokeType string
	Joke []string
	Flags map[string] bool
	Id float64
	Lang string
}



```
Params struct-
```go
type Params struct{
	Categories *[]string
	Blacklist *[]string
	JokeType *string
}
```

Usage with all parameters-
```go
import "github.com/icelain/jokeapi"

func main(){
  jt := "single"
  blacklist := []string{"nsfw"}
  ctgs := []string{"Programming","Dark"}
  
  api := new(jokeapi.JokeAPI)
  response := api.Get(jokeapi.Params{&ctgs, &blacklist, &jt})
}

```
Usage without all parameters(requires other params to be declared as nil)-
```go
jokeapi.Params{&ctgs, nil, nil}
```
