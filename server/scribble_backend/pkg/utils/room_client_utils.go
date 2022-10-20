package utils

//	Insert a new id so that there are equal no. of user in each group
func InsertClientInRoom(grp1 []string, grp2 []string, id string) ([]string, []string) {
	if len(grp1) > len(grp2) {
		grp2 = append(grp2, id)
		return grp1, grp2
	}
	grp1 = append(grp1, id)
	return grp1, grp2
}
