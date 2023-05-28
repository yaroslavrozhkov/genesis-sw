package GenesisProject

type SubscribeEmail struct {
	Email string `json:"email" binding:"required"`
}
