package main

import (
	"doc/testbottele/web"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
	"net"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	t := new(TelegramLog)

	t.LogRestartServer()
	Webhook(app)
	app.Listen(":3000")
}
func (te TelegramLog) LogRestartServer() {
	url := TELEGRAM_CONFIG["DOMAIN"] + "bot" + "5488767666:AAExtp7kn1yblEN1JdOGE6PUZeugWmEW1Yw" + "/sendMessage"
	bodyString, errB := json.Marshal(te.Body)
	fmt.Println(errB)
	resString, errR := json.Marshal(te.Response)
	fmt.Println(errR)
	html := fmt.Sprintf("<b>%s</b>\n<b>SOURCE: fiber+heloooo be lo %s</b>\n<b>Host :%s</b>\n<b>URL: </b>%s\n<b>Body:</b><code>%s</code>\n<b>Response:</b><code>%s</code>\n<b>Error Message: </b>%s\n<b>Response time: </b>%s\n", Env, te.Title, GetOutboundIP().String(), te.URL,
		string(bodyString), string(resString), te.ErrorMessage, te.ResponseTime.String())
	data := map[string]interface{}{
		"chat_id":    1001746636889,
		"text":       html,
		"parse_mode": "HTML",
	}
	dataSentByte, errJ := json.Marshal(data)
	if errJ != nil {
		fmt.Println("Error get data byte LogToTelegram: " + errJ.Error())
		return
	}
	// không call được hàm SendRequest vì sẽ bị lỗi gọi vòng tròn
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBodyRaw(dataSentByte)
	resp := fasthttp.AcquireResponse()

	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		fmt.Println("Error get data byte LogToTelegram: " + err.Error())
	} else {
		bodyBytes := resp.Body()
		fmt.Println("Sent data byte LogToTelegram")
		fmt.Println("Sent data byte LogToTelegram" + string(bodyBytes))
	}
}
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func Webhook(app *fiber.App) {
	api := app.Group("/")
	group := api.Group("/")
	group.Get("/getWebhook", web.Getwebhook)
	group.Post("/getWebhook", web.Getwebhook)
}
