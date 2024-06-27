package main

import "fmt"

func main() {
	pool := GetInstance()

	// Acquire a bullet and simulate firing it
	bullet := pool.Acquire()
	bullet.Position = 10 // Set starting position
	fmt.Printf("Bullet fired from position %d\n", bullet.Position)

	// Move the bullet and draw it
	bullet.Move()
	bullet.Draw()

	// Release the bullet back to the pool
	pool.Release(bullet)
	fmt.Println("Bullet released and returned to pool")
}
