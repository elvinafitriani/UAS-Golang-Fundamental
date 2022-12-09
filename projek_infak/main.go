package main

import (
	"fmt"
	"net/http"
	"projek_infak/controller"
)

func main() {
	//TABEL ALUMNI
	http.HandleFunc("/getAlumni", controller.GetAlumni)
	http.HandleFunc("/postAlumni", controller.PostAlumni)
	http.HandleFunc("/updateAlumni/", controller.UpdateAlumni)
	http.HandleFunc("/deleteAlumni/", controller.DeleteAlumni)

	//TABEL PRODI
	http.HandleFunc("/getProdi", controller.GetProdi)
	http.HandleFunc("/postProdi", controller.PostProdi)

	//TABEL ANGKATAN
	http.HandleFunc("/getAngkatan", controller.GetAngkatan)
	http.HandleFunc("/postAngkatan", controller.PostAngkatan)

	//TABEL REKENING
	http.HandleFunc("/getRekening", controller.GetRekening)
	http.HandleFunc("/postRekening", controller.PostRekening)

	//TABEL INFAK
	http.HandleFunc("/getInfak", controller.GetInfak)
	http.HandleFunc("/postInfak", controller.PostInfak)
	http.HandleFunc("/updateInfak/", controller.UpdateInfak)
	http.HandleFunc("/deleteInfak/", controller.DeleteInfak)

	//LAIN LAIN
	http.HandleFunc("/getLimit/", controller.Limit)
	http.HandleFunc("/getDetailInfak", controller.DetailInfak)
	http.HandleFunc("/getDetailAlumni", controller.DetailAlumni)
	http.HandleFunc("/getDetailRekening", controller.DetailRekening)

	fmt.Println("Succes Starting Sevice")

	if err := http.ListenAndServe(":5002", nil); err != nil {
		fmt.Println("Error starting service")
	}
}
