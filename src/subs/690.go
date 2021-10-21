package subs

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

type Employee struct {
	Id int
	Importance int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {

	m := make(map[int]Employee)
	for _, e := range employees {
		m[e.Id] = *e
	}

	helper := func(_id int) int {return 0}
	helper = func(_id int) int {
		_res := m[_id].Importance
		for _, i := range m[_id].Subordinates {
			_res += helper(i)
		}
		return _res
	}

	return helper(id)
}