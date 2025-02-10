package utils

import (
	"fmt"
	"time"
)

func DelayedLoader() {
	fmt.Println("\nProcessing...")
	time.Sleep(2 * time.Second) // Simulate a delay
}
