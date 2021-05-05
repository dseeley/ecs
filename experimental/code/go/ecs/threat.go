// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

import (
	"time"
)

//
type Threat struct {
	// The date and time when intelligence source first reported sighting this
	// indicator.
	IndicatorFirstSeen time.Time `ecs:"indicator.first_seen"`

	// The date and time when intelligence source last reported sighting this
	// indicator.
	IndicatorLastSeen time.Time `ecs:"indicator.last_seen"`

	// Number of times this indicator was observed conducting threat activity.
	IndicatorSightings int64 `ecs:"indicator.sightings"`

	// Type of indicator as represented by Cyber Observable in STIX 2.0.
	// Expected values
	//   * autonomous-system
	//   * artifact
	//   * directory
	//   * domain-name
	//   * email-addr
	//   * file
	//   * ipv4-addr
	//   * ipv6-addr
	//   * mac-addr
	//   * mutex
	//   * process
	//   * software
	//   * url
	//   * user-account
	//   * windows-registry-key
	//   * x-509-certificate
	IndicatorType string `ecs:"indicator.type"`

	// Describes the type of action conducted by the threat.
	IndicatorDescription string `ecs:"indicator.description"`

	// Count of AV/EDR vendors that successfully detected malicious file or
	// URL.
	IndicatorScannerStats int64 `ecs:"indicator.scanner_stats"`

	// Identifies the name of the intelligence provider.
	IndicatorProvider string `ecs:"indicator.provider"`

	// Identifies the confidence rating assigned by the provider using STIX
	// confidence scales.
	// Expected values:
	//   * Not Specified, None, Low, Medium, High
	//   * 0-10
	//   * Admirality Scale (1-6)
	//   * DNI Scale (5-95)
	//   * WEP Scale (Impossible - Certain)
	IndicatorConfidence string `ecs:"indicator.confidence"`

	// Identifies the name of specific module this data is coming from.
	IndicatorModule string `ecs:"indicator.module"`

	// Identifies the name of specific dataset from the intelligence source.
	IndicatorDataset string `ecs:"indicator.dataset"`

	// Identifies a threat indicator as an IP address (irrespective of
	// direction).
	IndicatorIP string `ecs:"indicator.ip"`

	// Identifies a threat indicator as a domain (irrespective of direction).
	IndicatorDomain string `ecs:"indicator.domain"`

	// Identifies a threat indicator as a port number (irrespective of
	// direction).
	IndicatorPort int64 `ecs:"indicator.port"`

	// Identifies a threat indicator as an email address (irrespective of
	// direction).
	IndicatorEmailAddress string `ecs:"indicator.email.address"`

	// Traffic Light Protocol sharing markings.
	// Expected values are:
	//   * White
	//   * Green
	//   * Amber
	//   * Red
	IndicatorMarkingTlp string `ecs:"indicator.marking.tlp"`

	// Identifies the atomic indicator that matched a local environment
	// endpoint or network event.
	IndicatorMatchedAtomic string `ecs:"indicator.matched.atomic"`

	// Identifies the field of the atomic indicator that matched a local
	// environment endpoint or network event.
	IndicatorMatchedField string `ecs:"indicator.matched.field"`

	// Identifies the type of the atomic indicator that matched a local
	// environment endpoint or network event.
	IndicatorMatchedType string `ecs:"indicator.matched.type"`
}
