package apisspanel_test

import (
	"testing"

	. "v2ray.com/core/common/apisspanelsspanel"
)

func CreateClient() *ApiClient {
	client := NewClient("http://127.0.0.1:667", 41, "NimaQu")

	return client
}

func TestApiClient_GetNodeInfost(t *testing.T) {
	client := CreateClient()
	nodeInfo, err := client.GetNodeInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(nodeInfo)
}

func TestApiClient_GetUserList(t *testing.T) {
	client := CreateClient()
	nodeInfo, err := client.GetUserList()
	if err != nil {
		t.Error(err)
	}
	t.Log(nodeInfo)
}
func TestApiClient_ReportNodeStatus(t *testing.T) {
	client := CreateClient()
	err := client.ReportNodeStatus(&NodeStatus{
		CPU:    "1",
		Mem:    "1",
		Net:    "1",
		Disk:   "1",
		Uptime: 0,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestApiClient_ReportNodeOnline(t *testing.T) {
	client := CreateClient()
	online := []*NodeOnline{
		&NodeOnline{
			1,
			"1.1.1.1",
		},
		&NodeOnline{
			2,
			"1.1.1.2",
		},
	}
	err := client.ReportNodeOnline(online)
	if err != nil {
		t.Error(err)
	}
}

func TestApiClient_ReportUserTraffic(t *testing.T) {
	client := CreateClient()
	traffic := []*UserTraffic{
		&UserTraffic{
			1,
			1,
			1,
		},
		&UserTraffic{
			2,
			2,
			2,
		},
	}
	err := client.ReportUserTraffic(traffic)
	if err != nil {
		t.Error(err)
	}
}
