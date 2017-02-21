#go run client-arg.go gps.go -runcommand gpsloc 3cd59439-98cf-4548-ba54-cb263f965dc7

#go run client-arg.go gps.go -runcommand=runlogin mindy haha
# now the expred time has been changed to infinity

#curl -v -X GET -H "clientuserid:b52c8eed-aefb-45dc-a807-a1f7c3bee78a" -H "Auth-Token:63ab9f90-d694-4c32-a847-fd89794be1c4" http://localhost:8080/getgpsloc
#curl -v -X GET -H "username:mindy" -H "password:haha" -H "Auth-Token:63ab9f90-d694-4c32-a847-fd89794be1c4" http://localhost:8080/getuserid
curl -X POST -H "Auth-Token:63ab9f90-d694-4c32-a847-fd89794be1c4" -H "lati:23.333" -H "long:123.7335" http://localhost:8080/setgpsloc


