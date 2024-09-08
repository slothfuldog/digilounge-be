package encryptor

import (
	com "digilounge/infrastructure/functions"
	"encoding/base64"
	"encoding/json"
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

	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)

	return string(encryptedBase64), nil
}

func FieldGenerator(customData map[string]interface{}) (result string, e error) {
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
	token.SetExpiration(time.Now().Add(24 * time.Hour))

	customJson, ers := json.Marshal(customData)

	if ers != nil {
		com.PrintLog(fmt.Sprintf("(GENERATOR:1002) %s", er))
		return "", fmt.Errorf("error marshaling custom data: %v", ers)
	}

	token.SetString(os.Getenv("secretKey"), string(customJson))

	encrypted := token.V4Encrypt(key, nil)

	return encrypted, nil
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
