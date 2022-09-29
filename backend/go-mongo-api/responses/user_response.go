package responses

// Creates a UserResponse struct with Status, Message, and Data property to represent the API response type.

type UserResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}

// json:"status", json:"message", and json:"data" are known as struct tags. Struct tags allow us to attach meta-information to corresponding struct properties. In other words, we use them to reformat the JSON response returned by the API.
