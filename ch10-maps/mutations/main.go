package main
import "errors"
/*
Insert an Element
m[key] = elem

Get an Element
elem = m[key]

Delete an Element
delete(m, key)

Check If a Key Exists
elem, ok := m[key]

If key is in m, then ok is true and
elem is the value as expected.
If key is not in the map, then ok is
false and elem is the zero value for the map's element type.

*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	elem, ok := users[name]
	if ok == false {
		return false, errors.New("not found")
	}
	if ok {
		if elem.scheduledForDeletion {
			delete(users,name)
			return true, nil
		}

		if !elem.scheduledForDeletion {
			return false, nil
		}
	}
	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

