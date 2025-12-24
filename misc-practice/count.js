let count = document.getElementById("count-text");
count.innerHTML = 0;

function countUp() {
    count.innerHTML++;
}

function countDown() {
    count.innerHTML--;
}

function reset() {
    count.innerHTML = 0;
}