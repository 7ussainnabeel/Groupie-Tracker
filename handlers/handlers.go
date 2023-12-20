package handlers

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var arr_cities [][]string

func Index(w http.ResponseWriter, r *http.Request) {
	searchInput := strings.TrimSpace(r.FormValue("searchInput"))
	searchInputLower := strings.ToLower(searchInput)
	fromCareerStartDate, e1 := strconv.Atoi(r.FormValue("fromCareerStartDate"))
	if e1 != nil {
		fromCareerStartDate = 1900
	}
	toCareerStartDate, e2 := strconv.Atoi(r.FormValue("toCareerStartDate"))
	if e2 != nil {
		toCareerStartDate = 2030
	}
	if fromCareerStartDate > toCareerStartDate {
		newVar := fromCareerStartDate
		fromCareerStartDate = toCareerStartDate
		toCareerStartDate = newVar
	}
	fromFirstAlbumDate, e3 := strconv.Atoi(r.FormValue("fromFirstAlbumDate"))
	if e3 != nil {
		fromFirstAlbumDate = 1900
	}
	toFirstAlbumDate, e2 := strconv.Atoi(r.FormValue("toFirstAlbumDate"))
	if e2 != nil {
		toFirstAlbumDate = 2030
	}
	if fromFirstAlbumDate > toFirstAlbumDate {
		newVar := fromFirstAlbumDate
		fromFirstAlbumDate = toFirstAlbumDate
		toFirstAlbumDate = newVar
	}
	member1 := strings.TrimSpace(r.FormValue("1member")) == "on"
	member2 := strings.TrimSpace(r.FormValue("2member")) == "on"
	member3 := strings.TrimSpace(r.FormValue("3member")) == "on"
	member4 := strings.TrimSpace(r.FormValue("4member")) == "on"
	member5 := strings.TrimSpace(r.FormValue("5member")) == "on"
	member6 := strings.TrimSpace(r.FormValue("6member")) == "on"
	member7 := strings.TrimSpace(r.FormValue("7member")) == "on"
	member8 := strings.TrimSpace(r.FormValue("8member")) == "on"
	location := r.FormValue("location")

	if r.URL.Path != "/" {
		Errorshandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorshandler(w, http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("./template/index.html")
	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	band, errr := JsonArtists()

	if errr != nil {
		Errorshandler(w, http.StatusInternalServerError)
		return
	}

	var suggestions []string

	locations, err2 := JsonLocations()
	if err2 != nil {
		Errorshandler(w, http.StatusInternalServerError)
	}
	var new_locations, res_locations []string
	var arr [][]reflect.Value
	for _, v := range locations.Index {
		arr = append(arr, reflect.ValueOf(v.DatesLocations).MapKeys())
	}

	var filteredBand []Artist
	for i, v := range band {
		searchLocation := false
		filterLocation := false
		isContain := false
		searchInputLowerArray := strings.Split(searchInputLower, "-")
		var array []string
		if !Contains(suggestions, v.Name) {
			suggestions = append(suggestions, v.Name)
		}
		for _, vv := range v.Members {
			suggestion := vv + " - " + v.Name
			if !Contains(suggestions, suggestion) && v.Name != vv {
				suggestions = append(suggestions, suggestion)
			}
			if strings.Contains(strings.ToLower(vv), searchInputLower) || ContainsSuggestions(searchInputLowerArray, strings.ToLower(vv)) {
				isContain = true
			}
		}
		for _, vv := range arr[i] {
			vv2 := strings.ReplaceAll(vv.String(), "_", " ")
			vv2 = strings.ReplaceAll(vv2, "-", ", ")
			if strings.Contains(strings.ToLower(vv2), strings.ToLower(searchInput)) || strings.Contains(strings.ToLower(vv.String()), strings.ToLower(searchInput)) {
				searchLocation = true
			}
			if strings.ToLower(location) == vv2 || location == "" {
				filterLocation = true
			}
			array = append(array, vv2)
			vv2 = strings.Title(vv2)
			replacer := strings.NewReplacer("Uk", "UK", "Usa", "USA")
			res_vv := replacer.Replace(vv2)
			country_only := strings.Split(res_vv, ", ")
			if !Contains(suggestions, country_only[1]) {
				suggestions = append(suggestions, country_only[1])
			}
			if !Contains(suggestions, res_vv) {
				suggestions = append(suggestions, res_vv)
			}
			new_locations = append(new_locations, res_vv)

		}
		arr_cities = append(arr_cities, array)
		firstAlbum, _ := strconv.Atoi(strings.Split(v.FirstAlbum, "-")[2])
		if !Contains(suggestions, v.FirstAlbum) {
			suggestions = append(suggestions, v.FirstAlbum)
		}
		if !Contains(suggestions, strconv.FormatInt(int64(v.CreationDate), 10)) {
			suggestions = append(suggestions, strconv.FormatInt(int64(v.CreationDate), 10))
		}
		if filterLocation && fromCareerStartDate <= v.CreationDate && toCareerStartDate >= v.CreationDate && fromFirstAlbumDate <= firstAlbum && toFirstAlbumDate >= firstAlbum && ((!member1 && !member2 && !member3 && !member4 && !member5 && !member6 && !member7 && !member8) || (member1 && len(v.Members) == 1) || (member2 && len(v.Members) == 2) || (member3 && len(v.Members) == 3) || (member4 && len(v.Members) == 4) || (member5 && len(v.Members) == 5) || (member6 && len(v.Members) == 6) || (member7 && len(v.Members) == 7) || (member8 && len(v.Members) == 8)) {
			if searchLocation || strings.Contains(strings.ToLower(v.Name), searchInputLower) || v.FirstAlbum == searchInputLower || strings.Contains(strconv.FormatInt(int64(v.CreationDate), 10), searchInputLower) {
				filteredBand = append(filteredBand, v)
			} else {
				if isContain {
					filteredBand = append(filteredBand, v)
				}
			}
		}
	}
	for _, v := range new_locations {
		if !Contains(res_locations, v) {
			res_locations = append(res_locations, v)
		}
	}
	res := SearchInput{
		Group:               filteredBand,
		Places:              res_locations,
		SearchInput:         searchInput,
		FromCareerStartDate: fromCareerStartDate,
		ToCareerStartDate:   toCareerStartDate,
		FromFirstAlbumDate:  fromFirstAlbumDate,
		ToFirstAlbumDate:    toFirstAlbumDate,
		Member1:             member1,
		Member2:             member2,
		Member3:             member3,
		Member4:             member4,
		Member5:             member5,
		Member6:             member6,
		Member7:             member7,
		Member8:             member8,
		Location:            location,
		Suggestions:         suggestions,
	}
	err = template.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
}

func Artists(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path
	linkList := strings.Split(link, "/")
	id, err := strconv.Atoi(linkList[len(linkList)-1])
	if len(linkList) > 3 || linkList[1] != "artist" || (id <= 0 || id > 52) {
		Errorshandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorshandler(w, http.StatusMethodNotAllowed)
		return
	}
	template, err := template.ParseFiles("./template/artist.html")
	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	mainPage, err3 := JsonArtists()
	if err3 != nil {
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	concerts, err4 := JsonConcerts(strconv.Itoa(id))
	if err4 != nil {
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
	MapData := map[string][]string{}

	for key, value := range concerts.DatesLocations {
		key = strings.ReplaceAll(key, "_", " ")
		key = strings.ReplaceAll(key, "-", ", ")
		MapData[key] = value
	}
	res := AllData{
		Main:     mainPage[id-1],
		Concerts: MapData,
	}
	err = template.Execute(w, res)

	if err != nil {
		log.Println(err.Error())
		Errorshandler(w, http.StatusInternalServerError)
		return
	}
}

func Errorshandler(w http.ResponseWriter, status int) {
	template, err := template.ParseFiles("./template/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var Result Errors
	Result.Status = status
	Result.Message = http.StatusText(status)
	w.WriteHeader(status)
	err = template.Execute(w, Result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
