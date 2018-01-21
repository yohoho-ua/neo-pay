# NEO-PAY tracker
Application generates payment address and tracks if payment was made
Button will refresh address if needed
![front](https://i.imgur.com/5pRptrl.png "Kartinochka")  

Running rpc client with open wallet required. Use neo-cli
```dotnet neo-cli.dll /rpc
   open wallet <path_to_wallet>
```  
enter your pass

dependencies:   
https://github.com/CityOfZion/neo-go-sdk  
https://github.com/gorilla/mux


prepare:
```
go get
go build
```
run **neo-pay.exe**