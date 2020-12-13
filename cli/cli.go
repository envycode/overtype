package cli

import (
	"context"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"overtype/appcontext"
	"overtype/router"
	"syscall"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "overtype",
	Short: "Overtype main application",
	Run: func(cmd *cobra.Command, args []string) {
		handler := cors.
			AllowAll().
			Handler(router.SetupRoute(appCtx))

		srv := &http.Server{
			Handler:      handler,
			Addr:         appCtx.Config.App.AppAddress(),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Printf("Starting %v app on %v\n", appCtx.Config.App.Name, appCtx.Config.App.AppAddress())
		go listenServer(srv)
		waitForShutdown(srv)
	},
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.WithField("error", err.Error()).Fatal("Server closed because of an error")
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	log.Warn("Api server shutting down")
	if err := apiServer.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	log.Warn("Api server shutting complete")
}

var appCtx = appcontext.NewApplication()

func Execute() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	if appCtx.Config.App.Env == "production" {
		log.SetLevel(log.InfoLevel)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
