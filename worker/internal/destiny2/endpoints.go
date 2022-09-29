package destiny2

import (
	"fmt"
	"net/url"
)

/* Source: https://gist.github.com/cortical-iv/a22ef122e771b994454e02b6b4e481c3 */

// Main point is to get the user's id from their username.
//        https://bungie-net.github.io/multi/operation_get_Destiny2-SearchDestinyPlayer.html
func searchDestinyPlayerURL(username, platform string) string {
	return fmt.Sprintf("/Destiny2/SearchDestinyPlayer/%s/%s/?_c=1", platform, url.QueryEscape(username))
}

// Get information about different aspects of user's character like equipped items.
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetProfile.html
//    Note components are just strings: '200,300' : you need at least one component.
func profileURL(userID, membership, components string) string {
	return fmt.Sprintf("/Destiny2/%s/Profile/%s/?components=%s", membership, userID, components)
}

// Similar to get_profile but does it for a single character. Note individual character
//    ids are returned by get_profile.
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetCharacter.html
func characterURL(userID, membership, characterID, components string) string {
	return fmt.Sprintf("/Destiny2/%s/Profile/%s/Character/%s/?components=%s", membership, userID, characterID, components)
}

// Pull item with item instance id; for instance if you have two instances of
//    Uriel's Gift, this will let you pull information about one particular instance.
//    You will get the itemInstanceId from item stats returned from get_profile or
//    get_character.
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetItem.html
func itemURL(userID, membership, itemInstanceID, components string) string {
	return fmt.Sprintf("/Destiny2/%s/Profile/%s/item/%s/?components=%s", membership, userID, itemInstanceID, components)
}

// Hooking up with the manifest!
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetDestinyEntityDefinition.html
//    If you've got a hash, and know the entity type, you can use this (or just download
//    the manifest and make your  own database to do it 10x faster)
func entityDefinitionURL(entityHash, entityType string) string {
	return fmt.Sprintf("/Manifest/%s/%s", entityHash, entityType)
}

// This is how I find the groupId for the clan a user is in. You need the group
//    id to do more interesting things like pull all members of a clan (see get_group_members)
//        https://bungie-net.github.io/multi/operation_get_GroupV2-GetGroupsForMember.html
func groupsForMemberURL(userID, membership string) string {
	return fmt.Sprintf("/GroupV2/User/%s/%s/0/1/", membership, userID) // 0/1 are the filter/groupType
}

// Pull all members of a clan. Note clans can only have 100 members max.
//        https://bungie-net.github.io/multi/operation_get_GroupV2-GetMembersOfGroup.html
func membersOfGroupURL(groupID string) string {
	return fmt.Sprintf("/GroupV2/%s/Members/?currentPage=1", groupID)
}

// Returns useful history of activities, filtered by tyupe if you want (e.g.,
//    pvp, pve, etc: see modes below)
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetActivityHistory.html
//    Queries:
//        count: total number of results to return:
//        mode: filter for activity mode to return (None returns all activities--see below)
//        page: page number of results to return, starting with 0.
//    Sample of modes (this is not all of them)
//      all: None; story:  2; strike: 3; raid: 4; PvP: 5; Patrol 6; PvE 7; Clash  12
//      Nightfall 16; Trials: 39; Social: 40 (returns nada)
func activityHistoryURL(userID, membership, characterID, mode string, page, count int) string {
	return fmt.Sprintf(
		"/Destiny2/%s/Account/%s/Character/%s/Stats/Activities/?mode=%s&page=%d&count=%d",
		membership, userID, characterID, mode, page, count,
	)
}

// Return tons of useful stats about a character (or set character_id = '0' for
//    all character data lumped together).
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetHistoricalStats.html
//    Note modes are from the same list as above with activityHistoryURL.
func historicalStatsURL(userID, membership, characterID, mode string) string {
	return fmt.Sprintf("/Destiny2/%s/Account/%s/Character/%s/Stats/?modes=%s", membership, userID, characterID, mode)
}

// Get lots of stats almost as useful as get historical stats for character, but not quite as
//    no raid data.
//        https://bungie-net.github.io/multi/operation_get_Destiny2-GetHistoricalStatsForAccount.html
func historicalStatsForAccountURL(userID, membership string) string {
	return fmt.Sprintf("/Destiny2/%s/Account/%s/Stats", membership, userID)
}
