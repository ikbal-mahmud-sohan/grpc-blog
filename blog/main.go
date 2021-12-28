package main

import (
	"myGo/Blogs/blog/storage/postgres"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/spf13/viper"

	bcc "myGo/Blogs/blog/core/category"
	bcp "myGo/Blogs/blog/core/post"
	"myGo/Blogs/blog/services/category"
	"myGo/Blogs/blog/services/post"
	bgvc "myGo/Blogs/gunk/v1/category"
	bgv "myGo/Blogs/gunk/v1/post"

	"google.golang.org/grpc"
)

func main() {

	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("blog/env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("error loading configuration: %v", err)
	}
	//connect grpc server
	grpcServer := grpc.NewServer()
	store, err := newDBFromConfig(config)
	// store, err := postgres.NewStorage("dbstring")
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}

	cs := bcp.NewCoreSvc(store)
	s := post.NewPostSvc(cs)
	bgv.RegisterPostServiceServer(grpcServer, s)
//category
	cc := bcc.NewCategoryCoreSvc(store)
	c := category.NewCategorySvc(cc)
	bgvc.RegisterCategoryServiceServer(grpcServer, c)
	host, port := config.GetString("server.host"), config.GetString("server.port")

	//setup port,host
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))

	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	log.Printf("Server is starting on: http://%s:%s\n", host, port)

	// reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func newDBFromConfig(config *viper.Viper) (*postgres.Storage, error) {
	cf := func(c string) string { return config.GetString("database." + c) }
	ci := func(c string) string { return strconv.Itoa(config.GetInt("database." + c)) }
	dbParams := " " + "user=" + cf("user")
	dbParams += " " + "host=" + cf("host")
	dbParams += " " + "port=" + cf("port")
	dbParams += " " + "dbname=" + cf("dbname")
	if password := cf("password"); password != "" {
		dbParams += " " + "password=" + password
	}
	dbParams += " " + "sslmode=" + cf("sslMode")
	dbParams += " " + "connect_timeout=" + ci("connectionTimeout")
	dbParams += " " + "statement_timeout=" + ci("statementTimeout")
	dbParams += " " + "idle_in_transaction_session_timeout=" + ci("idleTransacionTimeout")
	db, err := postgres.NewStorage(dbParams)
	if err != nil {
		return nil, err
	}
	return db, nil
}
