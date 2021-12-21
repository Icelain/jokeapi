package jokeapi

import (
	"testing"
)

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

func Test_Fetch_SetAll(t *testing.T) {

	api := New()
	api.SetBlacklist([]string{"nsfw"})
	api.SetLang("de")

	resp, err := api.Fetch()

	if err != nil {

		t.Fatal(err)

	}

	if resp.Flags["nsfw"] {

		t.Error("blacklist flag not set properly")

	}

	if resp.Lang != "de" {

		t.Error("language flag not set properly")

	}
}
