
# Introduction

This is a TO-DO webapp that allows users to login/register for an account and to add, delete, edit and update tasks as they wish.

As the user is using the app, the javascript fetch API makes various calls to the backend codebase written in GO. 
The backend has functions that deal with the user's commands using the GO fiber framework. Note that Gorm (an ORM library) and a Postgresql docker container are used to implement the database. As well it implements sessions with redis for authentication.

Note: Unit tests for repository functions can be found in the ```repo_test.go``` file.

# To get started

- Make sure you have Golang and Docker installed
- Clone the repository
- Run 
```"docker compose build"``` followed by ```"docker compose up" ```
- The app should now be hosted at ```localhost:3000```

Thats it!


Interface:

<img width="916" alt="image" src="https://user-images.githubusercontent.com/89322519/227745124-09802972-40f4-4240-8673-a807df41cb9b.png">

<img width="950" alt="image" src="https://user-images.githubusercontent.com/89322519/227745158-80a97072-9c50-4219-8e7b-82118b16652f.png">
