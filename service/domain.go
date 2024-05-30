package service

import (
	"education/sdkInit"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
)

type Education struct {
	ObjectType     string `json:"docType"`
	Name           string `json:"Name"`           // 记录人
	Gender         string `json:"Gender"`         // 街道名称
	Nation         string `json:"Nation"`         // 具体位置
	EntityID       string `json:"EntityID"`       // 记录ID号
	Place          string `json:"Place"`          // 楼号
	BirthDay       string `json:"BirthDay"`       // 记录日期
	EnrollDate     string `json:"EnrollDate"`     // 记录时间
	GraduationDate string `json:"GraduationDate"` // 毕（结）业日期
	SchoolName     string `json:"SchoolName"`     // 小区名称
	Major          string `json:"Major"`          // 坐标纬度
	QuaType        string `json:"QuaType"`        // 记录内容
	Length         string `json:"Length"`         // 备注
	Mode           string `json:"Mode"`           // 记录形式
	Level          string `json:"Level"`          // 记录类型
	Graduation     string `json:"Graduation"`     // 毕（结）业
	CertNo         string `json:"CertNo"`         // 记录流水号

	Photo string `json:"Photo"` // 照片

	Historys []HistoryItem // 当前edu的历史记录
}

type HistoryItem struct {
	TxId      string
	Education Education
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

func InitService(chaincodeID, channelID string, org *sdkInit.OrgInfo, sdk *fabsdk.FabricSDK) (*ServiceSetup, error) {
	handler := &ServiceSetup{
		ChaincodeID: chaincodeID,
	}
	//prepare channel client context using client context
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(org.OrgUser), fabsdk.WithOrg(org.OrgName))
	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new channel client: %s", err)
	}
	handler.Client = client
	return handler, nil
}
