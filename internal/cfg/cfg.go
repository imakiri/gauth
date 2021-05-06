package cfg

type Config struct {
	Init ConfigInit
}

type ConfigInit struct {
	Messenger ConfigInitMessenger
}

type ConfigInitMessenger struct {
	Email    *ConfigInitMessengerEmail
	Telegram *ConfigInitMessengerTelegram
}

// --

type ConfigInitMessengerEmail struct {
	Basic *ConfigInitMessengerEmailBasic
}
type ConfigInitMessengerTelegram struct {
	Basic *ConfigInitMessengerTelegramBasic
}

// --

type ConfigInitMessengerEmailBasic struct {
	Identity string
	Address  string
	User     string
	Password string
	From     string
}
type ConfigInitMessengerTelegramBasic struct {
}
