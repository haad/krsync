module github.com/haad/krsync

go 1.18

require (
	github.com/spf13/cobra v1.5.0
	github.com/zloylos/grsync v1.6.4
	go.uber.org/zap v1.21.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

replace github.com/zloylos/grsync => github.com/haad/grsync v1.6.4
