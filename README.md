
# Introduction

This is a TO-DO webapp that allows users to login/register for an account and to add, delete, edit and update tasks as they wish.

As the user is using the app, the javascript fetch API makes various calls to the backend codebase written in GO. 
The backend has functions that deal with the user's commands using the GO fiber framework. Note that Gorm (an ORM library) and a Postgresql docker container are used to implement the database. As well it implements sessions with redis for authentication.

Unit tests can be found in the file methods_test.go

# To get started

- Make sure you have Golang and Docker installed
- Clone the repository
- Run 
```"docker compose build"``` followed by ```"docker compose up" ```
- The app should now be hosted at ```localhost:3000```

Thats it!
