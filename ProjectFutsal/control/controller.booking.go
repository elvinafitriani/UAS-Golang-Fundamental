package control

import (
	"ProjectFutsal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PostBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data []models.Booking
		var id_lapangan int
		var jamInt int
		var err error
		var no_lap int

		err = json.NewDecoder(r.Body).Decode(&data)

		Db.Table("lapangans").Select("lapangans.harga").Where("no_lapangan = ?", data[len(data)-1].NoLap).Scan(&data[len(data)-1].Harga)
		Db.Table("transaksis").Select("transaksis.dp_status").Where("id_transaksi=?", data[len(data)-1].KodeTransaksi).Scan(&data[len(data)-1].DpStatus)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var cek []string = strings.Split(data[len(data)-1].Tanggal, " ")
		var layout string
		if cek[2] == "PM" {

			layout = "2006-01-02 3:04 PM"
			tanggalLayout, _ := time.Parse(layout, data[len(data)-1].Tanggal)
			var valJam []string = strings.Split(data[len(data)-1].Tanggal, " ")
			var valJam2 []string = strings.Split(valJam[1], ":")
			jamInt, err = strconv.Atoi(valJam2[0])

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if jamInt > 9 {
				http.Error(w, "Can't Booking", http.StatusInternalServerError)
				return
			}

			data[len(data)-1].TanggalPesan = time.Now()
			data[len(data)-1].TanggalMain = tanggalLayout
		} else {

			layout = "2006-01-02 3:04 AM"
			tanggalLayout, _ := time.Parse(layout, data[len(data)-1].Tanggal)
			var valJam []string = strings.Split(data[len(data)-1].Tanggal, " ")
			var valJam2 []string = strings.Split(valJam[1], ":")
			jamInt, err = strconv.Atoi(valJam2[0])

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if jamInt < 8 {
				http.Error(w, "Can't Booking", http.StatusInternalServerError)
				return
			}

			data[len(data)-1].TanggalPesan = time.Now()
			data[len(data)-1].TanggalMain = tanggalLayout
		}

		if data[len(data)-1].TanggalMain.Before(time.Now()) {
			http.Error(w, "Can't Booking", http.StatusInternalServerError)
			return
		}

		err = Db.First(&models.Booking{}, "tanggal_main = ?", data[len(data)-1].TanggalMain).Error
		if err == nil {
			Db.Table("bookings").Select("bookings.no_lap").Where("tanggal_main = ?", data[len(data)-1].TanggalMain).Scan(&id_lapangan)
			if data[len(data)-1].NoLap == id_lapangan {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		Db.Table("transaksis").Select("transaksis.no_lap").Where("id_transaksi=?", data[len(data)-1].KodeTransaksi).Scan(&no_lap)
		fmt.Println(no_lap)
		if data[len(data)-1].NoLap != no_lap {
			http.Error(w, "field numbers do not match", http.StatusInternalServerError)
			return
		}

		err = Db.First(&models.Booking{}, "kode_transaksi = ?", data[len(data)-1].KodeTransaksi).Error

		if err == nil {
			http.Error(w, "Code Have Used", http.StatusInternalServerError)
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

// Methode Get
func GetAllBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data []models.Booking
		Db.Order("tanggal_main ASC").Find(&data)
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

func GetBookingByName(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var nama []string = strings.Split(url, "/")

		if nama[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var data []models.Booking
		error := Db.Where("nama_team=?", nama[2]).Find(&data).Error

		if error != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		var dataJson, err = json.Marshal(data)

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

func GetBookingByTm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		url := r.URL.String()
		var tanggal []string = strings.Split(url, "/")
		var dataJson []byte
		var data []models.Booking

		if tanggal[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.Where("tanggal_main = ?", tanggal[2]).Order("tanggal_main ASC").Find(&data).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		dataJson, err = json.Marshal(data)

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
	http.Error(w, "Method Not Found", http.StatusNotFound)
	return
}

// Delete Booking
func DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		var dataR []models.Riwayat
		var data []models.Booking
		url := r.URL.String()
		var id []string = strings.Split(url, "/")

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err := Db.First(&models.Booking{}, "id_booking = ?", id[2]).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		Db.Table("bookings").Select("bookings.id_booking,bookings.nama_team,bookings.no_hp,bookings.no_lap,bookings.tanggal,bookings.tanggal_main,bookings.harga_lap, bookings.kode_transaksi, bookings.dp_status").Where("id_booking = ?", id[2]).Scan(&dataR)
		Db.Create(&dataR)
		err = Db.Where("id_booking = ?", id[2]).Delete(&data).Error

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

// Update Booking
func UpdatesBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		url := r.URL.String()
		var harga int
		var data models.Booking
		var id []string = strings.Split(url, "/")
		var id_lapangan int
		var jamInt int
		var err error

		decode := json.NewDecoder(r.Body).Decode(&data)
		Db.Model(&models.Lapangan{}).Select("lapangans.harga").Where("no_lapangan = ?", data.NoLap).Scan(&harga)
		var tanggal_main time.Time

		if decode != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if id[2] == "" {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var cek []string = strings.Split(data.Tanggal, " ")
		var layout string
		if cek[2] == "PM" {

			layout = "2006-01-02 3:04 PM"
			tanggalLayout, _ := time.Parse(layout, data.Tanggal)
			var valJam []string = strings.Split(data.Tanggal, " ")
			var valJam2 []string = strings.Split(valJam[1], ":")
			jamInt, err = strconv.Atoi(valJam2[0])

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if jamInt > 9 {
				http.Error(w, "Can't Booking", http.StatusInternalServerError)
				return
			}

			data.TanggalPesan = time.Now()
			data.TanggalMain = tanggalLayout
		} else {

			layout = "2006-01-02 3:04 AM"
			tanggalLayout, _ := time.Parse(layout, data.Tanggal)
			var valJam []string = strings.Split(data.Tanggal, " ")
			var valJam2 []string = strings.Split(valJam[1], ":")
			jamInt, err = strconv.Atoi(valJam2[0])

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if jamInt < 8 {
				http.Error(w, "Can't Booking", http.StatusInternalServerError)
				return
			}

			data.TanggalPesan = time.Now()
			data.TanggalMain = tanggalLayout
		}

		err = Db.First(&models.Booking{}, "id_booking = ?", id[2]).Error

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		Db.Table("bookings").Select("bookings.tanggal_main").Where("id_booking = ?", id[2]).Scan(&tanggal_main)

		if tanggal_main != data.TanggalMain {
			err = Db.First(&models.Booking{}, "tanggal_main = ?", data.TanggalMain).Error
			if err == nil {
				Db.Table("bookings").Select("bookings.no_lap").Where("tanggal_main = ?", data.TanggalMain).Scan(&id_lapangan)
				if data.NoLap == id_lapangan {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = Db.Where("id_booking = ?", id[2]).Updates(&data).Error

		if err != nil {
			fmt.Println("hallo")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success Update Data"))
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
