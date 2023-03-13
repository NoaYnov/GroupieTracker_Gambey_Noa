package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Monster struct {
	Category        string   `json:"category"`
	CommonLocations []string `json:"common_locations"`
	Description     string   `json:"description"`
	Drops           []string `json:"drops"`
	ID              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Type            string
}

type MonsterRequest struct {
	Monsters   []Monster `json:"data"`
	Formulaire string
	Capa       string
	Typ        string
}

func main() {

	var b MonsterRequest
	b.Init()
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", b.OpenPage)
	http.ListenAndServe(":8080", nil)
}

func (b *MonsterRequest) OpenPage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := MonsterRequest{
		Formulaire: r.FormValue("nom"),
		Capa:       r.FormValue("information"),
		Monsters:   b.Monsters,
	}
	details.TypeMonsters()
	tmp.Execute(w, details)
}

func (b *MonsterRequest) Init() {
	url := "https://botw-compendium.herokuapp.com/api/v2/category/monsters"

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

	Jsonerr := json.Unmarshal(body, &b)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
}

func (b *MonsterRequest) TypeMonsters() {
	for _, v := range b.Monsters {
		if strings.Contains(v.Name, "fire") {
			v.Type = "Fire"

		}
	}
}
