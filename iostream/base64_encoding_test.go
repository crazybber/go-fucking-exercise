package iostream

import (
	b64 "encoding/base64"
	"fmt"
	"testing"
)

// BASE64URL是一种在BASE64的基础上编码形成新的加密方式，为了编码能在网络中安全顺畅传输，需要对BASE64进行的编码，特别是互联网中。
//编码表中不包含：+/ 但是报包含：- _
func TestBase64Functions(t *testing.T) {
	data := "crazy_123!?$bber*&()'-=@~!go♧ ☂ ♨ ☎ ☏ ☆ ★ △ ▲ ♠ ♣♟ ✖ ♂ ♀"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	assertEq("Y3JhenlfMTIzIT8kYmJlciomKCknLT1AfiFnb+KZpyDimIIg4pmoIOKYjiDimI8g4piGIOKYhSDilrMg4payIOKZoCDimaPimZ8g4pyWIOKZgiDimYA=", sEnc)

	//decode it
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)

	assertEq(data, string(sDec))

	// encode IN URL (with "- and _") without / +
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))

	assertEq("Y3JhenlfMTIzIT8kYmJlciomKCknLT1AfiFnb-KZpyDimIIg4pmoIOKYjiDimI8g4piGIOKYhSDilrMg4payIOKZoCDimaPimZ8g4pyWIOKZgiDimYA=", uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)

	assertEq([]byte(data), uDec)

}
func base64Functions() {

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
