package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {

	//Arrange
	assert := assert.New(t)
	name := "Campaign Test"
	content := "html body"
	contacts := []string{"email-1@email.com", "email-2@email.com"}

	//Action
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign Test"
	content := "html body"
	contacts := []string{"email-1@email.com", "email-2@email.com"}

	//Action
	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign Test"
	content := "html body"
	contacts := []string{"email-1@email.com", "email-2@email.com"}
	now := time.Now().Add(-time.Minute)
	//Action
	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}
