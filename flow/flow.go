package flow

// Flow defines an interface for overall process interaction.
type Flow interface {
	// Validate check if schema pass the required fields.
	Validate()

	// Render displays HTML forms to the user.
	Render()

	// PostHookFlows represents additional flow(s) to execute after finalizing intended flow.
	PostHookFlows([]Flow)
}
