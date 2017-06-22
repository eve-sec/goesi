/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.5.1
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

package goesiv3

import (
	"time"
)

/* 200 ok object */
type GetCorporationsCorporationIdOk struct {
	AllianceId             int32     `json:"alliance_id,omitempty"`             /* id of alliance that corporation is a member of, if any */
	CeoId                  int32     `json:"ceo_id,omitempty"`                  /* ceo_id integer */
	CorporationDescription string    `json:"corporation_description,omitempty"` /* corporation_description string */
	CorporationName        string    `json:"corporation_name,omitempty"`        /* the full name of the corporation */
	CreationDate           time.Time `json:"creation_date,omitempty"`           /* creation_date string */
	CreatorId              int32     `json:"creator_id,omitempty"`              /* creator_id integer */
	Faction                string    `json:"faction,omitempty"`                 /* faction string */
	MemberCount            int32     `json:"member_count,omitempty"`            /* member_count integer */
	TaxRate                float32   `json:"tax_rate,omitempty"`                /* tax_rate number */
	Ticker                 string    `json:"ticker,omitempty"`                  /* the short name of the corporation */
	Url                    string    `json:"url,omitempty"`                     /* url string */
}
