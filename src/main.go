package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type AutoGenerated struct {
	Data struct {
		Category        string      `json:"category"`
		CommonLocations interface{} `json:"common_locations"`
		Description     string      `json:"description"`
		Drops           []string    `json:"drops"`
		ID              int         `json:"id"`
		Image           string      `json:"image"`
		Name            string      `json:"name"`
	} `json:"data"`
}

type BOTW struct {
	Data struct {
		Creatures struct {
			Food []struct {
				Category        string   `json:"category"`
				CommonLocations []string `json:"common_locations"`
				CookingEffect   string   `json:"cooking_effect"`
				Description     string   `json:"description"`
				HeartsRecovered int      `json:"hearts_recovered"`
				ID              int      `json:"id"`
				Image           string   `json:"image"`
				Name            string   `json:"name"`
			} `json:"food"`
			NonFood []struct {
				Category        string   `json:"category"`
				CommonLocations []string `json:"common_locations"`
				Description     string   `json:"description"`
				Drops           []string `json:"drops"`
				ID              int      `json:"id"`
				Image           string   `json:"image"`
				Name            string   `json:"name"`
			} `json:"non_food"`
		} `json:"creatures"`
		Equipment []struct {
			Attack          int      `json:"attack"`
			Category        string   `json:"category"`
			CommonLocations []string `json:"common_locations"`
			Defense         int      `json:"defense"`
			Description     string   `json:"description"`
			ID              int      `json:"id"`
			Image           string   `json:"image"`
			Name            string   `json:"name"`
		} `json:"equipment"`
		Materials []struct {
			Category        string   `json:"category"`
			CommonLocations []string `json:"common_locations"`
			CookingEffect   string   `json:"cooking_effect"`
			Description     string   `json:"description"`
			HeartsRecovered int      `json:"hearts_recovered"`
			ID              int      `json:"id"`
			Image           string   `json:"image"`
			Name            string   `json:"name"`
		} `json:"materials"`
		Monsters []struct {
			Category        string   `json:"category"`
			CommonLocations []string `json:"common_locations"`
			Description     string   `json:"description"`
			Drops           []string `json:"drops"`
			ID              int      `json:"id"`
			Image           string   `json:"image"`
			Name            string   `json:"name"`
		} `json:"monsters"`
		Treasure []struct {
			Category        string   `json:"category"`
			CommonLocations []string `json:"common_locations"`
			Description     string   `json:"description"`
			Drops           []string `json:"drops"`
			ID              int      `json:"id"`
			Image           string   `json:"image"`
			Name            string   `json:"name"`
		} `json:"treasure"`
	} `json:"data"`
}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", OpenPage(w,r))
	http.ListenAndServe(":8080", nil)
	fmt.Println(moblin())
}

func OpenPage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := AutoGenerated{
		Image: ,
	}
	tmp.Execute(w, details)
}

func moblin() AutoGenerated {
	url := "https://botw-compendium.herokuapp.com/api/v2/master_mode/entry/golden_bokoblin"

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	JsonFile := AutoGenerated{}
	Jsonerr := json.Unmarshal(body, &JsonFile)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
	return JsonFile

}
