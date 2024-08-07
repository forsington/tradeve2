/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 200 ok object
type GetUniverseSystemKills200Ok struct {
	// Number of NPC ships killed in this system
	NpcKills int32 `json:"npc_kills"`
	// Number of pods killed in this system
	PodKills int32 `json:"pod_kills"`
	// Number of player ships killed in this system
	ShipKills int32 `json:"ship_kills"`
	// system_id integer
	SystemId int32 `json:"system_id"`
}
