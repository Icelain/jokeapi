package jokeapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL string = "https://v2.jokeapi.dev/joke/"
)

type jokeConsumer struct {
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

// Params is the config struct be used by JokeAPI{}.Fetch()
type Params struct {
	Categories []string
	Blacklist  []string
	JokeType   string
	Lang       string
}

// JokesResp is the response to be sent by JokeAPI{}.Fetch()
type JokesResp struct {
	Error    bool
	Category string
	JokeType string
	Joke     []string
	Flags    map[string]bool
	Id       float64
	Lang     string
}

// JokeAPI struct
type JokeAPI struct {
	ExportedParams Params
}

func contextifyError(context string, err error) error {
	return fmt.Errorf("%s: %w", context, err)
}

// FetchWithContext gets the content with respect to the parameters. Accepts custom context from user.
func (j *JokeAPI) FetchWithContext(ctx context.Context) (JokesResp, error) {

	var (
		//response = map[string]interface{}{}
		jokeConsumer jokeConsumer
		mainURL      string
		reqUrl       *url.URL
		err          error
	)

	//param handling begins here
	if len(j.ExportedParams.Categories) > 0 {
		mainURL = baseURL + strings.Join(j.ExportedParams.Categories, ",")
	} else {
		mainURL = baseURL + "Any"
	}

	reqUrl, err = url.Parse(mainURL)
	if err != nil {

		if err != nil {
			return JokesResp{}, contextifyError("Request failed as url could not be generated", err)
		}

	}

	query := reqUrl.Query()

	query.Set("blacklistFlags", strings.Join(j.ExportedParams.Blacklist, ","))
	query.Set("type", j.ExportedParams.JokeType)
	query.Set("lang", j.ExportedParams.Lang)

	reqUrl.RawQuery = query.Encode()

	// param handling ends here
	
	// create client
	client := &http.Client{}
	
	// create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {

		return JokesResp{}, err

	}

	resp, err := client.Do(req)
	if err != nil {
		return JokesResp{}, contextifyError("Request failed", err)
	}

	info, err := io.ReadAll(resp.Body)
	if err != nil {
		return JokesResp{}, contextifyError("Failed to decode request response", err)
	}

	if err = json.Unmarshal(info, &jokeConsumer); err != nil {

		return JokesResp{}, contextifyError("no joke found for your configuration", err)

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

// Fetch gets the content with respect to the parameters. Use FetchWithContext to add your custom context.
func (j *JokeAPI) Fetch() (JokesResp, error) {

	return j.FetchWithContext(context.TODO())

}

// SetParams sets parameters to JokeAPI struct instance. This method only exists because I don't want to make breaking changes to the existing api by removing it. I would recommend using Jokeapi{}.Set() or the singular methods instead
func (j *JokeAPI) SetParams(ctgs []string, blacklist []string, joketype string, lang string) {

	j.ExportedParams.Categories = ctgs
	j.ExportedParams.Blacklist = blacklist
	j.ExportedParams.JokeType = joketype
	j.ExportedParams.Lang = lang

}

// Set sets custom Params struct
func (j *JokeAPI) Set(params Params) {

	j.ExportedParams = params
}

// SetCategories sets joke categories. Common categories are Programming | Misc | Spooky | Dark | Fun
func (j *JokeAPI) SetCategories(ctgs []string) {

	j.ExportedParams.Categories = ctgs

}

// SetBlacklist sets joke blacklist. Common blacklists are nsfw | religious | political | racist | sexist | explicit
func (j *JokeAPI) SetBlacklist(b []string) {

	j.ExportedParams.Blacklist = b

}

// SetLang sets language. Go to https://v2.jokeapi.dev/languages?format=txt to select your preferable language format. By default its en (English). Note that (as of now) most jokes are available in en and de only and setting other languages will give a corresponding error
func (j *JokeAPI) SetLang(lang string) {

	j.ExportedParams.Lang = lang
}

// SetJokeType sets joke type
func (j *JokeAPI) SetJokeType(s string) {

	j.ExportedParams.JokeType = s

}

// New Generates instance of JokeAPI struct
func New() *JokeAPI {
	return &JokeAPI{Params{}}
}
