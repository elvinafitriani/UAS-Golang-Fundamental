package main

import (
	"fmt"
	"net/http"
	"project/controller"
)

func main() {
	http.HandleFunc("/getJadwal", controller.GetJadwal)
	http.HandleFunc("/postJadwal", controller.PostJadwal)
	http.HandleFunc("/updateJadwal/", controller.UpdateJadwal)
	http.HandleFunc("/deleteJadwal/", controller.DeleteJadwal)	
	http.HandleFunc("/viewJadwal", controller.ViewJadwal)

	http.HandleFunc("/getDosen", controller.GetDosen)
	http.HandleFunc("/postDosen", controller.PostDosen)
	http.HandleFunc("/updateDosen/", controller.UpdateDosen)
	http.HandleFunc("/deleteDosen/", controller.DeleteDosen)

	http.HandleFunc("/getMatkul", controller.GetMatkul)
	http.HandleFunc("/postMatkul", controller.PostMatkul)
	http.HandleFunc("/updateMatkul/", controller.UpdateMatkul)
	http.HandleFunc("/deleteMatkul/", controller.DeleteMatkul)

	http.HandleFunc("/getKhs", controller.GetKhs)
	http.HandleFunc("/postKhs", controller.PostKhs)
	http.HandleFunc("/updateKhs/", controller.UpdateKhs)
	http.HandleFunc("/deleteKhs/", controller.DeleteKhs)
	http.HandleFunc("/viewKhs", controller.ViewKhs)

	http.HandleFunc("/getMahasiswa", controller.GetMahasiswa)
	http.HandleFunc("/postMahasiswa", controller.PostMahasiswa)
	http.HandleFunc("/updateMahasiswa/", controller.UpdateMahasiswa)
	http.HandleFunc("/deleteMahasiswa/", controller.DeleteMahasiswa)

	println("Running Service")

	if err := http.ListenAndServe(":5006", nil); err != nil {
		fmt.Println("Error Starting Service")
	}
}
