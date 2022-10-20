package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func getDeviceId(l int) string {
	str := "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}

	ok1, _ := regexp.MatchString(".[1|2|3|4|5|6|7|8|9]", string(result))
	ok2, _ := regexp.MatchString(".[Z|X|C|V|B|N|M|A|S|D|F|G|H|J|K|L|Q|W|E|R|T|Y|U|I|P]", string(result))
	if ok1 && ok2 {
		return string(result)
	} else {
		return getDeviceId(l)
	}
}

func getDsRandStr(l int) string {
	str := "123456789abcdefghijklmnpqrstuvwxyz"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}

	ok1, _ := regexp.MatchString(".[1|2|3|4|5|6|7|8|9]", string(result))
	ok2, _ := regexp.MatchString(".[z|x|c|v|b|n|m|a|s|d|f|g|h|j|k|l|q|w|e|r|t|y|u|i|p]", string(result))
	if ok1 && ok2 {
		return string(result)
	} else {
		return getDeviceId(l)
	}
}

func getDs() string {
	t := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(999999)
	origin := fmt.Sprintf("salt=z8DRIUjNDT7IT5IZXvrUAxyupA1peND9&t=%d&r=%d", t, r)
	m := md5.New()
	m.Write([]byte(origin))
	encStr := hex.EncodeToString(m.Sum(nil))
	return fmt.Sprintf("%d,%d,%s", t, r, encStr)
}

func GetDSToken(isBBS bool) string {
	var salt string
	if isBBS {
		salt = "N50pqm7FSy2AkFz2B3TqtuZMJ5TOl3Ep"
	} else {
		salt = "z8DRIUjNDT7IT5IZXvrUAxyupA1peND9"
	}
	t := time.Now().Unix()
	rs := getDsRandStr(6)
	origin := fmt.Sprintf("salt=%s&t=%d&r=%s", salt, t, rs)
	s := md5Str(origin)
	return fmt.Sprintf("%d", t) + "," + rs + "," + s
}

func md5Str(s string) string {
	h := md5.New()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
