package router

import "backend/controller"

func setType(rtr *router) {
	typRtr := rtr.engine.Group("api/type")
	{
		ctrl := controller.NewTypeController()
		typRtr.GET("/list/:limit/:productLimit", ctrl.List)
	}
}

func setEvent(rtr *router) {
	evtRtr := rtr.engine.Group("api/event")
	{
		ctrl := controller.NewEventController()
		evtRtr.GET("/list/:limit", ctrl.List)
	}
}

func setProduct(rtr *router) {
	pdtRtr := rtr.engine.Group("api/product")
	{
		ctrl := controller.NewProductController()
		pdtRtr.GET("/list/:typeId/:eventIds/:limit/:page", ctrl.List)
		pdtRtr.GET("/list/:typeId/:eventIds/:limit/:page/:keyword", ctrl.List)
		pdtRtr.GET("/detail/:id", ctrl.FindID)
	}
}

func setNews(rtr *router) {
	newRtr := rtr.engine.Group("api/news")
	{
		ctrl := controller.NewNewsController()
		newRtr.GET("/list/:limit/:page", ctrl.List)
		newRtr.GET("/detail/:id", ctrl.FindID)
	}
}
