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
	Type_bis        string
}

type MonsterRequest struct {
	Monsters   []Monster `json:"data"`
	Formulaire string
	Capa       string
}

func main() {

	var b MonsterRequest
	b.Init()
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", OpenPageIndex)
	http.HandleFunc("/mob", b.OpenPage)
	http.ListenAndServe(":8080", nil)
}
func OpenPageIndex(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	tmp.Execute(w, nil)
}

func (b *MonsterRequest) OpenPage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("mob.html"))
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
	for i := 0; i < len(b.Monsters); i++ {
		if strings.Contains(b.Monsters[i].Name, "fire") || strings.Contains(b.Monsters[i].Name, "igneo") || strings.Contains(b.Monsters[i].Name, "meteo") {
			b.Monsters[i].Type = "fire"
		} else if strings.Contains(b.Monsters[i].Name, "ice") || strings.Contains(b.Monsters[i].Name, "snow") || strings.Contains(b.Monsters[i].Name, "frost") || strings.Contains(b.Monsters[i].Name, "blizz") || strings.Contains(b.Monsters[i].Name, "waterblight") || strings.Contains(b.Monsters[i].Name, "octorok") {
			b.Monsters[i].Type = "ice"
		} else if strings.Contains(b.Monsters[i].Name, "electric") || strings.Contains(b.Monsters[i].Name, "thunder") {
			b.Monsters[i].Type = "electric"
		} else if strings.Contains(b.Monsters[i].Name, "cursed") {
			b.Monsters[i].Type = "cursed"
		} else if strings.Contains(b.Monsters[i].Name, "guardian") {
			b.Monsters[i].Type = "guardian"
		} else if strings.Contains(b.Monsters[i].Name, "stal") {
			b.Monsters[i].Type = "stal"
		}

	}
}
