package singleton

import (
	"os"
	"sync"
)

type ConfigManager struct {
	data map[string]string
	Rmtx sync.RWMutex
}

var (
	configInstance *ConfigManager
	configOnce     sync.Once
)

func GetConfigManager() *ConfigManager {
	configOnce.Do(func() {
		configInstance = &ConfigManager{
			data: make(map[string]string),
		}
		configInstance.loadFromEnv()
	})
	return configInstance
}

func (c *ConfigManager) loadFromEnv() {
	envs := []string{
		"DATABASE_URL",
		"SERVER_PORT",
		"LOG_LEVEL",
		"APP_NAME",
		"APP_VERSION",
	}

	c.Rmtx.Lock()
	defer c.Rmtx.Unlock()

	for _, key := range envs {
		if value := os.Getenv(key); value != "" {
			c.data[key] = value
		}
	}
}

// Get - безопасное чтение
func (c *ConfigManager) Get(key string) string {
	c.Rmtx.RLock()
	defer c.Rmtx.RUnlock()
	return c.data[key]
}

// GetOrDefault - безопасное чтение с дефолтом
func (c *ConfigManager) GetOrDefault(key, defaultValue string) string {
	c.Rmtx.RLock()
	defer c.Rmtx.RUnlock()

	if value, ok := c.data[key]; ok && value != "" {
		return value
	}
	return defaultValue
}

// Has - проверка наличия
func (c *ConfigManager) Has(key string) bool {
	c.Rmtx.RLock()
	defer c.Rmtx.RUnlock()

	_, ok := c.data[key]
	return ok
}

// All - возвращает КОПИЮ карты (безопасно)
func (c *ConfigManager) All() map[string]string {
	c.Rmtx.RLock()
	defer c.Rmtx.RUnlock()

	copy := make(map[string]string)
	for k, v := range c.data {
		copy[k] = v
	}
	return copy
}
