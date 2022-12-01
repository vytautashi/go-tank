package server

var nextClientID func() int32

func init() {
	nextClientID = generateUniqueID()
}

func generateUniqueID() func() int32 {
	var id int32 = 0
	return func() int32 {
		id++
		return id
	}
}
