package config

var defaultValues = map[string]string{
	TelegramChatIDS: "",

	OLAPRestHost: "0.0.0.0",
	OLAPRestPort: "8230",

	IikoAPIURl:   "http://94.127.179.181:9081",
	IikoLogin:    "Yakupov",
	IikoPassword: "1Qz2876",

	MailHost:      "mail.nic.ru:587",
	MailLogin:     "robot@davidoffclub.ru",
	MailPassword:  "SwITERyCHEwConoT1",
	MailAddressee: "robot",
	//OLAP report
	OLAPMailSubject:    "OLAP report",
	OLAPMailRecipients: "[\"sergey@skoredin.pro\"]",
	//Amount report
	AmountMailSubject:    "mail_amount_subject",
	AmountMailRecipients: "mail_amount_recipients",
	//Daemons
	OlAPDaemonScheduler:   "0 9 * * *",
	AmountDaemonScheduler: "0 9 * * *",
}
