package controller

import (
	"encoding/json"
	"net/http"
	"projek-go/connection"
	"projek-go/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.Connect()
}

// ================================= tb kepengurusan ================================== //

func GetKepengurusan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Kepengurusan
		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostKepengurusan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Kepengurusan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		DB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateKep(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_kepengurusan []string = strings.Split(u, "/")
		if id_kepengurusan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Kepengurusan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		DB.First(&models.Kepengurusan{}, "id_kepengurusan = ?", id_kepengurusan[2])

		DB.Model(&models.Kepengurusan{}).Where("id_kepengurusan = ? ", id_kepengurusan[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteKep(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_kepengurusan []string = strings.Split(u, "/")
		if id_kepengurusan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Kepengurusan{}, "id_kepengurusan = ?", id_kepengurusan[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Kepengurusan{}, DB.Where("id_kepengurusan = ?", id_kepengurusan[2]))

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(200)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailKep(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id_kepengurusan []string = strings.Split(u, "/")
		if id_kepengurusan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var nama []models.Kepengurusan
		DB.Find(&nama, "id_kepengurusan=?", id_kepengurusan[2])
		datajson, err := json.Marshal(nama)
		if err != nil {
			http.Error(w, "error encode to json", 500)
			return
		}
		w.WriteHeader(200)
		w.Write(datajson)
		return
	}
}

// =============================================================================//

// // ================================= tb anak ================================== //

func GetAnak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Anak
		DB.Model(&models.Anak{}).Preload("RiwayatKesehatan").Preload("Adopsi").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostAnak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Anak
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		cek := DB.Create(&data).Error
		if cek != nil {
			http.Error(w, "Error Post Data", 500)
			return
		}
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(http.StatusCreated)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateAnak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var no_induk []string = strings.Split(u, "/")
		if no_induk[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Anak
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DB.First(&models.Anak{}, "no_induk = ?", no_induk[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.Anak{}).Where("no_induk = ? ", no_induk[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteAnak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var no_induk []string = strings.Split(u, "/")
		if no_induk[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Anak{}, "no_induk = ?", no_induk[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Anak{}, DB.Where("no_induk = ?", no_induk[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailAnak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var no_induk []string = strings.Split(u, "/")
		if no_induk[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var nama []models.Anak
		DB.Find(&nama, "no_induk=?", no_induk[2])
		DB.Model(&models.Anak{}).Preload("RiwayatKesehatan").Preload("Adopsi").Find(&nama, "no_induk=?", no_induk[2])
		datajson, err := json.Marshal(nama)
		if err != nil {
			http.Error(w, "error encode to json", 500)
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
}

// =============================================================================//

// ================================= tb sekolah ================================== //

func GetSekolah(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Sekolah
		DB.Model(&models.Sekolah{}).Preload("Anak").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostSekolah(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Sekolah
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateSekolah(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_sekolah []string = strings.Split(u, "/")
		if id_sekolah[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Sekolah
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DB.First(&models.Sekolah{}, "id_sekolah = ?", id_sekolah[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.Sekolah{}).Where("id_sekolah = ? ", id_sekolah[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteSekolah(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_sekolah []string = strings.Split(u, "/")
		if id_sekolah[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Sekolah{}, "id_sekolah = ?", id_sekolah[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Sekolah{}, DB.Where("id_sekolah = ?", id_sekolah[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailSekolah(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id_sekolah []string = strings.Split(u, "/")
		put, err := strconv.Atoi(id_sekolah[2])
		type detail struct {
			Id_sekolah     int    `json:"id_sekolah"`
			Nama_sekolah   string `json:"nama_sekolah"`
			Alamat_sekolah string `json:"alamat_sekolah"`
			Telp_sekolah   string `json:"telp_sekolah"`
			NoInduk        int    `json:"no_induk"`
			Nama_anak      string `json:"nama_anak"`
			Kelas          int    `json:"kelas"`
		}
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		var buff []detail
		DB.Model(&models.Sekolah{}).Select("sekolahs.id_sekolah,sekolahs.nama_sekolah,sekolahs.alamat_sekolah,sekolahs.telp_sekolah,anaks.no_induk, anaks.nama_anak,anaks.kelas").Joins("inner join anaks on sekolahs.id_sekolah = anaks.sekolah_id").Where(&models.Sekolah{IdSekolah: put}).Scan(&buff)
		datajson, err := json.Marshal(buff)
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", http.StatusNotFound)

}

// =============================================================================//

// ================================= tb rikes ================================== //

func GetRiwayatKesehatan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.RiwayatKesehatan
		DB.Model(&models.RiwayatKesehatan{}).Preload("Anak").Find(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostRiwayatKesehatan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.RiwayatKesehatan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateRikes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_kesehatan []string = strings.Split(u, "/")
		if id_kesehatan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.RiwayatKesehatan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DB.First(&models.RiwayatKesehatan{}, "id_kesehatan = ?", id_kesehatan[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.RiwayatKesehatan{}).Where("id_kesehatan = ? ", id_kesehatan[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteRikes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_kesehatan []string = strings.Split(u, "/")
		if id_kesehatan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.RiwayatKesehatan{}, "id_kesehatan = ?", id_kesehatan[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.RiwayatKesehatan{}, DB.Where("id_kesehatan = ?", id_kesehatan[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailRikes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var inp []string = strings.Split(url, "/")
		put, err := strconv.Atoi(inp[2])
		type detail struct {
			No_induk          int    `json:"no_induk"`
			Nama_anak         string `json:"nama_anak"`
			Jk_anak           string `json:"jk_anak"`
			Gol_darah         string `json:"gol_darah"`
			Tinggi_badan      int    `json:"tinggi_badan"`
			Berat_badan       int    `json:"berat_badan"`
			Tgl_periksa       string `json:"tgl_periksa"`
			Hasil_pemeriksaan string `json:"hasil_pemeriksaan"`
		}

		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}

		var buff []detail
		DB.Model(&models.RiwayatKesehatan{}).Select("anaks.no_induk,anaks.nama_anak,anaks.jk_anak,anaks.gol_darah,anaks.tinggi_badan,anaks.berat_badan,riwayat_kesehatans.tgl_periksa,riwayat_kesehatans.hasil_pemeriksaan").Joins("inner join anaks on riwayat_kesehatans.no_induk = anaks.no_induk ").Where(&models.RiwayatKesehatan{IdKesehatan: put}).Scan(&buff)
		datajson, err := json.Marshal(buff)
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(datajson)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Error Not Found", http.StatusNotFound)
}

// =============================================================================//

// ================================= tb donatur ================================== //

func GetDonatur(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Donatur
		DB.Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostDonatur(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Donatur
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateDonatur(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_donatur []string = strings.Split(u, "/")
		if id_donatur[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Donatur
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DB.First(&models.Donatur{}, "id_donatur = ?", id_donatur[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.Donatur{}).Where("id_donatur = ? ", id_donatur[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteDonatur(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_donatur []string = strings.Split(u, "/")
		if id_donatur[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Donatur{}, "id_donatur = ?", id_donatur[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Donatur{}, DB.Where("id_donatur = ?", id_donatur[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailDonatur(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id_donatur []string = strings.Split(u, "/")
		put, err := strconv.Atoi(id_donatur[2])
		type detail struct {
			IdDonatur        int    `json:"id_donatur"`
			NamaDonatur      string `json:"nama_donatur"`
			PekerjaanDonatur string `json:"pekerjaan_donatur"`
			TelpDonatur      int    `json:"telp_donatur"`
			IdSumbangan      int    `json:"id_sumbangan"`
			Tanggal          string `json:"tanggal"`
			Jumlah           int    `json:"jumlah"`
		}
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		var buff []detail
		DB.Model(&models.Donatur{}).Select("donaturs.id_donatur,donaturs.nama_donatur,donaturs.pekerjaan_donatur,donaturs.telp_donatur,sumbangans.id_sumbangan,sumbangans.tanggal,sumbangans.jumlah").Joins("inner join sumbangans on donaturs.id_donatur = sumbangans.id_donatur").Where(&models.Donatur{IdDonatur: put}).Scan(&buff)
		datajson, err := json.Marshal(buff)
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return
	}
	http.Error(w, "Error Not Found", http.StatusNotFound)

}

// =============================================================================//

// ================================= tb sumbangan ================================== //

func GetSumbangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Sumbangan
		DB.Model(&models.Sumbangan{}).Preload("Donatur").Find(&data)
		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostSumbangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Sumbangan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DB.Create(&data)
		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateSumbangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_sumbangan []string = strings.Split(u, "/")
		if id_sumbangan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data models.Sumbangan
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}

		err := DB.First(&models.Sumbangan{}, "id_sumbangan = ?", id_sumbangan[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.Sumbangan{}).Where("id_sumbangan = ? ", id_sumbangan[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return

	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteSumbangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_sumbangan []string = strings.Split(u, "/")
		if id_sumbangan[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Sumbangan{}, "id_sumbangan = ?", id_sumbangan[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Sumbangan{}, DB.Where("id_sumbangan = ?", id_sumbangan[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailSumbangan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id_sumbangan []string = strings.Split(u, "/")
		put, err := strconv.Atoi(id_sumbangan[2])
		type detail struct {
			IdSumbangan int    `json:"id_sumbangan"`
			IdDonatur   int    `json:"id_donatur"`
			Tanggal     string `json:"tanggal"`
			NamaDonatur string `json:"nama_donatur"`
			Jumlah      int    `json:"jumlah"`
		}
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		var buff []detail
		DB.Model(&models.Sumbangan{}).Select("sumbangans.id_sumbangan,donaturs.id_donatur,sumbangans.tanggal,donaturs.nama_donatur,sumbangans.jumlah").Joins("inner join donaturs on sumbangans.id_donatur = donaturs.id_donatur").Where(&models.Sumbangan{IdSumbangan: put}).Scan(&buff)
		datajson, err := json.Marshal(buff)
		if err != nil {
			http.Error(w, "Error Decode Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Write(datajson)
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found", http.StatusNotFound)

}

// =============================================================================//

// ================================= tb adopsi ================================== //

func GetAdopsi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Adopsi
		DB.Model(&models.Adopsi{}).Preload("Sekolah").Preload("Anak").Find(&data)

		datajson, err := json.Marshal(data)
		if err != nil {
			http.Error(w, " error encode to json ", 500)
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
	http.Error(w, "error not found", 404)
}

func PostAdopsi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data []models.Adopsi
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		DB.Create(&data)

		DB.Model(&models.Anak{}).Where("no_induk = ?", data[len(data)-1].IndukNo).Update("status_panti", "DI ADOPSI")

		w.Write([]byte("Sukses post data"))
		w.WriteHeader(200)
		return

	}
	http.Error(w, "Error Not Found ", 404)
}

func UpdateAdopsi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		u := r.URL.String()
		var id_adopsi []string = strings.Split(u, "/")
		if id_adopsi[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var data models.Adopsi
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Error Decode JSON ", 500)
			return
		}
		err := DB.First(&models.Adopsi{}, "id_adopsi = ?", id_adopsi[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		DB.Model(&models.Adopsi{}).Where("id_adopsi = ? ", id_adopsi[2]).Updates(&data)

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func DeleteAdopsi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		u := r.URL.String()
		var id_adopsi []string = strings.Split(u, "/")
		if id_adopsi[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := DB.First(&models.Adopsi{}, "id_adopsi = ?", id_adopsi[2]).Error
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		DB.Delete(&models.Adopsi{}, DB.Where("id_adopsi = ?", id_adopsi[2]))

		w.Write([]byte(http.StatusText(http.StatusOK)))
		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func DetailAdopsi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		u := r.URL.String()
		var id_adopsi []string = strings.Split(u, "/")
		if id_adopsi[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var nama []models.Adopsi
		DB.Find(&nama, "id_adopsi=?", id_adopsi[2]).Preload("Anak")
		datajson, err := json.Marshal(nama)
		if err != nil {
			http.Error(w, "error encode to json", 500)
			return
		}
		w.WriteHeader(200)
		w.Write(datajson)
		return
	}
}

// =============================================================================//
func LimitAnak(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Anak
		DB.Find(&data)

		u := r.URL.String()
		var id []string = strings.Split(u, "/")
		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		id2, error := strconv.Atoi(id[2])
		if error != nil {
			http.Error(w, "Error Encode to JSON", 500)
			return
		}

		DB.Limit(id2).Find(&data)

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
