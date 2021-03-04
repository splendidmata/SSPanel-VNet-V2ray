package apisspanel

import "encoding/json"


type UserModel struct {
	UserID     uint   `json:"id"`
	Uuid       string `json:"uuid"`
	Email      string `json:"email"`
	Passwd     string `json:"passwd"`
	Method     string `json:"method"`
	Port       uint16 `json:"port"`
	AlterId    uint32
	PrefixedId string
}

type VMessUser struct {
	UID        int    `json:"id"`
	VmessUID   string `json:"uuid"`
	SpeedLimit uint64 `json:"node_speedlimit"`
}


type NodeInfo struct {
	Server_raw string `json:"server"`
	Sort       uint   `json:"sort"`
	Server     map[string]interface{}
	ID            int    
	IsUDP         bool   
	SpeedLimit    uint64 `json:"node_speedlimit"`
	ClientLimit   int    
	PushPort      int    
	Secret        string 
	Key           string 
	Cert          string 
	V2License     string 
	V2AlterID     int    
	V2Port        int    
	V2Method      string 
	V2Net         string 
	V2Type        string 
	V2Host        string 
	V2Path        string 
	V2TLS         bool   
	V2Cdn         bool   
	V2TLSProvider string 
	RedirectUrl   string 
}

type DisNodeInfo struct {
	Server_raw string `json:"dist_node_server"`
	Sort       uint   `json:"dist_node_sort"`
	Port       uint16 `json:"port"`
	Server     map[string]interface{}
	UserId     uint `json:"user_id"`
}

type NodeOnline struct {
	UID int    `json:"user_id"`
	IP  string `json:"ip"`
}

type NodeOnlinePost struct{
	Data []*NodeOnline `json:"data"`
}

type NodeinfoResponse struct {
	Ret  uint            `json:"ret"`
	Data *NodeInfo `json:"data"`
}
type PostResponse struct {
	Ret  uint   `json:"ret"`
	Data string `json:"data"`
}
type UsersResponse struct {
	Ret  uint              `json:"ret"`
	Data []*VMessUser `json:"data"`
}
type AllUsers struct {
	Ret  uint
	Data map[string]UserModel
}
type Webapi struct {
	WebToken   string
	WebBaseURl string
}

type DisNodenfoResponse struct {
	Ret  uint                 `json:"ret"`
	Data []*DisNodeInfo `json:"data"`
}

type UserTraffic struct {
	UID      int `json:"user_id"`
	Upload   int `json:"u"`
	Download int `json:"d"`
}


type UserTrafficPost struct{
	Data []*UserTraffic `json:"data"`
}

// Node status report
type NodeStatus struct {
	CPU    string `json:"cpu"`
	Mem    string `json:"mem"`
	Net    string `json:"net"`
	Disk   string `json:"disk"`
	Uptime int    `json:"uptime"`
}

type SystemLoad struct {
	Uptime string `json:"uptime"`
	Load string `json:"load"`
}

var id2string = map[uint]string{
	0: "server_address",
	1: "port",
	2: "alterid",
	3: "protocol",
	4: "protocol_param",
	5: "path",
	6: "host",
	7: "inside_port",
	8: "server",
}
var maps = map[string]interface{}{
	"server_address": "",
	"port":           "",
	"alterid":        "16",
	"protocol":       "tcp",
	"protocol_param": "",
	"path":           "",
	"host":           "",
	"inside_port":    "",
	"server":         "",
}

type NodeRule struct {
	Mode  string         `json:"mode"`
	Rules []NodeRuleItem `json:"rules"`
}

type NodeRuleItem struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Pattern string `json:"pattern"`
}

// IllegalReport
type IllegalReport struct {
	UID    int    `json:"uid"`
	RuleID int    `json:"rule_id"`
	Reason string `json:"reason"`
}

type Certificate struct {
	Key string `json:"key"`
	Pem string `json:"pem"`
}

type Response struct {
	Status  string          `json:"status"`
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}
