package gen

import (
	"crypto/rand"
	"strconv"
	"strings"
	"swiss/utils"
)

const hexCharacters = "0123456789abcdef"

func stringGenerator(generatedLength int) string {
	var builder strings.Builder
	builder.Grow(generatedLength)

	for i := 0; i < generatedLength; i++ {
		var randomByte [1]byte
		_, err := rand.Read(randomByte[:])
		if err != nil {
			panic("crypto/rand failed")
		}
		randomByteIndex := randomByte[0] % byte(len(hexCharacters))
		builder.WriteByte(hexCharacters[randomByteIndex])
	}

	return builder.String()
}

func GenerateUUID() {
	var builder strings.Builder
	builder.Grow(36)
	for i := 0; i < 32; i++ {
		var randomByte [1]byte
		_, err := rand.Read(randomByte[:])
		if err != nil {
			panic("crypto/rand failed")
		}
		b := randomByte[0]
		var twoChars [2]byte
		twoChars[0] = '0' + (b >> 4)
		if twoChars[0] > '9' {
			twoChars[0] += 'a' - '9' - 1
		}
		twoChars[1] = '0' + (b & 0x0F)
		if twoChars[1] > '9' {
			twoChars[1] += 'a' - '9' - 1
		}
		builder.Write(twoChars[:])
	}
	hexBytes := builder.String()
	uuid := hexBytes[:8] + "-" + hexBytes[8:12] + "-" + hexBytes[12:16] + "-" + hexBytes[16:20] + "-" + hexBytes[20:32]
	utils.Success("UUID generated.")
	utils.Output(uuid)
}

func GenerateSecret() {
	var secretLength string
	if len(utils.AdditionalArguments) < 1 {
		secretLength = "16"
	} else {
		secretLength = utils.AdditionalArguments[0]
	}

	intLength, err := strconv.Atoi(secretLength)
	utils.CrashCheck(err)

	secretCode := stringGenerator(intLength)
	utils.Success("Secret code generated.")
	utils.Output(secretCode)
}
