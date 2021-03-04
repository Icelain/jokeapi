# jokeapi-go
Official golang wrapper for Sv443's jokeapi.

Install-
```go get -u github.com/icelain/jokeapi```

Basic Usage Without Parameters-
```go
import "github.com/icelain/jokeapi"

func main(){
  api := jokeapi.New()
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
  
  api := jokeapi.New()
  api.SetParams(&ctgs, &blacklist, &jt)
  response := api.Get()
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
