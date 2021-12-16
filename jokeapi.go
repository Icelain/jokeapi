package jokeapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	baseURL string = "https://v2.jokeapi.dev/joke/"
)

type JokeConsumer struct {
	Error    bool            `json:"error"`
	Category string          `json:"category"`
	Type     string          `json:"type"`
	ID       float64         `json:"id"`
	Lang     string          `json:"lang"`
	Flags    map[string]bool `json:"flags"`
	Joke     string          `json:"joke"`
	Setup    string          `json:"setup"`
	Delivery string          `json:"delivery"`
}

// Parameters to be used by JokeAPI{}.Fetch()
type Params struct {
	Categories []string
	Blacklist  []string
	JokeType   string
	Lang       string
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

func setSign(sign *string) {

	if *sign == "?" {
		*sign = "&"
	}
}

// Fetches content with respect to the parameters
func (j *JokeAPI) Fetch() (JokesResp, error) {

	var (
		//response = map[string]interface{}{}
		jokeConsumer JokeConsumer
		mainURL      string
		sign         string = "?"
	)

	//param handling begins here
	if len(j.ExportedParams.Categories) > 0 {
		mainURL = baseURL + strings.Join(j.ExportedParams.Categories, ",")
	} else {
		mainURL = baseURL + "Any"
	}

	if len(j.ExportedParams.Blacklist) > 0 {

		mainURL += sign + "blacklistFlags=" + strings.Join(j.ExportedParams.Blacklist, ",")

		setSign(&sign)
	}

	if j.ExportedParams.JokeType != "" {

		mainURL += sign + "type=" + j.ExportedParams.JokeType
		setSign(&sign)
	}

	if j.ExportedParams.Lang != "" {

		mainURL += sign + "lang=" + j.ExportedParams.Lang

	}

	//param handling ends here
	resp, err := http.Get(mainURL)
	if err != nil {
		return JokesResp{}, err
	}

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return JokesResp{}, err
	}

	if err = json.Unmarshal(info, &jokeConsumer); err != nil {

		return JokesResp{}, errors.New("no joke found for your configuration: ")

	}

	jo := []string{}

	if jokeConsumer.Type == "" {

		return JokesResp{}, errors.New("no joke found for your configuration")

	}

	if jokeConsumer.Type == "single" {
		jo = append(jo, jokeConsumer.Joke)

	} else {
		jo = append(jo, jokeConsumer.Setup, jokeConsumer.Delivery)
	}

	return JokesResp{
		Error:    jokeConsumer.Error,
		Category: jokeConsumer.Category,
		JokeType: jokeConsumer.Type,
		Joke:     jo,
		Flags:    jokeConsumer.Flags,
		Id:       jokeConsumer.ID,
		Lang:     jokeConsumer.Lang,
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

//Sets language. Go to https://v2.jokeapi.dev/languages?format=txt to select your preferable language format. By default its en (English). Most jokes are available in en and de only and setting other languages will give a corresponding error
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
