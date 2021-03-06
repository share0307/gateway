package contracts

//Server is comet server interface
type Server interface {
	Run()
	SendToConnections(connections []string, msg string) ([]string, error)
	Broadcast(msg string)
	KickConnections(connections []string)
	KickAllConnections()
	CheckConnectionsOnline(connections []string) []string
	GetAllConnections() []string
	JoinCluster()
}
