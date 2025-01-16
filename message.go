package beidou

// SOSMessage 定义解析后的数据结构
type SOSMessage struct {
	SenderID                 int     `json:"sender_id"`                   // 发信方 ID
	ReceiverID               int     `json:"receiver_id"`                 // 收信方 ID
	ServiceTypeEmergencyFlag int     `json:"service_type_emergency_flag"` // 服务类型/紧急标识
	Time                     string  `json:"time"`                        // 时间
	Longitude                float64 `json:"longitude"`                   // 经度
	LongitudeSign            string  `json:"longitude_sign"`              // 经度标志
	Latitude                 float64 `json:"latitude"`                    // 纬度
	LatitudeSign             string  `json:"latitude_sign"`               // 纬度标志
	Altitude                 float64 `json:"altitude"`                    // 高程
	RescueType               int     `json:"rescue_type"`                 // 搜救类型
	RescueCenterFlag         int     `json:"rescue_center_flag"`          // 搜救中心指示类型
	RescueDataLength         int     `json:"rescue_data_length"`          // 搜救业务数据长度
	LocationReportDataLength int     `json:"location_report_data_length"` // 位置报告状态数据长度
	LocationReportData       string  `json:"location_report_data"`        // 位置报告状态数据
	CheckSum string `json:"checksum"`
}

type CommunicationReceiptMessage struct {
	SenderID             uint64   `json:"sender_id"`             // 发信方 ID
	ReceiverID           uint64   `json:"receiver_id"`           // 收信方 ID
	ReceiptCount         int      `json:"receipt_count"`         // 回执数 (0-7)
	CommunicationTime1   string   `json:"comm_time_1"`           // 通信发信时间1（年月日）
	CommunicationTime2   string   `json:"comm_time_2"`           // 通信发信时间2（时分秒）
	AdditionalTimes      []string `json:"additional_times"`      // 其它的通信时间（年月日、时分秒）
}

// CommunicationMessage 定义解析后的通信报文数据结构
type CommunicationMessage struct {
	SenderID     uint64 `json:"sender_id"`     // 发信方 ID
	ReceiverID   uint64 `json:"receiver_id"`   // 收信方 ID
	Time         string `json:"time"`          // 时间 (hhmmss)
	EncodingType int    `json:"encoding_type"` // 编码类别 (0-3)
	DataFlag     int    `json:"data_flag"`     // 数据标识 (0-1)
	CommLength   int    `json:"comm_length"`   // 通信长度 (比特数)
	CommData     string `json:"comm_data"`     // 通信数据
	Checksum string `json:"checksum"`
}

type Option[T any] func(types *T)

// WithSOSMessageSenderID 设置 SenderID
func WithSOSMessageSenderID(id int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.SenderID = id
	}
}

// WithSOSMessageReceiverID 设置 ReceiverID
func WithSOSMessageReceiverID(id int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.ReceiverID = id
	}
}

// WithSOSMessageServiceTypeEmergencyFlag 设置服务类型/紧急标识
func WithSOSMessageServiceTypeEmergencyFlag(flag int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.ServiceTypeEmergencyFlag = flag
	}
}

// WithSOSMessageTime 设置时间
func WithSOSMessageTime(time string) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.Time = time
	}
}

// WithSOSMessageLongitude 设置经度
func WithSOSMessageLongitude(longitude float64) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.Longitude = longitude
	}
}

// WithSOSMessageLongitudeSign 设置经度标志
func WithSOSMessageLongitudeSign(sign string) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.LongitudeSign = sign
	}
}

// WithSOSMessageLatitude 设置纬度
func WithSOSMessageLatitude(latitude float64) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.Latitude = latitude
	}
}

// WithSOSMessageLatitudeSign 设置纬度标志
func WithSOSMessageLatitudeSign(sign string) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.LatitudeSign = sign
	}
}

// WithSOSMessageAltitude 设置高程
func WithSOSMessageAltitude(altitude float64) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.Altitude = altitude
	}
}

// WithSOSMessageRescueType 设置搜救类型
func WithSOSMessageRescueType(rescueType int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.RescueType = rescueType
	}
}

// WithSOSMessageRescueCenterFlag 设置搜救中心指示类型
func WithSOSMessageRescueCenterFlag(flag int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.RescueCenterFlag = flag
	}
}

// WithSOSMessageRescueDataLength 设置搜救业务数据长度
func WithSOSMessageRescueDataLength(length int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.RescueDataLength = length
	}
}

// WithSOSMessageLocationReportDataLength 设置位置报告状态数据长度
func WithSOSMessageLocationReportDataLength(length int) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.LocationReportDataLength = length
	}
}

// WithSOSMessageLocationReportData 设置位置报告状态数据
func WithSOSMessageLocationReportData(data string) Option[SOSMessage] {
	return func(s *SOSMessage) {
		s.LocationReportData = data
	}
}
