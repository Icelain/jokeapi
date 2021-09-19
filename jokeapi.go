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

type Params struct {
	Categories *[]string
	Blacklist  *[]string
	JokeType   *string
}

type JokesResp struct {
	Error    bool
	Category string
	JokeType string
	Joke     []string
	Flags    map[string]bool
	Id       float64
	Lang     string
}

type JokeAPI struct {
	ExportedParams Params
}

func (j *JokeAPI) Fetch() (JokesResp, error) {
	
	var (
		response = map[string]interface{}{}
		mainURL string
		isBlacklist bool = false

	)

	//param handling begins here
	if j.ExportedParams.Categories != nil {
		mainURL = baseURL + strings.Join(*j.ExportedParams.Categories, ",")
	} else {
		mainURL = baseURL + "Any"
	}

	if j.ExportedParams.Blacklist != nil {
		isBlacklist = true
		mainURL = mainURL + "?blacklistFlags=" + strings.Join(*j.ExportedParams.Blacklist, ",")
	}

	if j.ExportedParams.JokeType != nil {
		if isBlacklist {
			mainURL = mainURL + "&type=" + *j.ExportedParams.JokeType
		} else {
			mainURL = mainURL + "?type=" + *j.ExportedParams.JokeType
		}
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

//Sets parameters to JokeAPI struct instance
func (j *JokeAPI) SetParams(ctgs *[]string, blacklist *[]string, joketype *string) {

	j.ExportedParams.Categories = ctgs
	j.ExportedParams.Blacklist = blacklist
	j.ExportedParams.JokeType = joketype

}

func (j *JokeAPI) SetCategories(ctgs *[]string) {

	j.ExportedParams.Categories = ctgs

}

func (j *JokeAPI) SetBlacklist(b *[]string) {

	j.ExportedParams.Blacklist = b

}

func (j *JokeAPI) SetJokeType(s *string) {

	j.ExportedParams.JokeType = s

}

// Generates instance of JokeAPI struct
func New() *JokeAPI {
	return &JokeAPI{Params{}}
}
