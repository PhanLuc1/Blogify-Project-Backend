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
## 1.2. Create post 

### 1.2.1. Create `post` table

- Endpoint: `/post/creating`
- Method: `POST`

#### Request header

```json
{
  "token" :"..." 
}
```
#### Request body
```json
{
   "caption" : "....."
   "postImage" : [
     {
                "id": 2,
                "imageURL": "abcxyz",
                "Description": "abcxyz"
            },
            {
                "id": 4,
                "imageURL": "123456",
                "Description": "1234444"
            },
   ] 
}
```
#### Response 

```json
{
  "message": "200" // int
}
```

### 1.2.2. Create `Comment` table

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
