package jokeapi

import (
	"context"
	"testing"
)

const (
	WRONG_BLACKLIST = "blacklist not set properly"
	WRONG_LANG      = "language not set properly"
	WRONG_TYPE      = "joke type not set properly"
	WRONG_CATEGORY  = "category not set properly"
)

func Test_Fetch_Context(t *testing.T) {

	api := New()
	_, err := api.FetchWithContext(context.Background())

	if err != nil {

		t.Fatal(err)

	}
}

func Test_Fetch_Parts(t *testing.T) {

	api := New()
	resp, err := api.Fetch()

	if err != nil {

		t.Fatal(err)

	}

	if len(resp.Joke) == 1 {

		if !(resp.JokeType == "single" && resp.Joke[0] != "") {

			t.FailNow()

		}
	}

	if len(resp.Joke) == 2 {

		if !(resp.JokeType == "twopart" && resp.Joke[0] != "" && resp.Joke[1] != "") {
			t.FailNow()

		}
	}
}

func Test_Fetch_Set_Functional(t *testing.T) {

	api := New()
	api.SetBlacklist([]string{"nsfw"})
	api.SetLang("de")
	api.SetJokeType("twopart")

	resp, err := api.Fetch()

	if err != nil {

		t.Fatal(err)

	}

	if resp.Flags["nsfw"] {

		t.Error(WRONG_BLACKLIST)

	}

	if resp.Lang != "de" {

		t.Error(WRONG_LANG)

	}

	if resp.JokeType != "twopart" {

		t.Error(WRONG_TYPE)

	}

}

func Test_Fetch_Type(t *testing.T) {

	api := New()
	api.SetCategories([]string{"Programming"})

	resp, err := api.Fetch()

	if err != nil {

		t.Fatal(err)

	}

	if resp.Category != "Programming" {

		t.Error(WRONG_CATEGORY)

	}

}

func Test_Fetch_Set_All(t *testing.T) {

	api := New()
	api.SetParams([]string{"Programming"}, []string{"nsfw"}, "twopart", "en")

	resp, err := api.Fetch()

	if err != nil {

		t.Fatal(err)

	}

	if resp.JokeType != "twopart" {

		t.Error(WRONG_TYPE)

	}

	if resp.Flags["nsfw"] {

		t.Error(WRONG_BLACKLIST)

	}

	if resp.Category != "Programming" {

		t.Error(WRONG_CATEGORY)

	}

	if resp.Lang != "en" {

		t.Error(WRONG_LANG)

	}

}

func Test_Fetch_Params(t *testing.T) {

	api := New()
	api.Set(Params{Categories: []string{"Programming"}, Blacklist: []string{"nsfw"}, JokeType: "single"})

	resp, err := api.Fetch()

	if err != nil {

		t.Fatal(err)

	}

	if resp.Category != "Programming" {

		t.Error()

	}

}
