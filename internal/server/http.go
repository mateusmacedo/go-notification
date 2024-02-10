package server

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/mateusmacedo/go-notification/api/v1"
	"github.com/mateusmacedo/go-notification/docs"
	"github.com/mateusmacedo/go-notification/internal/middleware"
	"github.com/mateusmacedo/go-notification/pkg/jwt"
	"github.com/mateusmacedo/go-notification/pkg/log"
	"github.com/mateusmacedo/go-notification/pkg/server/http"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("health")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Its on!",
		})
	})

	return s
}
