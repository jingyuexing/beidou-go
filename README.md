# BeiDou Short Message Parser Module (北斗短报文解析)

## Translate

[Chinese](./README_zh-CN.md)

## Overview

This module is designed for parsing communication receipt query results from BeiDou short messages. According to the BeiDou short message communication protocol, the message format includes sender ID, receiver ID, receipt count, multiple time fields (year-month-day, hour-minute-second), and other communication information. This module extracts this information from raw messages and organizes it into easy-to-use structures.

## Main Features

- **Communication Receipt Query Message Parsing**: Parse contents including sender, receiver, receipt count, and time information according to BeiDou communication protocol.
- **Time Field Processing**: Support parsing of various time formats, including year-month-day and hour-minute-second, capable of handling multiple time period information.
- **Data Encapsulation**: Encapsulate parsed data in the `CommunicationReceiptMessage` structure for convenient subsequent use.

## Message Format

The message format for BeiDou communication receipt query results is as follows:

```
$BDRTI,x.x,x.x,x,xxxxxxxx,hhmmss,......,xxxxxxxx,hhmmss*hh<CR><LF>
```

### Statement Format Description

| No. | Name                     | Value Range     | Unit     | Description                            |
|-----|--------------------------|-----------------|----------|----------------------------------------|
| 1   | Sender ID               | 1-16,777,215    | --       | Query initiator's address              |
| 2   | Receiver ID             | 1-16,777,215    | --       | Original communication receiver address |
| 3   | Receipt Count           | 0-7             | --       | Receipt count (0 means no receipt)     |
| 4   | Comm Time 1 (YMD)       | Year Month Day  | --       | Example: 20210409                      |
| 5   | Comm Time 2 (HMS)       | Hour Min Sec    | --       | Sender's station entry time            |
| 6   | ...                     | --              | --       | Other time fields (YMD, HMS)           |
| 7   | Comm Time 7 (YMD)       | Year Month Day  | --       | Last communication time field          |
| 8   | Comm Time 7 (HMS)       | Hour Min Sec    | --       | Last communication time field          |

## Structure Definition

## `CommunicationReceiptMessage` Structure

CommunicationReceiptMessage represents the parsed communication receipt query result information. This structure contains sender ID, receiver ID, receipt count, communication time, and other time information.

```go
type CommunicationReceiptMessage struct {
    SenderID             uint64   `json:"sender_id"`             // Sender ID
    ReceiverID           uint64   `json:"receiver_id"`           // Receiver ID
    ReceiptCount         int      `json:"receipt_count"`         // Receipt count (0-7)
    CommunicationTime1   string   `json:"comm_time_1"`           // Communication time 1 (YMD)
    CommunicationTime2   string   `json:"comm_time_2"`           // Communication time 2 (HMS)
    AdditionalTimes      []string `json:"additional_times"`      // Additional communication times (YMD, HMS)
}
```

## Dependencies

This module depends on Go standard library, particularly the strconv and strings packages for string processing and type conversion.

## Error Handling

If the number of tokens is insufficient to parse the complete message, the parsing function will return an error.
If a token cannot be parsed correctly (e.g., incorrect number format), the parsing function will also return an error.

## Extended Functionality

This module currently only parses communication receipt query messages. Future extensions could support more types of BeiDou short messages, such as emergency rescue information, position report information, etc.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please submit an issue or create a pull request.
