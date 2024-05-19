package interactor

import (
	"context"

	"github.com/nao1215/honeycomb/app/model"
	"github.com/nao1215/honeycomb/app/service"
	"github.com/nao1215/honeycomb/app/usecase"
)

var _ usecase.AuthorGetter = (*AuthorGetter)(nil)

// AuthorGetter is an interactor that provides the GetAuthor usecase.
type AuthorGetter struct {
	service.RelayFinder
	service.EventsLister
}

// NewAuthorGetter is a constructor for the AuthorGetter interactor.
func NewAuthorGetter(rf service.RelayFinder, el service.EventsLister) *AuthorGetter {
	return &AuthorGetter{
		RelayFinder:  rf,
		EventsLister: el,
	}
}

// GetAuthor gets the author.
func (a *AuthorGetter) GetAuthor(ctx context.Context, input *usecase.AuthorGetterInput) (*usecase.AuthorGetterOutput, error) {
	// TODO: Get Relays data fron upper layer.
	relay, err := a.RelayFinder.FindRelay(ctx, &service.RelayFinderInput{
		Relays: map[model.WSS]model.Relay{
			"wss://relay-jp.nostr.wirednet.jp": {
				WSS:    "wss://relay-jp.nostr.wirednet.jp",
				Read:   true,
				Write:  true,
				Search: true,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	pub, err := input.NSecretKey.ToPublicKey()
	if err != nil {
		return nil, err
	}

	npub, err := pub.ToNPublicKey()
	if err != nil {
		return nil, err
	}

	privateKey, err := input.NSecretKey.ToPrivateKey()
	if err != nil {
		return nil, err
	}

	events, err := a.EventsLister.ListEvents(ctx, &service.EventsListerInput{
		Filter: model.ProfileFilter(pub),
		Relay:  relay.Relay,
	})
	if err != nil {
		return nil, err
	}

	profiles, err := toProfiles(events.Events)
	if err != nil {
		return nil, err
	}

	return &usecase.AuthorGetterOutput{
		Author: &model.Author{
			NSecretKey: input.NSecretKey,
			PrivateKey: privateKey,
			PublicKey:  pub,
			NPublicKey: npub,
			Profile:    profiles[0],
			Relays: map[model.WSS]model.Relay{
				"wss://relay-jp.nostr.wirednet.jp": {
					WSS:    "wss://relay-jp.nostr.wirednet.jp",
					Read:   true,
					Write:  true,
					Search: true,
				},
			},
			ConnectedRelay: relay.Relay,
		},
	}, nil
}
