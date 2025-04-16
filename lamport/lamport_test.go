package lamport

import (
	"testing"
	"fmt"
)

func ExampleLamportSignature_SignMessage_validatesignature() {
	privateKeySize := 32
	hashAlgorithm := "SHA256"
	message := []byte("This is a test message for Lamport signature.")

	// Build the Lamport signature scheme
	l := LamportBuilder(hashAlgorithm, privateKeySize)

	// Sign the message
	signature := l.SignMessage(message)

	// Verify the signature
	isValid := ValidateSignature(message, hashAlgorithm, signature)

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Hash Algorithm: %s\n", hashAlgorithm)
	fmt.Printf("Private Key Size: %d\n", privateKeySize)
	fmt.Printf("Signature is valid: %t\n", isValid)

	// Output:
	// Message: This is a test message for Lamport signature.
	// Hash Algorithm: SHA256
	// Private Key Size: 32
	// Signature is valid: true
}


type validSignatureInput struct {
  privateKeySize int
  hashAlgorithm string
  message string
}

func TestSignMessageValidateSuccess(t *testing.T)  {
  parametrize := [] validSignatureInput{
    {32, "SHA256", "Hello World"},
    {64, "SHA256", "Hello World"},

    {32, "SHA512", "Hello World"},
    {64, "SHA512", "Hello World"},
  }
  for _, input := range parametrize {
    privateKeySize := input.privateKeySize
    hashAlgorithm := input.hashAlgorithm
    message := []byte(input.message)

    l := LamportBuilder(hashAlgorithm, privateKeySize)

    signature := l.SignMessage(message)
    if !ValidateSignature(message, hashAlgorithm, signature) {
      t.Errorf("Lamport signature fails, sign a message and validate it returned false!")
    }
  }
}

type invalidSignatureInput struct {
  privateKeySize int
  hashAlgorithm string
  message string
  messageTained string
}

func TestSignMessageValidateFail(t *testing.T)  {
  parametrize := [] invalidSignatureInput{
    {32, "SHA256", "Hello World", "Hello World "},
    {64, "SHA256", "Hello World", "Hello World "},

    {32, "SHA512", "Hello World", "Hello World "},
    {64, "SHA512", "Hello World", "Hello World "},
  }
  for _, input := range parametrize {
    privateKeySize := input.privateKeySize
    hashAlgorithm := input.hashAlgorithm
    message := []byte(input.message)
    messageTained := []byte(input.messageTained)

    l := LamportBuilder(hashAlgorithm, privateKeySize)

    signature := l.SignMessage(message)
    if ValidateSignature(messageTained, hashAlgorithm, signature) {
      t.Errorf("Different messages got a valid signature")
    }
  }
}

type panicLamportBuildInput struct {
	hashAlgorithmName string
	privateKeySize    int
	expectedPanic     bool
}

func TestLamportBuilderPanicsOnInvalidPrivateKeySize(t *testing.T) {
	parametrize := []panicLamportBuildInput {
		{"SHA256", 16, true},
		{"SHA512", 128, true},
		{"SHA256", 32, false},
		{"SHA512", 64, false},
	}

	for _, input := range parametrize {
		panicked := recoverFunc(func() {
			LamportBuilder(input.hashAlgorithmName, input.privateKeySize)
		})

		if input.expectedPanic && !panicked {
			t.Errorf("LamportBuilder(%q, %d) should have panicked but did not", input.hashAlgorithmName, input.privateKeySize)
		}

		if !input.expectedPanic && panicked {
			t.Errorf("LamportBuilder(%q, %d) should not have panicked but did", input.hashAlgorithmName, input.privateKeySize)
		}
	}
}

func recoverFunc(f func()) (panicked bool) {
  // golang initializated panicked with 'zero' false.
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	f()
	// return of recoverFunc is named, GoLang by his self found the var and returned it.
	return
}
