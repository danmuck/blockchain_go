package interfaces


type RoutingTable interface {

	K() int

	InsertNode(node *Message)

	RemoveNode(key []byte) error

	Lookup(key []byte) (node *Message, ok bool)

	GetNodes(bucket int) []*Message

	ClosestK(key []byte) []*Message

	Buckets() int
}