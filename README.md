## Patient Registration App in GoLang and React.js

### Inspiration

This app was created for the [Done.](https://donefirst.com) software engineering internship process. I could've chosen any tech stack, but since their tech stack is in Go and React, I chose it for this project. I also wanted to use Go because I had never used it before prior to this.

### Backend/API

The API has 2 routes and 3 ways to utilize them

1. `/api/patients`
   - GET: fetches all the patients in the databasem returned as JSON
   - POST: creates a patient
     - Requires a name, date of birth, phone number, and email (name, dob, phone, email)
     - Example (raw body data):`{"name": "Randy Carlson","dob": "2/8/1992","phone": "123-456-7890","email","rc23@test.net"}`
2. `/api/login`
   - POST: validates the user trying to log in (only the admin can log in)
     - Requires a Username and Password

### Requirements

- GoLang (I used `go1.15.6` for this project)
- MySQL
- React.js

### Environment Variables

There are only two environment variables required to run the backend, which are located in a `.env` file not available in this repository

1. MYSQL_USERNAME (the username for MySQL)
2. MSQL_PASSWORD (the password for MySQL)