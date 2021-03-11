package config

import "os"

var defaultValues = map[string]string{
	//Telegram
	TelegramChatIDS:    os.Getenv(TelegramChatIDS),
	TelegramXAuthToken: os.Getenv(TelegramXAuthToken),
	TelegramURL:        os.Getenv(TelegramURL),

	//FTP
	FTPHost:     os.Getenv(FTPHost),
	FTPPort:     os.Getenv(FTPPort),
	FTPUser:     os.Getenv(FTPUser),
	FTPPassword: os.Getenv(FTPPassword),
	FTPPath:     os.Getenv(FTPPath),
	FTPFile:     os.Getenv(FTPFile),

	//REST
	OLAPRestHost: os.Getenv(OLAPRestHost),
	OLAPRestPort: os.Getenv(OLAPRestPort),
	//Iiko
	IikoAPIURl:   os.Getenv(IikoAPIURl),
	IikoLogin:    os.Getenv(IikoLogin),
	IikoPassword: os.Getenv(IikoPassword),

	//Mail
	MailHost:      os.Getenv(MailHost),
	MailPort:      os.Getenv(MailPort),
	MailLogin:     os.Getenv(MailLogin),
	MailPassword:  os.Getenv(MailPassword),
	MailAddressee: os.Getenv(MailAddressee),

	//OLAP report
	OLAPMailSubject:    os.Getenv(OLAPMailSubject),
	OLAPMailRecipients: os.Getenv(OLAPMailRecipients),

	//Amount report
	AmountMailSubject:    os.Getenv(AmountMailSubject),
	AmountMailRecipients: os.Getenv(AmountMailRecipients),

	//Daemons
	OlAPDaemonScheduler:        os.Getenv(OlAPDaemonScheduler),
	AmountDaemonScheduler:      os.Getenv(AmountDaemonScheduler),
	IntegrationDaemonScheduler: os.Getenv(IntegrationDaemonScheduler),
}
