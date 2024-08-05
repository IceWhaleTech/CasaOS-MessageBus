package ysk

import (
	"encoding/json"

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
			Label: label,
			Value: progress,
		}
		return yskCard
	}
	return yskCard
}

type YSKCardContent struct {
	TitleIcon        string                `json:"titleIcon"`
	TitleText        string                `json:"titleText"`
	BodyProgress     *YSKCardProgress      `json:"bodyProgress,omitempty"`
	BodyIconWithText *YSKCardIconWithText  `json:"bodyIconWithText,omitempty"`
	BodyList         []YSKCardListItem     `json:"bodyList,omitempty"`
	FooterActions    []YSKCardFooterAction `json:"footerActions,omitempty"`
}

type YSKCardProgress struct {
	Label string
	Value int
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
