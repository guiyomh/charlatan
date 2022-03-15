//go:build spec || test

package dto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/guiyomh/charlatan/internal/dto"
)

func TestSpecConvertSetToRow(t *testing.T) {

	t.Run("Should convert a set to a row", func(t *testing.T) {
		var row *dto.Row

		s := dto.Set{
			dto.Field("foo"): dto.EntryValue("bar"),
			dto.Field("baz"): dto.EntryValue(123),
		}
		expectRow := dto.Row{
			Fields: dto.Fields{
				dto.Field("foo"): "bar",
				dto.Field("baz"): "123",
			},
			Meta: dto.Meta{
				RecordID: dto.RecordID("fooID"),
				Table:    dto.TableName("biloute"),
			},
		}
		row = dto.ConvertSetToRow(s, expectRow.Meta)
		assert.Equal(t, expectRow, *row)
	})

}

func TestSpecConvertFixtureToRecord(t *testing.T) {

	t.Run("Should convert a simple fixture to a Record", func(t *testing.T) {
		var record dto.Record
		var err error

		f := dto.Fixture{
			dto.SetID("abc"): dto.Set{
				dto.Field("foo"): dto.EntryValue("bar"),
				dto.Field("baz"): dto.EntryValue(123),
			},
		}

		expectRecord := dto.Record{
			dto.RecordID("abc"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc"),
					Table:    dto.TableName("biloute"),
				},
			},
		}

		record, err = dto.ConvertFixtureToRecord(f, dto.FixtureName("biloute"))

		if assert.NoError(t, err) {
			assert.Equal(t, expectRecord, record)
		}
	})

	t.Run("Should convert a numeric range fixture to a Record", func(t *testing.T) {
		var record dto.Record
		var err error

		f := dto.Fixture{
			dto.SetID("abc_{1..3}"): dto.Set{
				dto.Field("foo"): dto.EntryValue("bar"),
				dto.Field("baz"): dto.EntryValue(123),
			},
		}

		expectRecord := dto.Record{
			dto.RecordID("abc_1"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc_1"),
					Table:    dto.TableName("biloute"),
				},
			},
			dto.RecordID("abc_2"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc_2"),
					Table:    dto.TableName("biloute"),
				},
			},
			dto.RecordID("abc_3"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc_3"),
					Table:    dto.TableName("biloute"),
				},
			},
		}

		record, err = dto.ConvertFixtureToRecord(f, dto.FixtureName("biloute"))

		if assert.NoError(t, err) {
			assert.Equal(t, expectRecord, record)
		}
	})

	t.Run("Should exclude template from Record", func(t *testing.T) {
		var record dto.Record
		var err error

		f := dto.Fixture{
			dto.SetID("abc_{1..3}"): dto.Set{
				dto.Field("foo"): dto.EntryValue("bar"),
				dto.Field("baz"): dto.EntryValue(123),
			},
			dto.SetID("bdc (template)"): dto.Set{
				dto.Field("foo"): dto.EntryValue("bar"),
				dto.Field("baz"): dto.EntryValue(123),
			},
		}

		expectRecord := dto.Record{
			dto.RecordID("abc_1"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc_1"),
					Table:    dto.TableName("biloute"),
				},
			},
			dto.RecordID("abc_2"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc_2"),
					Table:    dto.TableName("biloute"),
				},
			},
			dto.RecordID("abc_3"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("foo"): "bar",
					dto.Field("baz"): "123",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("abc_3"),
					Table:    dto.TableName("biloute"),
				},
			},
		}

		record, err = dto.ConvertFixtureToRecord(f, dto.FixtureName("biloute"))

		if assert.NoError(t, err) {
			assert.Equal(t, expectRecord, record)
		}
	})

	t.Run("Should extends a template", func(t *testing.T) {
		var record dto.Record
		var err error

		f := dto.Fixture{
			dto.SetID("user_{1..3} (extends myTpl)"): dto.Set{
				dto.Field("email"): dto.EntryValue("foo@bar.com"),
			},
			dto.SetID("myTpl (template)"): dto.Set{
				dto.Field("first_name"): dto.EntryValue("john"),
				dto.Field("last_name"):  dto.EntryValue("doe"),
			},
		}

		expectRecord := dto.Record{
			dto.RecordID("user_1"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("email"):      "foo@bar.com",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("user_1"),
					Table:    dto.TableName("customer"),
				},
			},
			dto.RecordID("user_2"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("email"):      "foo@bar.com",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("user_2"),
					Table:    dto.TableName("customer"),
				},
			},
			dto.RecordID("user_3"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("email"):      "foo@bar.com",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("user_3"),
					Table:    dto.TableName("customer"),
				},
			},
		}

		record, err = dto.ConvertFixtureToRecord(f, dto.FixtureName("customer"))

		if assert.NoError(t, err) {
			assert.Equal(t, expectRecord, record)
		}
	})
}

