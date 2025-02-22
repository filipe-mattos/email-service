package campaign

import (
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign Test"
	content  = "html body"
	contacts = []string{"email-1@email.com", "email-2@email.com"}
	fake     = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	//Arrange
	assert := assert.New(t)

	//Action
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	//Action
	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	//Action
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	//Action
	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 6", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	//Action
	_, err := NewCampaign(fake.Lorem().Text(31), content, contacts)

	assert.Equal("name is required with max 30", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	//Action
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 6", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	//Action
	_, err := NewCampaign(name, fake.Lorem().Text(2000), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	//Action
	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContactsInvalid(t *testing.T) {
	assert := assert.New(t)

	//Action
	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
