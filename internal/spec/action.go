package spec

const (
	InvalidAction = iota
	ActionCreate  // ActionCreate represents a create action.
	ActionRead    // ActionRead represents a read action.
	ActionUpdate  // ActionUpdate represents an update action.
	ActionDelete  // ActionDelete represents a delete action.
)

// Action represents an action.
type Action int
