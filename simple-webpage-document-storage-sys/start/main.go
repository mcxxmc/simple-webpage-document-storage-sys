package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"simple-webpage-document-storage-sys/common"
	"simple-webpage-document-storage-sys/controller"
	"simple-webpage-document-storage-sys/logging"
	"simple-webpage-document-storage-sys/manager"
	"syscall"
	"time"
)

func addMiddleware(router *gin.Engine) {
	router.Use(controller.SetHeader())
	router.Use(controller.DecodeToken())  //todo: disable this for login
}

func bindUrl(router *gin.Engine) {
	// view the directories and files
	router.GET("/filesystem/view", controller.ViewHierarchy)

	// login
	router.POST("filesystem/login", controller.Login)
	// logout
	router.GET("filesystem/logout", controller.Logout)

	// new user
	router.POST("filesystem/register", controller.NewUser)

	// url for fetching file content
	router.POST("/filesystem/read", controller.GetFile)
	// url for modifying file content
	router.POST("/filesystem/rewrite", controller.ModifyFile)
	// url for renaming file && dir
	router.POST("/filesystem/rename", controller.Rename)
	// url for creating a new file || a new dir
	router.POST("/filesystem/create", controller.Create)
	// url for deleting a file || a dir
	router.POST("/filesystem/remove", controller.Delete)
	// url for moving a file || dir
	router.POST("/filesystem/move", controller.Move)
}

func createServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr: addr,
		Handler: handler,
	}
}

func run(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.Fatal(err, logging.S("detail: ", "failed to start the server."))
	}
}

func gracefulShutDown(server *http.Server, delay time.Duration) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<- quit  // block here

	logging.Info("Server is closing...")

	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logging.Fatal(err, logging.S("detail: ", "Server shutdown."))
	}

	logging.Info("Server exits.")
}


func main() {
	// config the logger
	defer logging.Sync()

	// start the manager which maintains the user info
	manager.StartManager(common.Path_index_of_users)
	defer manager.SaveWhenShuttingDown()

	// config the router
	router := gin.Default()
	addMiddleware(router)
	bindUrl(router)
	server := createServer(common.Port, router)
	go run(server)
	gracefulShutDown(server, 5 * time.Second)
}

//TODO: 现在的数据库是单向连接的，只能从user到具体文档，而不能从具体文档到user;可以考虑升级成双向
//TODO: index of users can be replaced by MySQL instead of JSON for better performance when the scale is large
