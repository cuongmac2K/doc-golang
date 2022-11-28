package main

import (
	"fmt"

	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/sulat"
)

// HTMLHeader is the HTML skeletal framework head section of the standard HTML structure that serves as an email content
const HTMLHeader = ``

// HTMLFooter is the HTML skeletal framework footer section of the standard HTML structure that serves as an email content
const HTMLFooter = ``

// bodyHTML is your custom email contents, this is just an example of a password reset auto email from your Go's project
var bodyHTML = ``

// FullHTML use this when you preferred the full HTML template as your content
var FullHTML = `<!DOCTYPE html>
<html>
<head>
<title>Page Title</title>
</head>
<body>

<h1>This is a Heading</h1>
<p>This is a paragraph.</p>

</body>
</html>`

// SGC to initialize SendGrid API
var SGC = sulat.SGC{}

func init() {
	SGC = sulat.SGC{
		SendGridAPIKey:   "YOUR_SENDGRID_API_KEY_HERE",
		SendGridEndPoint: "/v3/mail/send",
		SendGridHost:     "https://api.sendgrid.com",
	}
}

func main() {
	fmt.Println("Hello Maharlikans!")

	mailOpt := &sulat.SendMail{
		Subject: "Inquiry from Maharlikans Code!",
		From:    sulat.NewEmail("Cuong", "cuongmac2k@gmail.com"),
		To:      sulat.NewEmail("Cuong", "cuongmac2k@gmail.com"),
	}

	// Method 1 by using the Full HTML or Static Content
	htmlContent, err := sulat.SetHTML(&sulat.EmailHTMLFormat{
		IsFullHTML:       true,
		FullHTMLTemplate: FullHTML,
	})

	isSend, err := sulat.SendEmailSG(mailOpt, htmlContent, &SGC)
	if err != nil {
		itrlog.Fatal(err)
	}
	fmt.Println("is sent: ", isSend)
}
