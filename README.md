# Summary
Various libraries associate to a main.go in order to try them.
* check branches.
	* '0/<branch name>': reusable context
	* 'wip/':
 

# Libs
## File handler
### List files
`func ListFiles(directory string, filter string) []string`
## Crypto
### symetric with runes or strings
`func applyKey(target string, key string) string`
## Cache
### map
### expiration map
## DB
### CRUD
## Stream endpoint
`func listen(endpoint string, port int32, baseDir string, tls bool) func(httpTruc)??`
## Serve files
`func listen(endpoint string, port int32, baseDir string, tls bool)`
