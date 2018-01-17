# NEO-PAY tracker
Just some code to fetch transactions which have used desired (our deposit)address
Check if payment was made

![alt text](https://i.imgur.com/wqAKLIe.png "Kartinochka")

Running rpc client with open wallet required. Use neo-cli
```dotnet neo-cli.dll /rpc
   open wallet <path_to_wallet>
```  
enter your pass

dependencies: 
**neo-go-sdk**
```
go get github.com/CityOfZion/neo-go-sdk
```
run:
```
go run main.go rpcClieng.go
```
