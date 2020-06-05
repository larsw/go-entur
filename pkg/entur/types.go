package entur

import (
  "time"
)

type SiriResponse struct {
  Siri Siri `json:Siri`
}

type Siri struct {
  ServiceDelivery *ServiceDelivery `json:ServiceDelivery`
  Version string `json:version`
}

type ServiceDelivery struct {

  ResponseTimestamp time.Time `json:ResponseTimestamp`

  ProducerRef struct {
    Value string `json:value`
  } `json:ProducerRef`

  RequestMessageRef struct {
    Value string `json:value` // guid?
  } `json:RequestMessageRef`

  MoreData bool `json:MoreData`

  VehicleMonitoringDelivery []VehicleMonitoringDelivery `json:VehicleMonitoringDelivery`
}


type VehicleMonitoringDelivery struct {
  ValidUntilTime time.Time `json:ValidUntilTime`
  ProgressBetweenStops struct {
    LinkDistance int32 `json:LinkDistance`
    Percentage float32 `json:Percentage`
  } `json:ProgressBetweenStops`

  MonitoredVehicleJourney []*MonitoredVehicleJourney `json:MonitoredVehicleJourney`

  RecordedAtTime time.Time `json:RecordedAtTime`
  ItemIdentifier string `json:ItemIdentifier`
}

type MonitoredVehicleJourney struct {
  LineRef struct {
    Value string `json:value`
  } `json:LineRef`

  DirectionRef struct {
    Value string `json:value`
  } `json:DirectionRef`

  FramedVehicleJourneyRef struct {
    DataFrameRef struct {
      Value string `json:value`
    } `json:DataFrameRef`
    DatedVehicleJourneyRef string `json:DatedVehicleJourneyRef`
  } `json:FramedVehicleJourneyRef`

  JourneyPatternRef struct {
    Value string `json:value`
  } `json:JourneyPatternRef`

  PublishedLineName []struct {
    Value string `json:value`
  } `json:PublishedLineName`

  OperatorRef struct {
    Value string `json:value`
  } `json:OperatorRef`

  OriginRef struct {
    Value string `json:value`
  } `json:OriginRef`

  OriginName struct {
    Value string `json:value`
  } `json:OriginName`

  DestinationRef struct {
    Value string `json:value`
  } `json:DestinationRef`

  DestinationName []struct {
    Value string `json:value`
  } `json:DestinationName`

  HeadwayService bool `json:HeadwayService`

  OriginAimedDepartureTime time.Time `json:OriginAimedDepartureTime`

  DestinationAimedArrivalTime time.Time `json:DestinationAimedArrivalTime`

  Monitored bool `json:Monitored`

  InCongestion bool `json:InCongestion`

  InPanic bool `json:InPanic`

  DataSource string `json:DataSource`

  VehicleLocation struct {
    Longitude float32 `json:Longitude`
    Latitude float32 `json:Latitude`
    SrsName string `json:srsName`
  } `json:VehicleLocation`

  Delay string `json:Delay`

  BlockRef struct {
    Value string `json:BlockRef`
  } `json:BlockRef`

  VehicleRef struct {
    Value string `json:value`
  } `json:VehicleRef`

  PreviousCalls struct {
    PreviousCall []struct {
      StopPointRef struct {
        value string `json:StopPointRef`
      }
      VisitNumber int32 `json:VisitNumber`
      StopPointName []string `json:StopPointName`
    } `json:PreviousCall`
  } `json:PreviousCalls`

  MonitoredCall struct {
    VehicleAtStop bool `json:VehicleAtStop`
    DestinationDisplay []struct {
      Value string `json:value`
    } `json:DestinationDisplay`
    StopPointRef struct {
      value string `json:StopPointRef`
    }
    VisitNumber int32 `json:VisitNumber`
    StopPointName []string `json:StopPointName`
  } `json:MonitoredCall`

  OnwardCalls struct {
    OnwardCall []* struct {
      StopPointRef struct {
        Value string `json:value`
      }
      VisitNumber int32 `json:VisitNumber`
      StopPointName []* struct {
        Value string `json:value`
      }
    } `json:OnwardCall`
  } `json:OnwardCalls`
}

