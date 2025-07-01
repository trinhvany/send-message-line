package redis

import (
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Danh sách mặc định
var defaultUsers = []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 3, Name: "Charlie"},
}
