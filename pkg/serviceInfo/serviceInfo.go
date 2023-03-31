package serviceInfo

import (
	"os"

	"github.com/xiaohubai/go-grpc-layout/configs"
)

type ServiceInfo struct {
	Name     string
	Env      string
	Version  string
	Id       string
	Metadata map[string]string
}

func NewServiceInfo(g *configs.Global) ServiceInfo {
	id, _ := os.Hostname()
	return ServiceInfo{
		Name:     g.AppName,
		Env:      g.Env,
		Version:  g.Version,
		Id:       id,
		Metadata: map[string]string{},
	}
}

func (s *ServiceInfo) GetInstanceId() string {
	return s.Id + "." + s.Name
}

func (s *ServiceInfo) SetMataData(k, v string) {
	s.Metadata[k] = v
}
