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
```
func GetElement(id uint32) Element
{
	elem, exist := cache.Get(id)
	// exist mais nil = introuvable
	
}
```

## Stream endpoint
`func listen(endpoint string, port int32, baseDir string, tls bool) func(httpTruc)??`
## Serve files
`func listen(endpoint string, port int32, baseDir string, tls bool)`

## TLS handler
gen avec openssl internally, ask path de destination des certifs.

## Config
json.
Soit dans .config/appName/config.json
soit dans répertoire courante de l'exe.
Il faut pouvoir generer un config file ap artir de rien. 

param d'init interatif ou pas
interactif: repertoire ou se trouve les fichiers, puis videos,
generer config file local ou dans home.
Ports d'écoute de chaque systeme
ssl ou pas. Si oui path des certificats.

ou pas: 
- certif: ./cert
- fichiers divers: ./static
- videos: ./videos
- ports:
	- static: 8080
	- videos: 8081
- bdd: db
- symkey: ???
- init structure DB file: creation.sql
- init value DB file: default-values.sql
	
