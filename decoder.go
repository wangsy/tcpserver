package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

// Hexdec ...
func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

//Hex2Bin ...
func Hex2Bin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%016b", ui), nil
}

func main() {
	byt := []byte("{\"registertime\":\"2019-11-13T01:51:13.392+09:00\",\"heartbeattime\":\"2019-11-16T11:41:19.466+09:00\",\"client_id\":\"7f0000010bba000003a2\",\"need_device_detail\":false,\"hw_server_addr\":\"wss://hw1.a-ing.kr\",\"hw_server_port\":\"7272\",\"loginhex\":\"00206001fc112233445566778802330011534d424830383139303834393032353600\",\"last_stock_detail_time\":\"2019-11-16 11:40:29 +0900\",\"stocknum\":\"6\",\"detail\":{\"2\":{\"batteryid\":\"SMBH98005001\",\"power\":4},\"3\":{\"batteryid\":\"SMBH98001262\",\"power\":4},\"4\":{\"batteryid\":\"SMBH98001922\",\"power\":4},\"6\":{\"batteryid\":\"SMBH98005787\",\"power\":4},\"7\":{\"batteryid\":\"SMBH98008024\",\"power\":4},\"8\":{\"batteryid\":\"SMBH98005769\",\"power\":4}},\"clientip\":\"195.233.151.112\",\"heartbeathex\":\"000761010011223344\"}")
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("hello")
	fmt.Println(dat["registertime"])

	/*
				$length = hexdec(substr($hex,0,4));
		        $cmd = hexdec(substr($hex,4,2));
		        $Version = hexdec(substr($hex,6,2));
		        $CheckSum = hexdec(substr($hex,8,2));
		        $Token = hexdec(substr($hex,10,8));//Authentication //287454020 //1979/2/10 8:20:20
	*/

	hexStr := "00206001f0112233445566778802330011534d424830383139303838383037313600"
	length, _ := Hexdec(hexStr[0:4])
	cmd, _ := Hexdec(hexStr[4:6])
	version, _ := Hexdec(hexStr[6:8])
	checksum, _ := Hexdec(hexStr[8:10])
	token, _ := Hexdec(hexStr[10:18])

	fmt.Println(length)
	fmt.Println(cmd)
	fmt.Println(version)
	fmt.Println(checksum)
	fmt.Println(token)

	/*
			      Get the cabinet ID length
		      $BoxIDLen = hexdec(substr($hex,30,4)) - 1;//Subtract the length of 00 at the end to get the correct cabinet id length 16
		      //Get the cabinet ID
		      $BoxID = hex2bin(substr($hex,34,$BoxIDLen*2));

	*/

	boxidlen, _ := Hexdec(hexStr[30:34])
	fmt.Println(boxidlen)
	fmt.Println(hexStr[34 : 34+(boxidlen-1)*2])
	// boxid, err := Hex2Bin(hexStr[34 : 34+(boxidlen-1)*2])
	hexByte, err := hex.DecodeString(hexStr[34 : 34+(boxidlen-1)*2])

	fmt.Println(err)
	fmt.Println(string(hexByte))
}
