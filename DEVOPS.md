# DevOps

## Commit Lint

To lint commit messages in this repository was used [commitlint](https://github.com/conventionalcommit/commitlint).

The binary [commitlint-0.10.1](commitlint-0.10.1) is used locally by git hooks.

The github pipeline check messages either.

Configs for commit lint can be found in file [.commitlint.yml](.commitlint.yml).

## Makefile

About the repository configuration there is a few commands available in [Makefile](Makefile).

### coverage

Check the test code covered, it generates file coverage.out.

### coverage-inspect-html

Generate an HTML file to programmers easily open and view lines covered or not.

### coverage-inspect-text

Used by CI to validate the amount of code covered.

### format

Formats all *.go files found in this repository.

### lint && lint-fix

Execute golangci-lint-2.0.2 binary, it was took from [repository](https://github.com/golangci/golangci-lint).

Configuration reference can be found [here](https://github.com/golangci/golangci-lint/blob/main/.golangci.reference.yml).

Configs for this repository are in [.golangci.yml](.golangci.yml).

#### lint false positives

If you really believe lint is returning a false positive use a commentary informing wich lint should be ignore in the desired line.

sample:

```go

func (l *LamportSignature) genPrivateKey() {
	for i := 0; i < l.algorithmSize*2; i++ {
		number := make([]byte, l.privateKeySize/8)
		rand.Read(number) // nolint: gosec
		l.privateKey = append(l.privateKey, number)
	}
}
```

In this case gosec was complaining about a possible error treatment, however reading rand.Read docs you may find the
following comment.

```go
// Read fills b with cryptographically secure random bytes. It never returns an
// error, and always fills b entirely.
//
// Read calls [io.ReadFull] on [Reader] and crashes the program irrecoverably if
// an error is returned. The default Reader uses operating system APIs that are
// documented to never return an error on all but legacy Linux systems.
func Read(b []byte) (n int, err error) {
```

Therefore disable lint for this specific line is reasoably.
