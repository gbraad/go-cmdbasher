/*
Copyright (C) 2017 Gerard Braad <me@gbraad.nl>

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

package basher

import (
	"fmt"
	"time"
)

const DefaultInterval = 2 * time.Second
type bashingFunc func(handler chan bool)

type Basher struct {
	interval time.Duration
	handler  chan bool
	bashing  bashingFunc
}

// New Basher creates the channel to handle progress dots
func New(bashing bashingFunc) *Basher {
	return &Basher{
		interval: DefaultInterval,
		handler:  make(chan bool),
		bashing:  bashing,
	}
}

// Start starts the basher
func (b *Basher) Start() {
	go func() {
		for {
			select {
			case <- b.handler:
				return
			default:
				fmt.Print(".")
				time.Sleep(b.interval)
			}
						
			//b.bashing(b.handler)
		}
	}()
}

// Stop ends the basher
func (b *Basher) Stop() {
	b.handler <- true
}

// SetInterval sets the interval
func (b *Basher) SetInterval(interval time.Duration) {
	b.interval = interval
}

// SetInterval sets the interval
func (b *Basher) SetHandler(handler chan bool) {
	b.handler = handler
}
