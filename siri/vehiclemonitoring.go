package siri

import (
	"time"
)

type ProducerRef struct {
	Value string `json:"value"`
}
type ProgressBetweenStops struct {
	LinkDistance int     `json:"LinkDistance"`
	Percentage   float64 `json:"Percentage"`
}
type LineRef struct {
	Value string `json:"value"`
}
type DirectionRef struct {
	Value string `json:"value"`
}
type DataFrameRef struct {
	Value string `json:"value"`
}
type FramedVehicleJourneyRef struct {
	DataFrameRef           DataFrameRef `json:"DataFrameRef"`
	DatedVehicleJourneyRef string       `json:"DatedVehicleJourneyRef"`
}
type JourneyPatternRef struct {
	Value string `json:"value"`
}
type PublishedLineName struct {
	Value string `json:"value"`
}
type OperatorRef struct {
	Value string `json:"value"`
}
type OriginRef struct {
	Value string `json:"value"`
}
type OriginName struct {
	Value string `json:"value"`
}
type DestinationRef struct {
	Value string `json:"value"`
}
type DestinationName struct {
	Value string `json:"value"`
}
type VehicleLocation struct {
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
	SrsName   string  `json:"srsName"`
}
type BlockRef struct {
	Value string `json:"value"`
}
type VehicleRef struct {
	Value string `json:"value"`
}
type StopPointRef struct {
	Value string `json:"value"`
}
type StopPointName struct {
	Value string `json:"value"`
}
type PreviousCall struct {
	StopPointRef  StopPointRef    `json:"StopPointRef"`
	VisitNumber   int             `json:"VisitNumber"`
	StopPointName []StopPointName `json:"StopPointName"`
}
type PreviousCalls struct {
	PreviousCall []PreviousCall `json:"PreviousCall"`
}
type DestinationDisplay struct {
	Value string `json:"value"`
}
type MonitoredCall struct {
	VehicleAtStop      bool                 `json:"VehicleAtStop"`
	DestinationDisplay []DestinationDisplay `json:"DestinationDisplay"`
	StopPointRef       StopPointRef         `json:"StopPointRef"`
	VisitNumber        int                  `json:"VisitNumber"`
	StopPointName      []StopPointName      `json:"StopPointName"`
}
type OnwardCall struct {
	StopPointRef  StopPointRef    `json:"StopPointRef"`
	VisitNumber   int             `json:"VisitNumber"`
	StopPointName []StopPointName `json:"StopPointName"`
}
type OnwardCalls struct {
	OnwardCall []OnwardCall `json:"OnwardCall"`
}

type MonitoredVehicleJourney struct {
	LineRef                     LineRef                 `json:"LineRef"`
	DirectionRef                DirectionRef            `json:"DirectionRef"`
	FramedVehicleJourneyRef     FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef"`
	JourneyPatternRef           JourneyPatternRef       `json:"JourneyPatternRef"`
	PublishedLineName           []PublishedLineName     `json:"PublishedLineName"`
	OperatorRef                 OperatorRef             `json:"OperatorRef"`
	OriginRef                   OriginRef               `json:"OriginRef"`
	OriginName                  []OriginName            `json:"OriginName"`
	DestinationRef              DestinationRef          `json:"DestinationRef"`
	DestinationName             []DestinationName       `json:"DestinationName"`
	HeadwayService              bool                    `json:"HeadwayService"`
	OriginAimedDepartureTime    time.Time               `json:"OriginAimedDepartureTime"`
	DestinationAimedArrivalTime time.Time               `json:"DestinationAimedArrivalTime"`
	Monitored                   bool                    `json:"Monitored"`
	InCongestion                bool                    `json:"InCongestion"`
	InPanic                     bool                    `json:"InPanic"`
	DataSource                  string                  `json:"DataSource"`
	VehicleLocation             VehicleLocation         `json:"VehicleLocation"`
	Delay                       string                  `json:"Delay"`
	BlockRef                    BlockRef                `json:"BlockRef"`
	VehicleRef                  VehicleRef              `json:"VehicleRef"`
	MonitoredCall               MonitoredCall           `json:"MonitoredCall"`
	OnwardCalls                 OnwardCalls             `json:"OnwardCalls"`
}

type VehicleActivity struct {
	ValidUntilTime          time.Time                 `json:"ValidUntilTime"`
	ProgressBetweenStops    ProgressBetweenStops      `json:"ProgressBetweenStops"`
	RecordedAtTime          time.Time                 `json:"RecordedAtTime"`
	ItemIdentifier          string                    `json:"ItemIdentifier"`
	MonitoredVehicleJourney []MonitoredVehicleJourney `json:"MonitoredVehicleJourney,omitempty"`
}

type VehicleMonitoringDelivery struct {
	VehicleActivity   []VehicleActivity `json:"VehicleActivity"`
	ResponseTimestamp time.Time         `json:"ResponseTimestamp"`
	Version           string            `json:"version"`
}
type ServiceDelivery struct {
	ResponseTimestamp         time.Time                   `json:"ResponseTimestamp"`
	ProducerRef               ProducerRef                 `json:"ProducerRef"`
	MoreData                  bool                        `json:"MoreData"`
	VehicleMonitoringDelivery []VehicleMonitoringDelivery `json:"VehicleMonitoringDelivery"`
}
type Siri struct {
	ServiceDelivery ServiceDelivery `json:"ServiceDelivery"`
	Version         string          `json:"version"`
}
