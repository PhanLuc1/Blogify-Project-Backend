# API

# 1. post

## 1.1. Get all posts or post

- Endpoint: `/posts` 
- Endpoint: `/posts/{postid}`
- Method: `GET`
### Response

```json
[
    {
        "id": 1,
        "user": {
            "id": 1,
            "email": "...",
            "username": "...",
            "password": "...",
            "state": true,
            "avataImage": "123"
        },
        "caption": "...",
        "postImages": [
            {
                "id": 1,
                "imageURL": "...",
                "Description": "...."
            },
            ...
        ],
        "createAt": "0001-01-01T00:00:00Z",
        "countCommet":" 123 int ",
        "comments": [
            {
                "id": 1,
                "user": {
                    "id": 3,
                    "email": "...",
                    "username": "....",
                    "password": "",
                    "state": true,
                    "avataImage": "...."
                },
                "postId": 1,
                "parentCommentId": 
                    {
                        "id": 3,
                        "user": {
                            "id": 3,
                            "email": "lucptt123@gmail.com",
                            "username": "lucphan123",
                            "password": "",
                            "state": true,
                            "avataImage": "123"
                        },
                },
                ....
                ],
                "reaction": {
                    "users": null,
                    "countReaction": 0
                }
            },
            ....
        ],
        "reaction": {
            "users": [
                {
                    "id": 1,
                    "email": "...",
                    "username": "...",
                    "password": "",
                    "state": true,
                    "avataImage": "..."
                },
                {
                    "id": 3,
                    "email": "....",
                    "username": ".....",
                    "password": "",
                    "state": true,
                    "avataImage": "..."
                }
            ],
            "countReaction": 2
        }
    },
    .....
]
```
### Response 
``` json
{
  "message" : "200"
}
```
## 1.2. Create post  `NEW UPDATE`

### 1.2.1. Create `post` table

- Endpoint: `/posts/creating` 
  sửa endpoint sau
- Method: `POST`

#### Request header


```
{
  "token" :"..." 
}
```
#### Request body
```json 
{
   //form-data 
   "caption" :" .... " // key
   "images"  :" .... " // key and values upload image
}
```
#### Response 

```json
{
  "message": "created" // int
}
```

## 1.2.2. Create `Comment` table
- Endpoint : `/posts/{postid}/comment`
- Method : "Post"
### Request Header
```json
{
  "Token" : " .... "
}
```
### Request body if not subcomment 
```json  
{
  "content" : "....."
}
```
### Request body if comment is subcomment 
```json
{
    "content" : "đẹp vcl",
    "parentCommentId" :  {
        "Int64": 3,
        "Valid": false
        }
}
```
### Response 
```json 
{
  "message" : "created"
}
```
## 1.2.3 `Reaction` Post
- Endpoint : `/posts/{postid}/reaction`
- Method : `Get`
### Request header 
```json
{
  "token" :  " ... "
}
```
### Response status 200
## 1.2.4 `Reaction` Comment
- Endpoint : `/posts/{commentid}/reaction`
- Method : `GET`
```json
{
  "token" :  " ... "
}
```
### Response status `200`
## 1.2.4 Delete `POST` `NEWUPDATE`
- endpoint: `posts/{postid}/delete`
- method: `DELETE`
### Request header 
```json 
{
  "Token" : "..."
}
```
### Response `200` `403` when post not be created by currentUser
## 1.2.5 Delete `Comment` `NEWUPDATE`
- endpoint: `posts/{commentId}/delete`
- method: `DELETE`
### Request header 
```json 
{
  "Token" : "..."
}
```
### Response `200` `403` when comment not be created by currentUser

# 2.User
## 2.1 Registration
- Endpoint : `/user/registration/`
- Method : `Post`
### Request body 
```json
{
   "email": "..... ",
   "password" : "..."
}
```
### Response
```json
{
  "message" : "201"
}
```
## 2.2 Sign-in
- Endpoint :`/user/sign-in`
- Method: `Post`
### Request body 
```json
{
   "email": "..... ",
   "password" : "..."
}
```
### Response 
```json
{
  "token" : " ... "
}
```
## 2.3 GetCode
- Endpoint :`/auth/code`
- Method: `Post`
```json 
{
  "email"
}
```
### Response
```json 
{
  "status" : "200"
}
```
## 2.4 Auth code
- Endpoint :`/auth`
- Method: `post`
### Request body
```json
{
  "email" : "...",
  "code" : "..."
}
```
### Response 
- 400 or 200 
## 2.5 User Info
- Endpoint : `/user`
- Method: `Get`
### Request header
```json
{
  "token" : "... "
}
```
### Response 
```json
{
    "id": 6,
    "email": "lucptt22@gmail.com",
    "username": "",
    "password": "",
    "state": true,
    "avataImage": ""
}
```
## 2.6 User `update`
- Endpoint : `/user/update`
- Method : `POST`
### Request header 
```json
{
  "token" : " .... "
}
```
### Request body
- 1 or 2 or null properties
```json
{
  "username" : "...",
  "avataImage": "..."
}
```
### Response status `200`
## 2.7 Get `User` by ID
- endpoint : `users/{userid}`
- method : `GET`
### Request header 
```json 
{
  "token" : "..."
}
```
### Response status `200`
## 2.8 Follow `user` `UPDATEE`
- endpoint : `users/{userid}/follower`
- method : `GET`
### Request header 
```json 
{
  "token" : "..."
}
```
### Response status `200`

## Get 2.9 `Image` 
- endpoint : `image?image= ...`
- method : `GET`
### Request header
```json
{
  "token" : "..."
}
```
### Response image status 200
## 2.10 Add `Biography Username` `NEW UPDATE !!!`
- endpoint : `user/updation`
- medthod : `POST`
### Request header
```json
{
  "token" : "..."
}
```
### Request body 
```json
{
  "biography" : "....",
  "username" : "....",
}
```
### Status response `200`
## 2.11 Upload `Avatar Image` `NEW UPDATE`
- endpoint: `user/updation/avatar`
- medthod : `POST`
### RequestHeader
```json
{
  "token" : " ..."
}
```
### RequestBody
```json
{
  //form-data
  "images" : "..."
}
```
### Response status `200`
## 2.12 Get All `AnotherUser` `NEW UPDATE`
- endpoint : `/users`
- method : `GET`
### Request header 
```json
{
  "token" : " ..."
}
```
### Response status  `200` 
```json
[
    {
        "id": 1,
        "username": "lucphan uchihaha",
        "follower": 1,
        "following": 1,
        "avatarImage": "beauty_20201101212051.jpg",
        "state": true
    },
    {
        "id": 3,
        "username": "",
        "follower": 2,
        "following": 0,
        "avatarImage": "",
        "state": true
    }
]
```
## 2.13 Set `STATE` PRIVATE OR PUBLIC `NEWUPDATE`
- endpoint: `/user/update/state`
- method : `GET`
### Request header 
```json 
{
  "Token" : "..."
}
```
### Response `200`