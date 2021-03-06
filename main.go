package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/jcmuller/gozenity"
	"github.com/martinlindhe/notify"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const (
	YANDEXLOGO = `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
VGoq4DDo6OM/LS/0j3n4LO0zoWsvLmVFYH0tDzyRh8g+qOoe8/FW1Q0NNbqSOlpImxQsGGtb/wB4
lS8Gmm3tW7Od4txumCJxYJ3v5+Efd901NDR0sdNTxiOKNoaxo5AL1RFZuHmZmd5EREeCIiAiIgIi
ICIiAiIgIiICIiAiIgIiIPGppaetgdBUwsmidzY9uQVCLvsypJy6S11DqZx49FJ7TPA8x8VPUWF8
dL/3QlabW59LO+G230+Skq7Q+oKEnNCZ2D6UDg/Phz+C0s1DWU5InpJ4iOp8Zb9oXQ6KLbRV8JXm
L0nzRH/JSJ93T93OQY9xwGuJ7AFm01kutWQKe3VUmetsTseeMLoBF5Gijxs229KLzHs4oj47/wCo
VBbtnF6qyHVXRUcZ577t53kPvIU5suhbRZ3NldGauoHESTDIB9zeQ+JUmRb6afHTrEKnVcZ1epjl
tbaPKOn3ERFvVQiIgIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiAiI
gIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiAiIgIiICIiD/9k=`
)

var (
	appDir      = os.Getenv("HOME") + "/.yatranslate/"
	logoFile    = appDir + "yalogo.jpeg"
	yandexToken = appDir + "token"
)

// YandexResponse respose json from Yandex
type YandexResponse struct {
	Text []string `json:"text"`
}

// DecodeLogo decode logo from string base64
func DecodeLogo(logo string) ([]byte, error) {
	l, err := base64.StdEncoding.DecodeString(logo)
	if err != nil {
		log.Println("error in decoding logo")
		return nil, err
	}
	return l, nil
}

// SaveLogo save logo to file
func SaveLogo(f string, l []byte) (bool, error) {
	err := ioutil.WriteFile(logoFile, l, 0744)
	if err != nil {
		log.Printf("error write to file: %s", logoFile)
		return false, err
	}
	return true, nil
}

// GetTranslate get a translate from Yandex
func GetTranslate(text []string) string {
	v := url.Values{}

	v.Set("text", strings.Join(text, " "))

	t, err := ioutil.ReadFile(yandexToken)

	if err != nil {
		log.Fatalln("Can't read token from file ", yandexToken)
	}
	token := string(t)
	url := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/translate?key=%s&%s&lang=ru&format=plain", token, v.Encode())
	res, err := http.Get(url)

	if err != nil {
		panic("api yandex not available")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var ya YandexResponse

	json.Unmarshal(body, &ya)

	return ya.Text[0]
}

// CheckApplicationDir check a dirictory of application
func CheckApplicationDir() {
	_, err := os.Stat(appDir)
	if err != nil {
		log.Println("Create a directory of application " + appDir)
		err := os.Mkdir(appDir, 0700)
		if err != nil {
			log.Println("Can't create a directory of application " + appDir + " error: " + err.Error())
		}
	}
}

// ChecklogoFile check a file of logo
func ChecklogoFile() error {
	var err error
	if _, err := os.Stat(logoFile); os.IsNotExist(err) {
		l, er := DecodeLogo(YANDEXLOGO)
		if er != nil {
			log.Println("Error decode logo")
		}
		SaveLogo(logoFile, l)
	}
	return err
}

// CheckToken check a token for a request to api of Yandex
func CheckToken() error {
	_, err := os.Stat(yandexToken)
	if err != nil {
		requestToken()
	}
	return err
}

// request a token from a user
func requestToken() {
	var p string
	var err error
	for p == "" {
		p, err = gozenity.Password("Yandex API token:")
		if err != nil {
			return
		}
	}

	b, _ := regexp.MatchString("^trnsl.*", p)
	if b {
		ioutil.WriteFile(yandexToken, []byte(p), 0744)
	} else {
		gozenity.Error("Token isn't valid!")
		requestToken()
	}
	return
}

func main() {
	CheckApplicationDir()
	ChecklogoFile()
	CheckToken()
	lenghWords := len(os.Args[1:])
	if lenghWords != 0 {
		text := fmt.Sprintf("%s", GetTranslate(os.Args[1:]))
		notify.Notify("app name", "Translate", text, logoFile)
	}
}
