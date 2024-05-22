package tui

// status is the status of the honeycomb operation.
type status uint

const (
	// statusNone is the status when the honeycomb operation is not executed.
	statusNone status = iota
	// statusPrivateKeyInput is the status when the private key input view is displayed.
	statusPrivateKeyInput
	// statusPrivateKeyValidateErr is the status when the private key validation error occurs.
	statusPrivateKeyValidateErr
	// statusPrivateKeySaveErr is the status when the private key save error occurs.
	statusPrivateKeySaveErr
)

// currentView is the current view of the honeycomb operation.
type currentView int

const (
	// currentVierTimeline is the current view when the timeline view is displayed.
	currentViewTimeline currentView = iota
	// currentViewTrend is the current view when the trend view is displayed.
	currentViewTrend
	// currentViewFollow is the current view when the follow view is displayed.
	currentViewFollow
	// currentViewFollower is the current view when the follower view is displayed.
	currentViewFollower
	// currentViewProfile is the current view when the profile view is displayed.
	currentViewProfile
	// currentViewSetting is the current view when the setting view is displayed.
	currentViewSetting
)

// next moves the current view to the next view.
func (c *currentView) next() {
	*c++
	if *c > currentViewSetting {
		*c = currentViewTimeline
	}
}

// prev moves the current view to the previous view.
func (c *currentView) prev() {
	*c--
	if *c < currentViewTimeline {
		*c = currentViewSetting
	}
}

// string returns the string representation of the current view.
func (c currentView) string() string {
	switch c {
	case currentViewTimeline:
		return "Timeline"
	case currentViewTrend:
		return "Trend"
	case currentViewFollow:
		return "Follow"
	case currentViewFollower:
		return "Follower"
	case currentViewProfile:
		return "Profile"
	case currentViewSetting:
		return "Setting"
	default:
		return ""
	}
}

// stringWithBee returns the string representation of the current view with the bee.
func (c *currentView) stringWithBee() string {
	switch *c {
	case currentViewTimeline:
		return "🐝  Timeline  🐝"
	case currentViewTrend:
		return "🐝  Trend  🐝"
	case currentViewFollow:
		return "🐝  Follow  🐝"
	case currentViewFollower:
		return "🐝  Follower  🐝"
	case currentViewProfile:
		return "🐝  Profile  🐝"
	case currentViewSetting:
		return "🐝  Setting  🐝"
	default:
		return ""
	}
}
