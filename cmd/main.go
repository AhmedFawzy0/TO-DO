package main
import(
	 "github.com/gofiber/fiber/v2"
	 "github.com/gofiber/template/html"
	"html/template"
	"github.com/AhmedFawzy0/TO-DO/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var tpl *template.Template

func main(){
	database.ConnectDb()
	database.ConnectRedis()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine, 
	})

	
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	setupRoutes(app)
	app.Static("/","./static")

	app.Listen(":3000")

}

