/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.3.dev6
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

package goesiv1

/* item object */
type GetKillmailsKillmailIdKillmailHashItem struct {
	Flag              int32 `json:"flag,omitempty"`               /* flag integer */
	ItemTypeId        int32 `json:"item_type_id,omitempty"`       /* item_type_id integer */
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"` /* quantity_destroyed integer */
	QuantityDropped   int64 `json:"quantity_dropped,omitempty"`   /* quantity_dropped integer */
	Singleton         int32 `json:"singleton,omitempty"`          /* singleton integer */
}
