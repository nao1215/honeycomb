package tui

// currentView is the current view of the honeycomb operation.
type currentView int

const (
	currentViewMin currentView = currentViewTimeline
	// currentVierTimeline is the current view when the timeline view is displayed.
	currentViewTimeline currentView = 1
	// currentViewFollow is the current view when the follow view is displayed.
	currentViewFollow currentView = 2
	// currentViewProfile is the current view when the profile view is displayed.
	currentViewProfile currentView = 3
	// currentViewMax is the maximum value of the current view.
	currentViewMax = currentViewProfile
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
		*c = currentViewProfile
	}
}

// string returns the string representation of the current view.
func (c currentView) string() string {
	switch c {
	case currentViewTimeline:
		return "TL"
	case currentViewFollow:
		return "Follow"
	case currentViewProfile:
		return "Profile"
	default:
		return ""
	}
}

// stringWithBee returns the string representation of the current view with the bee.
func (c *currentView) stringWithBee() string {
	switch *c {
	case currentViewTimeline:
		return "🐝  TL  🐝"
	case currentViewFollow:
		return "🐝  Follow  🐝"
	case currentViewProfile:
		return "🐝  Profile  🐝"
	default:
		return ""
	}
}

// visible is the visibility of the post form.
type visible bool

// invisible hides the post form.
func (p *visible) invisible() {
	*p = false
}

// visible shows the post form.
func (p *visible) visible() {
	*p = true
}

// isVisble returns true if the post form is visible.
func (p *visible) isVisible() bool {
	return bool(*p)
}
