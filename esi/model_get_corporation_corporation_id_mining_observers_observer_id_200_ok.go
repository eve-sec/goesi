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

/* A list of GetCorporationCorporationIdMiningObserversObserverId200Ok. */
//easyjson:json
type GetCorporationCorporationIdMiningObserversObserverId200OkList []GetCorporationCorporationIdMiningObserversObserverId200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationCorporationIdMiningObserversObserverId200Ok struct {
	CharacterId           int32  `json:"character_id,omitempty"`            /* The character that did the mining  */
	LastUpdated           string `json:"last_updated,omitempty"`            /* last_updated string */
	Quantity              int64  `json:"quantity,omitempty"`                /* quantity integer */
	RecordedCorporationId int32  `json:"recorded_corporation_id,omitempty"` /* The corporation id of the character at the time data was recorded.  */
	TypeId                int32  `json:"type_id,omitempty"`                 /* type_id integer */
}
