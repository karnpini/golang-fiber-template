package main

func main() {
	// cfg := config.LoadConfig()
	// app := fiber.New()

	// grpcClient, err := clients.NewIKMCoreClient(cfg.IKMCoreBackend.Endpoint)
	// if err != nil {
	// 	panic("Failed to connect to gRPC client")
	// }
	// defer grpcClient.Close()

	// Inject gRPC client into service layer
	// ikmContentService := service.NewIKMContentService(grpcClient)
	// ikmContentHandler := handlers.NewIKMContentHandler(ikmContentService)

	// prefix := app.Group("template")
	// routes.RegisterRoutes(prefix, ikmContentHandler)

	// Start server
	// err = app.Listen(":" + cfg.Port)
	// if err != nil {
	// 	panic("Server failed to start")
	// }
}
