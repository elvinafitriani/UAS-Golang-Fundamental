package doctor

import (
	"encoding/json"
	"fmt"

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

func Getdoc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		type view struct {
			Doc_id   int    `json:"doctor_id"`
			Doc_nama string `json:"Doctor_name"`
			Job_nama string `json:"Job_Name"`
			Jadwal   string `json:"jadwal"`
		}
		var data []view
		dbase.
			Model(&models.Doctor{}).
			Select("doctors.doc_id, doctors.doc_nama, Job_Spefs.Job_nama, doctors.jadwal").
			Joins("inner join Job_Spefs on doctors.Job_spes = Job_Spefs.Job_id").
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
			Doc_id     int    `json:"doctor_id"`
			Doc_nama   string `json:"Doctor_name"`
			Doc_age    string `json:"doc_age"`
			Job_spes   int    `json:"job"`
			Job_nama   string `json:"Job_Name"`
			Doc_pict   string `json:"doc_pict"`
			Jadwal     string `json:"jadwal"`
			Detail_ttl string `json:"detail_ttl"`
			Detail_edu string `json:"detail_edu"`
			Detail_sex string `json:"detail_sex"`
			Pas_id     int    `json:"pasien_id"`
			Pas_nama   string `json:"nama"`
			Detail_age string `json:"detail_age"`
		}
		var data []view
		dbase.
			Model(&models.Doctor{}).
			Select("doctors.doc_id, doctors.doc_nama, doctors.Job_spes, Job_Spefs.Job_nama, doctors.jadwal, doctors.doc_pict, detaildocs.*, doctors.doc_age, pasiens.pas_id, pasiens.pas_nama, pasiens.detail_age").
			Joins("inner join Job_Spefs on doctors.Job_spes = Job_Spefs.Job_id inner join detaildocs on detaildocs.detail_id = doctors.doc_id inner join pasiens on pasiens.pas_doc = doctors.doc_id").
			Where(&models.Doctor{Doc_id: put}).
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

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Doctor

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

func Postdetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Detaildoc

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
		var data models.Doctor
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
		err := dbase.First(&models.Doctor{}, "doc_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Println(id[2])
		fmt.Println(data)
		err = dbase.Where("doc_id = ?", id[2]).Updates(&data).Error
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

func Updatedetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var data models.Detaildoc
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
		err := dbase.First(&models.Detaildoc{}, "detail_id = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Println(id[2])
		fmt.Println(data)
		err = dbase.Where("detail_id = ?", id[2]).Updates(&data).Error
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

func Getjob(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		type view struct {
			Job_nama string `json:"Job_Name"`
		}
		var data []view
		dbase.
			Model(&models.JobSpef{}).
			Select("*").
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
