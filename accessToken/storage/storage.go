package storage

//InMemory means store in a variable
const InMemory = 1

//InFile means store in a specified file
const InFile = 2

//InRedis means store in the specified redis with given key
const InRedis = 3

//File ...
type File struct {
	name string
}

//Redis ...
type Redis struct {
	host     string
	port     int
	password string
}
