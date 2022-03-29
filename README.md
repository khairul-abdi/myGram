# myGram

## Tools
* Go Programming Language
* Gin Gonic Framework
* Postgresql
* Orm Gorm

## Prepare


# Usage
First thing first we have to create db :
```
    create database postgresql with the same name like dicrectory package/database/config.go
```
and then install all of the dependencies by typing this command in terminal :
```
    go mod tidy
```
and then we are ready to execute the server by this following command :
```
    go run main.go
```

# Features
## User
### Create User Register
```
curl --location --request POST 'http://localhost:9001/users/register' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-raw '{
    "email": "nia@gmail.com",
    "username": "nia",
    "password":"123456",
    "age":9
}'
```
### Response
```
{
  "code": 201,
  "data": {
    "age": 9,
    "email": "nia@gmail.com",
    "id": 8,
    "username": "nia"
  },
  "message": "Success insert data"
}
```
### Create User Login
```
curl --location --request POST 'http://localhost:9001/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "nia@gmail.com",
  "password":"123456"
}'
```
### Response
```
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pYUBnbWFpbC5jb20iLCJpZCI6OH0.bMWIYZRb7g78dgKr_0YfkWmYGLMFFfNMQcHMbfRJXww"
  },
  "message": "Success Create Token"
}
```
### Update User
```
curl --location --request PUT 'http://localhost:9001/users/8' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pYUBnbWFpbC5jb20iLCJpZCI6OH0.bMWIYZRb7g78dgKr_0YfkWmYGLMFFfNMQcHMbfRJXww' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "niakurniati@gmail.com",
    "username": "nia kurniati"
}'
```
### Response
```
{
  "code": 200,
  "data": {
    "age": 9,
    "email": "niakurniati@gmail.com",
    "id": 8,
    "updated_at": "2022-03-29T08:51:36.031953569+07:00",
    "username": "nia kurniati"
  },
  "message": "Success Update user"
}
```
### Delete User
```
curl --location --request DELETE 'http://localhost:9001/users/8' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pYUBnbWFpbC5jb20iLCJpZCI6N30.ZwNE9OAbPQCM3t18jkBf-jwLQxS4ppIWd0D1Wmt04IQ' \
--data-raw ''
```
### Response
```
{
  "code": 200,
  "message": "Your user has been successfully deleted"
}
```
## Photo
### Create Photo
```
curl --location --request POST 'http://localhost:9001/photos' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng' \
--header 'Content-Type: application/json' \
--data-raw '{
  "title":"Kerja",
  "caption": "Nuansa wfh dengan segelas kopi",
  "photourl":"www.image.com"
}'
```
### Response
```
{
  "code": 200,
  "data": {
    "caption": "Nuansa wfh dengan segelas kopi",
    "created_at": "2022-03-29T08:55:08.192377488+07:00",
    "id": 1,
    "photo_url": "www.image.com",
    "title": "Kerja",
    "user_id": 5
  },
  "message": "Success create photo"
}
```
### Get All Photo
```
curl --location --request GET 'http://localhost:9001/photos' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng'
```
### Response
```
{
  "code": 200,
  "data": [
    {
      "Id": 1,
      "Title": "Kerja",
      "Caption": "Nuansa wfh dengan segelas kopi",
      "PhotoUrl": "www.image.com",
      "UserId": 5,
      "CreatedAt": "2022-03-29T08:55:08.192377+07:00",
      "UpdatedAt": "2022-03-29T08:55:08.192377+07:00",
      "User": {
        "id": 5,
        "username": "iwan",
        "email": "iwan@gmail.com"
      }
    }
  ],
  "message": "Success"
}
```
### Update Photo
```
curl --location --request PUT 'http://localhost:9001/photos/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng' \
--header 'Content-Type: text/plain' \
--data-raw '{
  "title":"Kerja Remote (wfh)",
  "caption": "Nuansa wfh dengan segelas kopi",
  "photourl":"www.image.com"
}'
```
### Response
```
{
  "code": 200,
  "data": {
    "caption": "Nuansa wfh dengan segelas kopi",
    "id": 1,
    "photo_url": "www.image.com",
    "title": "Kerja Remote (wfh)",
    "updated_at": "2022-03-29T08:59:16.536162019+07:00",
    "user_id": 5
  },
  "message": "Success update photo"
}
```
### Delete Photo
```
curl --location --request DELETE 'http://localhost:9001/photos/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng'
```
# Response
```
{
  "code": 200,
  "message": "Your photo has been successfully deleted"
}
```

