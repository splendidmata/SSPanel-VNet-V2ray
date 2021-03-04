package apisspanel

import (
	"encoding/json"
	"fmt"
	_ "log"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	client  *resty.Client
	ApiHost string
	NodeID  int
	Key     string
}

func NewClient(apiHost string, nodeId int, key string) *ApiClient {
	apiClient := new(ApiClient)
	client := resty.New()
	// For Debug use
	// client.SetDebug(true)
	client.SetTimeout(5 * time.Second)
	client.SetHostURL(apiHost)
	// Create Key for each requests
	client.SetQueryParam("key", key)
	apiClient.client = client
	apiClient.NodeID = nodeId
	apiClient.Key = key
	apiClient.ApiHost = apiHost
	return apiClient
}


// GetNodeInfo will pull NodeInfo Config from sspanel
func (c *ApiClient) GetNodeInfo() (data *NodeInfo, err error) {
	path := fmt.Sprintf("/mod_mu/nodes/%d/info", c.NodeID)
	res, err := c.client.R().
		SetResult(&NodeinfoResponse{}).
		ForceContentType("application/json").
		Get(path)
	if err != nil {
		return nil, newError(fmt.Sprintf("request %s failed", c.AssembleUrl(path))).Base(err)
	}

	if res.StatusCode() > 400 {
		body := res.Body()
		return nil, newError(fmt.Sprintf("request %s failed: %s", c.AssembleUrl(path), string(body))).Base(err)
	}

	result := res.Result().(*NodeinfoResponse)

	if result.Ret != 1 {
		res, _ := json.Marshal(&result)
		return nil, newError(fmt.Sprintf("Ret %s invalid", string(res)))
	}
	// Parse node info into VNET mode
	nodeInfo := result.Data
	nodeInfo.Server_raw = strings.ToLower(nodeInfo.Server_raw)
	serverConf := strings.Split(nodeInfo.Server_raw, ";")

	V2Port, err := strconv.Atoi(serverConf[1])
	V2AlterID, err := strconv.Atoi(serverConf[2])
	V2Net := serverConf[3]
	var V2TLS bool
	if serverConf[4] == "tls" {
		V2TLS = true
	} else {
		V2TLS = false
	}
	extraServerConf := strings.Split(serverConf[5], "|")
	var V2Path, V2Host string
	for _, item := range extraServerConf {
		conf := strings.Split(item, "=")
		key := conf[0]
		value := conf[1]
		switch key {
		case "path":
			V2Path = value
		case "host":
			V2Host = value
		}
	}
	nodeInfo.V2Port = V2Port
	nodeInfo.V2AlterID = V2AlterID
	nodeInfo.V2Net = V2Net
	nodeInfo.V2TLS = V2TLS
	nodeInfo.V2Path = V2Path
	nodeInfo.V2Host = V2Host

	return nodeInfo, nil
}

// GetUserList pull user list from vnet-panel
func (c *ApiClient) GetUserList() (data []*VMessUser, err error) {
	path := "/mod_mu/users"
	res, err := c.client.R().
		SetQueryParam("node_id", strconv.Itoa(c.NodeID)).
		SetResult(&UsersResponse{}).
		ForceContentType("application/json").
		Get(path)

	if err != nil {
		return nil, newError(fmt.Sprintf("request %s failed", c.AssembleUrl(path))).Base(err)
	}

	if res.StatusCode() > 400 {
		body := res.Body()
		return nil, newError(fmt.Sprintf("request %s failed: %s", c.AssembleUrl(path), string(body))).Base(err)
	}

	result := res.Result().(*UsersResponse)

	if result.Ret != 1 {
		res, _ := json.Marshal(&result)
		return nil, newError(fmt.Sprintf("Ret %s invalid", string(res)))
	}

	return result.Data, err
}

// Report Node Status to vnet-panel
func (c *ApiClient) ReportNodeStatus(nodeStatus *NodeStatus) error {
	if c.NodeID == 0 {
		return newError("NodeId is 0")
	}

	path := fmt.Sprintf("/mod_mu/nodes/%d/info", c.NodeID)
	systemload := SystemLoad{
		Uptime: strconv.Itoa(nodeStatus.Uptime),
		Load:   nodeStatus.CPU,
	}

	res, err := c.client.R().
		SetBody(systemload).
		SetResult(&PostResponse{}).
		ForceContentType("application/json").
		Post(c.AssembleUrl(path))
	if err != nil {
		return newError(fmt.Sprintf("report node status error: %s", res.Body()))
	}

	response := res.Result().(*PostResponse)
	if response.Ret != 1 {
		return newError(fmt.Sprintf("report node status failed: %s", response.Data))
	}

	return nil
}

