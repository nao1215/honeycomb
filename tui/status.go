package tui

// currentView is the current view of the honeycomb operation.
type currentView int

const (
	currentViewMin currentView = currentViewTimeline
	// currentVierTimeline is the current view when the timeline view is displayed.
	currentViewTimeline currentView = 1
	// currentViewTrend is the current view when the trend view is displayed.
	currentViewTrend currentView = 2
	// currentViewFollow is the current view when the follow view is displayed.
	currentViewFollow currentView = 3
	// currentViewFollower is the current view when the follower view is displayed.
	currentViewFollower currentView = 4
	// currentViewProfile is the current view when the profile view is displayed.
	currentViewProfile currentView = 5
	// currentViewSetting is the current view when the setting view is displayed.
	currentViewSetting currentView = 6
	// currentViewMax is the maximum value of the current view.
	currentViewMax = currentViewSetting
)

// next moves the current view to the next view.
func (c *currentView) next() {
	*c++
	if *c > currentViewMax {
		*c = currentViewTimeline
	}
}

// prev moves the current view to the previous view.
func (c *currentView) prev() {
	*c--
	if *c < currentViewMin {
		*c = currentViewSetting
	}
}

// string returns the string representation of the current view.
func (c currentView) string() string {
	switch c {
	case currentViewTimeline:
		return "TL"
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
		return "🐝  TL  🐝"
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

// postFormVisible is the visibility of the post form.
type postFormVisible bool

// invisible hides the post form.
func (p *postFormVisible) invisible() {
	*p = false
}

// visible shows the post form.
func (p *postFormVisible) visible() {
	*p = true
}

// isVisble returns true if the post form is visible.
func (p *postFormVisible) isVisible() bool {
	return bool(*p)
}
