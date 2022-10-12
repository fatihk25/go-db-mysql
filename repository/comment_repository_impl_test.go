package repository

import (
	"context"
	"fmt"
	gomysql "go_mysql"
	"go_mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnections())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "johan@mail,com",
		Comment: "Ini comment mas yohan2",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnections())

	comment, err := commentRepository.FindById(context.Background(), 44)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnections())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
