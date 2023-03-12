package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

type Bokoblin struct {
	Monsters   []Monster `json:"data"`
	Bokob      *Bokoblin
	Formulaire string
	Capa       string
	Typ        *Bokoblin
}

func main() {

	var b *Bokoblin
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", b.OpenPage)
	http.ListenAndServe(":8080", nil)
	// fmt.Println(b.type_mob())
}

func (b *Bokoblin) OpenPage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := Bokoblin{
		Bokob:      b.Boko(),
		Formulaire: r.FormValue("nom"),
		// Typ:        b.type_mob(),
		Capa: r.FormValue("information"),
	}
	tmp.Execute(w, details)
}

func (b *Bokoblin) Boko() *Bokoblin {
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

	JsonFile := &Bokoblin{}
	Jsonerr := json.Unmarshal(body, &JsonFile)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}

	return JsonFile

}

// func (b *Bokoblin) type_mob() *Bokoblin {
// 	for _, v := range b.Monsters {
// 		if strings.Contains(v.Name, "fire") {
// 			v.Type = "feu"
// 		}
// 	}

// }
