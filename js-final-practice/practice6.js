/* Question 6
RapidText Inc. specializes in text processing and analytics. They have developed a new tool
that can analyze a large text string and perform a series of operations on it. The operations are
provided as an array of functions that each take a string as input and return a modified string as
output. The functions are applied in the order they are given in the array.
The problem is that some of these operations are computationally intensive and can take a long
time to run, so RapidText wants to be able to apply these operations only to the portions of the
text that actually need them.
Your task is to write a JavaScript function that takes a text string, an array of operations
(functions), and an array of ranges (each range being an array of two numbers indicating start
and end positions in the text). The function should apply each operation to only the text within
the specified ranges, and return the modified text.
Each operation will be a function of the form function(text) {  ...  } and will return a string.
Note: Ranges are inclusive and are based on zero-indexing. If a range is [2, 5], it refers to the
part of the string from the 3rd to the 6th character, inclusive. */

function processText(text, operations, ranges) {
    let newText = text;
    for (let range of ranges) {
        let slice = newText.slice(range[0], range[1] + 1);
        for (let operation of operations) {
            slice = operation(slice);
        }
        newText = newText.replace(newText.slice(range[0], range[1] + 1), slice);
    }
    return newText;
}

let text = "Hello, World!";
let operations = [
    
    function(text) { return text.replace(/o/g, "0"); },
    function(text) { return text.toUpperCase(); },
];
let ranges = [[0, 4], [7, 12]];
console.log(processText(text, operations, ranges)); // HELL0, W0RLD!
//Note: The sample output provided in the instructions has the function order reversed, but still expects "HELL0, W0RLD!" even though this wouldn't be the case because all the lowercase 'o's would be converted to uppercase before they were replaced with '0's, causing the replace function to do nothing. I reversed the order here to make the exercise more meaningful