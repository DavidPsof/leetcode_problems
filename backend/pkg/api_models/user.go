package api_models

type User struct {
	Name       string `json:"name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
}
