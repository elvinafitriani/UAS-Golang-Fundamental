package controller

import (
	"encoding/json"
	"net/http"
	"projek_infak/connection"
	"projek_infak/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func GetAngkatan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Angkatan
		DB.Model(&models.Angkatan{}).Preload("Alumni").Preload("Prodi").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func PostAngkatan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Angkatan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Create", 500)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetProdi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Prodi
		DB.Model(&models.Prodi{}).Preload("Alumni").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func PostProdi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Prodi
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Create", 500)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetAlumni(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Alumni
		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}

	http.Error(w, "Error Not Found", 404)
}

func PostAlumni(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Alumni
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Create", 500)
			return
		}
		w.Write([]byte("Sukses Post Data"))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func UpdateAlumni(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var datarequest models.Alumni
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}

		err := DB.First(&models.Alumni{}, "id_alumni = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Model(&models.Alumni{}).Where("id_alumni = ?", id[2]).Updates(&datarequest)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DeleteAlumni(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Alumni{}, "id_alumni=?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Alumni{}, DB.Where("id_alumni=?", id[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func GetRekening(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Rekening
		DB.Model(&models.Rekening{}).Preload("Infak").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func PostRekening(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Rekening
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Create", 500)
		}

		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func GetInfak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Infak
		DB.Model(&models.Infak{}).Preload("Alumni").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func PostInfak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Infak
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Create", 500)
			return
		}

		//trigger
		DB.Model(&models.Alumni{}).Where("id_alumni = ?", data[len(data)-1].Id_alumni).Update("total_infak", gorm.Expr("total_infak + ?", 1))
		DB.Model(&models.Rekening{}).Where("id_rekening = ?", data[len(data)-1].Id_rekening).Update("saldo", gorm.Expr("saldo + ?", data[len(data)-1].Jumlah_infak))

		w.Write([]byte("Sukses Post Data"))
		w.WriteHeader(200)

		return
	}
	http.Error(w, "Error Not Found", 404)
}

func UpdateInfak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var datarequest models.Infak

		err := DB.First(&models.Infak{}, "id_infak = ?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//trigger update saldo di tabel rekening
		DB.Table("infaks").Select("infaks.id_rekening").Where("id_infak = ?", id[2]).Scan(&datarequest.Id_rekening)
		DB.Table("infaks").Select("infaks.jumlah_infak").Where("id_rekening = ?", datarequest.Id_rekening).Scan(&datarequest.Jumlah_infak)

		DB.Model(&models.Rekening{}).Where("id_rekening = ?", datarequest.Id_rekening).Update("saldo", gorm.Expr("saldo - ?", datarequest.Jumlah_infak))

		//nge update
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&datarequest); err != nil {
			http.Error(w, "Error Decode JSON", 500)
			return
		}
		DB.Model(&models.Infak{}).Where("id_infak = ?", id[2]).Updates(&datarequest)

		//triger
		DB.Model(&models.Rekening{}).Where("id_rekening = ?", datarequest.Id_rekening).Update("saldo", gorm.Expr("saldo + ?", datarequest.Jumlah_infak))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DeleteInfak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Infak{}, "id_infak=?", id[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var datarequest models.Infak

		//trigger update saldo di tabel rekening
		DB.Table("infaks").Select("infaks.id_rekening").Where("id_infak = ?", id[2]).Scan(&datarequest.Id_rekening)
		DB.Table("infaks").Select("infaks.jumlah_infak").Where("id_rekening = ?", datarequest.Id_rekening).Scan(&datarequest.Jumlah_infak)

		DB.Model(&models.Rekening{}).Where("id_rekening = ?", datarequest.Id_rekening).Update("saldo", gorm.Expr("saldo - ?", datarequest.Jumlah_infak))

		//trigger total_infak di tabel  alumni
		DB.Table("infaks").Select("infaks.id_alumni").Where("id_infak = ?", id[2]).Scan(&datarequest.Id_alumni)
		DB.Model(&models.Alumni{}).Where("id_alumni = ?", datarequest.Id_alumni).Update("total_infak", gorm.Expr("total_infak - ?", 1))

		//DELETE
		DB.Delete(&models.Infak{}, DB.Where("id_infak=?", id[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
}

func Limit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Infak
		DB.Find(&data)

		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		idd, error := strconv.Atoi(id[2]) //castinya
		if error != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		DB.Limit(idd).Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}

	http.Error(w, "Error Not Found", 404)
}

func DetailInfak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Infak
		DB.Find(&data)
		type hasil struct {
			Tanggal       string `json:"tanggal"`
			Id_infak      int    `json:"id_infak"`
			Id_alumni     int    `json:"id_alumni"`
			Nama_alumni   string `json:"nama_alumni"`
			Nama_Rekening string `json:"nama_rek"`
			Jumlah_infak  int    `json:"jumlah_infak"`
		}

		var tampHasil []hasil
		DB.Model(&data).Select("infaks.tanggal, infaks.id_infak, infaks.id_alumni, alumnis.nama_alumni, rekenings.nama_rekening, infaks.jumlah_infak").
			Joins("INNER JOIN alumnis ON infaks.id_alumni=alumnis.id_alumni INNER JOIN rekenings ON infaks.id_rekening=rekenings.id_rekening").Scan(&tampHasil)
		datajson, err := json.Marshal(tampHasil)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func DetailAlumni(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Alumni
		DB.Find(&data)
		type hasil struct {
			Id_alumni   int    `json:"id_alumni"`
			Foto        string `json:"foto"`
			Nama_alumni string `json:"nama_alumni"`
			Angkatan_ke int    `json:"angkatan_ke"`
			Nama_prodi  string `json:"nama_prodi"`
			Total_infak int    `json:"total_infak"`
			Keterangan  string `json:"keterangan"`
		}

		var tampHasil []hasil
		DB.Model(&data).Select("alumnis.id_alumni, alumnis.foto, alumnis.foto, alumnis.nama_alumni, angkatans.angkatan_ke, prodis.nama_prodi, alumnis.total_infak, alumnis.keterangan").Joins("INNER JOIN angkatans ON alumnis.id_angkatan=angkatans.id_angkatan INNER JOIN prodis ON alumnis.id_prodi=prodis.id_prodi").Scan(&tampHasil)
		datajson, err := json.Marshal(tampHasil)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}

func DetailRekening(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Rekening
		DB.Find(&data)
		type hasil struct {
			Total_infak   int `json:"total_infak"`
			Saldo_bca     int `json:"saldo_bca"`
			Saldo_mandiri int `json:"saldo_mandiri"`
		}

		var tampHasil hasil
		DB.Model(&data).Select("SUM(saldo)").Scan(&tampHasil.Total_infak)
		DB.Model(&data).Select("saldo").Where("id_rekening = ?", 1).Scan(&tampHasil.Saldo_bca)
		DB.Model(&data).Select("saldo").Where("id_rekening = ?", 2).Scan(&tampHasil.Saldo_mandiri)

		datajson, err := json.Marshal(tampHasil)
		if err != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", 404)
}
