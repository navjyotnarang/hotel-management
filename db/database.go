package db

// This can be a placeholder for a persistent database in a real-world scenario
var roomsDB map[int]Room

func init() {
	roomsDB = make(map[int]Room)
}

func GetAllRooms() []Room {
	var roomList []Room
	for _, room := range roomsDB {
		roomList = append(roomList, room)
	}
	return roomList
}

func GetRoomByID(id int) (Room, bool) {
	room, exists := roomsDB[id]
	return room, exists
}

func CreateRoom(newRoom Room) {
	roomsDB[newRoom.ID] = newRoom
}

func UpdateRoom(updatedRoom Room) bool {
	_, exists := roomsDB[updatedRoom.ID]
	if exists {
		roomsDB[updatedRoom.ID] = updatedRoom
	}
	return exists
}

func DeleteRoom(id int) bool {
	_, exists := roomsDB[id]
	if exists {
		delete(roomsDB, id)
	}
	return exists
}
