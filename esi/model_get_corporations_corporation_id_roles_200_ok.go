/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.6
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetCorporationsCorporationIdRoles200Ok. */
//easyjson:json
type GetCorporationsCorporationIdRoles200OkList []GetCorporationsCorporationIdRoles200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdRoles200Ok struct {
	CharacterId           int32    `json:"character_id,omitempty"`             /* character_id integer */
	GrantableRoles        []string `json:"grantable_roles,omitempty"`          /* grantable_roles array */
	GrantableRolesAtBase  []string `json:"grantable_roles_at_base,omitempty"`  /* grantable_roles_at_base array */
	GrantableRolesAtHq    []string `json:"grantable_roles_at_hq,omitempty"`    /* grantable_roles_at_hq array */
	GrantableRolesAtOther []string `json:"grantable_roles_at_other,omitempty"` /* grantable_roles_at_other array */
	Roles                 []string `json:"roles,omitempty"`                    /* roles array */
	RolesAtBase           []string `json:"roles_at_base,omitempty"`            /* roles_at_base array */
	RolesAtHq             []string `json:"roles_at_hq,omitempty"`              /* roles_at_hq array */
	RolesAtOther          []string `json:"roles_at_other,omitempty"`           /* roles_at_other array */
}
