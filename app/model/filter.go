package model

import "github.com/nbd-wtf/go-nostr"

// ProfileFilter returns the filter for getting the user's profile.
func ProfileFilter(publicKey PublicKey) nostr.Filter {
	return nostr.Filter{
		Kinds:   []int{nostr.KindProfileMetadata},
		Authors: []string{publicKey.String()},
		Limit:   1,
	}
}

// ProfilesFilter returns the filter for getting the user's profiles.
func ProfilesFilter(publicKeys []PublicKey) nostr.Filter {
	pks := make([]string, len(publicKeys))
	for i, pk := range publicKeys {
		pks[i] = pk.String()
	}
	return nostr.Filter{
		Kinds:   []int{nostr.KindProfileMetadata},
		Authors: pks,
	}
}

// MyFollowFilter returns the filter for the followers.
func MyFollowFilter(publicKey PublicKey) nostr.Filter {
	return nostr.Filter{
		Kinds:   []int{nostr.KindContactList},
		Authors: []string{publicKey.String()},
		Limit:   1,
	}
}

// FollowsTimelineFilter returns the filter for the timeline.
func FollowsTimelineFilter(publicKeys []PublicKey, since, until *nostr.Timestamp, limit int) nostr.Filter {
	pks := make([]string, len(publicKeys))
	for i, pk := range publicKeys {
		pks[i] = pk.String()
	}
	return nostr.Filter{
		Kinds:   []int{nostr.KindTextNote},
		Authors: pks,
		Since:   since,
		Until:   until,
		Limit:   limit,
	}
}

// LikesFilter returns the filter for the likes.
func LikesFilter(id string) nostr.Filter {
	return nostr.Filter{
		Kinds: []int{nostr.KindTextNote},
		IDs:   []string{id},
	}
}
