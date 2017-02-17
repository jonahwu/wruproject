#go run client-arg.go gps.go -runcommand gpsloc
#go run client-arg.go gps.go -runcommand=runlogin kala kala
# now the expred time has been changed to infinity
curl -X POST -H "Auth-Token:cbcc7327-5222-4d89-a9f8-47ea40d357b3" -H "lati:23.333" -H "long:123.3335" http://localhost:8080/setgpsloc
