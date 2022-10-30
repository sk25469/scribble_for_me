package utils

import "github.com/sk25469/scribble_backend/pkg/model"

//	Insert a new id so that there are equal no. of user in each group
//
//	returns both groups
func InsertClientInRoom(grp1 []model.ClientInfo, grp2 []model.ClientInfo, client model.ClientInfo) ([]model.ClientInfo, []model.ClientInfo) {
	if len(grp1) > len(grp2) {
		grp2 = append(grp2, client)
		return grp1, grp2
	}
	grp1 = append(grp1, client)
	return grp1, grp2
}
