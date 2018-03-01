# URL Shortener

This is a small application that shorten your URL in order to share it easily.

## How it works

The software persist the original URL in the database and retrieves the inserted id. Then, we create a base62 encode on top of the returned id. After that, we update the row with the encoded version.

## Stack

- Go
- MySQL
- React
- Node

## Up and Running

First of all, you must clone and download this repository.

After that, you must do the following steps to get the MySQL working properly:
- Create a MySQL database named `url_shortener`.
- Run the `dump.sql` script in order to create the table structure

Assuming that you already have Go installed, you must type the following:
- `$ go run main.go`

This will create a server listening on port 3000.

At this point, you can play around with the application using Postman or a similar service.

Nos, you can go to the static folder and run:

- `$ npm i`
- `$ npm start`

The lines above will install the frontend dependencies, build the frontend and start watching for changes. 

