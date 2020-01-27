package config

import (
	"log"
)

type Config map[string]interface{}

const (
	DRIVER     = "driver"
	HOST       = "host"
	PORT       = "port"
	USERNAME   = "username"
	PASSWORD   = "password"
	DBNAME     = "dbname"
	COLLECTION = "collection"
	STAT       = "stat"
	TYPE       = "type"
	BEAUTY     = "beauty"
	PROMPT     = "prompt"
)

func (c Config) GetString(key string) string {
	v, ok := c[key].(string)
	if !ok {
		log.Printf("unable to assert type to string (%s)", key)
	}
	return v
}

func (c Config) GetInt(key string) int {
	// conversion here
	v, ok := c[key].(int)
	if !ok {
		log.Printf("unable to assert type to int (%s)", key)
	}
	return v
}

func (c Config) GetFloat(key string) float64 {
	v, ok := c[key].(float64)
	if !ok {
		log.Printf("unable to assert type to float64 (%s)", key)
	}
	return v
}

func (c Config) GetBool(key string) bool {
	v, ok := c[key].(bool)
	if !ok {
		log.Printf("unable to assert type to string (%s)", key)
	}
	return v
}
