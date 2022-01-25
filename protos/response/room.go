package response

type Room struct {
}

type ListRooms struct {
	PageResponse
	Rooms []Room
}
