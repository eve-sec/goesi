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

/* A list of GetCorporationsCorporationIdKillmailsRecent200Ok. */
//easyjson:json
type GetCorporationsCorporationIdKillmailsRecent200OkList []GetCorporationsCorporationIdKillmailsRecent200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdKillmailsRecent200Ok struct {
	KillmailHash string `json:"killmail_hash,omitempty"` /* A hash of this killmail */
	KillmailId   int32  `json:"killmail_id,omitempty"`   /* ID of this killmail */
}
