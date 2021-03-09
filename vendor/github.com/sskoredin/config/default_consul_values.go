package config

var defaultValues = map[string]string{
	//Telegram
	TelegramChatIDS:    "[492307185]",
	TelegramXAuthToken: "",
	TelegramURL:        "logger:3250",

	//FTP
	FTPHost:     "ftp.davido01.nichost.ru",
	FTPPort:     "21",
	FTPUser:     "davido01_skor",
	FTPPassword: "uLigERaOUpEliSChI1",
	FTPPath:     "davidoff-bronnaya.ru/auto",
	FTPFile:     "import.csv",

	//REST
	OLAPRestHost: "0.0.0.0",
	OLAPRestPort: "8230",
	//Iiko
	IikoAPIURl:   "http://94.127.179.181:9081",
	IikoLogin:    "Yakupov",
	IikoPassword: "1Qz2876",

	//Mail
	MailHost:      "mail.nic.ru",
	MailPort:      "587",
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
	OlAPDaemonScheduler:        "0 9 * * *",
	AmountDaemonScheduler:      "0 9 * * *",
	IntegrationDaemonScheduler: "0 10 * * *",
}
