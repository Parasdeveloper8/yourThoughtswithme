package main

import (
	"net/http"

	routes "github.com/Parasdeveloper8/myexpgoweb/Routes"
	"github.com/Parasdeveloper8/myexpgoweb/auth"
	"github.com/Parasdeveloper8/myexpgoweb/cors"
	"github.com/Parasdeveloper8/myexpgoweb/limiter"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")

	//Middlewares start
	router.Use(limiter.RateLimit())

	router.Use(auth.SessionMiddleware())
	//Middlewares end
	// Set up CORS
	cors.SetupCORS(router)

	router.GET("/", routes.HandleHome)
	router.GET("/favpro", auth.CheckEmail(), serveFavProForm)
	router.GET("/logpage", serveLogPage)
	router.GET("/regpage", serveRegPage)
	router.GET("/admin", serveAdminPage)
	router.GET("/surveypage", auth.CheckEmail(), serveSurveyPageRoute)
	router.POST("/register", auth.HandleRegister)
	router.POST("/login", auth.HandleLogin)
	router.GET("/afterlog", auth.CheckEmail(), serveAfterLog)
	router.POST("/logout", auth.HandleLogout)
	router.GET("/feedback", serveGiveFeedBack)
	router.POST("/submitfav", routes.HandleSurveySubmission)
	router.POST("/subfeed", auth.CheckEmail(), routes.HandleFeedSubmission)
	router.Run(":4700")
}

func serveLogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "loginpage.html", nil)
}

func serveRegPage(c *gin.Context) {
	c.HTML(http.StatusOK, "registerpage.html", nil)
}

func serveSurveyPageRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "surveypages.html", nil)
}

func serveAfterLog(c *gin.Context) {
	c.HTML(http.StatusOK, "afterlog.html", nil)
}
func serveFavProForm(c *gin.Context) {
	c.HTML(http.StatusOK, "favproform.html", nil)
}
func serveAdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}
func serveGiveFeedBack(c *gin.Context) {
	c.HTML(http.StatusOK, "givefeedback.html", nil)
}
