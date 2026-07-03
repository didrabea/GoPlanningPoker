package state

import (
	"planning-poker/models"
	"sync"
)

// in memory DB and lock system
var (
	Rooms = map[string]*models.Room{}
	Mu    sync.Mutex
)
