package middlewares

import (
	"context"
	"github.com/labstack/echo"
	"github.com/makinap/osu/service"
	"log"
	"net/http"
)

type authString string

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	log.Println("Using auth middleware ")
	return func(c echo.Context) error {

		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(c.Response().Writer, c.Request())
			//return echo.NewHTTPError(http.StatusUnauthorized, "Please provide credentials")
			return next(c)
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := service.JwtValidate(context.Background(), auth)

		if err != nil || !validate.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
			//return next(c)
		}

		customClaim, _ := validate.Claims.(*service.JwtCustomClaim)

		ctx := context.WithValue(c.Request().Context(), authString("auth"), customClaim)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func CtxValue(ctx context.Context) *service.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*service.JwtCustomClaim)
	return raw
}