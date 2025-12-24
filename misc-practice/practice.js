//Question 1
/*
input = Number(prompt("Input a number"));
if(input % 2 === 0){
    alert("Even");
}
else{
    alert("Odd");
}
*/

//Question 2
/*
let num1 = Number(prompt("Input first number"));
let num2 = Number(prompt("Input second number"));
let num3 = Number(prompt("Input third number"));
let nums = [num1, num2, num3];
alert("The largest number was: " + Math.max(...nums));
*/

//Question 3
/*
const min = -3;
const max = -3;
let arr = [];
for(let i = min; i <= max; i++){
    arr.push(i);
}
console.log(arr);
*/

//Question 4
/*
const target = "evaniscool";
const del = "|";
const space = 4;
let newStr = "";
for(let i = 0; i < target.length; i++){
    newStr += target[i];
    if(((i + 1) % space === 0) && (i != target.length-1)) newStr += del;
}
console.log(newStr);
*/

//Question 5
function mathIsFun(numberString){
    newArr = numberString.split("|");
    let num1 = -Infinity;
    let num2 = -Infinity;
    for(let i = 0; i < newArr.length - 1; i++){
        for(let j = i + 1; j < newArr.length; j++){
            if(Number(newArr[i]) + Number(newArr[j]) > num1 + num2){
                num1 = Number(newArr[i]);
                num2 = Number(newArr[j]);
            }
        }
    }
    return num1 + num2;
}

numberString = "-12|-13|-40|-50";
console.log(mathIsFun(numberString));