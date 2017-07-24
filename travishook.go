package travishook

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var travisciPubKey *rsa.PublicKey
var loadPubKeyOnce sync.Once
var defaultServer = "api.travis-ci.org"

func CheckSignature(rawSignature string, message []byte, server string) (err error) {

	/* base64-decode Signature HTTP Header */
	signature, err := base64.StdEncoding.DecodeString(rawSignature)
	if err != nil {
		return err
	}

	/* obtain public key from Travis CI's servers once */
	loadPubKeyOnce.Do(func() {
		travisciPubKey, err = GetPubKey(defaultServer)
		if err != nil {
			log.Fatalf("Failed to obtain TravisCI Public Key from %s: %v", defaultServer, err)
		}
	})
	if travisciPubKey == nil {
		return err
	}

	/* SHA1-hash the payload */
	hash := sha1.Sum(message)

	/* verify signature with public key and SHA1-hash */
	err = rsa.VerifyPKCS1v15(travisciPubKey, crypto.SHA1, hash[:], signature)
	if err != nil {
		return err
	}

	return nil
}

func GetPubKey(server string) (*rsa.PublicKey, error) {
	var rawPubKey string
	var data APIConfig

	/* request (json-encoded) config from API server */
	response, err := http.Get("https://" + server + "/config")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	/* read body */
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	/* decode json data */
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	/* extract public key string */
	rawPubKey = data.Config.Notifications.Webhook.PublicKey

	/* block encode string */
	block, _ := pem.Decode([]byte(rawPubKey))
	if block == nil {
		return nil, err
	}

	/* translate into Go rsaKey interface */
	rsaKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	/* retrieve public key */
	pubKey := rsaKey.(*rsa.PublicKey)

	return pubKey, nil
}
