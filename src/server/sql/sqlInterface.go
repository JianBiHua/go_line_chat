package sql

// SqlInterface is sql functions interface
type SqlInterface interface {

	// Open is crate tables
	//
	//	out:
	//	error: error info
	Open(path string) error

	//	close is close databse connect
	Close()

	// Exec is sql create table, insert, update, delete
	Exec(cmd string) (interface{}, error)

	// Exec2 is sql search
	Exec2(cmd string, f func(...interface{}), param ...interface{}) error
}
