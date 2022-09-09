module crawler

go 1.15

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/thehappymouse/ccmouse v0.1.9
	golang.org/x/net v0.0.0-20220728030405-41545e8bf201
	golang.org/x/text v0.3.7
	gopkg.in/olivere/elastic.v5 v5.0.86
	rpcc v0.0.0-00010101000000-000000000000
)

//replace crawler_dis => ../crawler_distributed

replace rpcc => ../crawler_distributed/rpc
