package encryptor

import (
	com "digilounge/infrastructure/functions"
	"fmt"
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/argon2"
)

func PasswordGenerator(password string) (encryptedPass string, e error) {
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(PASSGENERATOR: 1000) %s", err))
		return "", err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(PASSGENERATOR:1001) %s", er))
		return "", err
	}

	encrypted := argon2.IDKey([]byte(password), []byte(os.Getenv("salt")), 2, 64*1024, 8, 32)

	return string(encrypted), nil
}

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

func FieldGenerator(field string) (result string, e error) {
	key := GetStaticKey()
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(GENERATOR: 1000) %s", err))
		return "", err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(GENERATOR:1001) %s", er))
		return "", er
	}

	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())

	token.SetString(os.Getenv("secretKey"), field)

	encrypted := token.V4Encrypt(key, nil)

	return encrypted, nil
}

func VerifyField(encryptedToken, field string) (message string, e error) {
	key := GetStaticKey()
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1000) %s", err))
		return "", err
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY:1001) %s", er))
		return "", er
	}

	parser := paseto.NewParserWithoutExpiryCheck()

	token, er2 := parser.ParseV4Local(key, encryptedToken, nil)

	if er2 != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1003) %s", er2))
		return "", er2
	}

	json, err := token.GetString(os.Getenv("secretKey"))
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1004) %s\n", err))
		return "", err
	}

	return json, nil
}

func GetStaticKey() paseto.V4SymmetricKey {
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(GETSTATICKEY: 1000) %s", err))
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(GETSTATICKEY:1001) %s", er))
	}

	//HEX should be in 32-bit length

	hexKey := os.Getenv("hex")

	// Create the V4SymmetricKey from the fixed hex string
	key, ers := paseto.V4SymmetricKeyFromHex(hexKey)
	if ers != nil {
		// Handle the error if the hex string is invalid
		com.PrintLog(fmt.Sprintf("(GETSTATICKEY:1002) %s", er))
		return paseto.NewV4SymmetricKey() // Fallback to a random key
	}

	return key
}
