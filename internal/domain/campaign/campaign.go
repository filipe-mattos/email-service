// Dominio para campanha de email
package campaign

import (
	internalerrors "email-service/internal/internal-errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=6,max=30"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=6,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

// funcao de constructor
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}
	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}
