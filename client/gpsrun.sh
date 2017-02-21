#go run client-arg.go gps.go -runcommand gpsloc 3cd59439-98cf-4548-ba54-cb263f965dc7
#go run client-arg.go gps.go -runcommand=runlogin kala kala
# now the expred time has been changed to infinity
#curl -v -X GET -H "clientuserid:kala" -H "Auth-Token:3cd59439-98cf-4548-ba54-cb263f965dc7" http://localhost:8080/getgpsloc
curl -X POST -H "Auth-Token:3cd59439-98cf-4548-ba54-cb263f965dc7" -H "lati:23.333" -H "long:123.3335" http://localhost:8080/setgpsloc