func (c *ApiClient) ReportNodeOnline(online []*NodeOnline) error {
	if c.NodeID == 0 {
		return newError("NodeId is 0")
	}

	data := NodeOnlinePost{online}

	path := fmt.Sprintf("/mod_mu/users/aliveip")
	res, err := c.client.R().
		SetQueryParam("node_id", strconv.Itoa(c.NodeID)).
		SetBody(data).
		SetResult(&PostResponse{}).
		ForceContentType("application/json").
		Post(c.AssembleUrl(path))
	if err != nil {
		return newError(fmt.Sprintf("report node status error: %s", res.Body()))
	}

	response := res.Result().(*PostResponse)
	if response.Ret != 1 {
		return newError(fmt.Sprintf("report node status failed: %s", response))
	}

	return nil
}

func (c *ApiClient) ReportUserTraffic(traffics []*UserTraffic) error {
	if c.NodeID == 0 {
		return newError("NodeId is 0")
	}
	data := UserTrafficPost{traffics}
	path := "/mod_mu/users/traffic"
	res, err := c.client.R().
		SetQueryParam("node_id", strconv.Itoa(c.NodeID)).
		SetBody(data).
		SetResult(&PostResponse{}).
		ForceContentType("application/json").
		Post(c.AssembleUrl(path))
	if err != nil {
		return newError("report user traffic failed").Base(err)
	}
	response := res.Result().(*PostResponse)
	if response.Ret != 1 {
		return newError("report user traffic failed", response)
	}
	return nil
}

func (c *ApiClient) GetNodeRule() (rule *NodeRule, err error) {
	// if c.NodeID == 0 {
	// 	return nil, newError("NodeId is 0")
	// }

	// path := fmt.Sprintf("/api/v2ray/v1/nodeRule/%d", c.NodeID)
	// request := c.createCommonRequest()
	// request.SetResult(&Response{})
	// res, err := request.Get(c.AssembleUrl(path))
	// if err != nil {
	// 	return nil, newError("get node rule failed").Base(err)
	// }
	// response := res.Result().(*Response)
	// if response.Status != "success" {
	// 	return nil, newError("get node rule failed", response.Message)
	// }

	// nodeRule := new(NodeRule)
	// if err := json.Unmarshal(response.Data, &nodeRule); err != nil {
	// 	return nil, newError("get node rule failed").Base(err)
	// }
	nodeRule := &NodeRule{
		Mode: "all",
	}

	return nodeRule, nil
}

func (c *ApiClient) ReportIllegal(illegalReport *IllegalReport) error {
	// if c.NodeID == 0 {
	// 	return newError("NodeId is 0")
	// }

	// path := fmt.Sprintf("/api/v2ray/v1/trigger/%d", c.NodeID)
	// request := c.createCommonRequest()
	// request.SetBody(illegalReport)
	// request.SetResult(&Response{})
	// res, err := request.Post(c.AssembleUrl(path))
	// if err != nil {
	// 	return newError("illegal report failed").Base(err)
	// }

	// response := res.Result().(*Response)
	// if response.Status != "success" {
	// 	return newError("illegal report failed", response.Message)
	// }

	return nil
}

func (c *ApiClient) PushCertification(certificate *Certificate) error {
	// if c.NodeID == 0 {
	// 	return newError("NodeId is 0")
	// }

	// path := fmt.Sprintf("/api/v2ray/v1/certificate/%d", c.NodeID)
	// request := c.createCommonRequest()
	// request.SetBody(certificate)
	// request.SetResult(&Response{})
	// res, err := request.Post(c.AssembleUrl(path))
	// if err != nil {
	// 	return newError("push certificate failed").Base(err)
	// }

	// response := res.Result().(*Response)
	// if response.Status != "success" {
	// 	return newError("push certificate failed", response.Message)
	// }

	return nil
}

func (c *ApiClient) AssembleUrl(path string) string {
	return c.ApiHost + path
}

func (c *ApiClient) createCommonRequest() *resty.Request {
	request := c.client.R().EnableTrace()
	request.EnableTrace()
	request.SetHeader("key", c.Key)
	request.SetHeader("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	return request
}
