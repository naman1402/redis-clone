package resp

import "sync"

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

// Value represents a value in the RESP (Redis Serialization Protocol) format.
// It can be a string, number, bulk string, or array of other Values.
type Value struct {
	Type  string
	Str   string
	Num   int
	Bulk  string
	Array []Value
}

var Handlers = map[string]func([]Value) Value{
	"PING":    Ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
}

func Ping(args []Value) Value    {}
func set(args []Value) Value     {}
func get(args []Value) Value     {}
func hset(args []Value) Value    {}
func hget(args []Value) Value    {}
func hgetall(args []Value) Value {}
