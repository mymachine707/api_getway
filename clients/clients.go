package clients

import (
	"mymachine707/config"
	"mymachine707/protogen/blogpost"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author  blogpost.AuthorServiceClient
	Article blogpost.ArticleServiceClient
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	connAuthor, err := grpc.Dial(cfg.AuthorServiceGrpcHost+cfg.AuthorServiceGrpcPort, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	// defer conn.Close() // ulanishdan keyin yoppormasligi uchun hozircha defer keremas
	author := blogpost.NewAuthorServiceClient(connAuthor)

	connArticle, err := grpc.Dial(cfg.ArticleServiceGrpcHost+cfg.ArticleServiceGrpcPort, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	// defer conn.Close() // ulanishdan keyin yoppormasligi uchun hozircha defer keremas
	article := blogpost.NewArticleServiceClient(connArticle)

	return &GrpcClients{
		Author:  author,
		Article: article,
	}, nil
}


