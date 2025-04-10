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
  Publickey [][] byte
}

type LamportSignature struct {
  hashAlgorithm hash.Hash
  AlgorithmSize int

  PrivateKey [][] byte

  Message string
  MessageSignature signarute

  // use 4 to generate numbers of 32 bits
  // use 8 to generate numbers of 64 bits
  privateKeySize int
}

func (l *LamportSignature) genPrivateKey() {
  for i := 0; i < l.AlgorithmSize * 2; i++ {
    number := make([]byte, l.privateKeySize)
    rand.Read(number)
    l.PrivateKey = append(l.PrivateKey, number)
  }
}

// func (l LamportSignature) GenPublicKey() {
//   for v := range l.PrivateKey {
//     l.hashAlgorithm.Write(v)
//     l.MessageSignature.Publickey = append(l.MessageSignature.Publickey, hash)
//   }
// }

func (l LamportSignature) Signmessage(message string) {

}

func (l LamportSignature) Validatesignature(message string) {

}

func (l* LamportSignature) selectHashAlgorithm (hashAlgorithmName string) {
  switch hashAlgorithmName {
  case "SHA256":
      l.hashAlgorithm = sha256.New()
      l.AlgorithmSize = 256
  case "SHA512":
      l.hashAlgorithm = sha512.New()
      l.AlgorithmSize = 512
  }

}

func LamportBuilder(hashAlgorithmName string, privateKeySize int) LamportSignature {
  lamport := LamportSignature{}

  if !(privateKeySize == 4 || privateKeySize == 8) {
    panic("privatekeysize must be eq to 4 or 8 to use numbers of 32 or 64 bits")
  }
  lamport.privateKeySize = privateKeySize

  (&lamport).selectHashAlgorithm(hashAlgorithmName)
  (&lamport).genPrivateKey()


  return lamport
}

func LamportBuilderInformingPrivateKey(hashAlgorithmName string,  privateKey [][]byte) LamportSignature {
  lamport := LamportSignature{}

  (&lamport).selectHashAlgorithm(hashAlgorithmName)
  if len(privateKey) != lamport.AlgorithmSize * 2 {
    panic(fmt.Sprintf("invalid number of bytes given for privateKey, for %v you must a private key with %v numbers", hashAlgorithmName, lamport.AlgorithmSize * 2))
  }

  lamport.PrivateKey = privateKey

  lamport.privateKeySize = len(privateKey[0])

  privateKeySize := lamport.privateKeySize
  if !(privateKeySize == 4 || privateKeySize == 8) {
    panic("privatekeysize must be eq to 4 or 8 to use numbers with 32 or 64 bits")
  }


  return lamport
}
