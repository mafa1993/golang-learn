module crawler_dis

go 1.15

require (
	crawler v0.0.0-00010101000000-000000000000
	gopkg.in/olivere/elastic.v5 v5.0.86
	rpcc v0.0.0-00010101000000-000000000000
)

replace crawler => ../crawler

replace rpcc => ./rpc
