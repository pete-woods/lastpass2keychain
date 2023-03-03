module github.com/pete-woods/lastpass2keychain

go 1.20

require (
	github.com/keybase/go-keychain v0.0.0-20221221221913-9be78f6c498b
	gotest.tools/v3 v3.4.0
)

require github.com/google/go-cmp v0.5.5 // indirect

replace github.com/keybase/go-keychain v0.0.0-20221221221913-9be78f6c498b => github.com/pete-woods/go-keychain v0.0.0-20230303113558-9ff426b86657
