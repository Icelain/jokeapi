package jokeapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

var baseURL string="https://sv443.net/jokeapi/v2/joke/"

type Params struct{
	Categories *[]string
	Blacklist *[]string
	JokeType *string
}

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
func (j *JokeAPI) Get(params Params) JokesResp{

	var mainURL string=""
	var isBlacklist bool=false

//param handling begins here
	if params.Categories !=nil{
		mainURL = baseURL + strings.Join(*params.Categories,",")
	} else{
		mainURL = baseURL + "Any"
	}

	if params.Blacklist!=nil{
		isBlacklist=true
		mainURL = mainURL + "?blacklistFlags=" + strings.Join(*params.Blacklist,",")
	}

	if params.JokeType !=nil{
		if isBlacklist{
		mainURL=mainURL +"&type="+ *params.JokeType
	} else{
		mainURL=mainURL +"?type="+ *params.JokeType
	}
	}
	
//param handling ends here
	resp, err:= http.Get(mainURL)
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

	flagInterface := j.response["flags"].(map[string]interface{})

	flags := map[string]bool{
		"nsfw": flagInterface["nsfw"].(bool),
		"religious": flagInterface["religious"].(bool),
		"racist" : flagInterface["racist"].(bool),
		"sexist": flagInterface["sexist"].(bool),
		"political": flagInterface["political"].(bool),
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

