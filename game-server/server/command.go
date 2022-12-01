package server

// Codes of commands that are sent to clients
const (
	CommandToClientInitPlayer       uint8 = 0
	CommandToClientInitMap          uint8 = 1
	CommandToClientPlayersPositions uint8 = 2
	CommandToClientBulletsPositions uint8 = 3
)

// Codes of commands that are received from clients
const (
	CommandFromClientUpdateInput uint8 = 100
)
