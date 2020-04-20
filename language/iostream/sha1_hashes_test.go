package iostream

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestShaHashes(t *testing.T) {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	str := fmt.Sprintf("%x", bs)

	//from other place to get a sha1 string
	assertEq("cf23df2207d99a74fbe169e3eba035e633b65d94", str)

	pShow(bs)

	contens := []byte("sha256 this string")

	s256 := sha256.New()

	s256.Write(contens)

	bs = s256.Sum(nil)

	str = fmt.Sprintf("%x", bs)
	//from other place to get a sha1 string
	assertEq("1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb", str)
	pShow(bs)

}

func sha1Hashes() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
