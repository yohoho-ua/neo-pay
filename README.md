# NEO-PAY tracker
Just some code to fetch transactions which have used desired (our deposit)address
Check if payment was made

![front](https://i.imgur.com/lqy0wOa.png "Kartinochka")  
![console log](https://i.imgur.com/wqAKLIe.png "Kartinochka2")

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
**gorilla/mux**
```
go get -u github.com/gorilla/mux
```

run:
```
go run main.go customer.go checkPayment.go neo_sdk_wrapper_util.go
```
