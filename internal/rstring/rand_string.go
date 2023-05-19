/*
    Copyright 2021 Rabia Research Team and Developers

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
/*
	The rstring package contains only one function, which generates a random string in a fast way. Possible characters
	of the string are from a to z and from A to Z.

	Note:

	1. Random strings are used as the primary key of clients' read and write requests to the database.
	2. The random string generation function follows the RandStringBytesMaskImprSrcUnsafe() function
	at https://stackoverflow.com/a/31832326
*/
package rstring

import (
	"math/rand"
	"strconv"
	"unsafe"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // low conflict
	//letterBytes   = "abcde" // high conflict
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

/*
	Generates a random string

	src: a pointer to a rand.Rand object, so this function can be called by multiple Goroutines
	n: the length of the desired random string, including characters from a to z and A tot Z.

	It returns the generated string
*/
func RandString(src *rand.Rand, n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func RandPriceString(src *rand.Rand, n int) string {
	min := 3708
	max := 3910
	// Generate Random
	//valCurrency := float64(rand.Intn(max-min)+min) / 1000
	valCurrency := float64(src.Intn(max-min)+min) / 1000
	valPrice := strconv.FormatFloat(valCurrency, 'f', 6, 64)
	if len(valPrice) != n {
		//log.Panicf("got error creating random price value %v in ExchangeStore: this length is not %v bytes", valPrice, n)
		panic("got error creating random price value in ExchangeStore: this length is not 8 bytes")
	}
	return valPrice
}
