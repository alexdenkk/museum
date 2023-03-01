function buyClick (){
    console.log(this.parentElement.parentElement.classList.add("clicked"))
}

function removeClick (){
    console.log(this.parentElement.parentElement.classList.remove("clicked"))
}


document.querySelectorAll('.buy').forEach( elem => elem.addEventListener("click", buyClick))

document.querySelectorAll('.remove').forEach( elem => elem.addEventListener("click", removeClick))
