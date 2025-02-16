package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCampaign(t *testing.T) {

	//Arrange
	assert := assert.New(t)
	name := "Campaign Test"
	content := "html body"
	contacts := []string{"email-1@email.com", "email-2@email.com"}

	//Action
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
