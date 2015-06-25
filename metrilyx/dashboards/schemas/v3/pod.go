package v3

type PodContentOrientation string

const (
	POD_CONTENT_VERTICAL   PodContentOrientation = "vertical"
	POD_CONTENT_HORIZONTAL PodContentOrientation = "horizontal"
)

type Pod struct {
	Name        string                `json:"name"`
	Orientation PodContentOrientation `json:"orientation"`
	Panels      []*PodPanel           `json:"panels"`
}

func NewPod(skeleton bool) *Pod {
	if skeleton {
		return &Pod{"", POD_CONTENT_VERTICAL, make([]*PodPanel, 0)}
	} else {
		return &Pod{"", POD_CONTENT_VERTICAL, []*PodPanel{NewPodPanel()}}
	}
}
