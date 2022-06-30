function loadData() {
    document.getElementById("year-text").innerHTML = new Date().getFullYear();
    document.getElementById("newItemForm").addEventListener("submit", submitData);
}

// function submitData() {
//   alert(window.location.href);
// }


