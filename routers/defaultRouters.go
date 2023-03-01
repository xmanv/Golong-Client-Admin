package routers

import (
	"ginshop57/controllers/itying"
	"ginshop57/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/category:id", itying.ProductController{}.Category)
		defaultRouters.GET("/detail", itying.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", itying.ProductController{}.GetImgList)

		defaultRouters.GET("/cart", itying.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", itying.CartController{}.AddCart)

		defaultRouters.GET("/cart/successTip", itying.CartController{}.AddCartSuccess)
		defaultRouters.GET("/cart/decCart", itying.CartController{}.DecCart)
		defaultRouters.GET("/cart/incCart", itying.CartController{}.IncCart)

		defaultRouters.GET("/cart/changeOneCart", itying.CartController{}.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", itying.CartController{}.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", itying.CartController{}.DelCart)

		defaultRouters.GET("/pass/login", itying.PassController{}.Login)
		defaultRouters.GET("/pass/captcha", itying.PassController{}.Captcha)

		defaultRouters.GET("/pass/registerStep1", itying.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", itying.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", itying.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", itying.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", itying.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", itying.PassController{}.DoRegister)
		defaultRouters.POST("/pass/doLogin", itying.PassController{}.DoLogin)
		defaultRouters.GET("/pass/loginOut", itying.PassController{}.LoginOut)
		//判断用户权限
		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, itying.BuyController{}.Checkout) //判断用户权限
		defaultRouters.POST("/buy/doCheckout", middlewares.InitUserAuthMiddleware, itying.BuyController{}.DoCheckout)
		defaultRouters.GET("/buy/pay", middlewares.InitUserAuthMiddleware, itying.BuyController{}.Pay)
		defaultRouters.GET("/buy/orderPayStatus", middlewares.InitUserAuthMiddleware, itying.BuyController{}.OrderPayStatus)

		defaultRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, itying.AddressController{}.AddAddress)
		defaultRouters.POST("/address/editAddress", middlewares.InitUserAuthMiddleware, itying.AddressController{}.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", middlewares.InitUserAuthMiddleware, itying.AddressController{}.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", middlewares.InitUserAuthMiddleware, itying.AddressController{}.GetOneAddressList)

		defaultRouters.GET("/alipay", middlewares.InitUserAuthMiddleware, itying.AlipayController{}.Alipay)
		defaultRouters.POST("/alipayNotify", itying.AlipayController{}.AlipayNotify)
		defaultRouters.GET("/alipayReturn", middlewares.InitUserAuthMiddleware, itying.AlipayController{}.AlipayReturn)

		defaultRouters.GET("/wxpay", middlewares.InitUserAuthMiddleware, itying.WxpayController{}.Wxpay)
		defaultRouters.POST("/wxpay/notify", itying.WxpayController{}.WxpayNotify)

		defaultRouters.GET("/user", middlewares.InitUserAuthMiddleware, itying.UserController{}.Index)
		defaultRouters.GET("/user/order", middlewares.InitUserAuthMiddleware, itying.UserController{}.OrderList)
		defaultRouters.GET("/user/orderinfo", middlewares.InitUserAuthMiddleware, itying.UserController{}.OrderInfo)

		defaultRouters.GET("/search", itying.SearchController{}.Index)
		defaultRouters.GET("/search/getOne", itying.SearchController{}.GetOne)
		defaultRouters.GET("/search/addGoods", itying.SearchController{}.AddGoods)
		defaultRouters.GET("/search/updateGoods", itying.SearchController{}.UpdateGoods)
		defaultRouters.GET("/search/deleteGoods", itying.SearchController{}.DeleteGoods)
		defaultRouters.GET("/search/query", itying.SearchController{}.Query)
		defaultRouters.GET("/search/filterQuery", itying.SearchController{}.FilterQuery)
		defaultRouters.GET("/search/pagingQuery", itying.SearchController{}.PagingQuery)
	}

}
