var data = document.getElementById("tbodyt");
let color = ""

async function table(){
    let no = 1
    await fetch("http://localhost:5000/gettrans")
    .then((Response) => Response.json())
    .then((result)=>{
        for(let index in result){

            if(index %2==0){
                color = "grey"
            }else{
                color="white"
            }

            var tr = document.createElement("tr")
            tr.className = "trowst"
            tr.bgColor = color
            data.appendChild(tr)

            var tb = document.createElement("td")
            tb.className = "tdatat"
            tb.innerText = no;
            tr.appendChild(tb)

            var tb1 = document.createElement("td")
            tb1.className = "tdatat"
            tb1.innerText = result[index].id_transaksi;
            tr.appendChild(tb1)
            
            var tb2 = document.createElement("td")
            tb2.className = "tdatat"
            tb2.innerText = result[index].no_lap;
            tr.appendChild(tb2)

            var tb3 = document.createElement("td")
            tb3.className = "tdatat"
            tb3.innerText = result[index].dp_status;
            tr.appendChild(tb3)

            var tb4 = document.createElement("td")
            tb4.className = "tdatat"
            tb4.innerText = result[index].tanggal_transaksi;
            tr.appendChild(tb4)

            no++
        }
    })
    console.log(no)
    if(no === 1){
        data.innerHTML="Not Found Data"
    }
}

const value = document.getElementById("formt")

async function Click(){
    let no = 1
    await fetch("http://localhost:5000/gettransbyid/"+value.value)
    .then((Response) => Response.json())
    .then((result)=>{
        data.innerHTML =""
        for(let index in result){

            if(index %2==0){
                color = "grey"
            }else{
                color="white"
            }

            var tr = document.createElement("tr")
            tr.className = "trowst"
            tr.bgColor = color
            data.appendChild(tr)

            var tb = document.createElement("td")
            tb.className = "tdatat"
            tb.innerText = no;
            tr.appendChild(tb)

            var tb1 = document.createElement("td")
            tb1.className = "tdatat"
            tb1.innerText = result[index].id_transaksi;
            tr.appendChild(tb1)
            
            var tb2 = document.createElement("td")
            tb2.className = "tdatat"
            tb2.innerText = result[index].no_lap;
            tr.appendChild(tb2)

            var tb3 = document.createElement("td")
            tb3.className = "tdatat"
            tb3.innerText = result[index].dp_status;
            tr.appendChild(tb3)

            var tb4 = document.createElement("td")
            tb4.className = "tdatat"
            tb4.innerText = result[index].tanggal_transaksi;
            tr.appendChild(tb4)

            no++
        }
    })
    console.log(no)
    if(no === 1){
        data.innerHTML="Not Found Data"
    }
}

table()

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