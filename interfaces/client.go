package interfaces


type Client interface {

	Ping(id []byte, message []byte) error

	Store(value []byte) error

	FindNode(id []byte) ([]*Message, error)

	FindValue(id []byte) ([]byte, Message, error)

	Connect() error
	
	Shutdown() error

	Neighbors() []*Message

}





