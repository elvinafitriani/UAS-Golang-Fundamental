const elemenOutput = document.querySelector(".isi");
view();

function view(){
    fetch("http://localhost:5006/getMahasiswa")
    .then(function(response){
        return response.json();
    })
    .then(function(json){
        console.log(json);
        elemenOutput.innerHTML = "";
        for(let a in json){
            let j1 = document.createElement("tr");
            elemenOutput.appendChild(j1);

            let j2 = document.createElement("td");
            j2.textContent = json[a].nim;
            j1.appendChild(j2);

            let j3 = document.createElement("td");
            let j33 = document.createElement("img");
            j33.className = 'gg';
            j3.appendChild(j33);
            j33.src= json[a].foto;
            j1.appendChild(j33);
            

            let j4 = document.createElement("td");
            j4.textContent = json[a].nama_mahasiswa;
            j1.appendChild(j4);
            
            let j5 = document.createElement("td");
            j5.textContent = json[a].alamat_mahasiswa;
            j1.appendChild(j5);

            let j6 = document.createElement("td");
            j6.textContent = json[a].jk_mhs;
            j1.appendChild(j6);
        }

    });
}