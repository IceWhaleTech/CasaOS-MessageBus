package ysk

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
	ID         string
	CardType   CartType
	RenderType RenderType
	Content    YSKCardContent
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
	TitleIcon        string
	TitleText        string
	BodyProgress     *YSKCardProgress
	BodyIconWithText *YSKCardIconWithText
	BodyList         []YSKCardListItem
	FooterActions    []YSKCardFooterAction
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
	Key     string
	Payload string
}

type YSKCardIcon = string
