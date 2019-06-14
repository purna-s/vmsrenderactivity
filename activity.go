package vmsrenderactivity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-vmsrenderactivity")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//XSD
type VMSInfo struct {
	XMLName      xml.Name   `xml:"VMSInfo" json:"-"`
	VMSmsgList []VMSmsg `xml:"VMSmsg" json:"VMSmsg"`
}

type VMSmsg struct {
	XMLName xml.Name `xml:"VMSmsg" json:"-"`
	Date   string   `xml:"Date" json:"Date"`
	EquipmentID   string   `xml:"EquipmentID" json:"EquipmentID"`
	LinkID    string   `xml:"LinkID" json:"LinkID"`
	Attribute    string   `xml:"Attribute" json:"Attribute"`	
	Message string   `xml:"Message" json:"Message"`
}

// end of XSD

// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {

	JsonString := ctx.GetInput("data").(string)

	activityLog.Debugf("Json String is : [%s]", JsonString)
	//fmt.Println("Json String is : ", JsonString)

	if len(JsonString) == 0 {
		activityLog.Debugf("value in  the field is empty ")
		fmt.Println("value in  the field is empty ")
		return
	}
	//fmt.Println(" JSON String " + JsonString)

	jdata := VMSInfo{}
	err = json.Unmarshal([]byte(JsonString), &jdata)
	xmlData, _ := xml.Marshal(jdata)

	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}

	//fmt.Println(" XML String ")
	//fmt.Println(string(xmlData))

	// Set the output as part of the context
	activityLog.Debugf("Activity has rendered VMS json Successfully")
	//fmt.Println("Activity has rendered VMS json Successfully")

	ctx.SetOutput("output", string(xmlData))

	return true, nil
}
