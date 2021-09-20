# jokeapi-go
Official golang wrapper for Sv443's jokeapi.

 [![GoDoc](https://godoc.org/github.com/icelain?status.png)](https://godoc.org/github.com/icelain/jokeapi)

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
  
  api:= jokeapi.New()
  
  api.Set(jokeapi.Params{Blacklist: blacklist, JokeType: jt, Categories: ctgs})
  response, err := api.Fetch()
}

```
Usage with config struct-
```go
api.Set(jokeapi.Params{})
```
Or-
```go
api.SetBlacklist(blacklist)
api.SetCategory(ctgs)
api.SetJokeType(joketype)
api.SetLang(language)
```
