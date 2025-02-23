package models

import (
	"database/sql"
	"time"

	"github.com/PhanLuc1/Blogify-Project-Backend/src/database"
	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	Id              int           `json:"id"`
	User            User          `json:"user"`
	PostId          int           `json:"postId"`
	ParentCommentID sql.NullInt64 `json:"parentCommentId"`
	Content         string        `json:"content"`
	CreateAt        time.Time     `json:"creatAt"`
	Replies         []Comment     `json:"parentComment"`
	Reaction        Reaction      `json:"reaction"`
	IsCurrentUser   bool          `json:"isCommentUser"`
}

func GetCommentsForPost(postID int, currentUserID int) ([]Comment, error) {
	var comments []Comment

	err := getCommentsRecursive(database.Client, &comments, postID, nil, currentUserID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
func getCommentsRecursive(db *sql.DB, comments *[]Comment, postID int, parentCommentID *int, currentUserID int) error {
	query := "SELECT id, userId, postId, parentCommentId, content, createAt FROM comment WHERE postId = ? AND parentCommentId "
	if parentCommentID == nil {
		query += "IS NULL"
	} else {
		query += "= ?"
	}
	query += " ORDER BY id DESC"
	var rows *sql.Rows
	var err error
	if parentCommentID != nil {
		rows, err = db.Query(query, postID, parentCommentID)
		if err != nil {
			return err
		}
	} else {
		rows, err = db.Query(query, postID)
		if err != nil {
			return err
		}
	}

	defer rows.Close()

	for rows.Next() {
		var comment Comment
		var UserId int
		err := rows.Scan(&comment.Id, &UserId, &comment.PostId, &comment.ParentCommentID, &comment.Content, &comment.CreateAt)
		if err != nil {
			return err
		}
		comment.User, err = GetInfoUser(UserId)
		if err != nil {
			return err
		}
		comment.Reaction, err = GetReactionComment(comment.Id)
		if err != nil {
			return err
		}
		err = getCommentsRecursive(db, &comment.Replies, postID, &comment.Id, UserId)
		if err != nil {
			return err
		}
		comment.IsCurrentUser = (UserId == currentUserID)
		*comments = append(*comments, comment)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
func GetAmountCommentPost(postId int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM comment WHERE postId = ? "
	err := database.Client.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, err
}
