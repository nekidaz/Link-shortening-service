Shorten URL Service
This is a simple service that shortens URLs.

Getting Started
Clone this repository.
Create a .env file in the root directory with the following contents:
makefile
Copy code
DATABASE_URL=<your-database-url>
Replace <your-database-url> with the URL of your PostgreSQL database.
Run the following command to install the required dependencies:
go
Copy code
go mod download
Run the following command to start the service:
go
Copy code
go run main.go
The service will be accessible at http://localhost:8080.
API Endpoints
POST /api/shorten
This endpoint takes a JSON object with a single property url, which is the URL to be shortened. It returns a JSON object with a single property short_url, which is the shortened URL.

Example request:

bash
Copy code
POST /api/shorten HTTP/1.1
Content-Type: application/json

{
  "url": "https://www.google.com/search?q=golang"
}
Example response:

css
Copy code
HTTP/1.1 200 OK
Content-Type: application/json

{
  "short_url": "http://localhost:8080/a1b2c3d4"
}
GET /:short_id
This endpoint takes a shortened ID and redirects the user to the original URL.

Example request:

bash
Copy code
GET /a1b2c3d4 HTTP/1.1
Example response:

arduino
Copy code
HTTP/1.1 301 Moved Permanently
Location: https://www.google.com/search?q=golang
Dependencies
This service depends on the following packages:

github.com/gin-gonic/gin
github.com/joho/godotenv
github.com/jinzhu/gorm
github.com/jinzhu/gorm/dialects/postgres
github.com/rs/xid
These dependencies will be automatically installed when running go mod download.

Database Setup
This service uses PostgreSQL as its database. Make sure you have PostgreSQL installed and running on your machine before running the service.

To create the required database schema, run the following command:

graphql
Copy code
psql <your-database-url> < schema.sql
Replace <your-database-url> with the URL of your PostgreSQL database. The schema.sql file can be found in the db directory.