## Comments
### Create Comment
```
curl --location --request POST 'http://localhost:9001/comments' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiZGlAZ21haWwuY29tIiwiaWQiOjZ9.No-Ze2Pnr-P96dE-c_gcNLhgUW2QTEiDaZ3ZES8m6jA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message":"keren banget photonya iwan",
    "photo_id": 2
}'
```
# Response
```
{
  "code": 200,
  "data": {
    "created_at": "2022-03-29T09:12:27.191400465+07:00",
    "id": 2,
    "message": "keren banget photonya iwan",
    "photo_id": 2,
    "user_id": 6
  },
  "message": "Success create comment"
}
```
### Gets Comment
```
curl --location --request GET 'http://localhost:9001/comments' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiZGlAZ21haWwuY29tIiwiaWQiOjZ9.No-Ze2Pnr-P96dE-c_gcNLhgUW2QTEiDaZ3ZES8m6jA'
```
### Response
```
{
  "code": 200,
  "data": [
    {
      "Id": 2,
      "UserId": 6,
      "PhotoId": 2,
      "Message": "keren banget photonya iwan",
      "CreatedAt": "2022-03-29T09:12:27.189855+07:00",
      "UpdatedAt": "2022-03-29T09:12:27.189855+07:00",
      "User": {
        "id": 6,
        "username": "abdi",
        "email": "abdi@gmail.com"
      },
      "Photo": {
        "id": 2,
        "title": "Kerja",
        "caption": "Nuansa wfh dengan segelas kopi",
        "photo_url": "www.image.com",
        "user_id": 5
      }
    }
  ],
  "message": "Success"
}
```
### Update Comment
```
curl --location --request PUT 'http://localhost:9001/comments/2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiZGlAZ21haWwuY29tIiwiaWQiOjZ9.No-Ze2Pnr-P96dE-c_gcNLhgUW2QTEiDaZ3ZES8m6jA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message":"keren photonya broooo iwan"
}'
```
### Response
```
{
  "code": 200,
  "data": {
    "caption": "",
    "id": 2,
    "message": "keren photonya broooo iwan",
    "photo_url": "",
    "title": "",
    "updated_at": "2022-03-29T09:14:45.521264973+07:00",
    "user_id": 6
  },
  "message": "Success update comment"
}
```
### Delete Comment
```
curl --location --request DELETE 'http://localhost:9001/comments/2' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFiZGlAZ21haWwuY29tIiwiaWQiOjZ9.No-Ze2Pnr-P96dE-c_gcNLhgUW2QTEiDaZ3ZES8m6jA'
```
### Response
```
{
  "code": 200,
  "message": "Your comment has been successfully deleted"
}
```
## Social Media
### Create Social Media
```
curl --location --request POST 'http://localhost:9001/socialmedias' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name": "Nia",
  "socialmediaurl": "www.fb.com"
}'
```
# Response
```
{
  "code": 200,
  "data": {
    "created_at": "2022-03-29T09:03:34.024120829+07:00",
    "id": 1,
    "name": "Nia",
    "social_media_url": "www.fb.com",
    "user_id": 5
  },
  "message": "Success create photo"
}
```
### Gets Social Media
```
curl --location --request GET 'http://localhost:9001/socialmedias' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng'
```
# Response
```
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "Name": "Nia Kurniari",
      "SocialMediaUrl": "www.fb.com/nia_kurniati",
      "UserId": 5,
      "CreatedAt": "2022-03-29T09:03:34.02412+07:00",
      "UpdatedAt": "2022-03-29T09:05:49.39292+07:00",
      "User": {
        "id": 5,
        "username": "iwan",
        "email": "iwan@gmail.com"
      }
    }
  ],
  "message": "Success"
}
```
### Update Social Media
```
curl --location --request PUT 'http://localhost:9001/socialmedias/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Nia Kurniari",
    "socialmediaurl": "www.fb.com/nia_kurniati"
}  '
```
### Response
```
{
  "code": 200,
  "data": {
    "created_at": "2022-03-29T09:05:49.392920492+07:00",
    "id": 1,
    "name": "Nia Kurniari",
    "social_media_url": "www.fb.com/nia_kurniati",
    "user_id": 5
  },
  "message": "Success update social media"
}
```
### Delete Social Media
```
curl --location --request DELETE 'http://localhost:9001/socialmedias/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml3YW5AZ21haWwuY29tIiwiaWQiOjV9.PGODkrPszlKSOQxgbX20LUyKAqsGIwOAjN8UpR8f7ng' \
--data-raw ''
```
### Response
```
{
  "code": 200,
  "message": "Your social media has been successfully deleted"
}
```










