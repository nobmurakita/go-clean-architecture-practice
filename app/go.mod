module go-clean-architecture-practice

go 1.22.4

replace go-clean-architecture-practice/go-pkg => ../pkg

require (
	github.com/google/go-cmp v0.6.0
	go-clean-architecture-practice/go-pkg v0.0.0-00010101000000-000000000000
)

require github.com/oklog/ulid/v2 v2.1.0 // indirect
