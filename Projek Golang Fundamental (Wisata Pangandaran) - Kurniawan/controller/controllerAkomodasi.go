package controller

import (
	"encoding/json"
	"net/http"
	"projekan/connection"
	"projekan/models"
	"strings"
)

func init() {
	DB = connection.ConnectToDb()
}

func PostAkomodasi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Akomodasi
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Create(&data)
		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func GetAkomodasi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Akomodasi
		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//cors
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func UpdateAkomodasi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var datarequest models.Akomodasi
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Akomodasi{}, "id_akomodasi = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.Akomodasi{}).Where("id_akomodasi = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DeleteAkomodasi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Akomodasi{}, "id_akomodasi = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//delete
		DB.Where("akomodasi_wisata = ?", id[2]).Delete(&models.PaketWisata{})
		DB.Where("id_akomodasi = ?", id[2]).Delete(&models.Akomodasi{})

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
