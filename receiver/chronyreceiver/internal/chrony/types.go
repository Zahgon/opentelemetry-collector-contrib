// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) Facebook, Inc. and its affiliates.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package chrony // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/chrony"

import (
	"net"
	"time"

	"github.com/facebook/time/ntp/chrony"
)

/*
	Notice:

	This file's contents has been taken and modified from:
		https://github.com/facebook/time/blob/f7a0d7702d50399208f352704780bc2fccf29eab/ntp/chrony/helpers.go
		https://github.com/facebook/time/blob/f7a0d7702d50399208f352704780bc2fccf29eab/ntp/chrony/packet.go

	The original implementation relies on the package `logrus` which is
	not compatible with open telemetry collector and easy to convert to
	the adopted logging package.

	Changes:
	- Replaced IPAddr named padding variable with a blank assignment
	- Renamed chronyFloat to binaryFloat to avoid name stuttering
	- Renamed decodePacket to newTrackingData
	- Removed any code paths that did not relate to tracking data
*/

const (
	// magic numbers to convert Chrony's binaryFloat to normal float
	floatExpBits  = 7
	floatCoefBits = (4*8 - floatExpBits)

	ipAddrInet4 uint16 = 1

	// This is used in timeSpec.SecHigh for 32-bit timestamps
	noHighSec uint32 = 0x7fffffff

	successfulRequest = 0

	replyTrackingCode = 5

	maxDataLen = 396
)

type (
	Tracking  = chrony.Tracking
	ReplyHead = chrony.ReplyHead
)

type ipAddr struct {
	IP     [16]uint8
	Family uint16
	_      uint16 // Padding
}

func (ip *ipAddr) ToNetIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

type timeSpec struct {
	SecHigh uint32
	SecLow  uint32
	Nsec    uint32
}

func (bt *timeSpec) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type binaryFloat int32

func (bf binaryFloat) Float() float64 { _ = "STUB: not implemented"; return 0 }

type requestTrackingContent struct {
	chrony.RequestHead

	Data [maxDataLen]uint8
}

type replyTrackingContent struct {
	RefID              uint32
	IPAddr             ipAddr // our current sync source
	Stratum            uint16
	LeapStatus         uint16
	RefTime            timeSpec
	CurrentCorrection  binaryFloat
	LastOffset         binaryFloat
	RMSOffset          binaryFloat
	FreqPPM            binaryFloat
	ResidFreqPPM       binaryFloat
	SkewPPM            binaryFloat
	RootDelay          binaryFloat
	RootDispersion     binaryFloat
	LastUpdateInterval binaryFloat
}

// newTrackingData is simplified implementation of:
//
//	https://github.com/facebook/time/blob/cdd8191cdec2aaa46ee72aa28cf5777c9d6f898b/ntp/chrony/packet.go#L695
//
// this client doesn't perform any other actions and due to the logrus logger being part of that code path,
// it was simpler to port the logic here and reference the original.
func newTrackingData(data []uint8) (*Tracking, error) { _ = "STUB: not implemented"; return nil, nil }

// Convert the data from the chrony c representation of the data to a more go idiomatic value
