package resp

import (
	"sync"
)

// RWMutex is a reader/writer mutual exclusion lock.
// Provides two locking mechanisms: RLock() and Lock()
// Maps in Go aren't thread-safe, so we need to use a mutex to protect the maps during concurrent access.
var SETs = map[string]string{}
var SETsMu = sync.RWMutex{} // For SETs map
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{} // For HSETs map

// Value represents a value in the RESP (Redis Serialization Protocol) format.
// It can be a string, number, bulk string, or array of other Values.
type Value struct {
	Type  string
	Str   string
	Num   int
	Bulk  string
	Array []Value
}

// Handlers is a map of command names to their corresponding handler functions.
// Each handler function takes a slice of Values as input and returns a single Value.
// The available commands are:
// - "PING": Responds with "PONG" or the first argument as a string.
// - "SET": Sets the value for the given key.
// - "GET": Retrieves the value for the given key.
// - "HSET": Sets the value for the given hash key and field.
// - "HGET": Retrieves the value for the given hash key and field.
// - "HGETALL": Retrieves all the fields and values for the given hash.
var Handlers = map[string]func([]Value) Value{
	"PING":    Ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
}

func Ping(args []Value) Value {
	if len(args) == 0 {
		return Value{Type: "string", Str: "PONG"}
	}
	return Value{Type: "string", Str: args[0].Bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{Type: "error", Str: "wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{Type: "string", Str: "Command successfully executed"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{Type: "error", Str: "wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return Value{Type: "nil"}
	}
	return Value{Type: "string", Bulk: value}
}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{Type: "error", Str: "wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{Type: "string", Str: "Command successfully executed"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{Type: "error", Str: "wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{Type: "nil"}
	}

	return Value{Type: "bulk", Bulk: value}
}

func hgetall(args []Value) Value {

	if len(args) != 1 {
		return Value{Type: "error", Str: "wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return Value{Type: "nil"}
	}

	var array []Value
	for _, v := range value {
		array = append(array, Value{Type: "bulk", Bulk: v})
	}
	return Value{Type: "array", Array: array}
}
