package api

import (
	"github.com/godot-services/gdar/internal/api/asset"
	"github.com/godot-services/gdar/internal/api/auth"
	"github.com/labstack/echo/v4"
)

func Routes(apiGroup *echo.Group) {
	apiGroup.GET("/configure", auth.ConfigureHandler)
	apiGroup.POST("/register", auth.RegisterHandler)
	apiGroup.POST("/login", auth.LoginHandler)
	apiGroup.POST("/logout", auth.LogoutHandler)
	apiGroup.POST("/change_password", auth.ChangePasswordHandler)

	// TODO authentication middleware
	assetGroup := apiGroup.Group("/api/asset")
	assetGroup.GET("/", asset.ListAssetsHandler)
	assetGroup.GET("/:id", asset.GetAssetHandler)
	assetGroup.POST("/:id/delete", asset.DeleteAssetHandler)
	assetGroup.POST("/:id/undelete", asset.UndeleteAssetHandler)
	assetGroup.POST("/:id/support_level", asset.UpdateSupportLevelHandler)
	assetGroup.POST("/", asset.CreateAssetHandler)
	assetGroup.POST("/:id", asset.UpdateAssetHandler)

	assetGroup.GET("/edit/:id", asset.GetEditHandler)
	assetGroup.POST("/edit/:id", asset.UpdateEditHandler)
	assetGroup.POST("/edit/:id/review", asset.ReviewEditHandler)
	assetGroup.POST("/edit/:id/accept", asset.AcceptEditHandler)
	assetGroup.POST("/edit/:id/reject", asset.RejectEditHandler)
}
