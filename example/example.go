/**
 * An example of producing timestamp in a format that is popular in MySQL
 * @author: rafalgolarz.com
 */
package main

import (
	"fmt"
	"time"

	"github.com/rafalgolarz/blog-examples/timestamp"
)

func main() {

	now := timestamp.FromNow{}
	fmt.Println("current timestamp: ", now)

	//timestamp of 30 min from now
	expiresAt := timestamp.FromNow{
		Offset:   30,
		TimeUnit: time.Minute,
	}
	fmt.Println("expires at: ", expiresAt)

	//timestamp of 2 hours from now
	reminder := timestamp.FromNow{
		Offset:   2,
		TimeUnit: time.Hour,
	}
	//let's put the string into a sample SQL query
	exampleQuery := "UPDATE reminders SET when = " + reminder.String()
	exampleQuery += " WHERE reminders.id = 1;"

	fmt.Println("example SQL query: ", exampleQuery)

	//timestamp of 15 min ago
	expiredAt := timestamp.FromNow{-15, time.Minute}
	fmt.Println("expired at: ", expiredAt)
}
