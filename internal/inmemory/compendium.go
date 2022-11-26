package inmemory

import (
	"sync"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type WeaponCache struct {
	mu      sync.RWMutex
	weapons []domain.Weapon
}

func (c *WeaponCache) Save(l []domain.Weapon) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.weapons = l
}

func (c *WeaponCache) Get() []domain.Weapon {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.weapons
}

type WeaponClassCache struct {
	mu      sync.RWMutex
	classes []domain.WeaponClass
}

func (c *WeaponClassCache) Save(l []domain.WeaponClass) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.classes = l
}

func (c *WeaponClassCache) Get() []domain.WeaponClass {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.classes
}
