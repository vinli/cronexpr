/*!
 * Copyright 2013 Raymond Hill
 *
 * Project: github.com/gorhill/example_test.go
 * File: example_test.go
 * Version: 1.0
 * License: GPL v3 see <https://www.gnu.org/licenses/gpl.html>
 *
 */

package cronexpr

/******************************************************************************/

import (
	"fmt"
	"time"
)

/******************************************************************************/

// ExampleMustParse
func ExampleMustParse() {
	t := time.Date(2013, time.August, 31, 0, 0, 0, 0, time.UTC)
	nextTimes := MustParse("0 0 29 2 *").NextN(t, 5)
	for i := range nextTimes {
		fmt.Println(nextTimes[i].Format(time.RFC1123))
		// Output:
		// Fri, 28 Feb 2014 00:00:00 UTC
		// Sat, 28 Feb 2015 00:00:00 UTC
		// Mon, 29 Feb 2016 00:00:00 UTC
		// Tue, 28 Feb 2017 00:00:00 UTC
		// Wed, 28 Feb 2018 00:00:00 UTC

		// This cron expression was meant every 29th of Feb,
		// i.e. leap years, below is what the test was.
		// However, I don't think it makes sense for Vinli,
		// who wants to schedule something every leap year
		// Feb?  I think the above makes more sense.
		//
		// Mon, 29 Feb 2016 00:00:00 UTC
		// Sat, 29 Feb 2020 00:00:00 UTC
		// Thu, 29 Feb 2024 00:00:00 UTC
		// Tue, 29 Feb 2028 00:00:00 UTC
		// Sun, 29 Feb 2032 00:00:00 UTC
	}
}
