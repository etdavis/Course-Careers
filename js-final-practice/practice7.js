/* Question 7 (No Test Cases)
The TinyTown Library wants to digitize their book inventory. They are looking for a system that
allows them to keep track of all the books they have, who borrowed them, and when they are
due back.
Your task is to create a JavaScript class named Library. The Library class should have:
- A constructor that takes no arguments and initializes an empty array of books.
- A books property that is an array of objects, where each object represents a book and
has properties for title (string), author (string), borrower (string, initially null).
- A method addBook(title, author) that creates a new book object and adds it to the books
array.
- A method checkoutBook(title, borrower) that finds the book by title and sets the
borrower.
- A method returnBook(title) that finds the book by title and resets the borrower to null.
- A method listBorrowedBooks() that returns an array of books that are currently
borrowed. */

class Library {
    books;

    constructor() {
        this.books = [];
    }

    addBook(title, author) {
        this.books.push({title: title, author: author, borrower: null});
    }

    checkoutBook(title, borrower) {
        for (let book of this.books) {
            if (book.title === title) {
                book.borrower = borrower;
                return;
            }
        }
        console.log(`Sorry, ${title} is not available in this library.`);
    }

    returnBook(title) {
        for (let book of this.books) {
            if (book.title === title) {
                book.borrower = null;
                return;
            }
        }
        console.log(`Sorry, ${title} does not belong to this library.`);
    }

    listBorrowedBooks() {
        // My solution
        let borrowedBooks = [];
        for(let book of this.books) {
            if(book.borrower !== null) {
                borrowedBooks.push(book);
            }
        }
        return borrowedBooks;
        // Another solution I found that I wish I thought of
        //return this.books.filter(book => book.borrower !== null);
        //return this.books.filter(book => book.borrower);
    }
}

let library = new Library();
library.addBook("The Great Gatsby", "F. Scott Fitzgerald");
library.addBook("Moby Dick", "Herman Melville");
library.addBook("1984", "George Orwell");
library.checkoutBook("The Great Gatsby", "John Doe");
library.checkoutBook("1984", "Jane Doe");
console.log(library.listBorrowedBooks());
// Outputs: 
// [{title: "The Great Gatsby", author: "F. Scott Fitzgerald", borrower: "John Doe"},
// {title: "1984", author: "George Orwell", borrower: "Jane Doe"}]
library.returnBook("The Great Gatsby");
console.log(library.listBorrowedBooks());
// Outputs: [{title: "1984", author: "George Orwell", borrower: "Jane Doe"}]