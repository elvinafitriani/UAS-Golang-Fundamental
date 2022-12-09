
const mn = document.getElementById("mn")

function akomodasi(){
    fetch("http://localhost:5000/getAkomodasi")
    .then((Response)=> Response.json())
    .then((result)=>{
        console.log(result)
        for(let index in result){
            const content = document.createElement("div")
            content.className ="konten"
            mn.appendChild(content)

            const img = document.createElement("img")
            img.className="gambar"
            img.src = result[index].gambar_akomodasi
            content.appendChild(img)

            const container = document.createElement("div")
            container.className ="content"
            content.appendChild(container)

            const label1 = document.createElement("span")
            label1.className = "leb"
            label1.innerText = "Nama Akomodasi : "
            container.appendChild(label1) 

            const name = document.createElement("span")
            name.className = "ket"
            name.innerText = result[index].nama_akomodasi
            container.appendChild(name) 

            const br = document.createElement("br")
            container.appendChild(br) 

            const label2 = document.createElement("span")
            label2.className = "leb"
            label2.innerText = "Lokasi Akomodasi : "
            container.appendChild(label2) 

            const lok = document.createElement("span")
            lok.className = "ket"
            lok.innerText = result[index].lokasi_akomodasi
            container.appendChild(lok) 

            const br2 = document.createElement("br")
            container.appendChild(br2) 

            const label3 = document.createElement("span")
            label3.className = "leb"
            label3.innerText = "Harga Tiket : "
            container.appendChild(label3) 

            const harga = document.createElement("span")
            harga.className = "ket"
            harga.innerText = result[index].harga_kamar
            container.appendChild(harga) 

            const br3 = document.createElement("br")
            container.appendChild(br3) 

            const label4 = document.createElement("span")
            label4.className = "leb"
            label4.innerText = "Fasilitas : "
            container.appendChild(label4) 

            const fas = document.createElement("span")
            fas.className = "ket"
            fas.innerText = result[index].fasilitas
            container.appendChild(fas) 
        }

    })
}

akomodasi()