package jokeapi

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

var baseURL string="https://sv443.net/jokeapi/v2/joke/"

//Parameters to be accepted by a function
type Params struct{
	Categories *[]string
	Blacklist *[]string
	JokeType *string
}

//Response returned to the user
type JokesResp struct{
	Error bool
	Category string
	JokeType string
	Joke []string
	Flags map[string] bool
	Id float64
	Lang string
}

//Main struct
type JokeAPI struct{

	response map[string] interface{}

}
//Get function. Accepts a Params struct
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
	//Sends request
	resp, err:= http.Get(mainURL)
	if err!=nil{
		panic(err)
	}
	//Converts to readable bytecode
	info, err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}
	//Unmarshalls json content to map[string] interface
	json.Unmarshal(info, &j.response)

	jo := []string{}
	
	//Checks if joke type is single or twopart and saves to string slice accordingly
	if j.response["type"].(string)=="single"{
		jo = append(jo, j.response["joke"].(string))

	} else {
		jo = append(jo, j.response["setup"].(string), j.response["delivery"].(string))
	}
	//Convert flags response to a map[string]interface type
	flagInterface := j.response["flags"].(map[string]interface{})
	
	//Converts flagInterface to map[string] bool
	flags := map[string]bool{
		"nsfw": flagInterface["nsfw"].(bool),
		"religious": flagInterface["religious"].(bool),
		"racist" : flagInterface["racist"].(bool),
		"sexist": flagInterface["sexist"].(bool),
		"political": flagInterface["political"].(bool),
	}

	
	//Returns response
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

