
const mn = document.getElementById("mn")

function wisata(){
    fetch("http://localhost:5000/getDestinasi")
    .then((Response)=> Response.json())
    .then((result)=>{
        console.log(result)
        for(let index in result){
            const content = document.createElement("div")
            content.className ="konten"
            mn.appendChild(content)

            const img = document.createElement("img")
            img.className="gambar"
            img.src = result[index].gambar
            content.appendChild(img)

            const container = document.createElement("div")
            container.className ="content"
            content.appendChild(container)

            const label1 = document.createElement("span")
            label1.className = "leb"
            label1.innerText = "Nama Wisata : "
            container.appendChild(label1) 

            const name = document.createElement("span")
            name.className = "ket"
            name.innerText = result[index].nama_destinasi
            container.appendChild(name) 

            const br = document.createElement("br")
            container.appendChild(br) 

            const label2 = document.createElement("span")
            label2.className = "leb"
            label2.innerText = "Lokasi Wisata : "
            container.appendChild(label2) 

            const lok = document.createElement("span")
            lok.className = "ket"
            lok.innerText = result[index].lokasi_wisata
            container.appendChild(lok) 

            const br2 = document.createElement("br")
            container.appendChild(br2) 

            const label3 = document.createElement("span")
            label3.className = "leb"
            label3.innerText = "Harga Tiket : "
            container.appendChild(label3) 

            const harga = document.createElement("span")
            harga.className = "ket"
            harga.innerText = result[index].harga_tiket
            container.appendChild(harga) 

            const br3 = document.createElement("br")
            container.appendChild(br3) 

            const label4 = document.createElement("span")
            label4.className = "leb"
            label4.innerText = "Deskripsi : "
            container.appendChild(label4) 

            const des = document.createElement("span")
            des.className = "ket"
            des.innerText = result[index].deskripsi
            container.appendChild(des) 
        }

    })
}

wisata()