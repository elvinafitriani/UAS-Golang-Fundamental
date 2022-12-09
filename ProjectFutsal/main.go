package main

import (
	"ProjectFutsal/control"
	"fmt"
	"net/http"
)

func main() {
	//Lapngan Router
	http.HandleFunc("/postlap", control.PostLapangan)
	http.HandleFunc("/getlap", control.GetAllLapngan)
	http.HandleFunc("/updlap/", control.DeleteLapangan)
	http.HandleFunc("/dellap/", control.UpdateLapangan)

	//Transaksi
	http.HandleFunc("/posttrans", control.PostTrans)
	http.HandleFunc("/gettrans", control.GetAllTransaksi)
	http.HandleFunc("/gettransbyid/", control.GetTransaksiById)
	http.HandleFunc("/updtrans/", control.DeleteTransaksi)
	http.HandleFunc("/deltrans/", control.UpdateTransaksi)

	//Booking Router
	http.HandleFunc("/postbook", control.PostBooking)
	http.HandleFunc("/getbook", control.GetAllBooking)
	http.HandleFunc("/getbookbyname/", control.GetBookingByName)
	http.HandleFunc("/getbookbytm/", control.GetBookingByTm)
	http.HandleFunc("/delbook/", control.DeleteById)
	http.HandleFunc("/updbook/", control.UpdatesBooking)

	//Riwayat Router
	http.HandleFunc("/getallriwayat", control.GetAllRiwayat)
	http.HandleFunc("/getriwayatbyid/", control.GetRiwayatById)

	fmt.Println("Start Running")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Can't Start Server")
		return
	}
}
