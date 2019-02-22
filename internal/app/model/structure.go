package model

type FixtureData map[string]interface{}

type FixtureFields map[string]interface{}
type FixtureRecords map[string]FixtureFields
type FixtureTables map[string]FixtureRecords
