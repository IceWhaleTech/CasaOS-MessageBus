//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.2 -package codegen api/message_bus/openapi.yaml > codegen/message_bus_api.go"

package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/external"
	"github.com/IceWhaleTech/CasaOS-Common/model"
	"github.com/IceWhaleTech/CasaOS-Common/utils/file"
	util_http "github.com/IceWhaleTech/CasaOS-Common/utils/http"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
	"github.com/IceWhaleTech/CasaOS-MessageBus/common"
	"github.com/IceWhaleTech/CasaOS-MessageBus/config"
	"github.com/IceWhaleTech/CasaOS-MessageBus/repository"
	"github.com/IceWhaleTech/CasaOS-MessageBus/route"
	"github.com/IceWhaleTech/CasaOS-MessageBus/service"
	"github.com/coreos/go-systemd/daemon"
	"go.uber.org/zap"
)

const localhost = "127.0.0.1"

var (
	//go:embed api/index.html
	_docHTML string

	//go:embed api/message_bus/openapi.yaml
	_docYAML string
)

func main() {
	// arguments
	configFlag := flag.String("c", "", "config file path")
	versionFlag := flag.Bool("v", false, "version")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("v%s\n", common.MessageBusVersion)
		os.Exit(0)
	}

	// initialization
	config.InitSetup(*configFlag)

	logger.LogInit(config.AppInfo.LogPath, config.AppInfo.LogSaveName, config.AppInfo.LogFileExt)

	// repository
	if err := file.IsNotExistMkDir(config.CommonInfo.RuntimePath); err != nil {
		panic(err)
	}

	databaseFilePath := filepath.Join(config.CommonInfo.RuntimePath, "message-bus.db")

	repository, err := repository.NewDatabaseRepository(databaseFilePath)
	if err != nil {
		panic(err)
	}
	defer repository.Close()

	// service
	services := service.NewServices(&repository)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	services.Start(&ctx)

	// route
	swagger, err := codegen.GetSwagger()
	if err != nil {
		panic(err)
	}

	apiRouter, err := route.NewAPIRouter(swagger, &services)
	if err != nil {
		panic(err)
	}

	docRouter, err := route.NewDocRouter(swagger, _docHTML, _docYAML)
	if err != nil {
		panic(err)
	}

	mux := &util_http.HandlerMultiplexer{
		HandlerMap: map[string]http.Handler{
			"v2":  apiRouter,
			"doc": docRouter,
		},
	}

	// http listener
	listener, err := net.Listen("tcp", net.JoinHostPort(localhost, "0"))
	if err != nil {
		panic(err)
	}

	// register at gateway
	u, err := url.Parse(swagger.Servers[0].URL)
	if err != nil {
		panic(err)
	}

	apiPath := strings.TrimRight(u.Path, "/")
	apiPaths := []string{apiPath, "/doc" + apiPath}

	gatewayManagement, err := external.NewManagementService(config.CommonInfo.RuntimePath)
	if err != nil {
		panic(err)
	}

	for _, apiPath := range apiPaths {
		err = gatewayManagement.CreateRoute(&model.Route{
			Path:   apiPath,
			Target: "http://" + listener.Addr().String(),
		})

		if err != nil {
			panic(err)
		}
	}

	// write address file
	addressFilePath, err := writeAddressFile(config.CommonInfo.RuntimePath, external.MessageBusAddressFilename, "http://"+listener.Addr().String())
	if err != nil {
		panic(err)
	}

	// notify systemd
	if supported, err := daemon.SdNotify(false, daemon.SdNotifyReady); err != nil {
		logger.Error("Failed to notify systemd that message bus service is ready", zap.Error(err))
	} else if supported {
		logger.Info("Notified systemd that message bus service is ready")
	} else {
		logger.Info("This process is not running as a systemd service.")
	}

	// start http server
	logger.Info("MessageBus service is listening...", zap.Any("address", listener.Addr().String()), zap.String("filepath", addressFilePath))

	server := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	err = server.Serve(listener)
	logger.Info("MessageBus service is stopped", zap.Error(err))
}

func writeAddressFile(runtimePath string, filename string, address string) (string, error) {
	err := os.MkdirAll(runtimePath, 0o755)
	if err != nil {
		return "", err
	}

	filepath := filepath.Join(runtimePath, filename)
	return filepath, os.WriteFile(filepath, []byte(address), 0o600)
}