func TestSpecConvertFixtureSetToTable(t *testing.T) {
	fs := dto.FixtureSet{
		dto.FixtureName("customers"): dto.Fixture{
			dto.SetID("people (template)"): dto.Set{
				dto.Field("first_name"): dto.EntryValue("john"),
				dto.Field("last_name"):  dto.EntryValue("doe"),
			},
			dto.SetID("user_{1..3} (extends people)"): dto.Set{
				dto.Field("is_admin"): dto.EntryValue(false),
			},
			dto.SetID("admin (extends people)"): dto.Set{
				dto.Field("is_admin"): dto.EntryValue(true),
			},
		},
	}
	expectTable := dto.Tables{
		dto.TableName("customers"): dto.Record{
			dto.RecordID("user_1"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "false",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("user_1"),
					Table:    dto.TableName("customers"),
				},
			},
			dto.RecordID("user_2"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "false",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("user_2"),
					Table:    dto.TableName("customers"),
				},
			},
			dto.RecordID("user_3"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "false",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("user_3"),
					Table:    dto.TableName("customers"),
				},
			},
			dto.RecordID("admin"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "true",
					dto.Field("first_name"): "john",
					dto.Field("last_name"):  "doe",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("admin"),
					Table:    dto.TableName("customers"),
				},
			},
		},
	}

	tables, err := dto.ConvertFixtureSetToTable(fs)

	if assert.NoError(t, err) {
		assert.Equal(t, expectTable, tables)
	}
}

func TestFakeData(t *testing.T) {
	tables := dto.Tables{
		dto.TableName("customers"): dto.Record{
			dto.RecordID("user_1"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "false",
					dto.Field("first_name"): "{firstname}",
					dto.Field("last_name"):  "{lastname}",
					dto.Field("age"):        "{number:30,60}",
				},
			},
			dto.RecordID("user_2"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "false",
					dto.Field("first_name"): "{firstname}",
					dto.Field("last_name"):  "{lastname}",
					dto.Field("age"):        "{number:30,60}",
				},
			},
			dto.RecordID("admin"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("is_admin"):   "true",
					dto.Field("first_name"): "{firstname}",
					dto.Field("last_name"):  "{lastname}",
					dto.Field("age"):        "{number:30,60}",
				},
			},
		},
	}

	tables = dto.FakeData(tables)
	customer := tables[dto.TableName("customers")]
	assert.NotEqual(t, "{firstname}", customer[dto.RecordID("admin")].Fields[dto.Field("first_name")])
	assert.NotEqual(t, "{firstname}", customer[dto.RecordID("user_1")].Fields[dto.Field("first_name")])
	assert.NotEqual(t, "{firstname}", customer[dto.RecordID("user_2")].Fields[dto.Field("first_name")])

	assert.NotEqual(t, "{lastname}", customer[dto.RecordID("admin")].Fields[dto.Field("last_name")])
	assert.NotEqual(t, "{lastname}", customer[dto.RecordID("user_1")].Fields[dto.Field("last_name")])
	assert.NotEqual(t, "{lastname}", customer[dto.RecordID("user_2")].Fields[dto.Field("last_name")])
}

func TestSpecReplaceCurrent(t *testing.T) {
	fixtures := dto.FixtureSet{
		dto.FixtureName("phone"): dto.Fixture{
			dto.SetID("phone_2_{bob,george}"): dto.Set{
				dto.Field("number"):  "{phone}",
				dto.Field("user_id"): "@user_{current}",
			},
			dto.SetID("phone_{bob,harry,george}"): dto.Set{
				dto.Field("number"):  "{phone}",
				dto.Field("user_id"): "@user_{current}",
			},
		},
	}
	expectedTables := dto.Tables{
		dto.TableName("phone"): dto.Record{
			dto.RecordID("phone_2_bob"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("number"):  "{phone}",
					dto.Field("user_id"): "@user_bob",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("phone_2_bob"),
					Table:    dto.TableName("phone"),
				},
			},
			dto.RecordID("phone_2_george"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("number"):  "{phone}",
					dto.Field("user_id"): "@user_george",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("phone_2_george"),
					Table:    dto.TableName("phone"),
				},
			},
			dto.RecordID("phone_bob"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("number"):  "{phone}",
					dto.Field("user_id"): "@user_bob",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("phone_bob"),
					Table:    dto.TableName("phone"),
				},
			},
			dto.RecordID("phone_george"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("number"):  "{phone}",
					dto.Field("user_id"): "@user_george",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("phone_george"),
					Table:    dto.TableName("phone"),
				},
			},
			dto.RecordID("phone_harry"): &dto.Row{
				Fields: dto.Fields{
					dto.Field("number"):  "{phone}",
					dto.Field("user_id"): "@user_harry",
				},
				Meta: dto.Meta{
					RecordID: dto.RecordID("phone_harry"),
					Table:    dto.TableName("phone"),
				},
			},
		},
	}

	tables, err := dto.ConvertFixtureSetToTable(fixtures)

	if assert.NoError(t, err) {
		assert.Equal(t, expectedTables, tables)
	}

}
