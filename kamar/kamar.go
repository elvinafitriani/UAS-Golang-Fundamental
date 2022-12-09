package kamar

import (
	"encoding/json"
	"net/http"
	"projek/connection"
	"projek/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

var dbase *gorm.DB

func init() {
	dbase = connection.ConnectToDb()
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		type view struct {
			Room_id   int    `json:"room_id"`
			Room_nama string `json:"Room_name"`
			Room_spes string `json:"space"`
		}
		var data []view
		dbase.
			Model(&models.Room{}).
			Select("* from rooms").
			Scan(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Enconde to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Error Not Found", http.StatusNotFound)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var inp []string = strings.Split(url, "/")
		put, err := strconv.Atoi(inp[2])

		type view struct {
			Room_id   int    `json:"room_id"`
			Room_nama string `json:"Room_name"`
			Room_spes string `json:"space"`
		}
		var data []view
		dbase.
			Model(&models.Room{}).
			Select("* from rooms").
			Where(&models.Room{Room_id: put}).
			Scan(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Enconde to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Error Not Found", http.StatusNotFound)
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Room

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
			return
		}

		if err := dbase.Create(&data).Error; err != nil {
			http.Error(w, "Error Post", http.StatusNotAcceptable)
			return
		}

		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, "Error Not Found", http.StatusNotFound)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var data models.Room
		url := r.URL.String()
		var id []string = strings.Split(url, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		decode := json.NewDecoder(r.Body)
		errors := decode.Decode(&data)

		if errors != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err := dbase.First(&models.Room{}, "id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if err := dbase.Find(&models.Room{}, "id=?", id[2]).Updates(&data).Error; err != nil {
			http.Error(w, "Error Update", http.StatusNotAcceptable)
		}
		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}
