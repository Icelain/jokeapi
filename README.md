# jokeapi-go

 [![GoDoc](https://godoc.org/github.com/icelain?status.png)](https://godoc.org/github.com/icelain/jokeapi)
 [![Go Report Card](https://goreportcard.com/badge/github.com/icelain/jokeapi)](https://goreportcard.com/report/github.com/icelain/jokeapi)
 [![Test Coverage](https://gocover.io/_badge/github.com/icelain/jokeapi)](https://gocover.io/github.com/icelain/jokeapi)
 [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

Official golang wrapper for Sv443's jokeapi.

Install-
```go get -u github.com/icelain/jokeapi```

Basic Usage Without Parameters-
```go
import "github.com/icelain/jokeapi"

func main(){
  api := jokeapi.New()
  response, err := api.Fetch()
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
Config struct-
```go
api.Set(jokeapi.Params{})
```
Functional config -
```go
api.SetBlacklist(blacklist)
api.SetCategories(ctgs)
api.SetJokeType(joketype)
api.SetLang(language)
```
