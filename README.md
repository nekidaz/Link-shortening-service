## Shorten URL Service

This is a simple service that shortens URLs.

### Getting Started

1. Clone this repository.
2. Create a `.env` file in the root directory with the following contents:

```
DATABASE_URL=<your-database-url>
```
Replace `<your-database-url>` with the URL of your PostgreSQL database.
3. Run the following command to install the required dependencies:
```
go mod download
```
4. Run the following command to start the service:
```
go run main.go
```

5. The service will be accessible at `http://localhost:8080`.

### API Endpoints

#### POST /shorten
This endpoint takes a JSON object with a single property `url`, which is the URL to be shortened. It returns a JSON object with a single property `short_url`, which is the shortened URL.

Example request:
```
POST /api/shorten HTTP/1.1
Content-Type: application/json

{
"url": "https://www.google.com/search?q=golang"
}
```
Example response:
```
HTTP/1.1 200 OK
Content-Type: application/json

{
"short_url": "http://localhost:8080/a1b2c3d4"
} 
```

#### GET /:short_id
This endpoint takes a shortened ID and redirects the user to the original URL.

Example request:

```GET /a1b2c3d4 HTTP/1.1```

Example response:

```
HTTP/1.1 301 Moved Permanently
Location: https://www.google.com/search?q=golang
```
### Dependencies

This service depends on the following packages:

- `github.com/gin-gonic/gin`
- `github.com/joho/godotenv`
- `github.com/jinzhu/gorm`
- `github.com/jinzhu/gorm/dialects/postgres`
- `github.com/rs/xid`

These dependencies will be automatically installed when running `go mod download`.

### Database Setup

This service uses PostgreSQL as its database. Make sure you have PostgreSQL installed and running on your machine before running the service.

To create the required database schema, run the following command:

```psql <your-database-url> < schema.sql```


Replace `<your-database-url>` with the URL of your PostgreSQL database. The `schema.sql` file can be found in the `db` directory.





