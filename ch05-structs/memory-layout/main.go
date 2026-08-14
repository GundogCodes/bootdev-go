package main

/*
In Go, structs sit in memory in a contiguous block, 
with fields placed one after another as defined in the struct.
For example this struct:

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

*/

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
}

