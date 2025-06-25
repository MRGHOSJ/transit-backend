# 📦 Package: `model`

> 📍 `C:\Users\DELL\Desktop\transit-backend-main\internal\transport\model\transport.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (1)](#-structs)

## 🧱 Structs

### `Connection`

```go
type Connection struct {
	StationKey string json:"station_id
	Distance float64 json:"distance
	Type TransportType json:"type
	Line string json:"line,omitempty
	Duration float64 json:"duration
}
```

**Summary:** Defines a Connection struct linking stations with distance and duration.

**Parameters:**
- `StationKey` (string): Connected station identifier
- `Distance` (float64): Distance between stations
- `Type` (TransportType): Connection transport type
- `Line` (string): Transport line (optional)
- `Duration` (float64): Travel duration in minutes

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
conn := Connection{StationKey: "s123", Distance: 1.2, Duration: 5.0}
```

**Edge Cases:**
- Zero/negative Distance/Duration
- Empty StationKey
- Invalid TransportType


---

