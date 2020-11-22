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
	Flags map[string] bool
	Id float64
	Lang string
}


type JokeAPI struct{

	response map[string] interface{}

}
func (j *JokeAPI) Get() JokesResp{

resp, err:= http.Get(baseURL+"Any")
if err!=nil{
	panic(err)
}

info, err:= ioutil.ReadAll(resp.Body)
if err!=nil{
	panic(err)
}

json.Unmarshal(info, &j.response)
flagInterface := j.response["flags"].(map[string]interface{})

flags := map[string]bool{
	"nsfw": flagInterface["nsfw"].(bool),
	"religious": flagInterface["religious"].(bool),
	"racist" : flagInterface["racist"].(bool),
	"sexist": flagInterface["sexist"].(bool),
	"political": flagInterface["political"].(bool),
}

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
	Flags : flags,
	Id : j.response["id"].(float64),
	Lang : j.response["lang"].(string),
}
}

