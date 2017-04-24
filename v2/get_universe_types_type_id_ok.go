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

package goesiv2

/* 200 ok object */
type GetUniverseTypesTypeIdOk struct {
	Capacity        float32                                `json:"capacity,omitempty"`         /* capacity number */
	Description     string                                 `json:"description,omitempty"`      /* description string */
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"` /* dogma_attributes array */
	DogmaEffects    []GetUniverseTypesTypeIdDogmaEffect    `json:"dogma_effects,omitempty"`    /* dogma_effects array */
	GraphicId       int32                                  `json:"graphic_id,omitempty"`       /* graphic_id integer */
	GroupId         int32                                  `json:"group_id,omitempty"`         /* group_id integer */
	IconId          int32                                  `json:"icon_id,omitempty"`          /* icon_id integer */
	Mass            float32                                `json:"mass,omitempty"`             /* mass number */
	Name            string                                 `json:"name,omitempty"`             /* name string */
	PortionSize     int32                                  `json:"portion_size,omitempty"`     /* portion_size integer */
	Published       bool                                   `json:"published,omitempty"`        /* published boolean */
	Radius          float32                                `json:"radius,omitempty"`           /* radius number */
	TypeId          int32                                  `json:"type_id,omitempty"`          /* type_id integer */
	Volume          float32                                `json:"volume,omitempty"`           /* volume number */
}
