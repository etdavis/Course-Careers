/*Question 2
Write a function that has two parameters, the first being an array of strings, the second being an
array of unique integers (whole numbers). Your function should return the number of strings in
the input array that have a length contained in the second array. If the array of numbers is empty
you should return 0. See the sample input for further explanation.*/

function stringyThingy(strArr, numArr) {
    if (numArr.length === 0) {
        return 0;
    }
    let count = 0;
    for (let str of strArr) {
        if (numArr.includes(str.length)) {
            count++;
        }
    }
    return count;
}

let strings = ["hello", "world", "tim", "a", "it"];
let numbers = [5, 1];
console.log(stringyThingy(strings, numbers)); // Output: 3