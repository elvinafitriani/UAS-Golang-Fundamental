package main

import (
	"fmt"
	"net/http"
	"projek/doctor"
	"projek/kamar"
	"projek/obat"
	"projek/pasien"
)

func main() {
	http.HandleFunc("/getdoc", doctor.Getdoc)
	http.HandleFunc("/getdoc/", doctor.GetRequest)
	http.HandleFunc("/postdoc", doctor.Post)
	http.HandleFunc("/updatedoc/", doctor.Update)
	http.HandleFunc("/postdetaildoc", doctor.Postdetail)
	http.HandleFunc("/updatedetail/", doctor.Updatedetail)

	http.HandleFunc("/getpasien", pasien.Get)
	http.HandleFunc("/getpasien/", pasien.GetRequest)
	http.HandleFunc("/postpasien", pasien.Post)
	http.HandleFunc("/updatepas/", pasien.Update)
	http.HandleFunc("/deletepas/", pasien.Delete)

	http.HandleFunc("/getroom", kamar.Get)
	http.HandleFunc("/getroom/", kamar.GetRequest)
	http.HandleFunc("/postroom", kamar.Post)
	http.HandleFunc("/updateroom/", doctor.Update)

	http.HandleFunc("/postmed", obat.Post)
	http.HandleFunc("/updatemed/", obat.Update)
	http.HandleFunc("/getjob", doctor.Getjob)

	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", nil); err != nil { //return listenandserve bertipe error
		fmt.Println("Error Starting Service")
	}

}
