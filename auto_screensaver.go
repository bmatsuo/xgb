package xgb

/*
	This file was generated by screensaver.xml on May 8 2012 11:03:24pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// Imports are not necessary for XGB because everything is 
// in one package. They are still listed here for reference.
// import "xproto"

// ScreensaverInit must be called before using the MIT-SCREEN-SAVER extension.
func (c *Conn) ScreensaverInit() error {
	reply, err := c.QueryExtension(16, "MIT-SCREEN-SAVER").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return errorf("No extension named MIT-SCREEN-SAVER could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["MIT-SCREEN-SAVER"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["MIT-SCREEN-SAVER"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range newExtErrorFuncs["MIT-SCREEN-SAVER"] {
		newErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["MIT-SCREEN-SAVER"] = make(map[int]newEventFun)
	newExtErrorFuncs["MIT-SCREEN-SAVER"] = make(map[int]newErrorFun)
}

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Id'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

const (
	ScreensaverKindBlanked  = 0
	ScreensaverKindInternal = 1
	ScreensaverKindExternal = 2
)

const (
	ScreensaverEventNotifyMask = 1
	ScreensaverEventCycleMask  = 2
)

const (
	ScreensaverStateOff      = 0
	ScreensaverStateOn       = 1
	ScreensaverStateCycle    = 2
	ScreensaverStateDisabled = 3
)

// Event definition ScreensaverNotify (0)
// Size: 32

const ScreensaverNotify = 0

type ScreensaverNotifyEvent struct {
	Sequence uint16
	Code     byte
	State    byte
	// padding: 1 bytes
	SequenceNumber uint16
	Time           Timestamp
	Root           Id
	Window         Id
	Kind           byte
	Forced         bool
	// padding: 14 bytes
}

// Event read ScreensaverNotify
func NewScreensaverNotifyEvent(buf []byte) Event {
	v := ScreensaverNotifyEvent{}
	b := 1 // don't read event number

	v.Code = buf[b]
	b += 1

	v.Sequence = Get16(buf[b:])
	b += 2

	v.State = buf[b]
	b += 1

	b += 1 // padding

	v.SequenceNumber = Get16(buf[b:])
	b += 2

	v.Time = Timestamp(Get32(buf[b:]))
	b += 4

	v.Root = Id(Get32(buf[b:]))
	b += 4

	v.Window = Id(Get32(buf[b:]))
	b += 4

	v.Kind = buf[b]
	b += 1

	if buf[b] == 1 {
		v.Forced = true
	} else {
		v.Forced = false
	}
	b += 1

	b += 14 // padding

	return v
}

// Event write ScreensaverNotify
func (v ScreensaverNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Code
	b += 1

	b += 2 // skip sequence number

	buf[b] = v.State
	b += 1

	b += 1 // padding

	Put16(buf[b:], v.SequenceNumber)
	b += 2

	Put32(buf[b:], uint32(v.Time))
	b += 4

	Put32(buf[b:], uint32(v.Root))
	b += 4

	Put32(buf[b:], uint32(v.Window))
	b += 4

	buf[b] = v.Kind
	b += 1

	if v.Forced {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 14 // padding

	return buf
}

func (v ScreensaverNotifyEvent) ImplementsEvent() {}

func (v ScreensaverNotifyEvent) SequenceId() uint16 {
	return v.Sequence
}

func (v ScreensaverNotifyEvent) String() string {
	fieldVals := make([]string, 0, 10)
	fieldVals = append(fieldVals, sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, sprintf("Code: %d", v.Code))
	fieldVals = append(fieldVals, sprintf("State: %d", v.State))
	fieldVals = append(fieldVals, sprintf("SequenceNumber: %d", v.SequenceNumber))
	fieldVals = append(fieldVals, sprintf("Time: %d", v.Time))
	fieldVals = append(fieldVals, sprintf("Root: %d", v.Root))
	fieldVals = append(fieldVals, sprintf("Window: %d", v.Window))
	fieldVals = append(fieldVals, sprintf("Kind: %d", v.Kind))
	fieldVals = append(fieldVals, sprintf("Forced: %t", v.Forced))
	return "ScreensaverNotify {" + stringsJoin(fieldVals, ", ") + "}"
}

func init() {
	newExtEventFuncs["MIT-SCREEN-SAVER"][0] = NewScreensaverNotifyEvent
}

// Request ScreensaverQueryVersion
// size: 8
type ScreensaverQueryVersionCookie struct {
	*cookie
}

func (c *Conn) ScreensaverQueryVersion(ClientMajorVersion byte, ClientMinorVersion byte) ScreensaverQueryVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.screensaverQueryVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return ScreensaverQueryVersionCookie{cookie}
}

func (c *Conn) ScreensaverQueryVersionUnchecked(ClientMajorVersion byte, ClientMinorVersion byte) ScreensaverQueryVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.screensaverQueryVersionRequest(ClientMajorVersion, ClientMinorVersion), cookie)
	return ScreensaverQueryVersionCookie{cookie}
}

// Request reply for ScreensaverQueryVersion
// size: 32
type ScreensaverQueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
	// padding: 20 bytes
}

// Waits and reads reply data from request ScreensaverQueryVersion
func (cook ScreensaverQueryVersionCookie) Reply() (*ScreensaverQueryVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return screensaverQueryVersionReply(buf), nil
}

// Read reply into structure from buffer for ScreensaverQueryVersion
func screensaverQueryVersionReply(buf []byte) *ScreensaverQueryVersionReply {
	v := new(ScreensaverQueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = Get16(buf[b:])
	b += 2

	v.ServerMinorVersion = Get16(buf[b:])
	b += 2

	b += 20 // padding

	return v
}

func (cook ScreensaverQueryVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for ScreensaverQueryVersion
func (c *Conn) screensaverQueryVersionRequest(ClientMajorVersion byte, ClientMinorVersion byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajorVersion
	b += 1

	buf[b] = ClientMinorVersion
	b += 1

	b += 2 // padding

	return buf
}

// Request ScreensaverQueryInfo
// size: 8
type ScreensaverQueryInfoCookie struct {
	*cookie
}

func (c *Conn) ScreensaverQueryInfo(Drawable Id) ScreensaverQueryInfoCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.screensaverQueryInfoRequest(Drawable), cookie)
	return ScreensaverQueryInfoCookie{cookie}
}

func (c *Conn) ScreensaverQueryInfoUnchecked(Drawable Id) ScreensaverQueryInfoCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.screensaverQueryInfoRequest(Drawable), cookie)
	return ScreensaverQueryInfoCookie{cookie}
}

// Request reply for ScreensaverQueryInfo
// size: 32
type ScreensaverQueryInfoReply struct {
	Sequence         uint16
	Length           uint32
	State            byte
	SaverWindow      Id
	MsUntilServer    uint32
	MsSinceUserInput uint32
	EventMask        uint32
	Kind             byte
	// padding: 7 bytes
}

// Waits and reads reply data from request ScreensaverQueryInfo
func (cook ScreensaverQueryInfoCookie) Reply() (*ScreensaverQueryInfoReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return screensaverQueryInfoReply(buf), nil
}

// Read reply into structure from buffer for ScreensaverQueryInfo
func screensaverQueryInfoReply(buf []byte) *ScreensaverQueryInfoReply {
	v := new(ScreensaverQueryInfoReply)
	b := 1 // skip reply determinant

	v.State = buf[b]
	b += 1

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.SaverWindow = Id(Get32(buf[b:]))
	b += 4

	v.MsUntilServer = Get32(buf[b:])
	b += 4

	v.MsSinceUserInput = Get32(buf[b:])
	b += 4

	v.EventMask = Get32(buf[b:])
	b += 4

	v.Kind = buf[b]
	b += 1

	b += 7 // padding

	return v
}

func (cook ScreensaverQueryInfoCookie) Check() error {
	return cook.check()
}

// Write request to wire for ScreensaverQueryInfo
func (c *Conn) screensaverQueryInfoRequest(Drawable Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// Request ScreensaverSelectInput
// size: 12
type ScreensaverSelectInputCookie struct {
	*cookie
}

// Write request to wire for ScreensaverSelectInput
func (c *Conn) ScreensaverSelectInput(Drawable Id, EventMask uint32) ScreensaverSelectInputCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.screensaverSelectInputRequest(Drawable, EventMask), cookie)
	return ScreensaverSelectInputCookie{cookie}
}

func (c *Conn) ScreensaverSelectInputChecked(Drawable Id, EventMask uint32) ScreensaverSelectInputCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.screensaverSelectInputRequest(Drawable, EventMask), cookie)
	return ScreensaverSelectInputCookie{cookie}
}

func (cook ScreensaverSelectInputCookie) Check() error {
	return cook.check()
}

// Write request to wire for ScreensaverSelectInput
func (c *Conn) screensaverSelectInputRequest(Drawable Id, EventMask uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Drawable))
	b += 4

	Put32(buf[b:], EventMask)
	b += 4

	return buf
}

// Request ScreensaverSetAttributes
// size: pad((24 + (4 + pad((4 * popCount(int(ValueMask)))))))
type ScreensaverSetAttributesCookie struct {
	*cookie
}

// Write request to wire for ScreensaverSetAttributes
func (c *Conn) ScreensaverSetAttributes(Drawable Id, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual Visualid, ValueMask uint32, ValueList []uint32) ScreensaverSetAttributesCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.screensaverSetAttributesRequest(Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList), cookie)
	return ScreensaverSetAttributesCookie{cookie}
}

func (c *Conn) ScreensaverSetAttributesChecked(Drawable Id, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual Visualid, ValueMask uint32, ValueList []uint32) ScreensaverSetAttributesCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.screensaverSetAttributesRequest(Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList), cookie)
	return ScreensaverSetAttributesCookie{cookie}
}

func (cook ScreensaverSetAttributesCookie) Check() error {
	return cook.check()
}

// Write request to wire for ScreensaverSetAttributes
func (c *Conn) screensaverSetAttributesRequest(Drawable Id, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual Visualid, ValueMask uint32, ValueList []uint32) []byte {
	size := pad((24 + (4 + pad((4 * popCount(int(ValueMask)))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Drawable))
	b += 4

	Put16(buf[b:], uint16(X))
	b += 2

	Put16(buf[b:], uint16(Y))
	b += 2

	Put16(buf[b:], Width)
	b += 2

	Put16(buf[b:], Height)
	b += 2

	Put16(buf[b:], BorderWidth)
	b += 2

	buf[b] = Class
	b += 1

	buf[b] = Depth
	b += 1

	Put32(buf[b:], uint32(Visual))
	b += 4

	Put32(buf[b:], ValueMask)
	b += 4
	for i := 0; i < popCount(int(ValueMask)); i++ {
		Put32(buf[b:], ValueList[i])
		b += 4
	}
	b = pad(b)

	return buf
}

// Request ScreensaverUnsetAttributes
// size: 8
type ScreensaverUnsetAttributesCookie struct {
	*cookie
}

// Write request to wire for ScreensaverUnsetAttributes
func (c *Conn) ScreensaverUnsetAttributes(Drawable Id) ScreensaverUnsetAttributesCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.screensaverUnsetAttributesRequest(Drawable), cookie)
	return ScreensaverUnsetAttributesCookie{cookie}
}

func (c *Conn) ScreensaverUnsetAttributesChecked(Drawable Id) ScreensaverUnsetAttributesCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.screensaverUnsetAttributesRequest(Drawable), cookie)
	return ScreensaverUnsetAttributesCookie{cookie}
}

func (cook ScreensaverUnsetAttributesCookie) Check() error {
	return cook.check()
}

// Write request to wire for ScreensaverUnsetAttributes
func (c *Conn) screensaverUnsetAttributesRequest(Drawable Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// Request ScreensaverSuspend
// size: 8
type ScreensaverSuspendCookie struct {
	*cookie
}

// Write request to wire for ScreensaverSuspend
func (c *Conn) ScreensaverSuspend(Suspend bool) ScreensaverSuspendCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.screensaverSuspendRequest(Suspend), cookie)
	return ScreensaverSuspendCookie{cookie}
}

func (c *Conn) ScreensaverSuspendChecked(Suspend bool) ScreensaverSuspendCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.screensaverSuspendRequest(Suspend), cookie)
	return ScreensaverSuspendCookie{cookie}
}

func (cook ScreensaverSuspendCookie) Check() error {
	return cook.check()
}

// Write request to wire for ScreensaverSuspend
func (c *Conn) screensaverSuspendRequest(Suspend bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["MIT-SCREEN-SAVER"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Suspend {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
