function loadData() {
    document.getElementById("year-text").innerHTML = new Date().getFullYear();
    document.getElementById("newItemForm").addEventListener("submit", myFunction);
}

function myFunction() {
  alert(window.location.href);
}


