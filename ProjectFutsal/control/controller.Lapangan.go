package control

import (
	"ProjectFutsal/database"
	"ProjectFutsal/models"
	"encoding/json"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	Db = database.Connect()
}

// Methode Post
func PostLapangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data []models.Lapangan
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = Db.Create(&data).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Success Post Data"))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

// Get Lapangan
func GetAllLapngan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Lapangan
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
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

// Update Lapangan
func UpdateLapangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		url := r.URL.String()
		var id []string = strings.Split(url, "/")
		var data models.Lapangan

		decode := json.NewDecoder(r.Body).Decode(&data)

		if decode != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.First(&models.Lapangan{}, "no_lapangan = ?", id[2]).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		err = Db.Where("no_lapangan = ?", id[2]).Updates(&data).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success Update Data"))
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

// Delete Lapangan
func DeleteLapangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		url := r.URL.String()
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.Where("no_lapangan = ?", id[2]).Delete(&models.Lapangan{}).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write([]byte("Success Delete Data"))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
