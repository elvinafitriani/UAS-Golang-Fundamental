
/////////////////////////////  DATA INFAK  /////////////////////////////////////////////////
const tableBody = document.querySelector('#myTable'); 

tampil();
function tampil(){
    fetch('http://localhost:5002/getInfak')
    .then(function (response) {
        return response.json();
    })
    .then(function (json) {
        console.log(json)
        for (let x in json) {
            const baris = document.createElement('tr');
            tableBody.appendChild(baris);

            const elemenTgl = document.createElement('td');
            elemenTgl.textContent = json[x].tanggal;
            baris.appendChild(elemenTgl);

            const elemenId = document.createElement('td');
            elemenId.textContent = json[x].id_infak;
            baris.appendChild(elemenId);

            const elemenAlumni = document.createElement('td');
            elemenAlumni.textContent = json[x].id_alumni;
            baris.appendChild(elemenAlumni);
            
            const elemenRek = document.createElement('td');
            elemenRek.textContent = json[x].id_rekening;
            baris.appendChild(elemenRek);

            const elemenJml = document.createElement('td');
            elemenJml.textContent = json[x].jumlah_infak;
            baris.appendChild(elemenJml);

            const elemenKet = document.createElement('td');
            elemenKet.textContent = json[x].keterangan;
            baris.appendChild(elemenKet);
        }
    });
}




////////////////////////////// DETAIL INFAK ////////////////////////////////////////////////////

//DATA INFAK
const tableBody2 = document.querySelector('#detailInfak'); 

tampil2();
function tampil2(){
    fetch('http://localhost:5002/getDetailInfak')
    .then(function (response) {
        return response.json();
    })
    .then(function (json) {
        console.log(json)
        // tableBody2.innerHTML = '';

        for (let x in json) {
            const baris = document.createElement('tr');
            tableBody2.appendChild(baris);

            const elemenTgl = document.createElement('td');
            elemenTgl.textContent = json[x].tanggal;
            baris.appendChild(elemenTgl);

            const elemenId = document.createElement('td');
            elemenId.textContent = json[x].id_infak;
            baris.appendChild(elemenId);

            const elemenAlumni = document.createElement('td');
            elemenAlumni.textContent = json[x].id_alumni;
            baris.appendChild(elemenAlumni);

            const elemenNama = document.createElement('td');
            elemenNama.textContent = json[x].nama_alumni;
            baris.appendChild(elemenNama);

            const elemenRek = document.createElement('td');
            elemenRek.textContent = json[x].nama_rek;
            baris.appendChild(elemenRek);

            const elemenJml = document.createElement('td');
            elemenJml.textContent = json[x].jumlah_infak;
            baris.appendChild(elemenJml);
        }
    });
}



////////////////////////////// DETAIL ALUMNI ////////////////////////////////////////////////////

const tableBody3 = document.querySelector('#detailAlumni'); 

tampil3();
function tampil3(){
    fetch('http://localhost:5002/getDetailAlumni')
    .then(function (response) {
        return response.json();
    })
    .then(function (json) {
        console.log(json)
        // tableBody3.innerHTML = '';

        for (let x in json) {
            const baris = document.createElement('tr');
            tableBody3.appendChild(baris);

            const elemenId = document.createElement('td');
            elemenId.textContent = json[x].id_alumni;
            baris.appendChild(elemenId);

            const elemenFoto = document.createElement('td');
            const elemenFoto2 = document.createElement('img');
            elemenFoto2.src = json[x].foto;
            elemenFoto2.className="foto";
            elemenFoto.appendChild(elemenFoto2);
            baris.appendChild(elemenFoto);

            const elemenNama = document.createElement('td');
            elemenNama.textContent = json[x].nama_alumni;
            baris.appendChild(elemenNama);

            const elemenAkt = document.createElement('td');
            elemenAkt.textContent = json[x].angkatan_ke;
            baris.appendChild(elemenAkt);

            const elemenProdi = document.createElement('td');
            elemenProdi.textContent = json[x].nama_prodi;
            baris.appendChild(elemenProdi);

            const elemenTotal = document.createElement('td');
            elemenTotal.textContent = json[x].total_infak;
            baris.appendChild(elemenTotal);

            const elemenKet = document.createElement('td');
            elemenKet.textContent = json[x].keterangan;
            baris.appendChild(elemenKet);
        }
    });
}




////////////////////////////// DETAIL ALUMNI ////////////////////////////////////////////////////

const tableBody4 = document.querySelector('#detaillRekening'); 

tampil4();
function tampil4(){
    fetch('http://localhost:5002/getDetailRekening')
    .then(function (response) {
        return response.json();
    })
    .then(function (json) {
        console.log(json.total_infak);
        // tableBody4.innerHTML = '';

            const baris = document.createElement('tr');
            tableBody4.appendChild(baris);

            const elemensaldo = document.createElement('td');
            elemensaldo.innerText =json.total_infak;
            elemensaldo.className="rekening";
            baris.appendChild(elemensaldo);

            const elemenbca = document.createElement('td');
            elemenbca.innerText = json.saldo_bca;
            elemenbca.className="rekening";
            baris.appendChild(elemenbca);

            const elemenmandiri = document.createElement('td');
            elemenmandiri.innerText = json.saldo_mandiri;
            elemenmandiri.className="rekening";
            baris.appendChild(elemenmandiri);
    }
    );
}








