package config

import (
	"encoding/json"
	"fmt"
	"hibernum-backend/main/models"

	"github.com/brianvoe/gofakeit/v7"
)

func FakeData() {
	n := 1
	numInts := 1000

	var users []models.User

	roles := []string{"admin", "user", "manager"}
	permissions := []string{"read", "write", "delete"}
	numbers := []int{}

	for i := 0; i < numInts; i++ {
		numbers = append(numbers, i)
	}

	for i := 0; i < n; i++ {
		var user models.User
		user.Name = gofakeit.Name()
		user.Email = gofakeit.Email()
		user.Password = gofakeit.Password(true, true, true, true, false, 12)
		user.Phone = gofakeit.Phone()
		user.RegistrationDate = gofakeit.Date().String()
		user.LastUpdate = gofakeit.Date().String()
		user.Token = gofakeit.UUID()
		user.ProfilePicture = gofakeit.URL()
		user.LanguagePreference = gofakeit.Language()
		user.EmailVerified = gofakeit.Date().String()
		user.PasswordStrengthIndicator = gofakeit.Word()
		user.Blocked = gofakeit.Date().String()

		// Selecionando aleatoriamente um papel (role) e permissões
		roleIndex := gofakeit.Number(0, len(roles)-1)
		permissionIndex := gofakeit.Number(0, len(permissions)-1)

		user.Role = roles[roleIndex]
		user.Permissions = []string{permissions[permissionIndex]}

		// Gerando dados de endereço
		user.Address = models.Address{
			Street: gofakeit.Street(),
			//FIXME
			// Number:       gofakeit.Number(1, 100),
			Complement:   gofakeit.Address().Address,
			Neighborhood: gofakeit.StreetSuffix(),
			City:         gofakeit.City(),
			State:        gofakeit.State(),
			ZipCode:      gofakeit.Address().Zip,
		}

		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
