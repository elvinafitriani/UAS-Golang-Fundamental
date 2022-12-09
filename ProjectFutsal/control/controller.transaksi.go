package control

import (
	"ProjectFutsal/models"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Transaksi
func PostTrans(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data []models.Transaksi
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		data[len(data)-1].TanggalTransaksi = time.Now()
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

// Transaksi
func GetAllTransaksi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Transaksi
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

func GetTransaksiById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Transaksi
		url := r.URL.String()
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		Db.Where("id_transaksi=?", id[2]).Find(&data)

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
func UpdateTransaksi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		url := r.URL.String()
		var id []string = strings.Split(url, "/")
		var data models.Transaksi

		decode := json.NewDecoder(r.Body).Decode(&data)

		if decode != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.First(&models.Transaksi{}, "no_lapangan = ?", id[2]).Error

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

// Delete Transaksi
func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		url := r.URL.String()
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.Where("no_lapangan = ?", id[2]).Delete(&models.Transaksi{}).Error

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
