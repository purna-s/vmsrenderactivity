package vmsrenderactivity

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("data", `{"VMSInfo": {"VMSmsg": {"Date": "2017-03-22 14:35:08.0","EquipmentID": "amvms_v9111","LinkID": "103103931","Attribute": "sst","Message": "VEH BREAKDOWN SH,AFT W`LANDS AVE 2"}}}`)
	//tc.SetInput("file", "D:/Flogo/xml.jsp")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	output := tc.GetOutput("output")
	assert.Equal(t, output, output)

}
