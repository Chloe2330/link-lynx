package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"os"
	"math/big"
)

// Returns the hash value of a string input as a byte slice using SHA-256 hash function
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	
	// prepares the input data for calculation by converting string to byte slice and 
	// updating the state of the hash algorithm
	algorithm.Write([]byte(input))

	// performs actual calculation and returns hash value for future use
	return algorithm.Sum(nil)
}

// Returns Base58 encoded value as a string
func base58Encoded(bytes []byte) string {

	// specify method of encoding (Bitcoin variant of Base58)
	encoding := base58.BitcoinEncoding

	// encode byte slice
	encoded, err := encoding.Encode(bytes)

	// cannot encode
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// convert from byte slice to human readable string
	return string(encoded)
}

// Note: hashing vs encoding
// Hash functions are one-way functions, so it is impossible to reverse the function
// and retrieve the original value given the hash value. Encoding is a reversible process, 
// as encoded data can be decoded back to its original format. 

// In this program, the SHA-256 hash function serves a security purpose. The inclusion 
// of an user ID ensures that the output space is much larger than the input space and
// two different users that try to convert the same URL will not receive the same short
// URL output (no collisions). 

// The Base58 encode serves to turn the binary code into a human-readable string that is 
// more user-friendly. Specifically, Base58 omits confusing characters like 0, O, 1, I, l
// and punctuation.

// Therefore, the byte slice returned from sha256Of() cannot be reversed back into its 
// original string format with the initial URL and user ID, but the string returned from
// base58Encoded() can be decoded back to its byte slice format. 

// Generates the short URL by calling previous two functions and through truncation
func GenerateShortLink(initialLink string, userId string) string {

	// returns hash value (byte slice) from concatenation of initial URL, user ID
	urlHashBytes := sha256Of(initialLink + userId)

	// converts byte slice to *big.Int type and obtain unsigned integer value (Uint64()) 
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// converts number to string and back to byte slice, Base58 encodes byte slice and 
	// returns back as string
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	// truncate and return first eight characters of string
	return finalString[:8]
}

// Why can't I directly pass urlHashBytes into the base58 function? 
// current conversions: []byte -> *big.Int -> uint64 -> string -> []byte -> func returns string

// urlHashBytes can not be passed directly into the Base58 encoding function because it 
// is not designed to handle hash values. 

// Byte slice returned from SHA-256 algorithm is 32 bytes, so direct conversion
// to uint64 type (only 8 bytes) would result in data loss/rounding imprecisions, so 
// *big.Int provides arbitrary precision integers and allows for safe conversion from 
// a byte slice to a fixed-size integer type. 

// uint64 to byte slice conversions are possible, but it's easier and more succint to 
// first convert the fixed-size integer to a string, and then convert the string to a byte
// slice with the inbuilt []byte() function
