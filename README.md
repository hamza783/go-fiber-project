# go-fiber-project
Small project to learn go


Run the project
	go run main.go

home page
GET endpoint: http://localhost:3000


get all books
GET endpoint: http://localhost:3000/api/v1/book


get a book by id
GET endpoint: http://localhost:3000/api/v1/book/{id}


add a book
POST endpoint: http://localhost:3000/api/v1/book
BODY: 
	title: "Title of the book"
	author: "Author"
	rating: 1


delete a book
DELETE endpoint: http://localhost:3000/api/v1/book/{id}
