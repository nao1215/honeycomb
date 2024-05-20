// Package interactor contains the implementations of usecases.
package interactor

import (
	"context"
	"encoding/json"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
	"github.com/nbd-wtf/go-nostr"
)

var _ usecase.ProfileGetter = (*ProfileGetter)(nil)

// ProfileGetter implements the ProfileGetter interface.
type ProfileGetter struct {
	service.EventsLister
}

// NewProfileGetter creates a new ProfileGetter.
func NewProfileGetter(
	profileGetter service.EventsLister,
) *ProfileGetter {
	return &ProfileGetter{
		EventsLister: profileGetter,
	}
}

// GetProfile gets the user's profile.
func (p *ProfileGetter) GetProfile(ctx context.Context, input *usecase.ProfileGetterInput) (*usecase.ProfileGetterOutput, error) {
	eventsOutput, err := p.EventsLister.ListEvents(ctx, &service.EventsListerInput{
		Filter: model.ProfileFilter(input.PublicKey),
		Relay:  input.ConnectedRelay,
	})
	if err != nil {
		return nil, err
	}

	npub, err := input.PublicKey.ToNPublicKey()
	if err != nil {
		return nil, err
	}

	profiles, err := toProfiles(eventsOutput.Events)
	if err != nil {
		return nil, err
	}
	return &usecase.ProfileGetterOutput{
		Profile:    *profiles[0],
		NPublicKey: npub,
	}, nil
}

// toProfiles converts the events to profiles.
func toProfiles(events []*nostr.Event) ([]*model.Profile, error) {
	if len(events) == 0 {
		return nil, usecase.ErrNoProfile
	}

	profiles := make([]*model.Profile, 0, len(events))
	for _, event := range events {
		var profile model.Profile
		err := json.Unmarshal([]byte(event.Content), &profile)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}
	return profiles, nil
}
