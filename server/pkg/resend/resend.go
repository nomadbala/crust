package resend

import (
	"fmt"
	"github.com/nomadbala/crust/server/internal/config"
	"github.com/nomadbala/crust/server/pkg/log"
	"github.com/resend/resend-go/v2"
	"go.uber.org/zap"
)

var client *resend.Client

func ConfigureResendClient(cfg config.Resend) {
	client = resend.NewClient(cfg.ApiKey)
}

func SendResendMessage(receiver, message string) error {
	params := &resend.SendEmailRequest{
		From:    "crust@kiteo.app",
		To:      []string{receiver},
		Html:    fmt.Sprintf("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\"><html dir=\"ltr\" lang=\"en\"><head><meta content=\"width=device-width\" name=\"viewport\"/><link rel=\"preload\" as=\"image\" href=\"https://raw.githubusercontent.com/nomadbala/crust/refs/heads/main/docs/crust.jpg\"/><meta content=\"text/html; charset=UTF-8\" http-equiv=\"Content-Type\"/><meta name=\"x-apple-disable-message-reformatting\"/><meta content=\"IE=edge\" http-equiv=\"X-UA-Compatible\"/><meta name=\"x-apple-disable-message-reformatting\"/><meta content=\"telephone=no,address=no,email=no,date=no,url=no\" name=\"format-detection\"/><meta content=\"light\" name=\"color-scheme\"/><meta content=\"light\" name=\"supported-color-schemes\"/><!--$--><style>\n    @font-face {\n      font-family: 'Inter';\n      font-style: normal;\n      font-weight: 400;\n      mso-font-alt: 'sans-serif';\n      src: url(https://rsms.me/inter/font-files/Inter-Regular.woff2?v=3.19) format('woff2');\n    }\n\n    * {\n      font-family: 'Inter', sans-serif;\n    }\n  </style><style>blockquote,h1,h2,h3,img,li,ol,p,ul{margin-top:0;margin-bottom:0}</style></head><body><table align=\"center\" width=\"100%\" border=\"0\" cellPadding=\"0\" cellSpacing=\"0\" role=\"presentation\" style=\"max-width:600px;min-width:300px;width:100%;margin-left:auto;margin-right:auto;padding:0.5rem\"><tbody><tr style=\"width:100%\"><td><h2 style=\"text-align:left;color:rgb(17, 24, 39);margin-bottom:12px;margin-top:0px;font-size:30px;line-height:36px;font-weight:700\"></h2><table align=\"center\" width=\"100%\" border=\"0\" cellPadding=\"0\" cellSpacing=\"0\" role=\"presentation\" style=\"margin-top:0px;margin-bottom:32px\"><tbody style=\"width:100%\"><tr style=\"width:100%\"><td align=\"left\" data-id=\"__react-email-column\"><img title=\"Image\" alt=\"Image\" src=\"https://raw.githubusercontent.com/nomadbala/crust/608d49e0d51e14ceb174662321c8b497bca20892/docs/crust.svg\" style=\"display:block;outline:none;border:none;text-decoration:none;height:100%;width:100%;max-width:75px;max-height:75px\"/></td></tr></tbody></table><h2 style=\"text-align:left;color:rgb(17, 24, 39);margin-bottom:12px;margin-top:0px;font-size:30px;line-height:36px;font-weight:700\"><strong>Hello!</strong></h2><p style=\"font-size:15px;line-height:24px;margin:16px 0;text-align:left;margin-bottom:20px;margin-top:0px;color:rgb(55, 65, 81);-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale\">Welcome to Crust. We’re excited to have you in our community.</p><p style=\"font-size:15px;line-height:24px;margin:16px 0;text-align:left;margin-bottom:20px;margin-top:0px;color:rgb(55, 65, 81);-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale\">To complete your registration, please verify your email by entering the following 6-digit code:</p><p style=\"font-size:15px;line-height:24px;margin:16px 0;text-align:left;margin-bottom:20px;margin-top:0px;color:rgb(55, 65, 81);-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale\">" + message + "</p><p style=\"font-size:15px;line-height:24px;margin:16px 0;text-align:left;margin-bottom:20px;margin-top:0px;color:rgb(55, 65, 81);-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale\">If you have any questions or need assistance, feel free to reach out to us!</p><p style=\"font-size:15px;line-height:24px;margin:16px 0;text-align:left;margin-bottom:20px;margin-top:0px;color:rgb(55, 65, 81);-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale\">Regards,<br/>Crust</p></td></tr></tbody></table><!--/$--></body></html>"),
		Subject: "Crust",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Logger.Error("error occurred while sending message via resend", zap.Error(err))
		return err
	}

	return nil
}
