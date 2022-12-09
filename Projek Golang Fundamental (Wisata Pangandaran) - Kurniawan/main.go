package main

import (
	"fmt"
	"net/http"
	"projekan/controller"
)

func main() {

	// Destinasi Wisata
	http.HandleFunc("/postDestinasi", controller.PostDestinasi)
	http.HandleFunc("/getDestinasi", controller.GetDestinasi)
	http.HandleFunc("/updateDestinasi/", controller.UpdateDestinasi)
	http.HandleFunc("/deleteDestinasi/", controller.DeleteDestinasi)

	// Akomodasi
	http.HandleFunc("/postAkomodasi", controller.PostAkomodasi)
	http.HandleFunc("/getAkomodasi", controller.GetAkomodasi)
	http.HandleFunc("/updateAkomodasi/", controller.UpdateAkomodasi)
	http.HandleFunc("/deleteAkomodasi/", controller.DeleteAkomodasi)

	// Paket Wisata
	http.HandleFunc("/postPaket", controller.PostPaket)
	http.HandleFunc("/getPaket", controller.GetPaket)
	http.HandleFunc("/updatePaket/", controller.UpdatePaket)
	http.HandleFunc("/deletePaket/", controller.DeletePaket)

	fmt.Println("Starting Service")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error starting service")
	}
}
