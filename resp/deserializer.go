package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

/*
* Deserializer is used to read bytes and interprets type prefixes
RESP bytes -> Go bytes
*/

// Recursive descent parsing of RESP protocol data
// deserializer.Read()
//   → Reads '*' → calls readArray()
//     → Reads '3' → creates array of size 3
//       → First iteration: reads '$' → calls readBulk() → returns "SET"
//       → Second iteration: reads '$' → calls readBulk() → returns "key"
//       → Third iteration: reads '$' → calls readBulk() → returns "value"

type Deserializer struct {
	reader *bufio.Reader
	// provides buffered I/O operations
}

// constructor function
func NewDeserializer(_reader io.Reader) *Deserializer {
	return &Deserializer{reader: bufio.NewReader(_reader)}
}

func (r *Deserializer) Read() (Value, error) {
	// First byte is the type identifier
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch _type {
	case ARRAY:
		return r.readArray() // *
	case BULK:
		return r.readBulk() // $
	default:
		fmt.Printf("Unknown type: %v", string(_type))
		return Value{}, nil
	}
}

func (r *Deserializer) readArray() (Value, error) {
	v := Value{}
	v.Type = "array"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	// create new array attribute in Value
	// iterate to every character and call Read() to check the Type and call internal functions accordingly
	v.Array = make([]Value, 0)
	for i := 0; i < len; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		// Append characters into array attribute of value
		v.Array = append(v.Array, val)
	}

	return v, err
}

// Bulk format: <data>\r\n
func (r *Deserializer) readBulk() (Value, error) {
	// Defining Value struct, and Type attribute
	v := Value{}
	v.Type = "Bulk"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	// buffer
	bulk := make([]byte, len)
	r.reader.Read(bulk)   // read exact number of bytes
	v.Bulk = string(bulk) // convert to string
	r.readLine()          // read the trailing CRLF
	return v, err
}

func (r *Deserializer) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte() // read byte by byte
		if err != nil {
			return nil, 0, err
		}

		n += 1 // for the record
		line = append(line, b)

		// check ending (\r\n)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}

	return line[:len(line)-2], n, nil // remove the ending, and send the remaining data
}

func (r *Deserializer) readInteger() (x, n int, e error) {
	// read the full line
	line, n, e := r.readLine()
	if e != nil {
		return 0, 0, e
	}
	// string to int64
	integer64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}

	return int(integer64), n, nil
}
