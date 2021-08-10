package config

import (
	"fmt"

	"github.com/fatih/color"

	db "github.com/azizk17/go-app/db/sqlc"
	"github.com/azizk17/go-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/pkg/errors"
)

type Server struct {
	*fiber.App
	TemplateEngine *html.Engine
	Store          db.Store
	Name           string `mapstructure:"name" yaml:"name" env:"APP_NAME" env-default:"iSend.to"`
	Version        string `mapstructure:"APP_VERSION" yaml:"version" env:"APP_VERSION" env-default:"dev"`
	Mode           string `mapstructure:"APP_MODE" yaml:"mode" env:"APP_MODE" env-default:"app"`
	Env            string `mapstructure:"APP_ENV" yaml:"env" env:"APP_ENV" env-default:"dev"`
	Key            string `mapstructure:"key" yaml:"key" env:"APP_KEY" env-default:"1894cde6c936a294a478cff0a9227fd276d86df6573b51af5dc59c9064edf426"`
	Url            string `mapstructure:"url" yaml:"url" env:"APP_URL" env-default:"http://localhost"`
	Host           string `mapstructure:"host" yaml:"host" env:"APP_HOST" env-default:"localhost"`
	Port           string `mapstructure:"port" yaml:"port" env:"APP_PORT" env-default:"8080"`
	Path           string `mapstructure:"APP_PATH" yaml:"path" env:"APP_PATH"`
	ProxyHeader    string `mapstructure:"PROXY_HEADER" yaml:"PROXY_HEADER" env:"PROXY_HEADER" env-default:"*"`
	AssetPath      string `mapstructure:"ASSET_PATH" yaml:"asset_path" env:"ASSET_PATH" env-default:"assets"`
	PublicPath     string `mapstructure:"PUBLIC_PATH" yaml:"public_path" env:"PUBLIC_PATH" env-default:"public"`
	UploadPath     string `mapstructure:"UPLOAD_PATH" yaml:"upload_path" env:"UPLOAD_PATH" env-default:"uploads"`
	StoragePath    string `mapstructure:"STORAGE_PATH" yaml:"storage_path" env:"STORAGE_PATH" env-default:"storage"`
	LogPath        string `mapstructure:"LOG_PATH" yaml:"log_path" env:"LOG_PATH" env-default:"storage/logs"`
	ExecPath       bool   `mapstructure:"EXEC_PATH" yaml:"exec_path" env:"EXEC_PATH" env-default:"false"`
	Debug          bool   `mapstructure:"APP_DEBUG" yaml:"debug" env:"APP_DEBUG" env-default:"true"`
	UploadSize     int    `mapstructure:"UPLOAD_SIZE" yaml:"upload_size" env:"UPLOAD_SIZE" env-default:"400"`
}

func (s *Server) Setup() {

	s.App = fiber.New(fiber.Config{
		// Views:                 s.TemplateEngine,
		Concurrency:  256 * 1024 * 1024,
		ServerHeader: s.Name,
		// BodyLimit:             s.UploadSize,
		ReduceMemoryUsage:     true,
		ErrorHandler:          CustomErrorHandler,
		DisableStartupMessage: false,
		// ProxyHeader:           s.ProxyHeader,
	})
}

//
func (s *Server) Serve(addr ...string) error {
	a := s.Host + ":" + s.Port
	if len(addr) != 0 {
		a = addr[0]
	}

	// Load routes and ...
	load(s)
	error404(s)
	// s.App.Use(recover.New())
	s.startupMessage(a)
	return s.App.Listen(a)

}

//
// func (s *Server) ServeWithGraceFullShutdown(addr ...string) error {
// 	a := s.Host + ":" + s.Port
// 	if len(addr) != 0 {
// 		a = addr[0]
// 	}
// 	s.startupMessage(a)

// 	// Listen from a different goroutine
// 	go func() {
// 		if err := s.Listen(a); err != nil {
// 			log.Print(err)
// 			os.Exit(0)
// 		}
// 	}()

// 	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
// 	signal.Notify(c,
// 		syscall.SIGINT,
// 		syscall.SIGTERM,
// 		syscall.SIGABRT,
// 		syscall.SIGQUIT,
// 	) // When an interrupt is sent, notify the channel
// 	<-c // This blocks the main thread until an interrupt is received
// 	fmt.Println("Shutting down!")
// 	return s.Shutdown()
// }

func (s *Server) startupMessage(a string) {
	color.New(color.FgGreen, color.Bold).Printf("\nStarting server on: %v\n", a)
}

//
func (s *Server) Stop() {
	_ = s.App.Shutdown()
}

func load(s *Server) {
	// routes
	routes.LoadRoutes(s.App.Group(""))
	// middlewares
}
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	if err != nil {
		// In case the SendFile fails
		return ctx.Status(500).SendString("Internal Server Error")
	}

	// Return from handler
	return nil
}
func CustomErrorHandler(c *fiber.Ctx, err error) error {
	// StatusCode defaults to 500
	code := fiber.StatusInternalServerError
	//nolint:misspell    // Retrieve the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	} //nolint:gofmt,wsl
	er := errors.WithStack(err)
	fmt.Printf("%+v", er)
	fmt.Printf("%+v", err)
	if c.Is("json") {
		return c.Status(code).JSON(err)
	}
	return c.Status(code).Render(fmt.Sprintf("errors/%d", code), fiber.Map{ //nolint:nolintlint,errcheck
		"error": err,
	})
}

func error404(s *Server) {
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Page not Found")
	})
}
