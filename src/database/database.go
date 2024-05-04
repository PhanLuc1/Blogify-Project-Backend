package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/models"
	_ "github.com/go-sql-driver/mysql"
)

func DBSet() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/threadsapplication?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Susscessfully connected to MYSQL !!")
	return db
}

var Client *sql.DB = DBSet()

func GetImageProduct(postId int) (postImages []models.PostImage, err error) {
	result1, err := Client.Query("SELECT postimage.id, postimage.imageURL, postimage.description FROM postimage WHERE postId = ?", postId)
	if err != nil {
		return nil, err
	}
	for result1.Next() {
		var imageTemp models.PostImage
		err := result1.Scan(&imageTemp.Id, &imageTemp.ImageURL, &imageTemp.Description)
		if err != nil {
			return nil, err
		}
		postImages = append(postImages, imageTemp)
	}
	return postImages, nil
}
func GetInfoUser(userId int) (user models.User, err error) {
	err = Client.QueryRow("SELECT user.id, user.email, user.username, user.state, user.avataImage FROM user").Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.State,
		&user.AvataImage,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
