package dal

type Heap struct {
	data map[int]User
}

func NewHeap(users []User) *Heap {
	dataMap := make(map[int]User)
	for _, user := range users {
		dataMap[user.ID] = user
	}
	return &Heap{
		data: dataMap,
	}
}

func (h *Heap) Iterate(f func(*User)) {
	for _, user := range h.data {
		f(&user)
	}
}

func (h *Heap) Get(ctid int) User {
	return h.data[ctid]
}

func (h *Heap) GetPtr(ctid int) *User {
	user := h.data[ctid]
	return &user
}
