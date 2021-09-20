package jokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
        baseURL string = "https://sv443.net/jokeapi/v2/joke/"	
)

// Parameters to be used by JokeAPI{}.Fetch()
type Params struct {
	Categories []string
	Blacklist  []string
	JokeType   string
	Lang string
}

// Response to be sent by JokeAPI{}.Fetch()
type JokesResp struct {
	Error    bool
	Category string
	JokeType string
	Joke     []string
	Flags    map[string]bool
	Id       float64
	Lang     string
}

// Base JokeAPI struct
type JokeAPI struct {
	ExportedParams Params
}

// Fetches content with respect to the parameters
func (j *JokeAPI) Fetch() (JokesResp, error) {
	
	var (
		response = map[string]interface{}{}
		mainURL string
		isBlacklist bool
	)

	//param handling begins here
	if len(j.ExportedParams.Categories) > 0 {
		mainURL = baseURL + strings.Join(j.ExportedParams.Categories, ",")
	} else {
		mainURL = baseURL + "Any"
	}

	if len(j.ExportedParams.Blacklist) > 0{
		isBlacklist = true
		mainURL = mainURL + "?blacklistFlags=" + strings.Join(j.ExportedParams.Blacklist, ",")
	}

	if j.ExportedParams.JokeType == "" {
		if isBlacklist {
			mainURL = mainURL + "&type=" + j.ExportedParams.JokeType
		} else {
			mainURL = mainURL + "?type=" + j.ExportedParams.JokeType
		}
	}
	
	if j.ExportedParams.Lang != "" {

		mainURL += "?lang=" + j.ExportedParams.Lang
		
	}

	//param handling ends here
	resp, err := http.Get(mainURL)
	if err != nil {
		return JokesResp{}, err
	}

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return JokesResp{} ,err	
	}

	json.Unmarshal(info, &response)

	jo := []string{}

	if response["type"].(string) == "single" {
		jo = append(jo, response["joke"].(string))

	} else {
		jo = append(jo, response["setup"].(string), response["delivery"].(string))
	}

	flagInterface := response["flags"].(map[string]interface{})

	flags := map[string]bool{
		"nsfw":      flagInterface["nsfw"].(bool),
		"religious": flagInterface["religious"].(bool),
		"racist":    flagInterface["racist"].(bool),
		"sexist":    flagInterface["sexist"].(bool),
		"political": flagInterface["political"].(bool),
	}

	return JokesResp{
		Error:    response["error"].(bool),
		Category: response["category"].(string),
		JokeType: response["type"].(string),
		Joke:     jo,
		Flags:    flags,
		Id:       response["id"].(float64),
		Lang:     response["lang"].(string),
	}, nil
}

//Sets parameters to JokeAPI struct instance. This method only exists because I don't want to make breaking changes to the existing api by removing it. I would recommend using Jokeapi{}.Set() or the singular methods instead  
func (j *JokeAPI) SetParams(ctgs []string, blacklist []string, joketype string, lang string) {

	j.ExportedParams.Categories = ctgs
	j.ExportedParams.Blacklist = blacklist
	j.ExportedParams.JokeType = joketype
	j.ExportedParams.Lang = lang

}

// Sets custom Params struct
func (j *JokeAPI) Set(params Params) {

	j.ExportedParams = params
}

// Sets joke categories
func (j *JokeAPI) SetCategories(ctgs []string) {

	j.ExportedParams.Categories = ctgs

}

// Sets joke blacklist
func (j *JokeAPI) SetBlacklist(b []string) {

	j.ExportedParams.Blacklist = b

}
//Sets language. Go to https://v2.jokeapi.dev/languages?format=txt to select your preferable language format. By default its en (English).
func (j *JokeAPI) SetLang(lang string) {

	j.ExportedParams.Lang = lang
}

// Sets joke type
func (j *JokeAPI) SetJokeType(s string) {

	j.ExportedParams.JokeType = s

}

// Generates instance of JokeAPI struct
func New() *JokeAPI {
	return &JokeAPI{Params{}}
}
