package main

func (p people) dataPersonUpdate(pe person) (people, error) {
	for _, p := range dataPeople {
		if p.Id == pe.Id {
			if pe.Name != "" {
				p.Name = pe.Name
			}
			if pe.LastName != "" {
				p.LastName = pe.LastName
			}
		}
	}

	return dataPeople, nil
}
