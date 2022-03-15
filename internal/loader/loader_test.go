package loader

import (
	"testing"

	"github.com/guiyomh/charlatan/internal/dto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFixtureLoaderLoad(t *testing.T) {
	content := `
user:
  user_tpl (template):
    first_name: "{firstname}"
    last_name: "{lastname}"
    pseudo: "{username}"
    email: "{email}"
    create_at: "{date}"
    password: "{password:false,false,true,false,false,6}>"
  admin_1:
    first_name: "William"
    last_name: "Wallace"
    pseudo: "WW"
    password: "freedommmmmmm"
    email: "freedom@gouv.co.uk"
    isAdmin: true
    create_at: "1305-08-23 06:06:06"
  admin_{2..5} (extends user_tpl):
    isAdmin: true
  user_{bob,harry,george} (extends user_tpl):
    first_name: "{current}"
    isAdmin: false
`
	want := dto.FixtureSet{
		dto.FixtureName("user"): dto.Fixture{
			dto.SetID("admin_1"): dto.Set{
				dto.Field("create_at"):  "1305-08-23 06:06:06",
				dto.Field("email"):      "freedom@gouv.co.uk",
				dto.Field("first_name"): "William",
				dto.Field("isAdmin"):    true,
				dto.Field("last_name"):  "Wallace",
				dto.Field("password"):   "freedommmmmmm",
				dto.Field("pseudo"):     "WW",
			},
			dto.SetID("admin_{2..5} (extends user_tpl)"): dto.Set{
				dto.Field("isAdmin"): true,
			},
			dto.SetID("user_tpl (template)"): dto.Set{
				dto.Field("create_at"):  "{date}",
				dto.Field("email"):      "{email}",
				dto.Field("first_name"): "{firstname}",
				dto.Field("last_name"):  "{lastname}",
				dto.Field("password"):   "{password:false,false,true,false,false,6}>",
				dto.Field("pseudo"):     "{username}",
			},
			dto.SetID("user_{bob,harry,george} (extends user_tpl)"): dto.Set{
				dto.Field("first_name"): "{current}",
				dto.Field("isAdmin"):    false,
			},
		},
	}
	l := fixtureLoader{
		logger: &zap.Logger{},
	}
	set, err := l.load([]byte(content))

	if assert.NoError(t, err) {
		assert.Equal(t, want, set)
	}
}
