package person

import "encoding/json"

// Person ...
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ValidatePerson(value []byte) (valid bool, errs []string, person Person) {
	err := json.Unmarshal(value, &person)
	if err != nil {
		errs = append(errs, "Unable to Unmarshal value")
		return
	}
	if person.Name == "" {
		errs = append(errs, "Name must not be blank")
	}
	if person.Age < 1 {
		errs = append(errs, "Age must be at least 1")
	}
	if len(errs) > 0 {
		return
	}
	valid = true
	return
}
