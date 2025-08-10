package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type PaymentRouter struct {
}

func (r PaymentRouter) InitPaymentRouter(Router *gin.RouterGroup) {
	paymentRouterPublic := Router.Group("/payment")
	{
		paymentRouterPublic.GET("/vnpay-return", controllers.Payment.VNPayReturn)
		paymentRouterPublic.GET("/vnpay-ipn", controllers.Payment.VNPayIPN)
	}

	paymentRouterPrivate := Router.Group("/payment")
	paymentRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		paymentRouterPrivate.GET("/")
		paymentRouterPrivate.GET("/create-payment-url")
		paymentRouterPrivate.POST("/create-payment-url", controllers.Payment.CreatePaymentURL)
		paymentRouterPrivate.GET("/querydr")
		paymentRouterPrivate.POST("/querydr")
		paymentRouterPrivate.GET("/refund")
		paymentRouterPrivate.POST("/refund")
	}

}
