# jokeapi-go
Official golang wrapper for Sv443's jokeapi.

 ![Go Reference](https://pkg.go.dev/github.com/icelain/jokeapi)

Install-
```go get -u github.com/icelain/jokeapi```

Basic Usage Without Parameters-
```go
import "github.com/icelain/jokeapi"

func main(){
  api, err := jokeapi.New()
  
  if err != nil {
  	panic(err)
  }
  
  response := api.Fetch()
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

Usage with all parameters-
```go
import "github.com/icelain/jokeapi"

func main(){
  jt := "single"
  blacklist := []string{"nsfw"}
  ctgs := []string{"Programming","Dark"}
  
  api, err := jokeapi.New()
  
  if err != nil {
  	panic(err)
  }
  
  api.SetParams(&ctgs, &blacklist, &jt)
  response := api.Fetch()
}

```
Usage without all parameters(requires other params to be declared as nil)-
```go
api.SetParams(&ctgs,nil, nil)
```
Or-
```go
api.SetBlacklist(&blacklist)
api.SetCategory(&ctgs)
api.SetType(&joketype)
```
