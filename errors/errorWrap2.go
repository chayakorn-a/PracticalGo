package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func insertInfo() error {
	// have to be %w, to be wraped
	return fmt.Errorf("can't insert duplicate info : %w", sql.ErrTxDone)
}

// have to run with GO version 13 or above
func main() {
	var err error
	err = insertInfo()
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("rollback coz", err)
	}
}
