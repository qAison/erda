// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dop

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/internal/tools/openapi/legacy/api/apis"
	"github.com/erda-project/erda/internal/tools/openapi/legacy/api/spec"
)

var CMDB_APPLICATION_UPDATE = apis.ApiSpec{
	Path:         "/api/applications/<applicationID>",
	BackendPath:  "/api/applications/<applicationID>",
	Host:         "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:       "http",
	Method:       "PUT",
	CheckLogin:   true,
	CheckToken:   true,
	RequestType:  apistructs.ApplicationUpdateRequest{},
	ResponseType: apistructs.ApplicationUpdateResponse{},
	IsOpenAPI:    true,
	Doc:          "summary: 更新应用",
	Audit: func(ctx *spec.AuditContext) error {
		appID, err := ctx.GetParamInt64("applicationId")
		if err != nil {
			return err
		}
		return ctx.CreateAudit(&apistructs.Audit{
			ScopeType:    "app",
			ScopeID:      uint64(appID),
			TemplateName: apistructs.UpdateAppTemplate,
			Context:      make(map[string]interface{}, 0),
		})
	},
}
