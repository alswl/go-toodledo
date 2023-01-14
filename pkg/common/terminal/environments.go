package terminal

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

func ListAll() ([]*models.EnvironmentWithKey, error) {
	var cks []*models.EnvironmentWithKey
	var cs map[string]models.Environment
	err := viper.UnmarshalKey("environments", &cs)
	if err != nil {
		return nil, err
	}
	for k, v := range cs {
		cks = append(cks, &models.EnvironmentWithKey{
			Key:         k,
			Environment: v,
		})
		// fmt.Printf("%s: %s, %s\n", k, v.Name, v.Project)
	}
	return cks, nil
}

func ListAllKeys() ([]string, error) {
	cks, err := ListAll()
	if err != nil {
		return []string{}, err
	}
	keys, _ := funk.Map(cks, func(x *models.EnvironmentWithKey) string {
		// TODO using description, v2 completions
		// return fmt.Sprintf("%s", x.Key, x.Name)
		return x.Key
	}).([]string)
	return keys, nil
}
