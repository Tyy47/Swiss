package gen

import (
	"crypto/rand"
	"os"
	"strconv"
	"strings"

	"github.com/Tyy47/clibox/argbin"
	"github.com/Tyy47/clibox/outbin"
)

const hexCharacters = "0123456789abcdef"

var output = outbin.NewOutput(os.Stdout, os.Stderr)

// stringGenerator builds a randomized string using hex characters and random bytes.
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

// GenerateCommand creates the "gen" command for swiss
//
// Command is stored in main.go
func GenerateCommand() *argbin.Command {
	return &argbin.Command{
		Name: "gen",
		Description: "generator command for swiss.",
		Subcommands: []*argbin.Command{
			generateSecretCommand(),
			generateUUIDCommand(),
		},
	}
}

// generateSecretCommand creates the "secret" subcommand for "gen".
func generateSecretCommand() *argbin.Command {
	return &argbin.Command{
		Name: "secret",
		Description: "secret code generator.",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			var secretLength string = ctx.ParsedValue

			intLength, err := strconv.Atoi(secretLength)
			if err != nil {
				return err
			}

			if intLength == 0 {
				intLength = 16
			}

			secretCode := stringGenerator(intLength)
			
			output.Successf("Secret code generated: %s", secretCode)
			return nil
		},
	}

}

func generateUUIDCommand() *argbin.Command {
	return &argbin.Command{
		Name: "uuid",
		Description: "uuid generator for swiss",
		Execute: func(ctx *argbin.Context) error {
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
			output.Successf("UUID generated: %s", uuid)
			return nil
		},
	}
}
