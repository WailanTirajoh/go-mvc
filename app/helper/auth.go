package helper

import (
	"encoding/json"

	"github.com/WailanTirajoh/go-simple-clean-architecture/app/model"
)

func AuthObject(authjson string) (model.User, error) {
	var user model.User

	if err := json.Unmarshal([]byte(authjson), &user); err != nil {
		return user, err
	}

	return user, nil
}
