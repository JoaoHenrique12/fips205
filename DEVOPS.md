# DevOps

Use make install-ci-tools to install locally all the explained tools below.

## Commit Lint

[Why use conventional commits ?](https://www.conventionalcommits.org/en/v1.0.0/#why-use-conventional-commits)

To lint commit messages in this repository was used [commitlint written in go](https://github.com/conventionalcommit/commitlint).

```bash
go install github.com/conventionalcommit/commitlint@latest
```

The github pipeline check messages either.

Configs for commit lint can be found in file [.commitlint.yml](.commitlint.yml).

Commands executed to create this hooks in .commitlint
```bash
commitlint init
commitlint hook create
```

This commands updates the following path
```bash
git config --get core.hooksPath
```

Because of it, put your pre-commit/git hooks inside .commitlint/, if you wish update them.

### Why use commitlint written in GoLang

At beginnig was considered to use [commitlint written in JavaScript](https://github.com/conventional-changelog/commitlint), wich is a 
popular project to check commit messages in github/gitlab pipelines. This tool already has a pre-defined [ci config](https://commitlint.js.org/guides/ci-setup.html).
The pre-defined file shows how to validate all commits in a PR, this functionallity is not directed covered by the go commit linter. Despite these advantages,
this project uses the go version with a bash code to ensure all commits in a PR are valid. To see how it's done, check it in [.github/workflows/ci.yaml](.github/workflows/ci.yaml).
The reason for this is that the JS version requires installing npm, npx, and Node; this installation was taking ~1:30min on pipeline, and the actual lint does not have this
amount of time. Furthermore, to validate commits locally, it would require these binaries installed (an unnecessary overhead).

## Gocyclo

```bash
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

*Cyclomatic complexity is a measurement developed by Thomas McCabe to determine the stability and level of confidence in a program.
It measures the number of linearly-independent paths through a program module. Programs with lower Cyclomatic complexity are easier
to understand and less risky to modify.* [IBM reference](https://www.ibm.com/docs/en/raa/6.1.0?topic=metrics-cyclomatic-complexity)

Given that, gocyclo is configured via golangci-lint to check this complexity, and gitooks do not allow commit codes with cyclomatic over 8.

The command make top5-cyclo is available as well.

## Go Semantic Release

[go-semantic-release](https://github.com/go-semantic-release/semantic-release?tab=readme-ov-file) take advantage of conventional commits
to create tags based on [semver](https://semver.org/), furthermore it automaitcly generates a changelog for each release.
Example: [v1.0.0](https://github.com/JoaoHenrique12/fips205/releases/tag/v1.0.0)

Present only on github CI pipeline.

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

This installlation is through binary from github, go install only find versions 1.x, and the config
file defined for this tool uses version 2. Use make install-ci-tools to get this binary.
[golangci-lint official repository](https://github.com/golangci/golangci-lint).

Configuration reference can be found [here](https://github.com/golangci/golangci-lint/blob/main/.golangci.reference.yml).

Configs for this repository are in [.golangci.yml](.golangci.yml).

#### lint false positives

If you really believe the linter is returning a false positive, then use a commentary informing which line the linter
should ignore.

Sample:

```go

func (l *LamportSignature) genPrivateKey() {
	for i := 0; i < l.algorithmSize*2; i++ {
		number := make([]byte, l.privateKeySize/8)
		rand.Read(number) // nolint: gosec
		l.privateKey = append(l.privateKey, number)
	}
}
```

In this case, gosec was complaining about a possible error treatment; however, reading rand.Read docs you may find the
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

Therefore, disable lint for this specific line is reasoably.
