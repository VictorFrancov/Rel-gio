package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var fusosHorarios = map[string]string{
	"Pacific/Midway":                 "Pacific/Midway",
	"Pacific/Pago_Pago":              "Pacific/Pago_Pago",
	"Pacific/Honolulu":               "Pacific/Honolulu",
	"America/Juneau":                 "America/Juneau",
	"America/Los_Angeles":            "America/Los_Angeles",
	"America/Tijuana":                "America/Tijuana",
	"America/Denver":                 "America/Denver",
	"America/Phoenix":                "America/Phoenix",
	"America/Chihuahua":              "America/Chihuahua",
	"America/Mazatlan":               "America/Mazatlan",
	"America/Chicago":                "America/Chicago",
	"America/Regina":                 "America/Regina",
	"America/Mexico_City":            "America/Mexico_City",
	"America/Monterrey":              "America/Monterrey",
	"America/Guatemala":              "America/Guatemala",
	"America/New_York":               "America/New_York",
	"America/Indiana/Indianapolis":   "America/Indiana/Indianapolis",
	"America/Bogota":                 "America/Bogota",
	"America/Lima":                   "America/Lima",
	"America/Halifax":                "America/Halifax",
	"America/Caracas":                "America/Caracas",
	"America/La_Paz":                 "America/La_Paz",
	"America/Santiago":               "America/Santiago",
	"America/St_Johns":               "America/St_Johns",
	"America/Sao_Paulo":              "America/Sao_Paulo",
	"America/Argentina/Buenos_Aires": "America/Argentina/Buenos_Aires",
	"America/Guyana":                 "America/Guyana",
	"America/Godthab":                "America/Godthab",
	"Atlantic/South_Georgia":         "Atlantic/South_Georgia",
	"Atlantic/Azores":                "Atlantic/Azores",
	"Atlantic/Cape_Verde":            "Atlantic/Cape_Verde",
	"Europe/Dublin":                  "Europe/Dublin",
	"Europe/Lisbon":                  "Europe/Lisbon",
	"Europe/London":                  "Europe/London",
	"Africa/Casablanca":              "Africa/Casablanca",
	"Africa/Monrovia":                "Africa/Monrovia",
	"Etc/UTC":                        "Etc/UTC",
	"Europe/Belgrade":                "Europe/Belgrade",
	"Europe/Bratislava":              "Europe/Bratislava",
	"Europe/Budapest":                "Europe/Budapest",
	"Europe/Ljubljana":               "Europe/Ljubljana",
	"Europe/Prague":                  "Europe/Prague",
	"Europe/Sarajevo":                "Europe/Sarajevo",
	"Europe/Skopje":                  "Europe/Skopje",
	"Europe/Warsaw":                  "Europe/Warsaw",
	"Europe/Zagreb":                  "Europe/Zagreb",
	"Europe/Brussels":                "Europe/Brussels",
	"Europe/Copenhagen":              "Europe/Copenhagen",
	"Europe/Madrid":                  "Europe/Madrid",
	"Europe/Paris":                   "Europe/Paris",
	"Europe/Amsterdam":               "Europe/Amsterdam",
	"Europe/Berlin":                  "Europe/Berlin",
	"Europe/Rome":                    "Europe/Rome",
	"Europe/Stockholm":               "Europe/Stockholm",
	"Europe/Vienna":                  "Europe/Vienna",
	"Africa/Algiers":                 "Africa/Algiers",
	"Europe/Bucharest":               "Europe/Bucharest",
	"Africa/Cairo":                   "Africa/Cairo",
	"Europe/Helsinki":                "Europe/Helsinki",
	"Europe/Kiev":                    "Europe/Kiev",
	"Europe/Riga":                    "Europe/Riga",
	"Europe/Sofia":                   "Europe/Sofia",
	"Europe/Tallinn":                 "Europe/Tallinn",
	"Europe/Vilnius":                 "Europe/Vilnius",
	"Europe/Athens":                  "Europe/Athens",
	"Europe/Istanbul":                "Europe/Istanbul",
	"Europe/Minsk":                   "Europe/Minsk",
	"Asia/Jerusalem":                 "Asia/Jerusalem",
	"Africa/Harare":                  "Africa/Harare",
	"Europe/Moscow":                  "Europe/Moscow",
	"Asia/Kuwait":                    "Asia/Kuwait",
	"Asia/Riyadh":                    "Asia/Riyadh",
	"Africa/Nairobi":                 "Africa/Nairobi",
	"Asia/Baghdad":                   "Asia/Baghdad",
	"Asia/Tehran":                    "Asia/Tehran",
	"Asia/Muscat":                    "Asia/Muscat",
	"Asia/Baku":                      "Asia/Baku",
	"Asia/Tbilisi":                   "Asia/Tbilisi",
	"Asia/Yerevan":                   "Asia/Yerevan",
	"Asia/Kabul":                     "Asia/Kabul",
	"Asia/Yekaterinburg":             "Asia/Yekaterinburg",
	"Asia/Karachi":                   "Asia/Karachi",
	"Asia/Tashkent":                  "Asia/Tashkent",
	"Asia/Kolkata":                   "Asia/Kolkata",
	"Asia/Kathmandu":                 "Asia/Kathmandu",
	"Asia/Dhaka":                     "Asia/Dhaka",
	"Asia/Colombo":                   "Asia/Colombo",
	"Asia/Almaty":                    "Asia/Almaty",
	"Asia/Novosibirsk":               "Asia/Novosibirsk",
	"Asia/Rangoon":                   "Asia/Rangoon",
	"Asia/Bangkok":                   "Asia/Bangkok",
	"Asia/Jakarta":                   "Asia/Jakarta",
	"Asia/Krasnoyarsk":               "Asia/Krasnoyarsk",
	"Asia/Shanghai":                  "Asia/Shanghai",
	"Asia/Chongqing":                 "Asia/Chongqing",
	"Asia/Hong_Kong":                 "Asia/Hong_Kong",
	"Asia/Urumqi":                    "Asia/Urumqi",
	"Asia/Kuala_Lumpur":              "Asia/Kuala_Lumpur",
	"Asia/Singapore":                 "Asia/Singapore",
	"Asia/Taipei":                    "Asia/Taipei",
	"Australia/Perth":                "Australia/Perth",
	"Asia/Irkutsk":                   "Asia/Irkutsk",
	"Asia/Ulaanbaatar":               "Asia/Ulaanbaatar",
	"Asia/Seoul":                     "Asia/Seoul",
	"Asia/Tokyo":                     "Asia/Tokyo",
	"Asia/Yakutsk":                   "Asia/Yakutsk",
	"Australia/Darwin":               "Australia/Darwin",
	"Australia/Adelaide":             "Australia/Adelaide",
	"Australia/Sydney":               "Australia/Sydney",
	"Australia/Brisbane":             "Australia/Brisbane",
	"Australia/Hobart":               "Australia/Hobart",
	"Asia/Vladivostok":               "Asia/Vladivostok",
	"Pacific/Guam":                   "Pacific/Guam",
	"Asia/Magadan":                   "Asia/Magadan",
	"Pacific/Noumea":                 "Pacific/Noumea",
	"Pacific/Fiji":                   "Pacific/Fiji",
	"Asia/Kamchatka":                 "Asia/Kamchatka",
	"Pacific/Majuro":                 "Pacific/Majuro",
	"Pacific/Auckland":               "Pacific/Auckland",
	"Pacific/Tongatapu":              "Pacific/Tongatapu",
}

