/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.5.5
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

import (
	"time"
)

/* A list of GetCharactersCharacterIdOrders200Ok. */
//easyjson:json
type GetCharactersCharacterIdOrders200OkList []GetCharactersCharacterIdOrders200Ok

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdOrders200Ok struct {
	AccountId    int32     `json:"account_id,omitempty"`    /* Wallet division for the buyer or seller of this order. Always 1000 for characters. Currently 1000 through 1006 for corporations */
	Duration     int32     `json:"duration,omitempty"`      /* Numer of days for which order is valid (starting from the issued date). An order expires at time issued + duration */
	Escrow       float32   `json:"escrow,omitempty"`        /* For buy orders, the amount of ISK in escrow */
	IsBuyOrder   bool      `json:"is_buy_order,omitempty"`  /* True for a bid (buy) order. False for an offer (sell) order */
	IsCorp       bool      `json:"is_corp,omitempty"`       /* is_corp boolean */
	Issued       time.Time `json:"issued,omitempty"`        /* Date and time when this order was issued */
	LocationId   int64     `json:"location_id,omitempty"`   /* ID of the location where order was placed */
	MinVolume    int32     `json:"min_volume,omitempty"`    /* For bids (buy orders), the minimum quantity that will be accepted in a matching offer (sell order) */
	OrderId      int64     `json:"order_id,omitempty"`      /* Unique order ID */
	Price        float32   `json:"price,omitempty"`         /* Cost per unit for this order */
	Range_       string    `json:"range,omitempty"`         /* Valid order range, numbers are ranges in jumps */
	RegionId     int32     `json:"region_id,omitempty"`     /* ID of the region where order was placed */
	State        string    `json:"state,omitempty"`         /* Current order state */
	TypeId       int32     `json:"type_id,omitempty"`       /* The type ID of the item transacted in this order */
	VolumeRemain int32     `json:"volume_remain,omitempty"` /* Quantity of items still required or offered */
	VolumeTotal  int32     `json:"volume_total,omitempty"`  /* Quantity of items required or offered at time order was placed */
}
