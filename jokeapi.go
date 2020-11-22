package jokeapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const baseURL="https://sv443.net/jokeapi/v2/joke/"

type JokesResp struct{
	Error bool
	Category string
	JokeType string
	Joke []string
	Id float64
	Lang string
}


type Jokes struct{

	response map[string] interface{}

}
func (j *Jokes) Get() JokesResp{

resp, err:= http.Get(baseURL+"Any")
if err!=nil{
	panic(err)
}

info, err:= ioutil.ReadAll(resp.Body)
if err!=nil{
	panic(err)
}

json.Unmarshal(info, &j.response)

jo := []string{}

if j.response["type"].(string)=="single"{
	jo = append(jo, j.response["joke"].(string))

} else {
	jo = append(jo, j.response["setup"].(string), j.response["delivery"].(string))
}

return JokesResp{
	Error : j.response["error"].(bool),
	Category: j.response["category"].(string),
	JokeType : j.response["type"].(string),
	Joke : jo,
	Id : j.response["id"].(float64),
	Lang : j.response["lang"].(string),
}
}
