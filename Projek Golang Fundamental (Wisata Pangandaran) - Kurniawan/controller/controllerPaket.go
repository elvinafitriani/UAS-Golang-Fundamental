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

func PostPaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.PaketWisata
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

func GetPaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Join
		DB.Table("paket_wisata").Select("paket_wisata.gambar_paket,paket_wisata.nama_paket,paket_wisata.harga_paket,destinasi_wisata.nama_destinasi,destinasi_wisata.lokasi_wisata,akomodasis.nama_akomodasi,akomodasis.lokasi_akomodasi,akomodasis.fasilitas").Joins("LEFT JOIN destinasi_wisata ON destinasi_wisata.id_destinasi = paket_wisata.destinasi").Joins("LEFT JOIN akomodasis ON akomodasis.id_akomodasi = paket_wisata.akomodasi_wisata").Scan(&data)

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

func UpdatePaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var datarequest models.PaketWisata
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.PaketWisata{}, "id_paket = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.PaketWisata{}).Where("id_paket = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DeletePaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		var data models.PaketWisata
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.PaketWisata{}, "id_paket = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//delete
		DB.Model(&models.PaketWisata{}).Where("id_paket = ?", id[2]).Delete(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
