package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {

	currentReality := ReadFile()
	currentInbound := currentReality.Inbounds[0]

	var newReality RealityJson
	newReality.Log = currentReality.Log
	newReality.Outbounds = currentReality.Outbounds

	publicKey := getPublicKey()

	serverIP := GetOutboundIP().String()

	StringConfigAll := ""

	ports := []int{443, 8081, 8082, 8083, 8084, 8085, 8086, 8087, 8088, 8089}
	domains := []string{"www.datadoghq.com",
		"apppash.ir",
		"rr2.ir",
		"bazarmag.ir",
		"shiralinia.ir",
		"000webhost.ir",
		"favakar.ir",
		"veket.ir",
		"salamat.ir",
		"tarhpro.ir"}

	newReality.Inbounds = make([]Inbound, len(domains))
	for i := 0; i < len(domains); i++ {

		inbound := currentInbound
		inbound.ListenPort = ports[i]
		inbound.TLS.ServerName = domains[i]
		inbound.TLS.Reality.Handshake.Server = domains[i]

		inbound.Users = []User{
			{
				NAME: "SB-" + domains[i],
				UUID: uuid.New().String(),
				Flow: "xtls-rprx-vision",
			},
		}

		newReality.Inbounds[i] = inbound

		StringConfig := "vless://" + inbound.Users[0].UUID + "@" + serverIP + ":" + strconv.Itoa(ports[i]) +
			"?encryption=none&flow=xtls-rprx-vision&security=reality&sni=" + domains[i] +
			"&fp=chrome&pbk=" + publicKey + "&sid=" + inbound.TLS.Reality.ShortID[0] + "&type=tcp&headerType=none#SB-" + domains[i]

		StringConfigAll += StringConfig + "\n"

	}

	//save new Reality in file
	err := WriteFile("./reality.json", newReality)
	if err != nil {
		log.Fatal("error during the WriteFile")
	}

	SaveSubscribe("./subscribe.txt", StringConfigAll)

}

func SaveSubscribe(filename string, StringConfigAll string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(StringConfigAll)
	if err2 != nil {
		log.Fatal(err2)
	}

}

func ReadFile() (currentReality RealityJson) {

	// Let's first read the `reality.json` file
	content, err := os.ReadFile("./reality.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshal the data into `currentReality`
	err = json.Unmarshal(content, &currentReality)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return currentReality
}

func WriteFile(filename string, newReality RealityJson) error {

	file, err := json.MarshalIndent(newReality, "", " ")
	if err != nil {
		log.Fatal("Error during MarshalIndent(): ", err)
		return err
	}

	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Error during WriteFile(): ", err)
		return err
	}

	return nil
}

func getPublicKey() string {
	dat, err := os.ReadFile("./public_key.txt")
	if err != nil {
		log.Fatal("error during the ReadFile")
	}
	publicKey := string(dat)

	publicKey = strings.TrimSpace(publicKey)

	return publicKey
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}