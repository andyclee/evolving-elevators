package main

type Person struct {
	current_floor 	int
	target_floor	int
	current_pain 	int
	lifetime 	int
	in_car		bool
}

type Floor struct {
	people		[]Person
	button_state	int // 0 None, 1 Up, 2 Down
}

type Car struct {
	passengers	[]Person
	open		bool
	buttons_pushed	[]int
}

type Building struct {
	floors		[]Floor
	cars		[]Car
}

func create_person(num_floors int, entry_weights []int, target_weights []int) Person {
	var new_curr_floor int = weighted_random_choice(0, num_floors, entry_weights)
	var new_target_floor int = weighted_random_choice(0, num_floors, target_weights)
	return Person{
		current_floor : new_curr_floor,
		target_floor : new_target_floor,
		current_pain  : 0,
		lifetime      : 0,
		in_car	      : false,
	}
}

func create_floor() Floor {
	var people []Person
	return Floor{
		people	      : people,
		button_state  : 0,
	}
}

func create_car() Car {
	var empty_passengers []Person
	var buttons []int
	return Car{
		passengers     : empty_passengers,
		open           : false,
		buttons_pushed : buttons,
	}
}

func create_building(Car) Building {
	var building_floors []Floor
	var building_cars []Car
	return Building{
		floors        : building_floors,
		cars          : building_cars,
	}
}
