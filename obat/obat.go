package obat

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
			Med_id    int    `json:"medicine_id"`
			Med_nama  string `json:"Medicine_name"`
			Med_stok  int    `json:"stok"`
			Med_rules string `json:"aturan"`
			Med_desc  string `json:"deskripsi"`
			Med_exp   int    `json:"med_exp"`
		}
		var data []view
		dbase.
			Model(&models.Medicine{}).
			Select("* from medicines").
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
			Med_id    int    `json:"medicine_id"`
			Med_nama  string `json:"Medicine_name"`
			Med_stok  int    `json:"stok"`
			Med_rules string `json:"aturan"`
			Med_desc  string `json:"deskripsi"`
			Med_exp   int    `json:"med_exp"`
		}
		var data []view
		dbase.
			Model(&models.Medicine{}).
			Select("* from medicines").
			Where(&models.Medicine{Med_id: put}).
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
		var data []models.Medicine

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
		var data models.Medicine
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
		err := dbase.First(&models.Medicine{}, "med_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if err := dbase.Find(&models.Medicine{}, "med_id=?", id[2]).Updates(&data).Error; err != nil {
			http.Error(w, "Error Update", http.StatusNotAcceptable)
		}
		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}
