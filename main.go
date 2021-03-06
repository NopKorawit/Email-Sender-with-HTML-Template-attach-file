package main

import (
	// "crypto/tls"
	"fmt"
	"send/config"

	"gopkg.in/gomail.v2"
)

func main() {

	//load config
	configModel := config.EmailConfig{}
	configModel.LoadConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", configModel.Email)
	m.SetHeader("To", "example.was@gmail.com") //ส่งไปอีเมลไหน
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan") //ส่ง cc
	m.SetHeader("Subject", "Hello! P'Steven Ball") //หัวข้อ

	body := ` 
	<table class="nl-container" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f8f3ed;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr>
    <td>
    <table class="row row-1" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td>
    <table class="row-content stack" style="color: #000000; width: 605px; height: 85px;" role="presentation" border="0" width="640" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td class="column column-1" style="font-weight: 400; text-align: left; vertical-align: top; padding-top: 30px; padding-bottom: 0px; border: 0px; width: 603px;">
    <table class="image_block" style="width: 109.141%; height: 10px;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr style="height: 10px;">
    <td style="width: 100%; padding-right: 0px; padding-left: 0px; padding-top: 60px; height: 10px;"><img class="big" style="display: block; height: auto; border: 0; width: 640px; max-width: 100%;" width="640" /></td>
    
    <table class="row row-2" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td>
    <table class="row-content stack" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #688a66; background-position: top center; color: #000000; width: 640px;" role="presentation" border="0" width="640" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td class="column column-1" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 0px; padding-bottom: 0px; border: 0px;" width="100%">
    <table class="image_block" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr>
    <td style="width: 100%; padding-right: 0px; padding-left: 0px; padding-top: 15px;">
    <div style="line-height: 10px;" align="center"><img class="big" style="display: block; height: 394px; border-width: 0px; border-color: initial; border-image: initial; width: 394px; max-width: 100%; margin-left: 50px; margin-right: 50px;" title="Your Logo" src="https://th.bing.com/th/id/OIP.xUIFDINSXBZiCLE2cavh-wHaHa?pid=ImgDet&amp;rs=1" alt="Your Logo" width="474" /></div>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="heading_block" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr>
    <td style="text-align: center; width: 100%; padding-top: 35px;">
    <h1 style="margin: 0; color: #ffffff; direction: ltr; font-family: Montserrat, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif; font-size: 50px; font-weight: 400; letter-spacing: 5px; line-height: 120%; text-align: center; margin-top: 0; margin-bottom: 0;"><strong>ThaiBev</strong></h1>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="heading_block" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr>
    <td style="text-align: center; width: 100%;">
    <h1 style="margin: 0; color: #ffffff; direction: ltr; font-family: Montserrat, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif; font-size: 50px; font-weight: 400; letter-spacing: 5px; line-height: 120%; text-align: center; margin-top: 0; margin-bottom: 0;"><strong>Invitation</strong></h1>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="button_block" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr>
    <td style="text-align: center; padding: 35px 20px 10px 20px;">
    <div align="center"><a style="text-decoration: none; display: inline-block; color: #ffffff; background-color: #08132f; border-radius: 6px; width: auto; font-weight: 400; padding-top: 8px; padding-bottom: 8px; font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; text-align: center; mso-border-alt: none; word-break: keep-all; border: 1px solid #08132f;" href="https://www.thaibev.com/career/" target="_blank"><strong>Join now</strong></a></div>
    </td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="row row-3" style="height: 22px; width: 100%;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr style="height: 22px;">
    <td style="height: 22px;">
    <table class="row-content stack" style="color: #000000; width: 640px; height: 18px;" role="presentation" border="0" width="640" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr style="height: 18px;">
    <td class="column column-1" style="font-weight: 400; text-align: left; vertical-align: top; padding-top: 0px; padding-bottom: 0px; border: 0px; height: 18px; width: 638px;"> </td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="row row-4" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td>
    <table class="row-content stack" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 640px;" role="presentation" border="0" width="640" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td class="column column-1" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 0px; padding-bottom: 0px; border: 0px;" width="100%">
    <table class="social_block" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr>
    <td style="text-align: center; padding: 70px 10px 10px 10px;">
    <table class="social-table" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="126px" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td style="padding: 0 5px 0 5px;"><a href="https://www.facebook.com/ThaiBeverage" target="_blank"><img style="display: block; height: auto; border-width: 0px; border-color: initial; border-image: initial;" title="" src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Facebook_Logo_%282019%29.png/768px-Facebook_Logo_%282019%29.png" alt="Facebook" width="32" height="32" /></a></td>
    <td style="padding: 0 5px 0 5px;"><a href="https://twitter.com/thai_bev" target="_blank"><img style="display: block; height: auto; border-width: 0px; border-color: initial; border-image: initial;" title="Twitter" src="https://openvisualfx.com/wp-content/uploads/2019/10/pnglot.com-twitter-bird-logo-png-139932.png" alt="Twitter" width="32" height="32" /></a></td>
    <td style="padding: 0 5px 0 5px;"><a href="https://www.youtube.com/user/Thaibeverage" target="_blank"><img style="display: block; height: auto; border-width: 0px; border-color: initial; border-image: initial;" title="YouTube" src="https://cdn.icon-icons.com/icons2/2429/PNG/512/youtube_logo_icon_147199.png" alt="YouTube" width="32" height="32" /></a></td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="text_block" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="10">
    <tbody>
    <tr>
    <td>
    <div style="font-family: Tahoma, Verdana, sans-serif;">
    <div class="txtTinyMce-wrapper" style="font-size: 12px; font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #3c3c3c; line-height: 1.2;">
    <p style="margin: 0; font-size: 12px; text-align: center;">https://www.thaibev.com/</p>
    </div>
    </div>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="text_block" style="word-break: break-word; height: 14px; width: 100%;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0">
    <tbody>
    <tr style="height: 14px;">
    <td style="padding: 10px 10px 40px; height: 14px;">
    <div style="font-family: Tahoma, Verdana, sans-serif;">
    <div class="txtTinyMce-wrapper" style="font-size: 12px; mso-line-height-alt: 14.399999999999999px; color: #3c3c3c; line-height: 1.2; font-family: 'Lato', Tahoma, Verdana, Segoe, sans-serif;">
    <p style="margin: 0; text-align: center;">Copyright © 2021 Thai Beverage Plc. All Rights Reserved.</p>
    </div>
    </div>
    </td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
    <table class="row row-5" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" role="presentation" border="0" width="100%" cellspacing="0" cellpadding="0" align="center">
    <tbody>
    <tr>
    <td>
    <div class="spacer_block" style="height: 0px; line-height: 0px; font-size: 1px;"> </div>
    </td>
    </tr>
    </tbody>
    </table>
    </td>
    </tr>
    </tbody>
    </table>
	`

	// fmt.Println(body)
	m.SetBody("text/html", body)
	// m.AddAlternative("text/html", "<p>Hello!</p>")
	pic1 := "E:/xxxx2.jpg"

	m.Attach(pic1)
	m.Attach("C:/xxxx.pdf")
	m.Attach("C:/xxxx.zip")
	m.Attach("C:/xxxx.xlsx")

	d := gomail.NewDialer(configModel.SmtpHost, configModel.SmtpPort, configModel.Email, configModel.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email Sent Successfully!")
}
