# 
## assignment

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

## Requirements

grpc based service, implements:
- stub method that simulates a long operation, the timeout is fixed
   multiple methods based on the received list of employees (name, yob, salary)
- returns a list sorted by the specified field
- return the oldest, most expensive
- average salary, median
   implement unit tests for service methods



## Protoc gen command

```protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 mgidEmployee.proto```

## How to run grpc server
Run example:
```sh
go run cmd/server/main.go -grpc-port=9090
```

## How to run grpc client
Run example:
```sh
go run cmd/client-grpc/main.go -server=localhost:9090
```

## How to run unit test of mgid service
Go to pkg/service/v1
```sh
cd pkg/service/v1/
```
and
```sh
go test -v
```






## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

[dill]: <https://github.com/joemccann/dillinger>
[git-repo-url]: <https://github.com/joemccann/dillinger.git>


[PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
[PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
[PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
[PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
[PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
[PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
