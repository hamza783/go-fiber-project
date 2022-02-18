# go-fiber-project
Small project to learn go


Run the project: <b>go run main.go</b>

home page
GET endpoint: http://localhost:3000
<br />
<br />

get all books
GET endpoint: http://localhost:3000/api/v1/book
<br />
<br />

get a book by id
GET endpoint: http://localhost:3000/api/v1/book/{id}
<br />
<br />

add a book
POST endpoint: http://localhost:3000/api/v1/book
<br />
BODY: 
	title: "Title of the book"
	author: "Author"
	rating: 1
<br />
<br />

delete a book
DELETE endpoint: http://localhost:3000/api/v1/book/{id}
