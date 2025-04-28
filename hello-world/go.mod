module example/hello

go 1.22.4

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/gookit/color v1.5.4
	rsc.io/quote v1.5.2
)

require (
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace example.com/greetings => ../greetings
