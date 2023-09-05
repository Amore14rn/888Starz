package app

import (
	"context"
	"fmt"
	"github.com/Amore14rn/888Starz/internal/config"
	v1 "github.com/Amore14rn/888Starz/internal/controllers/http/v1"
	"github.com/Amore14rn/888Starz/internal/domain/user/dao"
	"github.com/Amore14rn/888Starz/internal/domain/user/service"
	"github.com/Amore14rn/888Starz/pkg/common/core/clock"
	"github.com/Amore14rn/888Starz/pkg/common/core/closer"
	"github.com/Amore14rn/888Starz/pkg/common/core/identity"
	"github.com/Amore14rn/888Starz/pkg/common/logging"
	"github.com/Amore14rn/888Starz/pkg/errors"
	"github.com/Amore14rn/888Starz/pkg/graceful"
	"github.com/Amore14rn/888Starz/pkg/metric"
	psql "github.com/Amore14rn/888Starz/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg         *config.Config
	pgClient    *pgxpool.Pool
	router      *gin.Engine
	httpServer  *http.Server
	UserService v1.UserService
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {
	logging.L(ctx).Info("router initializing")

	router := gin.Default()

	logging.WithFields(ctx,
		logging.StringField("username", cfg.Postgres.User),
		logging.StringField("password", "<REMOVED>"),
		logging.StringField("host", cfg.Postgres.Host),
		logging.StringField("port", cfg.Postgres.Port),
		logging.StringField("database", cfg.Postgres.Database),
	).Info("PostgreSQL initializing")

	pgDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	pgClient, err := psql.NewClient(ctx, 5, 3*time.Second, pgDsn, false)
	if err != nil {
		return App{}, errors.Wrap(err, "psql.NewClient")
	}

	closer.AddN(pgClient)

	logging.L(ctx).Info("handlers initializing")

	logging.L(ctx).Info("heartbeat metric initializing")

	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	cl := clock.New()
	generator := identity.NewGenerator()

	userStorage := dao.NewUserStorage(pgClient)
	productService := service.NewUserService(userStorage)
	productPolicy := policy_product.NewProductPolicy(productService, generator, cl)
	productServiceServer := grpc_v1_product.NewServer(
		productPolicy,
		pb_prod_products.UnimplementedProductServiceServer{},
	)

	return App{
		cfg:    cfg,
		router: router,
	}, nil

}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP(ctx)
	})
	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {
	logger := logging.WithFields(ctx,
		logging.StringField("IP", a.cfg.Server.HOST),
		logging.StringField("Port", a.cfg.Server.PORT),
		logging.DurationField("WriteTimeout", a.cfg.Server.WriteTimeout),
		logging.DurationField("ReadTimeout", a.cfg.Server.ReadTimeout),
		logging.IntField("MaxHeaderBytes", a.cfg.Server.MaxHeaderBytes),
	)
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Server.HOST, a.cfg.Server.PORT))
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to create listener")
	}

	handler := a.router

	a.httpServer = &http.Server{
		Handler:        handler,
		WriteTimeout:   a.cfg.Server.WriteTimeout,
		ReadTimeout:    a.cfg.Server.ReadTimeout,
		MaxHeaderBytes: a.cfg.Server.MaxHeaderBytes,
	}
	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.With(logging.ErrorField(err)).Fatal("failed to start server")
		}
	}

	httpErrChan := make(chan error, 1)
	httpShutdownChan := make(chan struct{})

	graceful.PerformGracefulShutdown(a.httpServer, httpErrChan, httpShutdownChan)

	return err
}
