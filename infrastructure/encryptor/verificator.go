package encryptor

import (
	com "digilounge/infrastructure/functions"
	"encoding/json"
	"fmt"
	"os"

	"aidanwoods.dev/go-paseto"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/argon2"
)

func VerifyPassword(encrpytedPass, password string) (isMatch bool, e error) {
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFYPASS: 1000) %s", err))
		return false, err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(VERIFYPASS:1001) %s", er))
		return false, er
	}

	encrypted := argon2.IDKey([]byte(password), []byte(os.Getenv("salt")), 2, 64*1024, 8, 32)

	if encrpytedPass != string(encrypted) {
		com.PrintLog("(VERIFYPASS:1002) PASSWORD IS NOT MATCH")
		return false, fmt.Errorf("PASSWORD IS NOT MATCH")
	}

	return true, nil
}

func VerifyField(encryptedToken string) (payloads map[string]interface{}, e error) {
	var payload map[string]interface{}
	key := GetStaticKey()
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1000) %s", err))
		return nil, err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY:1001) %s", er))
		return nil, er
	}

	parser := paseto.NewParserWithoutExpiryCheck()

	token, er2 := parser.ParseV4Local(key, encryptedToken, nil)

	if er2 != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1003) %s", er2))
		return nil, er2
	}

	encrypted, err := token.GetString(os.Getenv("secretKey"))
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1004) %s\n", err))
		return nil, err
	}

	if errr := json.Unmarshal([]byte(encrypted), &payload); errr != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY:1005) %s", errr))
		return nil, fmt.Errorf("Error unmarshaling payload: %v", err)
	}

	return payload, nil
}
