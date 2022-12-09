var img1 = document.getElementById("lap1")
var img2 = document.getElementById("lap2")
var img3 = document.getElementById("lap3")

var harga1 = document.getElementById("text1")
var harga2 = document.getElementById("text2")
var harga3 = document.getElementById("text3")

var lap1 = document.getElementById("nolap1")
var lap2 = document.getElementById("nolap2")
var lap3 = document.getElementById("nolap3")

function project(){
    fetch("http://localhost:5000/getlap")
    .then((Response) => Response.json())
    .then((result) => {
        console.log(result)
        img1.src = result[0].image;
        img2.src = result[1].image;
        img3.src = result[2].image;
        
        harga1.innerText = result[0].harga;
        harga2.innerText = result[1].harga;
        harga3.innerText = result[2].harga;

        lap1.innerText = result[0].no_lapangan
        lap2.innerText = result[1].no_lapangan
        lap3.innerText = result[2].no_lapangan

    })
}

project()