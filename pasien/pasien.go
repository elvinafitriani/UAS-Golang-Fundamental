package pasien

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
			Pas_id    int    `json:"pasien_id"`
			Pas_nama  string `json:"nama"`
			Room_nama string `json:"Room_name"`
			Doc_nama  string `json:"Doctor_name"`
		}
		var data []view
		dbase.
			Model(&models.Pasien{}).
			Select("pasiens.pas_id, pasiens.pas_nama,rooms.room_nama, doctors.doc_nama").
			Joins("right join doctors on doctors.doc_id = pasiens.pas_doc inner join rooms on rooms.room_id = pasiens.pas_room").
			Order("pasiens.pas_id asc").
			Scan(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Enconde to JSON", 500)
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
			Pas_id        int    `json:"pasien_id"`
			Pas_nama      string `json:"nama"`
			Pas_obat      int    `json:"obat"`
			Med_nama      string `json:"Medicine_name"`
			Med_rules     string `json:"aturan"`
			Room_nama     string `json:"Room_name"`
			Doc_nama      string `json:"Doctor_name"`
			Detail_ttl    string `json:"detail_ttl"`
			Detail_age    string `json:"detail_age"`
			Detail_jk     string `json:"detail_jk"`
			Detail_addres string `json:"detail_address"`
			Detail_stat   string `json:"detail_stat"`
			Detail_desc   string `json:"detail_desc"`
		}
		var data []view
		dbase.
			Model(&models.Pasien{}).
			Select("pasiens.pas_id, pasiens.pas_nama, pasiens.pas_obat, medicines.med_nama, medicines.med_rules, rooms.room_nama, doctors.doc_nama, pasiens.detail_ttl, pasiens.detail_age, pasiens.detail_jk, pasiens.detail_addres, pasiens.detail_stat, pasiens.detail_desc").
			Joins("right join doctors on doctors.doc_id = pasiens.pas_doc inner join rooms on rooms.room_id = pasiens.pas_room inner join medicines on medicines.med_id = pasiens.pas_obat").
			Where(&models.Pasien{Pas_id: put}).Scan(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Enconde to JSON", 500)
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
		var data []models.Pasien

		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode Json", 500)
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
		var data models.Pasien
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
		err := dbase.First(&models.Pasien{}, "pas_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = dbase.Where("pas_id = ?", id[2]).Updates(&data).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		var data models.Pasien
		url := r.URL.String()
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err := dbase.First(&models.Pasien{}, "pas_id=?", id[2]).Error
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		dbase.Where("pas_id=?", id[2]).Delete(&data)
		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}
