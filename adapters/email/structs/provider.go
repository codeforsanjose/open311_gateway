package structs

//go:generate stringer -type=PayloadType

// PayloadType enumerates the valid payload types.
type PayloadType int

// NRT* are constants enumerating the valid request types.
const (
	PTUnknown PayloadType = iota
	PTJSON
	PTXML
	PTByte
	PTString
)

// ------------------------------------ Interfaces ------------------------------------

// Payloader is an interface to a Payload object to retrieve the Payload contents.
type Payloader interface {
	Get() (ptype PayloadType, contents interface{})
}

// ------------------------------------ Payload ------------------------------------

// Payload represents the payload for a call to a Provider.
type Payload struct {
	ptype   PayloadType
	content interface{}
}

// NewPayloadString returns a new Payload struct for the specified body of type string.
func NewPayloadString(body *string) *Payload {
	return &Payload{
		ptype:   PTString,
		content: body,
	}
}

// NewPayloadByte returns a new Payload struct for the specified body of type string.
func NewPayloadByte(body []byte) *Payload {
	return &Payload{
		ptype:   PTByte,
		content: body,
	}
}

// Get returns the type and contents of a Payload.
func (r Payload) Get() (ptype PayloadType, contents interface{}) {
	return r.ptype, r.content
}