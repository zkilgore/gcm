package gcm

// Message is used by the application server to send a message to
// the GCM server. See the documentation for GCM Architectural
// Overview for more information:
// http://developer.android.com/google/gcm/gcm.html#send-msg
type Message struct {
	RegistrationIDs       []string               `json:"registration_ids"`
	CollapseKey           string                 `json:"collapse_key,omitempty"`
	Notification          *Notification          `json:"notification,omitempty"`
	Data                  map[string]interface{} `json:"data,omitempty"`
	DelayWhileIdle        bool                   `json:"delay_while_idle,omitempty"`
	TimeToLive            int                    `json:"time_to_live,omitempty"`
	RestrictedPackageName string                 `json:"restricted_package_name,omitempty"`
	DryRun                bool                   `json:"dry_run,omitempty"`
}

// Notification @TOOD
type Notification struct {
	Title string `json:"title"`
	Body  string `json:"text"`
}

// NewMessage returns a new Message with the specified payload
// and registration IDs.
func NewMessage(title, body string, data map[string]interface{}, regIDs ...string) *Message {

	m := &Message{
		RegistrationIDs: regIDs,
		Data:            data,
	}

	if title != "" && body != "" {
		m.Notification = &Notification{
			Title: title,
			Body:  body,
		}
	}

	return m
}
