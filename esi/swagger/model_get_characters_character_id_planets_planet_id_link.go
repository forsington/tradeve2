/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// link object
type GetCharactersCharacterIdPlanetsPlanetIdLink struct {
	// destination_pin_id integer
	DestinationPinId int64 `json:"destination_pin_id"`
	// link_level integer
	LinkLevel int32 `json:"link_level"`
	// source_pin_id integer
	SourcePinId int64 `json:"source_pin_id"`
}
