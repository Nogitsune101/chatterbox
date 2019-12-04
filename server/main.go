package server

import (
	"chatterbox/server/handlers"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8080", "http service address")

// StartServer executes the chatterbox server
func StartServer(addr string) {

	// New Router
	router := mux.NewRouter()

	// Initialize Server Settings
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// API Server modules to load
	modules := []func(r *mux.Router){
		handlers.WebSocketModule,
		handlers.WebClientModule,
	}

	// Register Modules
	for _, module := range modules {
		log.Println(" - Registering Chatterbox module... [" + getModuleName(module) + "]")
		module(router)
	}

	// Create a channel for handling server exit
	signalInterrupt := make(chan os.Signal, 1)

	// Start http server
	go func() {
		log.Printf("[[[ Starting chatterbox on %s ]]]", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for a term signal
	signal.Notify(signalInterrupt, os.Interrupt)

	// Block until we receive our term signal.
	<-signalInterrupt

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutting Down Chatterbox")
	os.Exit(0)
}

// Since regexp is pre compiled, only do this once
var regexpModuleName = regexp.MustCompile("handlers\\.(.*)Module$")

// getModuleName uses Reflection and Regular Expression to grab the module name
func getModuleName(module interface{}) string {
	modulePath := runtime.FuncForPC(reflect.ValueOf(module).Pointer()).Name()
	moduleName := regexpModuleName.FindStringSubmatch(modulePath)[1]

	return moduleName
}
