package types

type PortAllocation struct {
	MainPort  int
	DevPort   int
	UserId    string
	Timestamp string
}

type PortAllocationRequest struct {
	usedPorts      map[int]bool
	PortAllocation map[string]PortAllocation
	startPort      int
	endPort        int
}
