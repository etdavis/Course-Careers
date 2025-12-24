/* Question 3
Write a function that accepts a plain text password and returns the relative strength of the
password based on the following criteria.
Types of Passwords:
- Weak Password: 0 - 2 points.
- Medium Password: 3 points.
- Strong Password: 4 points.
- Very Strong Password: 5 Points.
Criteria For Points:
- Length: 1 point if more than 10 characters
- Capital letters: 1 point if at least 1 capital letter is present
- Lowercase letters: 1 point if at least 1 lowercase letter is present
- Numbers: 1 point if at least 1 number is present
- Special characters: 1 point if any of the following characters are present: %, ^, &, *, (, ),
@, !, #, $ (i.e all the special characters accessible from the number keys 1, 2, 3, 4... etc)
Your function should return “weak”, “medium”, “strong” or “very strong” to classify the given
password.*/

function passwordStrength(password) {
    const uppers = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'];
    const lowers = ['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'];
    const numbers = ['0','1','2','3','4','5','6','7','8','9'];
    const specialChars = ['%','^','&','*','(',')','@','!','#','$'];
    const charsArr = [uppers, lowers, numbers, specialChars];
    
    let strength = 0;
    if(password.length > 10) {
        strength++;
    }
    for (chars of charsArr) {
        for (char of chars) {
            if(password.includes(char)) {
            strength++;
            break;
            }
        }
    }

    switch (strength) {
        case 3:
            return "medium";
        case 4:
            return "strong";
        case 5:
            return "very strong";
        default:
            return "weak";
    }
}

const pass = "helloworld123*";
console.log(passwordStrength(pass)); // Output: strong