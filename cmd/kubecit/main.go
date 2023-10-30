package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	nacosconfig "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"strings"

	"kubecit/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	if !strings.Contains(flagconf, "local") {

		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}

		var bct conf.RemoteStrap
		if err := c.Scan(&bct); err != nil {
			panic(err)
		}

		cc := constant.ClientConfig{
			NamespaceId:         bct.Nacos.Namespace,
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			LogLevel:            "debug",
		}

		client, err := clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig: &cc,
				ServerConfigs: []constant.ServerConfig{
					*constant.NewServerConfig(bct.Nacos.Ip, bct.Nacos.Port),
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		ccc := config.New(config.WithSource(nacosconfig.NewConfigSource(client, nacosconfig.WithGroup(bct.Nacos.Group), nacosconfig.WithDataID(bct.Nacos.DataId))))
		defer ccc.Close()

		if err := ccc.Load(); err != nil {
			panic(err)
		}
		var bc conf.Bootstrap
		if err := ccc.Scan(&bc); err != nil {
			panic(err)
		}

		spew.Dump(bc)
		app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err := app.Run(); err != nil {
			panic(err)
		}

	} else {
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}

		var bc conf.Bootstrap
		if err := c.Scan(&bc); err != nil {
			panic(err)
		}
		app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err := app.Run(); err != nil {
			panic(err)
		}
	}

}
