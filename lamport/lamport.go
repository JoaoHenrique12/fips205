package lamport

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

type Signarute struct {
	PickedNumbers [][]byte
	Publickey     [][]byte
}

type LamportSignature struct {
	hashAlgorithm hash.Hash
	algorithmSize int

	privateKey [][]byte

	Message          string
	MessageSignature Signarute

	// 32 or 64 to generate numbers with this amount of bits
	privateKeySize int
}

func (l *LamportSignature) genPrivateKey() {
	for i := 0; i < l.algorithmSize*2; i++ {
		number := make([]byte, l.privateKeySize/8)
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

func (l *LamportSignature) SignMessage(message string) Signarute {
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

	return l.MessageSignature
}

func compareBytes(a, b []byte) bool{
	if len(a) != len(b) {
		panic("could not compare []bytes with different sizes")
	}

	for i := 0; i< len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func ValidateSignature(message string, hash_algorithm hash.Hash, signarute Signarute) bool {
	is_valid := true
	hash_algorithm.Write([]byte(message))
	message_hash := hash_algorithm.Sum(nil)
	hash_algorithm.Reset()

	pair_indice := 0
	pick_idx := 0
	// This reverse ensure the picked numbers are reffering from LSB
	// to MSB in hash output order.
	for i := len(message_hash) - 1; i >= 0; i-- {
		byte := message_hash[i]
		for j := 0; j < 8; j++ {
			chose_first_number := byte & 1

			hash_algorithm.Write(signarute.PickedNumbers[pick_idx])
			hash_number_found := hash_algorithm.Sum(nil)
			hash_algorithm.Reset()

			if chose_first_number == 0 {
				is_valid = compareBytes(hash_number_found, signarute.Publickey[pair_indice])
			}else {
				is_valid = compareBytes(hash_number_found, signarute.Publickey[pair_indice + 1])
			}

			if ! is_valid {
				return is_valid
			}

			byte >>= 1
			pair_indice+=2
			pick_idx++
		}
	}


	return is_valid
}

func selectHashAlgorithm(hashAlgorithmName string) (hash.Hash, int) {
	var hashAlgorithm hash.Hash
	var hashAlgorithmSize int

	switch hashAlgorithmName {
	case "SHA256":
		hashAlgorithm = sha256.New()
		hashAlgorithmSize = 256
	case "SHA512":
		hashAlgorithm = sha512.New()
		hashAlgorithmSize = 512
	}
	return hashAlgorithm, hashAlgorithmSize
}

// hashAlgorithmName available options: SHA256; SHA512
//
// privateKeySize assume 32 or 64 to indicate the number of bits
func LamportBuilder(hashAlgorithmName string, privateKeySize int) LamportSignature {
	lamport := LamportSignature{}

	if !(privateKeySize == 32 || privateKeySize == 64) {
		panic("privateKeySize must be eq to 32 or 64")
	}
	lamport.privateKeySize = privateKeySize

	lamport.hashAlgorithm, lamport.algorithmSize = selectHashAlgorithm(hashAlgorithmName)

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

	lamport.hashAlgorithm, lamport.algorithmSize = selectHashAlgorithm(hashAlgorithmName)
	if len(privateKey) != lamport.algorithmSize*2 {
		panic(fmt.Sprintf("invalid number of bytes given for privateKey, for %v you must a private key with %v numbers", hashAlgorithmName, lamport.algorithmSize*2))
	}

	lamport.privateKey = privateKey

	lamport.privateKeySize = len(privateKey[0]) * 8

	privateKeySize := lamport.privateKeySize
	if !(privateKeySize == 32 || privateKeySize == 64) {
		panic("private key size must be eq 32 or 64")
	}

	lamport.genPublicKey()

	return lamport
}
