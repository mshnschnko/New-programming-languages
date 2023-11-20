package address

type Address struct {
	country  string
	city     string
	street   string
	building string
}

func NewAddress(country string, city string, street string, building string) *Address {
	address := Address{
		country:  country,
		city:     city,
		street:   street,
		building: building,
	}
	return &address
}

func (a *Address) GetAddress() (string, string, string, string) {
	return a.country, a.city, a.street, a.building
}

func (a *Address) GetCountry() string {
	return a.country
}

func (a *Address) GetCity() string {
	return a.city
}

func (a *Address) GetStreet() string {
	return a.street
}

func (a *Address) GetBuilding() string {
	return a.building
}

func (a *Address) SetCountry(country string) {
	a.country = country
}

func (a *Address) SetCity(city string) {
	// Проверка, что такой город есть в этой стране
	// ...
	a.city = city
}

func (a *Address) SetStreet(street string) {
	// Проверка, что такая улица есть в этом городе
	// ...
	a.street = street
}

func (a *Address) SetBuilding(building string) {
	// Проверка, что такой дом есть на этой улице
	// ...
	a.building = building
}

func (a *Address) SetAddress(country string, city string, street string, building string) {
	// Аналогичные проверки на существование адреса
	// ...
	a.country = country
	a.city = city
	a.street = street
	a.building = building
}
