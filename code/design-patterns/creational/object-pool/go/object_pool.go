package main

import (
	"fmt"
	"sync"
)

//////////////////////////////////// Bullet ////////////////////////////////////

type Bullet struct {
	Position int
	Active   bool
}

func NewBullet() *Bullet {
	return &Bullet{Position: 0, Active: false}
}

func (b *Bullet) Move() {
	b.Position += 1 // Simulates the bullet moving forward
}

func (b *Bullet) Draw() {
	fmt.Printf("Drawing bullet at position %d\n", b.Position)
}

//////////////////////////////////// BulletPool ////////////////////////////////////

type BulletPool struct {
	pool []*Bullet
	mu   sync.Mutex
}

var instance *BulletPool
var once sync.Once

// GetInstance returns the singleton instance of BulletPool
func GetInstance() *BulletPool {
	once.Do(func() {
		instance = &BulletPool{
			pool: make([]*Bullet, 0, 10), // Initialize with capacity of 10
		}
	})
	return instance
}

// Acquire retrieves a Bullet from the pool or creates a new one if none are available
func (p *BulletPool) Acquire() *Bullet {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.pool) > 0 {
		bullet := p.pool[len(p.pool)-1]
		p.pool = p.pool[:len(p.pool)-1]
		bullet.Active = true
		return bullet
	}
	return NewBullet() // Create a new bullet if the pool is empty
}

// Release adds a Bullet back to the pool if there is capacity
func (p *BulletPool) Release(bullet *Bullet) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.pool) < cap(p.pool) {
		bullet.Active = false
		bullet.Position = 0 // Reset position
		p.pool = append(p.pool, bullet)
	}
}
