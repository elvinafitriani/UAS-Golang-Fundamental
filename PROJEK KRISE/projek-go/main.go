package main

import (
	"fmt"
	"net/http"
	"projek-go/controller"
)

func main() {

	//===tb kepengurusan===//
	http.HandleFunc("/getpengurus", controller.GetKepengurusan)
	http.HandleFunc("/postpengurus", controller.PostKepengurusan)
	http.HandleFunc("/uppengurus/", controller.UpdateKep)
	http.HandleFunc("/delpengurus/", controller.DeleteKep)
	http.HandleFunc("/detailpengurus/", controller.DetailKep)
	//=============//

	//===tb anak===//
	http.HandleFunc("/getanak", controller.GetAnak)
	http.HandleFunc("/postanak", controller.PostAnak)
	http.HandleFunc("/upanak/", controller.UpdateAnak)
	http.HandleFunc("/delanak/", controller.DeleteAnak)
	http.HandleFunc("/detailanak/", controller.DetailAnak)
	http.HandleFunc("/limitanak/", controller.LimitAnak)
	//=============//

	//===tb sekolah===//
	http.HandleFunc("/getsekolah", controller.GetSekolah)
	http.HandleFunc("/postsekolah", controller.PostSekolah)
	http.HandleFunc("/upsekolah/", controller.UpdateSekolah)
	http.HandleFunc("/delsekolah/", controller.DeleteSekolah)
	http.HandleFunc("/detailsekolah/", controller.DetailSekolah)
	//=============//

	//===tb rikes===//
	http.HandleFunc("/getrikes", controller.GetRiwayatKesehatan)
	http.HandleFunc("/postrikes", controller.PostRiwayatKesehatan)
	http.HandleFunc("/uprikes/", controller.UpdateRikes)
	http.HandleFunc("/delrikes/", controller.DeleteRikes)
	http.HandleFunc("/detailrikes/", controller.DetailRikes)
	//=============//

	//===tb donatur===//
	http.HandleFunc("/getdonatur", controller.GetDonatur)
	http.HandleFunc("/postdonatur", controller.PostDonatur)
	http.HandleFunc("/updonatur/", controller.UpdateDonatur)
	http.HandleFunc("/deldonatur/", controller.DeleteDonatur)
	http.HandleFunc("/detaildonatur/", controller.DetailDonatur)
	//=============//

	//===tb sumbangan===//
	http.HandleFunc("/getsumbangan", controller.GetSumbangan)
	http.HandleFunc("/postsumbangan", controller.PostSumbangan)
	http.HandleFunc("/upsumbangan/", controller.UpdateSumbangan)
	http.HandleFunc("/delsumbangan/", controller.DeleteSumbangan)
	http.HandleFunc("/detailsumbangan/", controller.DetailSumbangan)
	//=============//

	//===tb adopsi===//
	http.HandleFunc("/getadopsi", controller.GetAdopsi)
	http.HandleFunc("/postadopsi", controller.PostAdopsi)
	http.HandleFunc("/upadopsi/", controller.UpdateAdopsi)
	http.HandleFunc("/deladopsi/", controller.DeleteAdopsi)
	http.HandleFunc("/detailadopsi/", controller.DetailAdopsi)
	//=============//

	fmt.Println("running service")

	if err := http.ListenAndServe(":5011", nil); err != nil {
		fmt.Println("Error Starting  Service")
	}

}
