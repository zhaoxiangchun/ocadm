// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netutils

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidNumber  = errors.New("invalid number")
	ErrOutOfRange     = errors.New("ip number out of range [0-255]")
	ErrInvalidIPAddr  = errors.New("invalid ip address")
	ErrInvalidMask    = errors.New("invalid mask")
	ErrOutOfRangeMask = errors.New("out of range masklen [0-32]")
)
