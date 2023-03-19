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
}

type Material struct {
	Data []struct {
		Category        string   `json:"category"`
		CommonLocations []string `json:"common_locations"`
		CookingEffect   string   `json:"cooking_effect"`
		Description     string   `json:"description"`
		HeartsRecovered float64  `json:"hearts_recovered"`
		ID              int      `json:"id"`
		Image           string   `json:"image"`
		Name            string   `json:"name"`
		Type            string
	} `json:"data"`
	Capa string
}
type Equipement struct {
	Data []struct {
		Attack          int      `json:"attack"`
		Category        string   `json:"category"`
		CommonLocations []string `json:"common_locations"`
		Defense         int      `json:"defense"`
		Description     string   `json:"description"`
		ID              int      `json:"id"`
		Image           string   `json:"image"`
		Name            string   `json:"name"`
		Type            string
	} `json:"data"`
	Capa string
}

type Creature struct {
	Data struct {
		Food []struct {
			Category        string   `json:"category"`
			CommonLocations []string `json:"common_locations"`
			CookingEffect   string   `json:"cooking_effect"`
			Description     string   `json:"description"`
			HeartsRecovered int      `json:"hearts_recovered"`
			ID              int      `json:"id"`
			Image           string   `json:"image"`
			Name            string   `json:"name"`
			Type            string
		} `json:"food"`
		NonFood []struct {
			Category        string      `json:"category"`
			CommonLocations []string    `json:"common_locations"`
			Description     string      `json:"description"`
			Drops           interface{} `json:"drops"`
			ID              int         `json:"id"`
			Image           string      `json:"image"`
			Name            string      `json:"name"`
			Type            string
		} `json:"non_food"`
	} `json:"data"`
	Capa string
}

type Treasure struct {
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
	Capa string
}

type Recherche struct {
	Mob *MonsterRequest
	Mat *Material
	Equ *Equipement
	Cre *Creature
	Tre *Treasure
}

func main() {
	var e Equipement
	var b MonsterRequest
	var ma Material
	var c Creature
	var t Treasure
	var rech Recherche
	b.InitMob()
	ma.InitMat()
	e.InitEquip()
	c.InitCrea()
	t.InitTreasure()
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.Handle("/Script/", http.StripPrefix("/Script/", http.FileServer(http.Dir("Script"))))
	http.HandleFunc("/", OpenPageIndex)
	http.HandleFunc("/mob", b.OpenPageMob)
	http.HandleFunc("/item", ma.OpenPageItem)
	http.HandleFunc("/equipement", e.OpenPageEquip)
	http.HandleFunc("/creature", c.OpenPageCrea)
	http.HandleFunc("/treasure", t.OpenPageTrea)
	http.HandleFunc("/tout", rech.OpenPageRecherche)
	http.HandleFunc("/404", NotFoundHandler)
	http.ListenAndServe(":8080", nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "404.html")
}

func (rech *Recherche) OpenPageRecherche(w http.ResponseWriter, r *http.Request) {
	rech.InitRecherche()
	tmp := template.Must(template.ParseFiles("recherche.html"))
	details := Recherche{
		Mob: rech.Mob,
		Mat: rech.Mat,
		Equ: rech.Equ,
		Cre: rech.Cre,
		Tre: rech.Tre,
	}
	tmp.Execute(w, details)
}

func (rech *Recherche) InitRecherche() {
	url := "https://botw-compendium.herokuapp.com/api/v2/category/treasure"

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

	Jsonerr := json.Unmarshal(body, &rech.Tre)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}

	url = "https://botw-compendium.herokuapp.com/api/v2/category/creatures"

	timeClient = http.Client{
		Timeout: time.Second * 2,
	}
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr = timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	Jsonerr = json.Unmarshal(body, &rech.Cre)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}

	url = "https://botw-compendium.herokuapp.com/api/v2/category/equipment"

	timeClient = http.Client{
		Timeout: time.Second * 2,
	}
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr = timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	Jsonerr = json.Unmarshal(body, &rech.Equ)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}

	url = "https://botw-compendium.herokuapp.com/api/v2/category/materials"

	timeClient = http.Client{
		Timeout: time.Second * 2,
	}
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr = timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	Jsonerr = json.Unmarshal(body, &rech.Mat)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}

	url = "https://botw-compendium.herokuapp.com/api/v2/category/monsters"

	timeClient = http.Client{
		Timeout: time.Second * 2,
	}
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr = timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}

	body, readErr = ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	Jsonerr = json.Unmarshal(body, &rech.Mob)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
}

func (t *Treasure) OpenPageTrea(w http.ResponseWriter, r *http.Request) {

	tmp := template.Must(template.ParseFiles("treasure.html"))
	details := Treasure{
		Data: t.Data,
		Capa: r.FormValue("information"),
	}
	tmp.Execute(w, details)
}

func (t *Treasure) InitTreasure() {
	url := "https://botw-compendium.herokuapp.com/api/v2/category/treasure"

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

	Jsonerr := json.Unmarshal(body, &t)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
}

