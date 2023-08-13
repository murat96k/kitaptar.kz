package service

import (
	"context"
	"one-lab/api"
	"one-lab/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, email, password string) (string, error)
	VerifyToken(token string) (string, error)
	UpdateUser(ctx context.Context, email string, req *api.UpdateUserRequest) error //test
	GetUser(ctx context.Context, email string) (*entity.User, error)                //test
	DeleteUser(ctx context.Context, email string) error                             //test

	CreateBook(ctx context.Context, req *api.BookRequest) error            //test
	GetUserBooks(email string) ([]entity.Book, error)                      //test
	GetAllBooks(ctx context.Context) ([]entity.Book, error)                //test
	GetBookById(ctx context.Context, id string) (*entity.Book, error)      //test
	DeleteBook(ctx context.Context, id string) error                       //test
	UpdateBook(ctx context.Context, id string, req *api.BookRequest) error //test

	CreateAuthor(ctx context.Context, req *api.AuthorRequest) error            //test
	GetAllAuthors(ctx context.Context) ([]entity.Author, error)                //test
	GetAuthorById(ctx context.Context, id string) (*entity.Author, error)      //test
	DeleteAuthor(ctx context.Context, id string) error                         //test
	UpdateAuthor(ctx context.Context, id string, req *api.AuthorRequest) error //test

	//
	//UpdateArticle(ctx context.Context, a *entity.Article) error
	//DeleteArticle(ctx context.Context, id int64) error
	//GetArticleByID(ctx context.Context, id int64) (*entity.Article, error)
	//GetAllArticles(ctx context.Context) ([]entity.Article, error)
	//GetArticlesByUserID(ctx context.Context, userID int64) ([]entity.Article, error)
	//
	//GetCategories(ctx context.Context) ([]entity.Category, error)
}
