package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Ananto30/go-grpc/config"
	"github.com/Ananto30/go-grpc/conn"
	"github.com/Ananto30/go-grpc/registry"
	"github.com/Ananto30/go-grpc/server"
	raven "github.com/getsentry/raven-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var debugPort int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Long:  `Start the gRPC API server of fortress service`,
	Run:   serve,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.Init()

		if err := conn.ConnectDB(); err != nil {
			return fmt.Errorf("Cant't connect database: %v", err)
		}

		// tcfg := config.Ties()
		// if _, err := service.Connect("ties", tcfg.Addr(), tcfg.Secure); err != nil {
		// 	return fmt.Errorf("Can't connect ties: %v", err)
		// }

		// if err := conn.ConnectEmitter(); err != nil {
		// 	return fmt.Errorf("Can't connect generic queue: %v", err)
		// }

		appCfg := config.App()
		if dsn := appCfg.Sentry; dsn != "" {
			if err := raven.SetDSN(dsn); err != nil {
				return err
			}
			raven.SetTagsContext(map[string]string{"service": "fortress"})
			raven.SetEnvironment(appCfg.Env)
			raven.SetRelease(appCfg.Version)
		}

		return nil
	},
}

func init() {
	serveCmd.PersistentFlags().IntP("port", "p", 8080, "port on which the server will listen")
	serveCmd.PersistentFlags().IntVarP(&debugPort, "debug-port", "d", 8081, "port on which the debug server will listen")
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	appCfg := config.App()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(appCfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srvr := &server.Server{}

	container := registry.BuildContainer()
	errC := container.Invoke(func(server *server.Server) {
		srvr = server
	})
	if errC != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		log.Println("Started " + appCfg.Env + " server")
		log.Println("Listening on " + strconv.Itoa(appCfg.Port))
		if err := srvr.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-stop

	log.Println("Shutting down server...")

	srvr.GracefulStop()

	log.Println("Server shut down gracefully")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// // apiS.Shutdown(ctx)
}
