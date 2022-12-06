package clients

import (
	"mymachine707/config"
	"mymachine707/protogen/blogpost"

	"google.golang.org/grpc"
)

type GrpcClients struct {
	Author  blogpost.AuthorServiceClient
	Article blogpost.ArticleServiceClient
	Users   blogpost.UserServiceClient

	conns []*grpc.ClientConn
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

	connUser, err := grpc.Dial(cfg.UsersServiceGrpcHost+cfg.UsersServiceGrpcPort, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	// defer conn.Close() // ulanishdan keyin yoppormasligi uchun hozircha defer keremas
	user := blogpost.NewUserServiceClient(connUser)

	conns := make([]*grpc.ClientConn, 0)
	return &GrpcClients{
		Author:  author,
		Article: article,
		Users:   user,
		conns:   append(conns, connAuthor, connArticle, connUser),
	}, nil
}

// Close for disconnection connection another serevice
func (c *GrpcClients) Close() {
	for _, v := range c.conns {
		v.Close()
	}
}
