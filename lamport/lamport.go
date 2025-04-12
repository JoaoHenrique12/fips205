package lamport

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

type signarute struct {
	PickedNumbers [][]byte
	Publickey     [][]byte
}

type LamportSignature struct {
	hashAlgorithm hash.Hash
	algorithmSize int

	privateKey [][]byte

	Message          string
	MessageSignature signarute

	// use 4 to generate numbers of 32 bits
	// use 8 to generate numbers of 64 bits
	privateKeySize int
}

func (l *LamportSignature) genPrivateKey() {
	for i := 0; i < l.algorithmSize*2; i++ {
		number := make([]byte, l.privateKeySize)
		rand.Read(number) // nolint: gosec
		l.privateKey = append(l.privateKey, number)
	}
}

func (l *LamportSignature) genPublicKey() {
	for _, v := range l.privateKey {
		l.hashAlgorithm.Write(v)
		hash := l.hashAlgorithm.Sum(nil)
		l.hashAlgorithm.Reset()
		l.MessageSignature.Publickey = append(l.MessageSignature.Publickey, hash)
	}
}

func (l *LamportSignature) SignMessage(message string) {
	l.Message = message
	l.hashAlgorithm.Write([]byte(message))
	message_hash := l.hashAlgorithm.Sum(nil)
	l.hashAlgorithm.Reset()

	pair_indice := 0
	// This reverse ensure the picked numbers are reffering from LSB
	// to MSB in hash output order.
	for i := len(message_hash) - 1; i >= 0; i-- {
		byte := message_hash[i]
		for j := 0; j < 8; j++ {
			chose_first_number := byte & 1
			if chose_first_number == 0 {
				l.MessageSignature.PickedNumbers = append(l.MessageSignature.PickedNumbers, l.privateKey[pair_indice])
			}else {
				l.MessageSignature.PickedNumbers = append(l.MessageSignature.PickedNumbers, l.privateKey[pair_indice + 1])
			}

			byte >>= 1
			pair_indice+=2
		}
	}
}

func (l *LamportSignature) ValidateSignature(message string) {

}

func (l *LamportSignature) selectHashAlgorithm(hashAlgorithmName string) {
	switch hashAlgorithmName {
	case "SHA256":
		l.hashAlgorithm = sha256.New()
		l.algorithmSize = 256
	case "SHA512":
		l.hashAlgorithm = sha512.New()
		l.algorithmSize = 512
	}
}

// hashAlgorithmName available options: SHA256; SHA512
//
// privateKeySize assume 4,8 to generate 32, 64 bits numbers
func LamportBuilder(hashAlgorithmName string, privateKeySize int) LamportSignature {
	lamport := LamportSignature{}

	if !(privateKeySize == 4 || privateKeySize == 8) {
		panic("privatekeysize must be eq to 4 or 8 to use numbers of 32 or 64 bits")
	}
	lamport.privateKeySize = privateKeySize

	(&lamport).selectHashAlgorithm(hashAlgorithmName)
	(&lamport).genPrivateKey()
	(&lamport).genPublicKey()

	return lamport
}

// hashAlgorithmName available options: SHA256; SHA512
//
// privateKey must be a matrix with arrays of size 4 or 8 to represent numbers with 32 or 64 bits.
// len(privateKey) must be eq 2 * size of output hashAlgorithmName.
func LamportBuilderInformingPrivateKey(hashAlgorithmName string, privateKey [][]byte) LamportSignature {
	lamport := LamportSignature{}

	(&lamport).selectHashAlgorithm(hashAlgorithmName)
	if len(privateKey) != lamport.algorithmSize*2 {
		panic(fmt.Sprintf("invalid number of bytes given for privateKey, for %v you must a private key with %v numbers", hashAlgorithmName, lamport.algorithmSize*2))
	}

	lamport.privateKey = privateKey

	lamport.privateKeySize = len(privateKey[0])

	privateKeySize := lamport.privateKeySize
	if !(privateKeySize == 4 || privateKeySize == 8) {
		panic("private key size must be eq to 4 or 8 to use numbers with 32 or 64 bits")
	}

	lamport.genPublicKey()

	return lamport
}
