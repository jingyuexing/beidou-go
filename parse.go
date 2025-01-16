package beidou

import (
	"errors"
	"strconv"
)


var ErrMessageTypeValid = errors.New("message type valid")
var ErrFieldsTypeValid = errors.New("field type valid")

// ParseSOSMessage 使用 Tokenizer 解析消息
func ParseSOSMessage(tokens ...Token) (*SOSMessage, error) {

	// 创建 SOSMessage 结构体实例
	message := new(SOSMessage)
	tokenList := Filter(tokens,func(token Token, index int) bool {
		return token.Type != Spiliter && token.Type != End && token.Type != Begin
	})
	if !(tokenList[0].Value == "BDEPI") {
		return nil,ErrMessageTypeValid
	}
	for index,token := range tokenList {
		switch index {
		case 1:
			senderID, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.SenderID = senderID
		case 2:
			receiverID, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.ReceiverID = receiverID
		case 3:
			serviceTypeEmergencyFlag, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.ServiceTypeEmergencyFlag = serviceTypeEmergencyFlag
		case 4:
			message.Time = token.Value
		case 5:
			longitude, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.Longitude = longitude
		case 6:
			message.LongitudeSign = token.Value
		case 7:
			latitude, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.Latitude = latitude
		case 8:
			message.LatitudeSign = token.Value
		case 9:
			altitude, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.Altitude = altitude
		case 10:
			rescueType, err := strconv.Atoi(tokenList[10].Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.RescueType = rescueType
		case 11:
			rescueCenterFlag, err := strconv.Atoi(tokenList[11].Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.RescueCenterFlag = rescueCenterFlag
		case 12:
			locationReportDataLength, err := strconv.Atoi(tokenList[12].Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.LocationReportDataLength = locationReportDataLength
		case 13:
			message.LocationReportData = token.Value
		case 15:
			message.CheckSum = token.Value
		}
	}
	return message, nil
}

// ParseCommunicationMessage 解析通信报文数据
func ParseCommunicationMessage(tokens ...Token) (*CommunicationMessage, error) {
	message := new(CommunicationMessage)
	tokenList := Filter(tokens,func(token Token, index int) bool {
		return token.Type != Spiliter && token.Type != End && token.Type != Begin
	})
	if (tokenList[0].Value != "BDTCI") {
		return nil,ErrMessageTypeValid
	}
	for index,token := range tokenList {
		switch index {
		case 1:
			senderID, err := strconv.ParseUint(token.Value, 10, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.SenderID = senderID
		case 2:
			// 收信方 ID
			receiverID, err := strconv.ParseUint(token.Value, 10, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.ReceiverID = receiverID
		case 3:
			// 时间
			message.Time = token.Value
		case 4:
			// 编码类别
			encodingType, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.EncodingType = encodingType
		case 5:
			// 数据标识
			dataFlag, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.DataFlag = dataFlag
		case 6:
			// 通信长度
			commLength, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.CommLength = commLength
		case 7:
			// 通信数据
			message.CommData = token.Value
		case 9:
			message.Checksum = token.Value;
		}
	}
	
	// 返回解析后的 CommunicationMessage
	return message, nil
}


func ParseCommunicationReceiptMessage(tokens ...Token) (*CommunicationReceiptMessage, error) {
	message := new(CommunicationReceiptMessage)
	tokenList := Filter(tokens,func(token Token, index int) bool {
		return token.Type != Spiliter && token.Type != End && token.Type != Begin
	})
	if (tokenList[0].Value != "BDRTI") {
		return nil,ErrMessageTypeValid
	}
	for index := 0;index <= 3;index++ {
		token := tokenList[index]
		switch index {
		case 1:
			senderID, err := strconv.ParseUint(token.Value, 10, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.SenderID = senderID

		case 2:
			// Parse Receiver ID (Token 2)
			receiverID, err := strconv.ParseUint(token.Value, 10, 64)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.ReceiverID = receiverID
		case 3:
			receiptCount, err := strconv.Atoi(token.Value)
			if err != nil {
				return nil, ErrFieldsTypeValid
			}
			message.ReceiptCount = receiptCount
		}
	}
	times := make([]string,message.ReceiptCount)
	for i := 4;i < message.ReceiptCount;i += 2 {
		times[i] = tokenList[i].Value
		times[i+1] = tokenList[i+1].Value
	}
	message.AdditionalTimes = times

	return message,nil
}