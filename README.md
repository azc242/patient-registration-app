# Patient Registration App in GoLang and React.js

## The App

This app was created for the [Done.](https://donefirst.com) software engineering internship process. I could've chosen any tech stack, but since their tech stack is in Go and React, I chose it for this project.

This app allows patients to submit their name, date of birth, phone number, email and address. Their registration time is automatically recorded.

The admin can log in using their credentials and view all the patients in the queue, ordered by the time they registered. By default, the admin has a username of `0`, and password of `0Admin`. However the MySQL database has a table for admins, so new admins can be registered.

Demo: https://youtu.be/pS9wZu33LmY

Patient registration:
![Registration view](https://i.imgur.com/J6tLW29.png)Admin view:
![Admin view](https://i.imgur.com/hW07mqg.png)

## Backend/API

The API has 2 routes and 3 ways to utilize them

1. `/api/patients`
   - GET: fetches all the patients in the databasem returned as JSON
   - POST: creates a patient
     - Requires a name, date of birth, phone number, and email (name, dob, phone, email)
     - Example (raw body data):`{"name": "Randy Carlson","dob": "2/8/1992","phone": "123-456-7890","email","rc23@test.net"}`
2. `/api/login`
   - POST: validates the user trying to log in (only the admin can log in)
     - Requires a Username and Password

## Requirements

- GoLang (I used `go1.15.6` for this project)
- MySQL
- npm and React.js

## Environment Variables

There are only two environment variables required to run the backend, which are located in a `.env` file not available in this repository

1. MYSQL_USERNAME (the username for MySQL)
2. MSQL_PASSWORD (the password for MySQL)

## Running The App

### Set up the server

To run the server locally, run `go build && ./Done`. This will build the application and execute it. The server will run on http://localhost:8080/

#### Running the client side

To run the React.js frontend, first navigate into the client directory by running `cd client` in the root directory. Then, run `npm start`. You will be able to access it by going to http://localhost:3000/

## What I Used

- [guuid](github.com/google/uuid) for generating unique patient IDs
- [handlers](github.com/gorilla/handlers) and [mux](github.com/gorilla/mux) for dealing with http requests and CORS
- [go-sql-driver](github.com/go-sql-driver/mysql) for MySQL queries
- [godotenv](github.com/joho/godotenv) for environment variables
- [Axios](https://www.npmjs.com/package/axios) for making GET/POST requests on the client side
- [Milligram](https://milligram.io/) for the clean, minimalistic CSS
