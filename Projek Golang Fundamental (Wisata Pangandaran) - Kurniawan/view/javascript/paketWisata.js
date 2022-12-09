
const mn = document.getElementById("mn")

function paketWisata(){
    fetch("http://localhost:5000/getPaket")
    .then((Response)=> Response.json())
    .then((result)=>{
        console.log(result)
        for(let index in result){
            const content = document.createElement("div")
            content.className ="konten"
            mn.appendChild(content)

            const img = document.createElement("img")
            img.className="gambar"
            img.src = result[index].gambar_paket
            content.appendChild(img)

            const container = document.createElement("div")
            container.className ="content"
            content.appendChild(container)

            const label1 = document.createElement("span")
            label1.className = "leb"
            label1.innerText = "Nama Paket : "
            container.appendChild(label1) 

            const name = document.createElement("span")
            name.className = "ket"
            name.innerText = result[index].nama_paket
            container.appendChild(name)

            const br = document.createElement("br")
            container.appendChild(br) 

            const label2 = document.createElement("span")
            label2.className = "leb"
            label2.innerText = "Harga Paket : "
            container.appendChild(label2) 

            const harga = document.createElement("span")
            harga.className = "ket"
            harga.innerText = result[index].harga_paket
            container.appendChild(harga)

            const br2 = document.createElement("br")
            container.appendChild(br2) 

            const label3 = document.createElement("span")
            label3.className = "leb"
            label3.innerText = "Nama Wisata : "
            container.appendChild(label3) 

            const name2 = document.createElement("span")
            name2.className = "ket"
            name2.innerText = result[index].nama_destinasi
            container.appendChild(name2) 

            const br3 = document.createElement("br")
            container.appendChild(br3) 

            const label4 = document.createElement("span")
            label4.className = "leb"
            label4.innerText = "Lokasi Wisata : "
            container.appendChild(label4) 

            const lok = document.createElement("span")
            lok.className = "ket"
            lok.innerText = result[index].lokasi_wisata
            container.appendChild(lok) 

            const br4 = document.createElement("br")
            container.appendChild(br4) 

            const label5 = document.createElement("span")
            label5.className = "leb"
            label5.innerText = "Nama Akomodasi : "
            container.appendChild(label5) 

            const name3 = document.createElement("span")
            name3.className = "ket"
            name3.innerText = result[index].nama_akomodasi
            container.appendChild(name3) 

            const br5 = document.createElement("br")
            container.appendChild(br5) 

            const label6 = document.createElement("span")
            label6.className = "leb"
            label6.innerText = "Lokasi Akomodasi : "
            container.appendChild(label6) 

            const lok2 = document.createElement("span")
            lok2.className = "ket"
            lok2.innerText = result[index].lokasi_akomodasi
            container.appendChild(lok2)

            const br6 = document.createElement("br")
            container.appendChild(br6) 

            const label7 = document.createElement("span")
            label7.className = "leb"
            label7.innerText = "Fasilitas : "
            container.appendChild(label7) 

            const fas = document.createElement("span")
            fas.className = "ket"
            fas.innerText = result[index].fasilitas
            container.appendChild(fas) 
        }
    })
}

paketWisata()