
var data = document.getElementById("tbodyr");
let color = ""
async function table(){
    let no = 1
    await fetch("http://localhost:5000/getallriwayat")
    .then((Response) => Response.json())
    .then((result)=>{
        for(let index in result){

            if(index %2==0){
                color = "grey"
            }else{
                color="white"
            }

            var tr = document.createElement("tr")
            tr.className = "trowsr"
            tr.bgColor = color
            data.appendChild(tr)

            var tb = document.createElement("td")
            tb.className = "tdatar"
            tb.innerText = no;
            tr.appendChild(tb)

            var tb1 = document.createElement("td")
            tb1.className = "tdatar"
            tb1.innerText = result[index].id_booking;
            tr.appendChild(tb1)
            
            var tb2 = document.createElement("td")
            tb2.className = "tdatar"
            tb2.innerText = result[index].nama_team;
            tr.appendChild(tb2)

            var tb3 = document.createElement("td")
            tb3.className = "tdatar"
            tb3.innerText = result[index].no_hp;
            tr.appendChild(tb3)

            var tb4 = document.createElement("td")
            tb4.className = "tdatar"
            tb4.innerText = result[index].id_lapangan;
            tr.appendChild(tb4)

            var tb5 = document.createElement("td")
            tb5.className = "tdatar"
            tb5.innerText = result[index].tanggal_main;
            tr.appendChild(tb5)

            var tb6 = document.createElement("td")
            tb6.className = "tdata"
            tb6.innerText = result[index].harga;
            tr.appendChild(tb6)

            var tb7 = document.createElement("td")
            tb7.className = "tdata"
            tb7.innerText = result[index].dp_status;
            tr.appendChild(tb7)

            no++
        }
    })
    console.log(no)
    if(no === 1){
        data.innerHTML="Not Found Data"
    }
}
table()
const btn = document.getElementById("btnr")
const form = document.getElementById("formr")

btn.addEventListener("click",async function(){
    let value = form.value
        const mn = (document.getElementById("tbodyr").innerHTML = "");

        let no = 1
        await fetch("http://localhost:5000/getriwayatbyid/" + value)
        .then((Response) => Response.json())
        .then((result)=>{
        for(let index in result){
            console.log(result)
            if(index %2==0){
                color = "grey"
            }else{
                color="white"
            }

            var tr = document.createElement("tr")
            tr.className = "trowsr"
            tr.bgColor = color
            data.appendChild(tr)

            var tb = document.createElement("td")
            tb.className = "tdatar"
            tb.innerText = no;
            tr.appendChild(tb)

            var tb1 = document.createElement("td")
            tb1.className = "tdatar"
            tb1.innerText = result[index].id_booking;
            tr.appendChild(tb1)
            
            var tb2 = document.createElement("td")
            tb2.className = "tdatar"
            tb2.innerText = result[index].nama_team;
            tr.appendChild(tb2)

            var tb3 = document.createElement("td")
            tb3.className = "tdatar"
            tb3.innerText = result[index].no_hp;
            tr.appendChild(tb3)

            var tb4 = document.createElement("td")
            tb4.className = "tdatar"
            tb4.innerText = result[index].id_lapangan;
            tr.appendChild(tb4)

            var tb5 = document.createElement("td")
            tb5.className = "tdatar"
            tb5.innerText = result[index].tanggal_main;
            tr.appendChild(tb5)

            var tb6 = document.createElement("td")
            tb6.className = "tdatar"
            tb6.innerText = result[index].harga;
            tr.appendChild(tb6)

            var tb7 = document.createElement("td")
            tb7.className = "tdatar"
            tb7.innerText = result[index].dp_status;
            tr.appendChild(tb7)
            no++
        }
    })
    console.log(no)
    if(no == 1){
        data.innerHTML="Not Found Data"
    }
})


// scroll
const scrl = document.getElementById("scrol");

window.addEventListener("scroll", () => {
  if (pageYOffset > 100) {
    scrl.classList.add("active");
  } else {
    scrl.classList.remove("active");
  }
});

scrl.addEventListener("click", () => {
  window.scrollTo({
    top: 0,
    behavior: "smooth",
  });
});
