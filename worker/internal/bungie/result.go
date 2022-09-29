package bungie

type (
	Result struct {
		ErrorCode          int               `json:"ErrorCode"`
		ErrorStatus        string            `json:"ErrorStatus"`
		Message            string            `json:"Message"`
		ThrottleSeconds    int               `json:"ThrottleSeconds"`
		MessageData        map[string]string `json:"MessageData"`
		DetailedErrorTrace string            `json:"DetailedErrorTrace"`
		Response           any               `json:"Response"`
	}
)
