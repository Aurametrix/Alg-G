func registerRoutes(env *common.Env) {

	r := isokit.NewRouter()
	r.Handle("/index", handlers.IndexHandler(env))
	r.Handle("/products", handlers.ProductsHandler(env))
	r.Handle("/product-detail/{productTitle}", handlers.ProductDetailHandler(env))
	r.Handle("/about", handlers.AboutHandler(env))
	r.Handle("/contact", handlers.ContactHandler(env))
	r.Handle("/shopping-cart", handlers.ShoppingCartHandler(env))
	r.Listen()
	env.Router = r
}
