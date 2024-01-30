module test.com/hello

go 1.21.6

replace test.com/hello/greetings => ../greetings

replace test.com/greetings => ../greetings

require test.com/greetings v0.0.0-00010101000000-000000000000
