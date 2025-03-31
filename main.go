package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// docs.SwaggerInfo.Title = "Backend service for eagle API"
	// docs.SwaggerInfo.Description = "This service is meant to manage eagle request"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost" + ":8087"
	// // docs.SwaggerInfo.Host = "jellyfish-app-gz2qc.ondigitalocean.app"
	// docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"https"}

	svc := gin.Default()
	// svc.Use(cors.Default())
	// svc.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{"https://localhost"}, // List of allowed origins https://jellyfish-app-gz2qc.ondigitalocean.app
	// 	AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders: []string{"Content-Length"},
	// }))
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// url := ginSwagger.URL("https://jellyfish-app-gz2qc.ondigitalocean.app/swagger/doc.json")

	// svc.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	svc.GET("/health", testservice)
	// svc.POST("/api/user/CreateNewUser", api.CreateNewUser)                             //done
	// svc.POST("/api/user/SignUp", api.SignUp)                                           //done
	// svc.GET("api/user/GetUserByEmailAddress/:EmailAddress", api.GetUserByEmailAddress) //done
	// svc.GET("api/user/LogIn/:UserName/:Password", api.LogIn)                           //done
	// svc.GET("api/user/GetUserByMobile/:MobileNumber", api.GetUserByMobile)             //done
	// //svc.POST("api/user/UploadImage", api.TestSevc)                               // to be done later
	// svc.POST("api/user/SendEmail", api.SendEmail)           //done
	// svc.POST("api/user/LogOutUser", testSvc)                //to be done later
	// svc.POST("api/user/CompleteSignUp", api.CompleteSignUp) //to be done later

	// svc.POST("api/cart/CreateCart", api.CreateCart)                 //done
	// svc.POST("api/cart/AddUserToCart", api.CreateCartMember)        //done
	// svc.POST("api/cart/AddItemToCart", api.AddItemToCart)           //done
	// svc.POST("api/cart/CreateCartType", testSvc)                    //to be done later
	// svc.POST("api/cart/RemoveItemFromCart", api.RemoveItemFromCart) //done
	// svc.PUT("api/cart/CloseCart", api.CloseCart)                    //done
	// svc.POST("api/cart/RemoveUserFromCart", api.RemoveUserFromCart) //done
	// // svc.GET("api/cart/GetCartByCartId/:CartId", api.GetCartByCartId)     //done
	// svc.GET("api/cart/GetCarts", api.RemoveUserFromCart)                     //done
	// svc.GET("api/cart/GetCartByUserId/:UserId", api.GetCartByUserId)         //done
	// svc.GET("api/cart/GetCartByEmail/:EmailAddress", api.GetCartByUserEmail) //done

	// svc.POST("api/admin/CreateTitle", api.CreateTitle)             //done
	// svc.GET("api/admin/GetTitles", api.GetTitles)                  //done
	// svc.GET("api/admin/GetAllStatus", api.GetAllStatus)            //done
	// svc.GET("/api/admin/GetToken/:clientname", api.GetToken)       //done
	// svc.POST("api/admin/RegisterNewClient", api.RegisterNewClient) //done

	// svc.POST("api/group/CreateGroupType", testSvc)
	// svc.POST("api/group/DeleteGroupType", testSvc)

	// svc.POST("api/admin/CreateRole", testSvc) //test service
	// svc.POST("api/admin/DeleteRole", testSvc)
	// svc.GET("api/admin/GetAllRoles", testSvc)
	// svc.GET("api/admin/GetRoleByRoleId", testSvc)
	// svc.POST("api/admin/CreateProduct", testSvc)
	// svc.GET("api/admin/GetProductById", testSvc)
	// svc.GET("api/admin/GetProducts", testSvc)
	// svc.POST("api/admin/DeleteProduct", testSvc)
	// svc.POST("api/admin/CreateStatus", testSvc)
	// svc.DELETE("api/admin/DeleteStatus", testSvc)
	ctx := gin.Default()
	ctx.GET("/", testservice)
	ctx.Run(":8091")
}

func testservice(c *gin.Context) {
	c.JSON(http.StatusOK, "response")
}
