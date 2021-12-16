package jokeapi

import (
	"testing"
)

func Test_Fetch(t *testing.T) {

	api := New()
	api.Set(Params{Categories: []string{"Programming"}, Lang: "de"})
	resp, err := api.Fetch()

	if err != nil {

		t.Error(err)

	}
	if len(resp.Joke) == 1 {

		t.Logf("Joke: %s\nID: %.0f", resp.Joke[0], resp.Id)
		return
	}
	t.Logf("Setup: %s\nDelivery: %s\n ID: %.0f", resp.Joke[0], resp.Joke[1], resp.Id)
}
