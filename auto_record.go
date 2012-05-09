package xgb

/*
	This file was generated by record.xml on May 8 2012 11:03:24pm EDT.
	This file is automatically generated. Edit at your peril!
*/

// RecordInit must be called before using the RECORD extension.
func (c *Conn) RecordInit() error {
	reply, err := c.QueryExtension(6, "RECORD").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return errorf("No extension named RECORD could be found on on the server.")
	}

	c.extLock.Lock()
	c.extensions["RECORD"] = reply.MajorOpcode
	for evNum, fun := range newExtEventFuncs["RECORD"] {
		newEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range newExtErrorFuncs["RECORD"] {
		newErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	c.extLock.Unlock()

	return nil
}

func init() {
	newExtEventFuncs["RECORD"] = make(map[int]newEventFun)
	newExtErrorFuncs["RECORD"] = make(map[int]newErrorFun)
}

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

// Skipping definition for base type 'Int8'

const (
	RecordHTypeFromServerTime     = 1
	RecordHTypeFromClientTime     = 2
	RecordHTypeFromClientSequence = 4
)

const (
	RecordCsCurrentClients = 1
	RecordCsFutureClients  = 2
	RecordCsAllClients     = 3
)

// Skipping resource definition of 'Context'

type RecordElementHeader byte

type RecordClientSpec uint32

// 'RecordRange8' struct definition
// Size: 2
type RecordRange8 struct {
	First byte
	Last  byte
}

// Struct read RecordRange8
func ReadRecordRange8(buf []byte, v *RecordRange8) int {
	b := 0

	v.First = buf[b]
	b += 1

	v.Last = buf[b]
	b += 1

	return b
}

// Struct list read RecordRange8
func ReadRecordRange8List(buf []byte, dest []RecordRange8) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = RecordRange8{}
		b += ReadRecordRange8(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write RecordRange8
func (v RecordRange8) Bytes() []byte {
	buf := make([]byte, 2)
	b := 0

	buf[b] = v.First
	b += 1

	buf[b] = v.Last
	b += 1

	return buf
}

// Write struct list RecordRange8
func RecordRange8ListBytes(buf []byte, list []RecordRange8) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// 'RecordRange16' struct definition
// Size: 4
type RecordRange16 struct {
	First uint16
	Last  uint16
}

// Struct read RecordRange16
func ReadRecordRange16(buf []byte, v *RecordRange16) int {
	b := 0

	v.First = Get16(buf[b:])
	b += 2

	v.Last = Get16(buf[b:])
	b += 2

	return b
}

// Struct list read RecordRange16
func ReadRecordRange16List(buf []byte, dest []RecordRange16) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = RecordRange16{}
		b += ReadRecordRange16(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write RecordRange16
func (v RecordRange16) Bytes() []byte {
	buf := make([]byte, 4)
	b := 0

	Put16(buf[b:], v.First)
	b += 2

	Put16(buf[b:], v.Last)
	b += 2

	return buf
}

// Write struct list RecordRange16
func RecordRange16ListBytes(buf []byte, list []RecordRange16) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// 'RecordExtRange' struct definition
// Size: 6
type RecordExtRange struct {
	Major RecordRange8
	Minor RecordRange16
}

// Struct read RecordExtRange
func ReadRecordExtRange(buf []byte, v *RecordExtRange) int {
	b := 0

	v.Major = RecordRange8{}
	b += ReadRecordRange8(buf[b:], &v.Major)

	v.Minor = RecordRange16{}
	b += ReadRecordRange16(buf[b:], &v.Minor)

	return b
}

// Struct list read RecordExtRange
func ReadRecordExtRangeList(buf []byte, dest []RecordExtRange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = RecordExtRange{}
		b += ReadRecordExtRange(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write RecordExtRange
func (v RecordExtRange) Bytes() []byte {
	buf := make([]byte, 6)
	b := 0

	{
		structBytes := v.Major.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.Minor.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	return buf
}

// Write struct list RecordExtRange
func RecordExtRangeListBytes(buf []byte, list []RecordExtRange) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// 'RecordRange' struct definition
// Size: 24
type RecordRange struct {
	CoreRequests    RecordRange8
	CoreReplies     RecordRange8
	ExtRequests     RecordExtRange
	ExtReplies      RecordExtRange
	DeliveredEvents RecordRange8
	DeviceEvents    RecordRange8
	Errors          RecordRange8
	ClientStarted   bool
	ClientDied      bool
}

// Struct read RecordRange
func ReadRecordRange(buf []byte, v *RecordRange) int {
	b := 0

	v.CoreRequests = RecordRange8{}
	b += ReadRecordRange8(buf[b:], &v.CoreRequests)

	v.CoreReplies = RecordRange8{}
	b += ReadRecordRange8(buf[b:], &v.CoreReplies)

	v.ExtRequests = RecordExtRange{}
	b += ReadRecordExtRange(buf[b:], &v.ExtRequests)

	v.ExtReplies = RecordExtRange{}
	b += ReadRecordExtRange(buf[b:], &v.ExtReplies)

	v.DeliveredEvents = RecordRange8{}
	b += ReadRecordRange8(buf[b:], &v.DeliveredEvents)

	v.DeviceEvents = RecordRange8{}
	b += ReadRecordRange8(buf[b:], &v.DeviceEvents)

	v.Errors = RecordRange8{}
	b += ReadRecordRange8(buf[b:], &v.Errors)

	if buf[b] == 1 {
		v.ClientStarted = true
	} else {
		v.ClientStarted = false
	}
	b += 1

	if buf[b] == 1 {
		v.ClientDied = true
	} else {
		v.ClientDied = false
	}
	b += 1

	return b
}

// Struct list read RecordRange
func ReadRecordRangeList(buf []byte, dest []RecordRange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = RecordRange{}
		b += ReadRecordRange(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write RecordRange
func (v RecordRange) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	{
		structBytes := v.CoreRequests.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.CoreReplies.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.ExtRequests.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.ExtReplies.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.DeliveredEvents.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.DeviceEvents.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	{
		structBytes := v.Errors.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}

	if v.ClientStarted {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if v.ClientDied {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// Write struct list RecordRange
func RecordRangeListBytes(buf []byte, list []RecordRange) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// 'RecordClientInfo' struct definition
// Size: (8 + pad((int(NumRanges) * 24)))
type RecordClientInfo struct {
	ClientResource RecordClientSpec
	NumRanges      uint32
	Ranges         []RecordRange // size: pad((int(NumRanges) * 24))
}

// Struct read RecordClientInfo
func ReadRecordClientInfo(buf []byte, v *RecordClientInfo) int {
	b := 0

	v.ClientResource = RecordClientSpec(Get32(buf[b:]))
	b += 4

	v.NumRanges = Get32(buf[b:])
	b += 4

	v.Ranges = make([]RecordRange, v.NumRanges)
	b += ReadRecordRangeList(buf[b:], v.Ranges)

	return b
}

// Struct list read RecordClientInfo
func ReadRecordClientInfoList(buf []byte, dest []RecordClientInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = RecordClientInfo{}
		b += ReadRecordClientInfo(buf[b:], &dest[i])
	}
	return pad(b)
}

// Struct write RecordClientInfo
func (v RecordClientInfo) Bytes() []byte {
	buf := make([]byte, (8 + pad((int(v.NumRanges) * 24))))
	b := 0

	Put32(buf[b:], uint32(v.ClientResource))
	b += 4

	Put32(buf[b:], v.NumRanges)
	b += 4

	b += RecordRangeListBytes(buf[b:], v.Ranges)

	return buf
}

// Write struct list RecordClientInfo
func RecordClientInfoListBytes(buf []byte, list []RecordClientInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += pad(len(structBytes))
	}
	return b
}

// Struct list size RecordClientInfo
func RecordClientInfoListSize(list []RecordClientInfo) int {
	size := 0
	for _, item := range list {
		size += (8 + pad((int(item.NumRanges) * 24)))
	}
	return size
}

// Error definition RecordBadContext (0)
// Size: 32

const BadRecordBadContext = 0

type RecordBadContextError struct {
	Sequence      uint16
	NiceName      string
	InvalidRecord uint32
}

// Error read RecordBadContext
func NewRecordBadContextError(buf []byte) Error {
	v := RecordBadContextError{}
	v.NiceName = "RecordBadContext"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = Get16(buf[b:])
	b += 2

	v.InvalidRecord = Get32(buf[b:])
	b += 4

	return v
}

func (err RecordBadContextError) ImplementsError() {}

func (err RecordBadContextError) SequenceId() uint16 {
	return err.Sequence
}

func (err RecordBadContextError) BadId() Id {
	return 0
}

func (err RecordBadContextError) Error() string {
	fieldVals := make([]string, 0, 1)
	fieldVals = append(fieldVals, "NiceName: "+err.NiceName)
	fieldVals = append(fieldVals, sprintf("Sequence: %d", err.Sequence))
	fieldVals = append(fieldVals, sprintf("InvalidRecord: %d", err.InvalidRecord))
	return "BadRecordBadContext {" + stringsJoin(fieldVals, ", ") + "}"
}

func init() {
	newExtErrorFuncs["RECORD"][0] = NewRecordBadContextError
}

// Request RecordQueryVersion
// size: 8
type RecordQueryVersionCookie struct {
	*cookie
}

func (c *Conn) RecordQueryVersion(MajorVersion uint16, MinorVersion uint16) RecordQueryVersionCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.recordQueryVersionRequest(MajorVersion, MinorVersion), cookie)
	return RecordQueryVersionCookie{cookie}
}

func (c *Conn) RecordQueryVersionUnchecked(MajorVersion uint16, MinorVersion uint16) RecordQueryVersionCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.recordQueryVersionRequest(MajorVersion, MinorVersion), cookie)
	return RecordQueryVersionCookie{cookie}
}

// Request reply for RecordQueryVersion
// size: 12
type RecordQueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	MajorVersion uint16
	MinorVersion uint16
}

// Waits and reads reply data from request RecordQueryVersion
func (cook RecordQueryVersionCookie) Reply() (*RecordQueryVersionReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return recordQueryVersionReply(buf), nil
}

// Read reply into structure from buffer for RecordQueryVersion
func recordQueryVersionReply(buf []byte) *RecordQueryVersionReply {
	v := new(RecordQueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = Get16(buf[b:])
	b += 2

	v.MinorVersion = Get16(buf[b:])
	b += 2

	return v
}

func (cook RecordQueryVersionCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordQueryVersion
func (c *Conn) recordQueryVersionRequest(MajorVersion uint16, MinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put16(buf[b:], MajorVersion)
	b += 2

	Put16(buf[b:], MinorVersion)
	b += 2

	return buf
}

// Request RecordCreateContext
// size: pad(((20 + pad((int(NumClientSpecs) * 4))) + pad((int(NumRanges) * 24))))
type RecordCreateContextCookie struct {
	*cookie
}

// Write request to wire for RecordCreateContext
func (c *Conn) RecordCreateContext(Context Id, ElementHeader RecordElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []RecordClientSpec, Ranges []RecordRange) RecordCreateContextCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.recordCreateContextRequest(Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return RecordCreateContextCookie{cookie}
}

func (c *Conn) RecordCreateContextChecked(Context Id, ElementHeader RecordElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []RecordClientSpec, Ranges []RecordRange) RecordCreateContextCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.recordCreateContextRequest(Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return RecordCreateContextCookie{cookie}
}

func (cook RecordCreateContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordCreateContext
func (c *Conn) recordCreateContextRequest(Context Id, ElementHeader RecordElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []RecordClientSpec, Ranges []RecordRange) []byte {
	size := pad(((20 + pad((int(NumClientSpecs) * 4))) + pad((int(NumRanges) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	buf[b] = byte(ElementHeader)
	b += 1

	b += 3 // padding

	Put32(buf[b:], NumClientSpecs)
	b += 4

	Put32(buf[b:], NumRanges)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		Put32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}
	b = pad(b)

	b += RecordRangeListBytes(buf[b:], Ranges)

	return buf
}

// Request RecordRegisterClients
// size: pad(((20 + pad((int(NumClientSpecs) * 4))) + pad((int(NumRanges) * 24))))
type RecordRegisterClientsCookie struct {
	*cookie
}

// Write request to wire for RecordRegisterClients
func (c *Conn) RecordRegisterClients(Context Id, ElementHeader RecordElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []RecordClientSpec, Ranges []RecordRange) RecordRegisterClientsCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.recordRegisterClientsRequest(Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return RecordRegisterClientsCookie{cookie}
}

func (c *Conn) RecordRegisterClientsChecked(Context Id, ElementHeader RecordElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []RecordClientSpec, Ranges []RecordRange) RecordRegisterClientsCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.recordRegisterClientsRequest(Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return RecordRegisterClientsCookie{cookie}
}

func (cook RecordRegisterClientsCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordRegisterClients
func (c *Conn) recordRegisterClientsRequest(Context Id, ElementHeader RecordElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []RecordClientSpec, Ranges []RecordRange) []byte {
	size := pad(((20 + pad((int(NumClientSpecs) * 4))) + pad((int(NumRanges) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	buf[b] = byte(ElementHeader)
	b += 1

	b += 3 // padding

	Put32(buf[b:], NumClientSpecs)
	b += 4

	Put32(buf[b:], NumRanges)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		Put32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}
	b = pad(b)

	b += RecordRangeListBytes(buf[b:], Ranges)

	return buf
}

// Request RecordUnregisterClients
// size: pad((12 + pad((int(NumClientSpecs) * 4))))
type RecordUnregisterClientsCookie struct {
	*cookie
}

// Write request to wire for RecordUnregisterClients
func (c *Conn) RecordUnregisterClients(Context Id, NumClientSpecs uint32, ClientSpecs []RecordClientSpec) RecordUnregisterClientsCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.recordUnregisterClientsRequest(Context, NumClientSpecs, ClientSpecs), cookie)
	return RecordUnregisterClientsCookie{cookie}
}

func (c *Conn) RecordUnregisterClientsChecked(Context Id, NumClientSpecs uint32, ClientSpecs []RecordClientSpec) RecordUnregisterClientsCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.recordUnregisterClientsRequest(Context, NumClientSpecs, ClientSpecs), cookie)
	return RecordUnregisterClientsCookie{cookie}
}

func (cook RecordUnregisterClientsCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordUnregisterClients
func (c *Conn) recordUnregisterClientsRequest(Context Id, NumClientSpecs uint32, ClientSpecs []RecordClientSpec) []byte {
	size := pad((12 + pad((int(NumClientSpecs) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	Put32(buf[b:], NumClientSpecs)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		Put32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}
	b = pad(b)

	return buf
}

// Request RecordGetContext
// size: 8
type RecordGetContextCookie struct {
	*cookie
}

func (c *Conn) RecordGetContext(Context Id) RecordGetContextCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.recordGetContextRequest(Context), cookie)
	return RecordGetContextCookie{cookie}
}

func (c *Conn) RecordGetContextUnchecked(Context Id) RecordGetContextCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.recordGetContextRequest(Context), cookie)
	return RecordGetContextCookie{cookie}
}

// Request reply for RecordGetContext
// size: (32 + RecordClientInfoListSize(InterceptedClients))
type RecordGetContextReply struct {
	Sequence      uint16
	Length        uint32
	Enabled       bool
	ElementHeader RecordElementHeader
	// padding: 3 bytes
	NumInterceptedClients uint32
	// padding: 16 bytes
	InterceptedClients []RecordClientInfo // size: RecordClientInfoListSize(InterceptedClients)
}

// Waits and reads reply data from request RecordGetContext
func (cook RecordGetContextCookie) Reply() (*RecordGetContextReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return recordGetContextReply(buf), nil
}

// Read reply into structure from buffer for RecordGetContext
func recordGetContextReply(buf []byte) *RecordGetContextReply {
	v := new(RecordGetContextReply)
	b := 1 // skip reply determinant

	if buf[b] == 1 {
		v.Enabled = true
	} else {
		v.Enabled = false
	}
	b += 1

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.ElementHeader = RecordElementHeader(buf[b])
	b += 1

	b += 3 // padding

	v.NumInterceptedClients = Get32(buf[b:])
	b += 4

	b += 16 // padding

	v.InterceptedClients = make([]RecordClientInfo, v.NumInterceptedClients)
	b += ReadRecordClientInfoList(buf[b:], v.InterceptedClients)

	return v
}

func (cook RecordGetContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordGetContext
func (c *Conn) recordGetContextRequest(Context Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// Request RecordEnableContext
// size: 8
type RecordEnableContextCookie struct {
	*cookie
}

func (c *Conn) RecordEnableContext(Context Id) RecordEnableContextCookie {
	cookie := c.newCookie(true, true)
	c.newRequest(c.recordEnableContextRequest(Context), cookie)
	return RecordEnableContextCookie{cookie}
}

func (c *Conn) RecordEnableContextUnchecked(Context Id) RecordEnableContextCookie {
	cookie := c.newCookie(false, true)
	c.newRequest(c.recordEnableContextRequest(Context), cookie)
	return RecordEnableContextCookie{cookie}
}

// Request reply for RecordEnableContext
// size: (32 + pad(((int(Length) * 4) * 1)))
type RecordEnableContextReply struct {
	Sequence      uint16
	Length        uint32
	Category      byte
	ElementHeader RecordElementHeader
	ClientSwapped bool
	// padding: 2 bytes
	XidBase        uint32
	ServerTime     uint32
	RecSequenceNum uint32
	// padding: 8 bytes
	Data []byte // size: pad(((int(Length) * 4) * 1))
}

// Waits and reads reply data from request RecordEnableContext
func (cook RecordEnableContextCookie) Reply() (*RecordEnableContextReply, error) {
	buf, err := cook.reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return recordEnableContextReply(buf), nil
}

// Read reply into structure from buffer for RecordEnableContext
func recordEnableContextReply(buf []byte) *RecordEnableContextReply {
	v := new(RecordEnableContextReply)
	b := 1 // skip reply determinant

	v.Category = buf[b]
	b += 1

	v.Sequence = Get16(buf[b:])
	b += 2

	v.Length = Get32(buf[b:]) // 4-byte units
	b += 4

	v.ElementHeader = RecordElementHeader(buf[b])
	b += 1

	if buf[b] == 1 {
		v.ClientSwapped = true
	} else {
		v.ClientSwapped = false
	}
	b += 1

	b += 2 // padding

	v.XidBase = Get32(buf[b:])
	b += 4

	v.ServerTime = Get32(buf[b:])
	b += 4

	v.RecSequenceNum = Get32(buf[b:])
	b += 4

	b += 8 // padding

	v.Data = make([]byte, (int(v.Length) * 4))
	copy(v.Data[:(int(v.Length)*4)], buf[b:])
	b += pad(int((int(v.Length) * 4)))

	return v
}

func (cook RecordEnableContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordEnableContext
func (c *Conn) recordEnableContextRequest(Context Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// Request RecordDisableContext
// size: 8
type RecordDisableContextCookie struct {
	*cookie
}

// Write request to wire for RecordDisableContext
func (c *Conn) RecordDisableContext(Context Id) RecordDisableContextCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.recordDisableContextRequest(Context), cookie)
	return RecordDisableContextCookie{cookie}
}

func (c *Conn) RecordDisableContextChecked(Context Id) RecordDisableContextCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.recordDisableContextRequest(Context), cookie)
	return RecordDisableContextCookie{cookie}
}

func (cook RecordDisableContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordDisableContext
func (c *Conn) recordDisableContextRequest(Context Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// Request RecordFreeContext
// size: 8
type RecordFreeContextCookie struct {
	*cookie
}

// Write request to wire for RecordFreeContext
func (c *Conn) RecordFreeContext(Context Id) RecordFreeContextCookie {
	cookie := c.newCookie(false, false)
	c.newRequest(c.recordFreeContextRequest(Context), cookie)
	return RecordFreeContextCookie{cookie}
}

func (c *Conn) RecordFreeContextChecked(Context Id) RecordFreeContextCookie {
	cookie := c.newCookie(true, false)
	c.newRequest(c.recordFreeContextRequest(Context), cookie)
	return RecordFreeContextCookie{cookie}
}

func (cook RecordFreeContextCookie) Check() error {
	return cook.check()
}

// Write request to wire for RecordFreeContext
func (c *Conn) recordFreeContextRequest(Context Id) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.extensions["RECORD"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}
