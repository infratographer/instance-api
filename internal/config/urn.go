package config

import (
	"github.com/google/uuid"
	"go.infratographer.com/x/urnx"
)

type URNType string

const (
	TypeInstance         URNType = "instance"
	TypeInstanceProvider URNType = "instance_provider"
)

func BuildURN(t URNType, id uuid.UUID) (*urnx.URN, error) {
	return urnx.Build(AppConfig.URN.Namespace, t.URNType(), id)
}

func (t URNType) URNType() string {
	switch t {
	case TypeInstance:
		return AppConfig.URN.Types.Instance
	case TypeInstanceProvider:
		return AppConfig.URN.Types.InstanceProvider
	default:
		return "Unknown"
	}

}
