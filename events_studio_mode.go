package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// PreviewSceneChangedEvent : The selected preview scene has changed (only available in Studio Mode).
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#previewscenechanged
type PreviewSceneChangedEvent struct {
	// Name of the scene being previewed.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// List of sources composing the scene.
	// Same specification as [`GetCurrentScene`](#getcurrentscene).
	// Required: Yes.
	Sources []*SceneItem `json:"sources"`
	_event  `json:",squash"`
}

// StudioModeSwitchedEvent : Studio Mode has been enabled or disabled.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.3-maintenance/docs/generated/protocol.md#studiomodeswitched
type StudioModeSwitchedEvent struct {
	// The new enabled state of Studio Mode.
	// Required: Yes.
	NewState bool `json:"new-state"`
	_event   `json:",squash"`
}
