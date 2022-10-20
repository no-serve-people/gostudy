package g20221020

import "fmt"

// https://mp.weixin.qq.com/s/UkX09FsBVop77-T1Xe7l_w?forceh5=1
// Go设计模式——观察者模式

type Platform string

const (
	WEIBO    Platform = "weibo"
	BILIBILI Platform = "bilibili"
	ZHIHU    Platform = "zhihu"
)

// MoYu 用来摸鱼的 APP 们的抽象接口
type MoYu interface {
	// Platform 用来摸鱼的平台名称
	Platform() Platform
	// GetName 获取平台热点事件发布人名称
	GetName() string
	// GetContent 获取热点事件内容
	GetContent() string
}

// Mortal 芸芸众生（用户）
type Mortal interface {
	// Accept 接收热点事件
	Accept()
}

type BaseMortal struct {
	// MoYu 持有摸鱼工具实例，便于获取当前发布的热点事件（摸鱼工具的状态）
	MoYu
}

type BaseMoYu struct {
	Name    string
	Content string
	// MortalList 热点事件发布后，需要通知的用户列表
	MortalList []Mortal
}

func (b *BaseMoYu) Attach(mortal Mortal) {
	b.MortalList = append(b.MortalList, mortal)
}

func (b *BaseMoYu) Publish(name, content string) {
	b.Name = name
	b.Content = content

	// 依次通知
	for _, item := range b.MortalList {
		item.Accept()
	}
}

func (b *BaseMoYu) GetName() string {
	return b.Name
}

func (b *BaseMoYu) GetContent() string {
	return b.Content
}

type MoYuByBilibili struct {
	BaseMoYu
}

func NewMoYuByBilibili() *MoYuByBilibili {
	return &MoYuByBilibili{}
}

func (m *MoYuByBilibili) Platform() Platform {
	return BILIBILI
}

type Tom struct {
	BaseMortal
}

func NewTom(yu MoYu) *Tom {
	return &Tom{BaseMortal{MoYu: yu}}
}

func (c Tom) Accept() {
	fmt.Printf("Tom accept from: %v, content: %v, published by: %v\n",
		c.BaseMortal.Platform(), c.BaseMortal.GetContent(), c.BaseMortal.GetName())
}

type Jerry struct {
	BaseMortal
}

func NewJerry(moYu MoYu) *Jerry {
	return &Jerry{BaseMortal{MoYu: moYu}}
}

func (j Jerry) Accept() {
	fmt.Printf("Jerry accept from: %v, content: %v, published by: %v\n",
		j.BaseMortal.Platform(), j.BaseMortal.GetContent(), j.BaseMortal.GetName())
}

func testObserver() {
	bilibili := NewMoYuByBilibili()
	tom := NewTom(bilibili)
	jerry := NewJerry(bilibili)

	bilibili.Attach(tom)
	bilibili.Attach(jerry)

	bilibili.Publish("张朝阳的物理课", "张朝阳、俞敏洪是怎样“摸鱼”的")
	bilibili.Publish("哔哩哔哩国创", "《暂停！让我查攻略》")
}
