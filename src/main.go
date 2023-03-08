package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Bokoblin struct {
	Data []struct {
		Category        string   `json:"category"`
		CommonLocations []string `json:"common_locations"`
		Description     string   `json:"description"`
		Drops           []string `json:"drops"`
		ID              int      `json:"id"`
		Image           string   `json:"image"`
		Name            string   `json:"name"`
	} `json:"data"`
	Bokob      *Bokoblin
	Formulaire string
	Type_mob   []string
}

func main() {

	var b *Bokoblin
	// fs := http.FileServer(http.Dir("css"))
	// http.Handle("/css/", http.StripPrefix("/css/", fs))
	// http.HandleFunc("/", b.OpenPage)
	// http.ListenAndServe(":8080", nil)
	fmt.Println(b.Create(b.type_mob()))
}

func (b *Bokoblin) OpenPage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	details := Bokoblin{
		Bokob:      b.Boko(),
		Formulaire: r.FormValue("nom"),
		Type_mob:   b.type_mob(),
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

func (b *Bokoblin) type_mob() []string {
	var result []string
	for _, v := range b.Boko().Data {
		if strings.Contains(v.Name, "fire") {
			result = append(result, "https://botw-compendium.herokuapp.com/api/v2/entry/"+strconv.Itoa(v.ID))
		}
	}
	return result
}

func (b *Bokoblin) Create(links []string) {
	// Mettre ici la liste des liens JSON à récupérer
	urls := links

	var combinedData []*Bokoblin

	for _, url := range urls {
		// Récupération du contenu JSON
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Erreur lors de la récupération du lien %s : %s\n", url, err)
			continue
		}
		defer resp.Body.Close()

		// Lecture du contenu JSON
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du contenu JSON du lien %s : %s\n", url, err)
			continue
		}

		// Décodage JSON
		var data []*Bokoblin
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Printf("Erreur lors du décodage JSON du lien %s : %s\n", url, err)
			continue
		}

		// Ajout des données décodées à la structure combinée
		combinedData = append(combinedData, data...)
	}

	// Faire quelque chose avec la structure combinée
	fmt.Printf("Nombre total d'éléments : %d\n", len(combinedData))
}
