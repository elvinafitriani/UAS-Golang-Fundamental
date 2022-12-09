package control

import (
	"ProjectFutsal/models"
	"encoding/json"
	"net/http"
	"strings"
)

func GetAllRiwayat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Riwayat
		Db.Find(&data)
		dataJson, err := json.Marshal(data)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(dataJson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

func GetRiwayatById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var data []models.Riwayat
		var dataJson []byte
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.Where("id_booking = ?", id[2]).Find(&data).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		dataJson, err = json.Marshal(&data)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "applicaton/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(dataJson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
