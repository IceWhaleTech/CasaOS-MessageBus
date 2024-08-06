package ysk

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/IceWhaleTech/CasaOS-MessageBus/codegen"
)

type CartType string

const (
	CardTypeTask      CartType = "task"
	CardTypeLongNote  CartType = "long-notice"
	CardTypeShortNote CartType = "short-notice"
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

type YSKCardContent struct {
	TitleIcon        string                `json:"titleIcon" gorm:"column:title_icon"`
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
	Label    string
	Progress int
}

type YSKCardIconWithText struct {
	Icon        string
	Description string
}

type YSKCardListItem struct {
	Icon        string
	Description string
	RightText   string
}

type YSKCardFooterAction struct {
	Side       ActionPosition
	Style      string
	Text       string
	MessageBus YSKCardMessageBusAction
}

type YSKCardMessageBusAction struct {
	Key     string `json:"key"`
	Payload string `json:"payload"`
}

type YSKCardIcon = string

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
