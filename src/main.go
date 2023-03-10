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
		Type            string
	} `json:"data"`
	Formulaire string
	Typ        *Mob
	Bokob      *Bokoblin
	Capa       string
}

type Mob struct {
	Data struct {
		Category        string   `json:"category"`
		CommonLocations []string `json:"common_locations"`
		Description     string   `json:"description"`
		Drops           []string `json:"drops"`
		ID              int      `json:"id"`
		Image           string   `json:"image"`
		Name            string   `json:"name"`
	} `json:"data"`
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
		Typ:        b.type_mob(),
		Capa:       r.FormValue("information"),
	}
	tmp.Execute(w, details)
	fmt.Println(r.FormValue("information"))
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

func (b *Bokoblin) type_mob() *Mob {
	var FinalResult *Mob
	var result []string
	for _, v := range b.Boko().Data {
		if strings.Contains(v.Name, "fire") {
			result = append(result, "https://botw-compendium.herokuapp.com/api/v2/entry/"+strconv.Itoa(v.ID))
		}
	}
	FinalResult = b.fusionnerFichiersJSON(result)
	// fmt.Println(result)
	// fmt.Println(FinalResult)
	return FinalResult
}

func (b *Bokoblin) fusionnerFichiersJSON(urls []string) *Mob {
	result := make(map[string]interface{})
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return nil
		}
		defer resp.Body.Close()
		var data map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil
		}
		for key, value := range data {
			result[key] = value
		}
		// fmt.Println(result)
	}
	bytesData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	finalResult := &Mob{}
	json.Unmarshal(bytesData, &finalResult)
	fmt.Println(finalResult)
	return finalResult
}
