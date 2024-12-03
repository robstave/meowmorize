#  Architecture Overview

The Project is composed of two parts,  a frontend and a backend.
The backend is golang.  Im using ECHO listening to port 8789.
There is a swagger as well /swagger/index.html

## backend golang

To run the backend I use my helper file.
`./helper run` but its basically
`go run cmd/main/main.go`
all the go code is in /internal outside of the cmd

### Database
The database is sqlite.  it currently is reading from meowmorize.db at the root directory

to start over, just delete the database.

## Frontend

The front end is in /meowmorize-frontend and is react
This is listening to port 8999 as defined in the .env file ( not 3000)

