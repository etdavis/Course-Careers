/*Question 1
Write a function that accepts an array of numbers and returns the 2 largest numbers in the
array. The largest values should be returned in an array where the first element in the array is
the larger of the two elements. You may assume you will always be given an array that contains
at least 2 numbers.*/

function twoLargestNumbersSort(arr) {
    const sortedArr = arr.sort((a, b) => b-a);
    return [sortedArr[0], sortedArr[1]];
}

console.log(twoLargestNumbersSort([7, -2, 3, 4, 5, 1, 23, 31, 23])); // [31, 23]
console.log(twoLargestNumbersSort([7, -2, 3, 4, 5, 1, 23, 23])); // [23, 23]
// This is definitely the simplest and easiest way to do it, but not the most efficient because sorting the whole array is unncecessary.


function twoLargestNumbersNoSort(arr) {
    let largest = -Infinity;
    let secondLargest = -Infinity;
    for (let num of arr) {
        if (num > largest) {
            secondLargest = largest;
            largest = num;
        }
        else if (num > secondLargest) {
            secondLargest = num;
        }
    }
    return [largest, secondLargest];
}
console.log(twoLargestNumbersNoSort([7, -2, 3, 4, 5, 1, 23, 31, 23])); // [31, 23]
console.log(twoLargestNumbersNoSort([7, -2, 3, 4, 5, 1, 23, 23])); // [23, 23]