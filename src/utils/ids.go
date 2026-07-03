package utils

// better would be uuid via package
import (
	"fmt"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRoomID() string {
	return fmt.Sprintf("room_%d", rng.Intn(100000))
}

func GenerateUserID() string {
	return fmt.Sprintf("user_%d", rng.Intn(100000))
}

func GenerateTopicID() string {
	return fmt.Sprintf("topic_%d", rng.Intn(100000))
}
