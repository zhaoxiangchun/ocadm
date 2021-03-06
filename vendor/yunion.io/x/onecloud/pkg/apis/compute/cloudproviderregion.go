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

package compute

import "yunion.io/x/onecloud/pkg/apis"

type CloudproviderregionDetails struct {
	apis.JoinModelBaseDetails

	Cloudprovider string
	Cloudregion   string

	// 云账号Id
	// example: fa4aaf88-aed8-422d-84e7-56dea533b364
	CloudaccountId string `json:"cloudaccount_id"`
	// 云账号名称
	// example: googl-account
	Cloudaccount string `json:"cloudaccount"`

	// 云账号所在域Id
	// example: 0df40413-eb69-49c7-895e-618ddeb80f55
	CloudaccountDomainId string `json:"cloudaccount_domain_id"`
	// 是否开启自动同步
	// example: false
	EnableAutoSync bool `json:"enable_auto_sync"`
	// 自动同步周期
	// example: 300
	SyncIntervalSeconds int `json:"sync_interval_seconds"`
}
