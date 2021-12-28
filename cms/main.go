package main

import (
	"fmt"
	"log"
	"myGo/Blogs/cms/handler"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	bgvc "myGo/Blogs/gunk/v1/category"
	bgv "myGo/Blogs/gunk/v1/post"
)

func main(){
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("cms/env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("error loading configuration: %v", err)
	}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	 store := sessions.NewCookieStore([]byte(config.GetString("session.secret")))
	
	conn, err:=grpc.Dial(
		fmt.Sprintf("%s:%s",config.GetString("blog.host"),config.GetString("blog.port")),
		grpc.WithInsecure(),
	)
	if err!=nil{
		log.Fatal(err)
	}
	cc:=bgvc.NewCategoryServiceClient(conn)
	tc:=bgv.NewPostServiceClient(conn)
	r:=handler.New( decoder, store,tc,cc)
	host,port:=config.GetString("server.host"),config.GetString("server.port")
	
	log.Printf("Server is starting on: http://%s:%s\n",host,port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s",host,port),r); err !=nil{
		log.Fatal(err)
	}
}