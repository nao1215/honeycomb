// Package model provides the data model of the application.
package model

// Profile is the data model of the user profile.
type Profile struct {
	Website     string `json:"website"`      // Website is the URL of the user's website.
	Nip05       string `json:"nip05"`        // Internet identifier (an email-like address); see https://github.com/nostr-protocol/nips/blob/master/05.md
	Picture     string `json:"picture"`      // Picture is the URL of the user's profile picture.
	Lud16       string `json:"lud16"`        // Lightning Address; see https://github.com/lnurl/luds/blob/luds/16.md
	DisplayName string `json:"display_name"` // Display name is the name of the user to be displayed.
	About       string `json:"about"`        // About is the user's description.
	Name        string `json:"name"`         // Name is the user's name.
	Bot         bool   `json:"bot"`          // Bot is true if the user is a bot.
}

// DisplayNameOrName returns the display name or name.
func (p *Profile) DisplayNameOrName() string {
	if p.DisplayName != "" {
		return p.DisplayName
	}
	return p.Name
}
