package main

import "encoding/xml"

/*
Message the message element
*/
type Message struct {
	XMLName xml.Name `xml:"message"`
	Header  Header   `xml:"header"`
	Body    Body     `xml:"body"`
}

/*
Header the header element
*/
type Header struct {
	XMLName xml.Name `xml:"header"`

	MessageClass  string `xml:"messageClass,attr"`
	TransactionID string `xml:"transactionID,attr"`
	Reply         string `xml:"reply,attr"`

	Location Location `xml:"location"`
}

/*
Location the location element
*/
type Location struct {
	XMLName xml.Name `xml:"location"`

	RouteID       string `xml:"routeID,attr"`
	RouteName     string `xml:"routeName,attr"`
	EquipmentID   string `xml:"equipmentID,attr"`
	EquipmentName string `xml:"equipmentName,attr"`
	ZoneID        string `xml:"zoneID,attr"`
	ZonePos       string `xml:"zonePos,attr"`
	ZoneName      string `xml:"zoneName,attr"`
	LaneNo        string `xml:"laneNo,attr"`
	ControllerG   string `xml:"controllerGuid,attr"`
}

/*
Body the body element
*/
type Body struct {
	XMLName xml.Name `xml:"body"`

	Pcb *Pcb `xml:"pcb,omitempty"`

	Panel *Panel `xml:"panel,omitempty"`

	ModelInfo *ModelInfo `xml:"modelInfo,omitempty"`
	Product   *Product   `xml:"product,omitempty"`
	BOM       *BOM       `xml:"BOM,omitempty"`
	Resource  *Resource  `xml:"resource,omitempty"`

	Result *Result `xml:"result,omitempty"`

	PCBAssignReq *PCBAssignReq `xml:"PCBAssignReq,omitempty"`
}

/*
Pcb the pcb element
*/
type Pcb struct {
	XMLName xml.Name `xml:"pcb"`

	Barcode          string `xml:"barcode,attr,omitempty"`
	ModelCode        string `xml:"modelCode,attr,omitempty"`
	SerialNo         string `xml:"serialNo,attr,omitempty"`
	PcbSide          string `xml:"pcbSide,attr,omitempty"`
	ScannerMountSide string `xml:"scannerMountSide,attr,omitempty"`
}

/*
Panel the panel element
*/
type Panel struct {
	XMLName xml.Name `xml:"panel"`

	PcbID       string `xml:"pcbID,attr,omitempty"`
	State       string `xml:"state,attr,omitempty"`
	Timestamp   string `xml:"timestamp,attr,omitempty"`
	ProductSide string `xml:"productSide,attr,omitempty"`
	RepairType  string `xml:"repairType,attr,omitempty"`
	ProductName string `xml:"productName,attr,omitempty"`

	SubPanel []SubPanel `xml:"subPanel,omitempty"`
}

/*
SubPanel the subPanel element
*/
type SubPanel struct {
	XMLName xml.Name `xml:"subPanel"`

	State string `xml:"state,attr,omitempty"`
	Pos   string `xml:"pos,attr,omitempty"`
}

/*
ModelInfo the modelInfo element
*/
type ModelInfo struct {
	XMLName xml.Name `xml:"modelInfo"`

	ModelNo string `xml:"modelNo,attr,omitempty"`
	Side    string `xml:"side,attr,omitempty"`
}

/*
Product the product element
*/
type Product struct {
	XMLName xml.Name `xml:"product"`

	Name string `xml:"name,attr,omitempty"`
	Side string `xml:"side,attr,omitempty"`
}

/*
BOM the BOM element
*/
type BOM struct {
	XMLName xml.Name `xml:"BOM"`

	Parts *Parts `xml:"parts,omitempty"`
	Comps *Comps `xml:"comps,omitempty"`
}

/*
Parts the Parts element
*/
type Parts struct {
	XMLName xml.Name `xml:"parts"`

	List []Part `xml:"part,omitempty"`
}

/*
Part the Part element
*/
type Part struct {
	XMLName xml.Name `xml:"part"`

	Name string `xml:"name,attr,omitempty"`
	Ref  string `xml:"ref,attr,omitempty"`
}

/*
Comps the Comps element
*/
type Comps struct {
	XMLName xml.Name `xml:"comps"`
	List    []Comp   `xml:"comp,omitempty"`
}

/*
Comp the Comp element
*/
type Comp struct {
	XMLName xml.Name `xml:"comp"`

	Pos    string `xml:"pos,attr,omitempty"`
	Ref    string `xml:"ref,attr,omitempty"`
	RefDes string `xml:"refDes,attr,omitempty"`
}

/*
Resource the resource element
*/
type Resource struct {
	XMLName xml.Name `xml:"resource"`
	Jigs    *Jigs    `xml:"jigs,omitempty"`
	Comps   *Comps   `xml:"comps,omitempty"`
}

/*
Jigs the jigs element
*/
type Jigs struct {
	XMLName xml.Name `xml:"jigs"`

	List []Jig `xml:"jig,omitempty"`
}

/*
Jig the jig element
*/
type Jig struct {
	XMLName xml.Name `xml:"jig"`

	Name string `xml:"name,attr,omitempty"`
	Ref  string `xml:"ref,attr,omitempty"`
}

/*
Result the result element
*/
type Result struct {
	XMLName xml.Name `xml:"result"`

	ErrorCode  string `xml:"errorCode,attr,omitempty"`
	ErrorText  string `xml:"errorText,attr,omitempty"`
	ActionCode string `xml:"actionCode,attr,omitempty"`
}

/*
PCBAssignReq the 10000 req element
*/
type PCBAssignReq struct {
	XMLName xml.Name `xml:"PCBAssignReq"`

	Manufacturelotno string `xml:"manufacturelotno,attr,omitempty"`
	Subpanelcount    string `xml:"subpanelcount,attr,omitempty"`
}

/*
MPPCBTagReq the 10000 response element
*/
type MPPCBTagReq struct {
	XMLName xml.Name `xml:"MPPCBTagReq"`

	Barcode          string `xml:"barcode,attr,omitempty"`
	Manufacturelotno string `xml:"manufacturelotno,attr,omitempty"`
	Timestamp        string `xml:"timestamp,attr,omitempty"`
	Subpanelcount    string `xml:"subpanelcount,attr,omitempty"`
	Position         string `xml:"position,attr,omitempty"`
	Result           string `xml:"result,attr,omitempty"`
	Camceraenable    string `xml:"camceraenable,attr,omitempty"`
}

/*
SPPCBTagReq the 10000 response element
*/
type SPPCBTagReq struct {
	XMLName xml.Name `xml:"SPPCBTagReq"`

	Barcode       string `xml:"barcode,attr,omitempty"`
	Position      string `xml:"position,attr,omitempty"`
	Result        string `xml:"result,attr,omitempty"`
	Camceraenable string `xml:"camceraenable,attr,omitempty"`
}
