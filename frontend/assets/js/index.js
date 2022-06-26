function loadData() {
    changeDate();
}

function changeDate() {
    document.getElementById("year-text").innerHTML = new Date().getFullYear();
}


