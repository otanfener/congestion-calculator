package models

type City struct {
	Name           string   `bson:"name"`
	Tariff         []Tariff `bson:"tariff"`
	ExemptVehicles []string `bson:"exempt_vehicles"`
	Rules          Rules    `bson:"rules"`
}

type Tariff struct {
	Fee  int    `bson:"fee"`
	From string `bson:"from"`
	To   string `bson:"to"`
}

type Rules struct {
	MaxChargePerDay      int      `bson:"max_charge_per_day"`
	SingleChargeInterval int      `bson:"single_charge_interval"`
	ChargeFreeMonths     []string `bson:"charge_free_months"`
	OfficialHolidays     []string `bson:"official_holidays"`
	WorkingDays          []string `bson:"working_days"`
}
