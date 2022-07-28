package main

import (
	"fmt"
	"github.com/san035/go_config_ini/config_ini"
)

func main() {
	telegram_token := config_ini.Load_all_params_from_ini("telegram", "token,webhook_url,webhook_port,ParseMode", "").(string)
	fmt.Printf("telegram_token = %s\nParseMode = %s",
		telegram_token,
		config_ini.Param_str["webhook_port"])
}
