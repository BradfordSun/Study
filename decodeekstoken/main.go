package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// aws eks get-token --cluster-name SupportChatbot --region cn-north-1
	// 去掉k8s-aws-v1.
	input := `aHR0cHM6Ly9zdHMuY24tbm9ydGgtMS5hbWF6b25hd3MuY29tLmNuLz9BY3Rpb249R2V0Q2FsbGVySWRlbnRpdHkmVmVyc2lvbj0yMDExLTA2LTE1JlgtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUV4cGlyZXM9NjAmWC1BbXotQ3JlZGVudGlhbD1BS0lBUEJCTUVBRzcyUkMzNFAzQSUy
RjIwMjIwMzI3JTJGY24tbm9ydGgtMSUyRnN0cyUyRmF3czRfcmVxdWVzdCZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QlM0J4LWs4cy1hd3MtaWQmWC1BbXotRGF0ZT0yMDIyMDMyN1QxMDI3MTRaJlgtQW16LVNpZ25hdHVyZT02NzJlNTY4NzM3YmUxYjJjNDUyOTk3MWVkNjI5ZTEyOTAwZmY5MTU4MjE5ODIxOWRlOWExMTFkNGVjMjQ0NDli`
	decodeBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodeBytes))
}
