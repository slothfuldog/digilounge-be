package encryptor

import (
	com "digilounge/infrastructure/functions"
	"fmt"
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/joho/godotenv"
)

func PasswordGenerator(password string) string {
	key := GetStaticKey()
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(GENERATOR: 1000) %s", err))
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(GENERATOR:1001) %s", er))
	}

	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())

	token.SetString(os.Getenv("secretKey"), password)

	encrypted := token.V4Encrypt(key, nil)

	return encrypted
}

func VerifyPassword(encryptedToken, password string) bool {
	key := GetStaticKey()
	path, err := os.Getwd()
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1000) %s", err))
		return false
	}
	currDir := fmt.Sprint(path, "/.env")
	er := godotenv.Load(currDir)

	if er != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY:1001) %s", er))
		return false
	}

	parser := paseto.NewParserWithoutExpiryCheck()

	token, er2 := parser.ParseV4Local(key, encryptedToken, nil)

	if er2 != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1003) %s", er2))
		return false
	}

	retrievedPassword, err := token.GetString(os.Getenv("secretKey"))
	if err != nil {
		com.PrintLog(fmt.Sprintf("(VERIFY: 1004) %s\n", err))
		return false
	}

	if retrievedPassword != password {
		fmt.Println("Passwords do not match.")
		return false
	}

	return true
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
