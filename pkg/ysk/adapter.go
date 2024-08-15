package ysk

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
)

type CartType string
type YSKCardIcon string

const (
	CardTypeTask      CartType = "task"
	CardTypeLongNote  CartType = "long-notice"
	CardTypeShortNote CartType = "short-notice"
)

const (
	FileIcon     YSKCardIcon = "/modules/icewhale_files/appicon.svg"
	DiskIcon     YSKCardIcon = "/src/assets/img/storage/disk.png"
	ZimaIcon     YSKCardIcon = "/src/assets/img/zima.svg"
	StorageIcon  YSKCardIcon = "/src/assets/img/storage/storage.svg"
	AppStoreIcon YSKCardIcon = "/src/assets/img/welcome/appstore.svg"
)

type RenderType string

const (
	RenderTypeCardTask           RenderType = "task"
	RenderTypeCardListNotice     RenderType = "list-notice"
	RenderTypeCardIconTextNotice RenderType = "icon-text-notice"
	RenderTypeCardMarkdownNotice RenderType = "markdown-notice"
)

type ActionPosition string

const (
	ActionPositionLeft  ActionPosition = "left"
	ActionPositionRight ActionPosition = "right"
)

type YSKCard struct {
	Id         string         `json:"id"`
	CardType   CartType       `json:"cardType"`
	RenderType RenderType     `json:"renderType"`
	Content    YSKCardContent `json:"content"`
}

func (ysk YSKCard) WithId(id string) YSKCard {
	ysk.Id = id
	return ysk
}

func (ysk YSKCard) WithTaskContent(TitleIcon YSKCardIcon, TitleText string) YSKCard {
	ysk.Content.TitleIcon = TitleIcon
	ysk.Content.TitleText = TitleText
	return ysk
}

func (yskCard YSKCard) WithProgress(label string, progress int) YSKCard {
	if yskCard.Content.BodyProgress != nil {
		yskCard.Content.BodyProgress = &YSKCardProgress{
			Label:    label,
			Progress: progress,
		}
		return yskCard
	}
	return yskCard
}

func (yskCard YSKCard) WithList(params []YSKCardListItem) YSKCard {
	yskCard.Content.BodyList = params
	return yskCard
}

// it will replace the old action by same side and style
func (YSKCard YSKCard) UpsertFooterAction(action YSKCardFooterAction) YSKCard {
	for i, a := range YSKCard.Content.FooterActions {
		if a.Side == action.Side && a.Style == action.Style {
			YSKCard.Content.FooterActions[i] = action
		}
	}
	return YSKCard
}

type YSKCardContent struct {
	TitleIcon        YSKCardIcon           `json:"titleIcon" gorm:"column:title_icon"`
	TitleText        string                `json:"titleText" gorm:"column:title_text"`
	BodyProgress     *YSKCardProgress      `json:"bodyProgress,omitempty" gorm:"serializer:json"`
	BodyIconWithText *YSKCardIconWithText  `json:"bodyIconWithText,omitempty" gorm:"serializer:json"`
	BodyList         []YSKCardListItem     `json:"bodyList,omitempty" gorm:"serializer:json"`
	FooterActions    []YSKCardFooterAction `json:"footerActions,omitempty" gorm:"serializer:json"`
}

func (p YSKCardContent) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *YSKCardContent) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}

type YSKCardProgress struct {
	Label    string `json:"label"`
	Progress int    `json:"progress"`
}

type YSKCardIconWithText struct {
	Icon        YSKCardIcon `json:"icon"`
	Description string      `json:"description"`
}

type YSKCardListItem struct {
	Icon        YSKCardIcon `json:"icon"`
	Description string      `json:"description"`
	RightText   string      `json:"rightText"`
}

type YSKCardFooterAction struct {
	Side       ActionPosition          `json:"side"`
	Style      string                  `json:"style"`
	Text       string                  `json:"text"`
	MessageBus YSKCardMessageBusAction `json:"messageBus"`
}

type YSKCardMessageBusAction struct {
	Key     string `json:"key"`
	Payload string `json:"payload"`
}

func ToCodegenYSKCard(card YSKCard) (codegen.YSKCard, error) {
	jsonBody, err := json.Marshal(card)
	if err != nil {
		return codegen.YSKCard{}, err
	}
	var yskCard codegen.YSKCard
	err = json.Unmarshal(jsonBody, &yskCard)

	return yskCard, err
}

func FromCodegenYSKCard(card codegen.YSKCard) (YSKCard, error) {
	jsonBody, err := json.Marshal(card)
	if err != nil {
		return YSKCard{}, err
	}
	var yskCard YSKCard
	err = json.Unmarshal(jsonBody, &yskCard)

	return yskCard, err
}
