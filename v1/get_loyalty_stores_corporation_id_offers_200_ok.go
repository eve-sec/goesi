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

/* 200 ok object */
type GetLoyaltyStoresCorporationIdOffers200Ok struct {
	IskCost       float32                                           `json:"isk_cost,omitempty"`       /* isk_cost number */
	LpCost        int32                                             `json:"lp_cost,omitempty"`        /* lp_cost integer */
	OfferId       int32                                             `json:"offer_id,omitempty"`       /* offer_id integer */
	Quantity      int32                                             `json:"quantity,omitempty"`       /* quantity integer */
	RequiredItems []GetLoyaltyStoresCorporationIdOffersRequiredItem `json:"required_items,omitempty"` /* required_items array */
	TypeId        int32                                             `json:"type_id,omitempty"`        /* type_id integer */
}