func (c *Creature) OpenPageCrea(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("creature.html"))
	details := Creature{
		Data: c.Data,
		Capa: r.FormValue("information"),
	}
	details.TypeCreature()

	tmp.Execute(w, details)
}

func (c *Creature) TypeCreature() {
	for i := 0; i < len(c.Data.Food); i++ {
		if strings.Contains(c.Data.Food[i].Name, "hearty") {
			c.Data.Food[i].Type = "hearty"
		} else if strings.Contains(c.Data.Food[i].Name, "mighty") {
			c.Data.Food[i].Type = "mighty"
		} else {
			c.Data.Food[i].Type = "undefined"
		}
	}
	for i := 0; i < len(c.Data.NonFood); i++ {
		if strings.Contains(c.Data.NonFood[i].Name, "horse") {
			c.Data.NonFood[i].Type = "horse"
		} else if strings.Contains(c.Data.NonFood[i].Name, "goat") {
			c.Data.NonFood[i].Type = "goat"
		} else if strings.Contains(c.Data.NonFood[i].Name, "sparrow") {
			c.Data.NonFood[i].Type = "sparrow"
		} else {

			c.Data.NonFood[i].Type = "undefined"
		}
	}
}

func (c *Creature) InitCrea() {
	url := "https://botw-compendium.herokuapp.com/api/v2/category/creatures"

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

	Jsonerr := json.Unmarshal(body, &c)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
}

func (e *Equipement) TypeEquipement() {
	for i := 0; i < len(e.Data); i++ {
		if strings.Contains(e.Data[i].Name, "shield") {
			e.Data[i].Type = "shield"
		} else if strings.Contains(e.Data[i].Name, "bow") {
			e.Data[i].Type = "bow"
		} else if strings.Contains(e.Data[i].Name, "spear") || strings.Contains(e.Data[i].Name, "harpoon") {
			e.Data[i].Type = "spear"
		} else if strings.Contains(e.Data[i].Name, "arrow") {
			e.Data[i].Type = "arrow"
		} else if strings.Contains(e.Data[i].Name, "sword") || strings.Contains(e.Data[i].Name, "blade") || strings.Contains(e.Data[i].Name, "cleaver") {
			e.Data[i].Type = "sword"
		} else if strings.Contains(e.Data[i].Name, "axe") {
			e.Data[i].Type = "axe"
		} else if strings.Contains(e.Data[i].Name, "boomerang") {
			e.Data[i].Type = "boomerang"
		} else if strings.Contains(e.Data[i].Name, "rod") {
			e.Data[i].Type = "rod"
		} else if strings.Contains(e.Data[i].Name, "crusher") {
			e.Data[i].Type = "crusher"
		}
	}
}

func (e *Equipement) OpenPageEquip(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("equipement.html"))
	details := Equipement{
		Data: e.Data,
		Capa: r.FormValue("information"),
	}
	details.TypeEquipement()
	tmp.Execute(w, details)
}

func (e *Equipement) InitEquip() {
	url := "https://botw-compendium.herokuapp.com/api/v2/category/equipment"

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

	Jsonerr := json.Unmarshal(body, &e)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
}

func (ma *Material) OpenPageItem(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("item.html"))
	details := Material{
		Data: ma.Data,
		Capa: r.FormValue("information"),
	}
	details.TypeItem()
	tmp.Execute(w, details)
}

func OpenPageIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	tmp := template.Must(template.ParseFiles("index.html"))
	tmp.Execute(w, nil)
}
func (ma *Material) InitMat() {
	url := "https://botw-compendium.herokuapp.com/api/v2/category/materials"

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

	Jsonerr := json.Unmarshal(body, &ma)
	if Jsonerr != nil {
		fmt.Println(Jsonerr)
	}
}
func (ma *Material) TypeItem() {
	for i := 0; i < len(ma.Data); i++ {
		if strings.Contains(ma.Data[i].Name, "hearty") {
			ma.Data[i].Type = "hearty"
		} else if strings.Contains(ma.Data[i].Name, "silent") {
			ma.Data[i].Type = "silent"
		} else if strings.Contains(ma.Data[i].Name, "mighty") {
			ma.Data[i].Type = "mighty"
		} else if strings.Contains(ma.Data[i].Name, "endura") {
			ma.Data[i].Type = "endura"
		} else if strings.Contains(ma.Data[i].Name, "swift") {
			ma.Data[i].Type = "swift"
		}
	}
}

func (b *MonsterRequest) OpenPageMob(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("mob.html"))
	details := MonsterRequest{
		Formulaire: r.FormValue("nom"),
		Capa:       r.FormValue("information"),
		Monsters:   b.Monsters,
	}
	details.TypeMonsters()
	tmp.Execute(w, details)
}

func (b *MonsterRequest) InitMob() {
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
		} else if strings.Contains(b.Monsters[i].Name, "guardian") || strings.Contains(b.Monsters[i].Name, "sentry") {
			b.Monsters[i].Type = "guardian"
		} else if strings.Contains(b.Monsters[i].Name, "stal") {
			b.Monsters[i].Type = "stal"
		}

	}
}
