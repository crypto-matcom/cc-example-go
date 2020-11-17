# Implementación de un Código de Cadena utilizando Go.

El **Código de Cadena** o **(Chaincode)** es un programa que corre sobre una infraestructura de [Hyperledger Fabric]() que provee la lógica de las transacciones hechas en la Blockchain.

Hyperledger Fabric da la opción de implementar los **Chaincode** en varios lenguajes de programación entre los que se encuentra [Go]().

Para implementar un **Chaincode** utilizando Go debe proveerse una estructura que implemente la interfaz `Chaincode` que provee Hyperledger en el `package "github.com/hyperledger/fabric-chaincode-go/shim"` en la cual la estructura debe proveer la implementación de dos metodos: `Init` e `Invoke` de la siguiente forma:

```go
package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type NewChainCode struct {}

func (s *NewChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	panic("implement me")
}

func (s *NewChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	panic("implement me")
}
// ...
``` 

El método `Init` será invocado solo en la primera ejecución del **Chaincode**. En el mismo debe hacerce la inicializacion de las estructuras y recursos que seran utilizados. Tener en cuenta que en Hyperledger Fabric cuando se actualiza un **Chaincode** el método `Init` vuelve a ser ejecutado.

El método `Invoke` será utilizado para realizar actualizaciones o ejecutar consultas sobre la cadena de bloques. Se debe tener en cuenta que las actualizaciones no se reflejarán hasta que la transacción no sea registrada en la cadena de bloques.

Ambos metodos reciben como único parámetro una estructura `stub` que implementa la interfaz `ChaincodeStubInterface` del mismo `package` de `Chaincore`. Dicha estructura contiene los metodos que proveen la interacción entre el código de cadena y la cadena de bloques. Entre los métodos más importantes se encuentran los siguentes:

- `GetArgs()` que retorna los argumentos contenidos en la transacción con los cuales será invocada la función. Usualmente contienen el formato: `["func to invoke", "arg0", "arg1", ...]`
- `GetState(key string)` que retorna la data almacenada en la blockchain bajo una llave específica. Notar que este método no retorna los datos de transacciones que no hayan sido registradas en la cadena. Dicha data, por motivos de eficiencia, es extraida del `StateDatabase`.
- `PutState(key string, value []byte)` que almacena nueva data bajo una llave dada en el `StateDatabase`.
- `GetCreator()` que retorna la identidad del agente o usuario que propone la transacción.

Para mas información revisar el [código de ambas interfaces](https://github.com/hyperledger/fabric-chaincode-go/blob/master/shim/interfaces.go)
