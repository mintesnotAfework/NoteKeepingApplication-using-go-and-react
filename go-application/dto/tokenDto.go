package dto

import "github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"

type TokenDto struct {
	Value        string `json:"value"`
	RefreshValue string `json:"refresh"`
}

func Combiner(refresh *models.Token, token *string) *TokenDto {
	var result TokenDto = TokenDto{
		RefreshValue: refresh.Refresh,
		Value:        *token,
	}

	return &result
}
