/*
https://github.com/san035/go_config_ini
go get github.com/san035/go_config_ini
*/
package config_ini

import (
	"encoding/json"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

var Param_str = map[string]string{}
var Param_bool = map[string]bool{}
var Param_int = map[string]int{}
var Param_int64 = map[string]int64{}
var Param_float64 = map[string]float64{}
var Param_ints = map[string][]int{}       // массив int
var Param_strings = map[string][]string{} // массив строк
var Param_map_string = map[string]map[string]string{}
var Сfg *ini.File

func init() {
	Load_config_ini() // Читаем config.ini
}

func Load_config_ini() (err error) {
	FileNameIni := "config.ini"
	Сfg, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment:        true, // игнор ;
		AllowPythonMultilineValues: true,
		UnescapeValueDoubleQuotes:  false,
	}, FileNameIni)
	CheckFatallError(err)
	return
}

//загружаем все строковые параметры из ini
//str_list_param_ini - список параметров через запятую
// type_param - пустое значение загружаемого параметре, по нему определяется тип
func Load_all_params_from_ini(section_ini, str_list_param_ini string, type_param interface{}) (return_value interface{}) {
	var err error
	for id_param, key_ini := range strings.Split(str_list_param_ini, ",") {
		// определение типа переменной
		switch type_param.(type) {
		case string:
			new_value := Сfg.Section(section_ini).Key(key_ini).String()
			if new_value == "" {
				log.Printf("Не найдено значение в config.ini [%s]%s\n", section_ini, key_ini)
			} else {
				Param_str[key_ini] = new_value
				if id_param == 0 {
					return_value = new_value
				}
			}

		case bool:
			Param_bool[key_ini], err = Сfg.Section(section_ini).Key(key_ini).Bool()
			if err != nil {
				log.Println(err.Error())
				Param_bool[key_ini] = type_param.(bool)
			}
			if id_param == 0 {
				return_value = Param_bool[key_ini]
			}

		case []string:
			Param_strings[key_ini] = Сfg.Section(section_ini).Key(key_ini).Strings(",")

		case float64:
			Param_float64[key_ini], _ = Сfg.Section(section_ini).Key(key_ini).Float64()
			if id_param == 0 {
				return_value = Param_float64[key_ini]
			}

		case int:
			Param_int[key_ini], _ = Сfg.Section(section_ini).Key(key_ini).Int()

		case int64:
			Param_int64[key_ini], _ = Сfg.Section(section_ini).Key(key_ini).Int64()
			if id_param == 0 {
				return_value = Param_int64[key_ini]
			}

		case []int:
			Param_ints[key_ini] = Сfg.Section(section_ini).Key(key_ini).Ints(",")

		case map[string]string:

			// строку в map
			var map_link = map[string]string{}
			value_key := Сfg.Section(section_ini).Key(key_ini).String()
			json.Unmarshal([]byte(value_key), &map_link)
			Param_map_string[key_ini] = map_link
			if len(map_link) == 0 && len(value_key) != 0 {
				log.Printf("Не прочитан параметр [%s]%s=%s", section_ini, key_ini, value_key)
			}

		default:
			log.Printf("Не известный тип %+v в config.ini[%s]%s\n", type_param, section_ini, key_ini)
		}
	}
	return
}

func CheckFatallError(err error, args ...interface{}) {
	if err != nil {
		log.Fatal(err, args)
	}
}