func mostrarHora(fuso string, localizacao string) {
	local, err := time.LoadLocation(localizacao)
	if err != nil {
		fmt.Printf("Erro ao carregar a localização: %v\n", err)
		return
	}
	for {
		horaAtual := time.Now().In(local)
		fmt.Printf("A hora atual em %s (%s) é: %s\n", fuso, localizacao, horaAtual.Format("15:04:05 MST"))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Relógio Mundial")
	fmt.Println("-----------------")
	fmt.Println("Escolha um fuso horário para ver a hora atual:")
	for fuso := range fusosHorarios {
		fmt.Println(" -", fuso)
	}

	fmt.Print("\nDigite os códigos dos fusos horários separados por vírgula (ou 'todos'): ")
	entradaFusosHorarios, _ := leitor.ReadString('\n')
	entradaFusosHorarios = strings.TrimSpace(entradaFusosHorarios)

	fusosSelecionados := strings.Split(entradaFusosHorarios, ",")

	if len(fusosSelecionados) == 1 && fusosSelecionados[0] == "todos" {
		for fuso, localizacao := range fusosHorarios {
			go mostrarHora(fuso, localizacao)
		}
	} else {
		for _, fuso := range fusosSelecionados {
			fusoTrim := strings.TrimSpace(fuso)
			localizacao, existe := fusosHorarios[fusoTrim]
			if !existe {
				fmt.Printf("Fuso horário não reconhecido: %s\n", fusoTrim)
				continue
			}
			go mostrarHora(fusoTrim, localizacao)
		}
	}

	select {}
}
