# Golang SocialMedia API application

This API was made for study purposes.

The entire application is contained within the `main.go` file.

## Dependencys

    apt-get install mysql-server

- Create and configure an user for the application;
- In the `migrations/` folder there are some queries to initialize the database and its tables, as some examples to populate it;


## Run the app

    go run main.go

# REST API

## Login

### Request

`POST /login`

### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Connection: close
    Content-Type: text/plain; charset=utf-8
    Content-Length: 2

    [USER_TOKEN_STRING]

## Create a new User

### Request

`POST /users`

### Body

  {
    "name": "User 1",
    "nick": "User 1",
    "email": "user@gmail.com",
    "password": "123456"
  }

### Response

    HTTP/1.1 201 CREATED
    Date: Wed, 03 Apr 2024 18:17:58 GMT
    Status: 201 CREATED
    Connection: close
    Content-Type: application/json

    {"ID": 1,"Name": "User 1","Nick": "User 1","Email": "user@gmail.com","Password": "$2a$10$thkaj0EOoTyeNKif3RjYr.9SDvvDLO460GT1wTJTnROyd2Mga.fY.","CreatedAt": "0001-01-01T00:00:00Z"}
    
## Get All Users (or filter by Name/Nick)

### Request

- `GET /users`
- `GET /users?user=[TEXT_TO_FILTER_FOR]` 

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json

    [{"ID":1,"Name":"User 1","Nick":"user_1","Email":"user_1@gmail.com","Password":"","CreatedAt":"2024-04-03T11:47:13-03:00"},{"ID":2,"Name":"User 2","Nick":"user_2","Email":"user_2@gmail.com","Password":"","CreatedAt":"2024-04-03T11:47:13-03:00"}]

## Get a User by ID

### Request

`GET /users/id`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json

    {"ID":1,"Name":"User 1","Nick":"user_1","Email":"user_1@gmail.com","Password":"","CreatedAt":"2024-04-03T11:47:13-03:00"}

## Update a User

### Request

`PUT /users/{userId}`

#### Authentication Required [Bearer Token]

### Body

  {
    "name": "User 1",
    "nick": "User 1",
    "email": "user@gmail.com",
    "password": "123456"
  }

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json


## Delete a User

### Request

`DELETE /users/{userId}`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json

## Follow a User

### Request

`POST /users/{userId}/follow`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json

    
## Unfollow a User

### Request

`POST /users/{userId}/unfollow`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json

## Get User's followers

### Request

`POST /users/{userId}/followers`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    
    [{"ID":2,"Name":"User 2","Nick":"user_2","Email":"user_2@gmail.com","Password":"","CreatedAt":"2024-04-03T11:47:13-03:00"},{"ID":3,"Name":"User 3","Nick":"user_3","Email":"user_3@gmail.com","Password":"","CreatedAt":"2024-04-03T11:47:13-03:00"}]

## Get who the User is following

### Request

`POST /users/{userId}/following`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    
    [{"ID":1,"Name":"User 1","Nick":"user_1","Email":"user_1@gmail.com","Password":"","CreatedAt":"2024-04-03T11:47:13-03:00"}]

## Update User's Password

### Request

`POST /users/{userId}/update-password`

#### Authentication Required [Bearer Token]

### Body

  {
    "current": "old password",
    "new": "new password"
  }

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json
    Content-Length: 0

## Create a new Post

### Request

`POST /posts`

### Body

  {
    "title": "Title text",
    "content": "content text"
  }

### Response

    HTTP/1.1 201 CREATED
    Date: Wed, 03 Apr 2024 18:17:58 GMT
    Status: 201 CREATED
    Connection: close
    Content-Type: application/json

    {"ID":4,"Title":"usuario2@gmail.com","Content":"user.2","AuthorID":4,"AuthorNick":"","Likes":0,"CreatedAt":"0001-01-01T00:00:00Z"}

## Get All Posts from a user and those he follows

### Request

`GET /posts`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json

    [{"ID":4,"Title":"usuario2@gmail.com","Content":"user.2","AuthorID":4,"AuthorNick":"User.2","Likes":0,"CreatedAt":"2024-04-03T15:56:44-03:00"}]

## Get a Post by ID

### Request

`GET /posts/id`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json

    {"ID":4,"Title":"usuario2@gmail.com","Content":"user.2","AuthorID":4,"AuthorNick":"User.2","Likes":0,"CreatedAt":"2024-04-03T15:56:44-03:00"}


## Update a Post

### Request

`PUT /posts/{postId}`

#### Authentication Required [Bearer Token]

### Body

  {
    "tile": "title text",
    "content": "content text"
  }

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json


## Delete a Post

### Request

`DELETE /posts/{postId}`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json

## Like a Post

### Request

`GET /posts/{postId}/like`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json

## Unlike a Post

### Request

`GET /posts/{postId}/unlike`

#### Authentication Required [Bearer Token]

### Response

    HTTP/1.1 204 NO CONTENT
    Date: Thu, 24 Feb 2011 12:36:31 GMT
    Status: 204 NO CONTENT
    Connection: close
    Content-Type: application/json
