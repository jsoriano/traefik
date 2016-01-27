package provider

import (
	log "github.com/Sirupsen/logrus"
	"github.com/emilevauge/traefik/types"
	api "k8s.io/kubernetes/pkg/api/v1"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

type Kubernetes struct {
	baseProvider
	APIServer string
	BasicAuth *KubernetesBasicAuth
	TLS       *KubernetesTLS
}

type KubernetesBasicAuth struct {
	User     string
	Password string
}

type KubernetesTLS struct {
	CA                 string
	Cert               string
	Key                string
	InsecureSkipVerify bool
}

type KubernetesBackend struct {
}

func (provider *Kubernetes) Provide(configurationChan chan<- types.ConfigMessage) error {
	var err error

	config := client.Config{
		Host: provider.APIServer,
	}
	c, err := client.New(&config)
	if err != nil {
		log.Errorf("Failed to create a Kubernetes client: %s", err)
		return err
	}

	log.Debug("Kubernetes API connection established")
	if provider.Watch {
		log.Debug("Watching events in Kubernetes")
		provider.watch(c, configurationChan)
	} else {
		provider.sendConfiguration(c, configurationChan)
	}
	return nil
}

func (p *Kubernetes) watch(c *client.Client) {
	ei := c.Events(api.NamespacesAll)
	go func() {
		// listener, err := ei.Watch
	}()
}

func (p *Kubernetes) sendConfiguration(c *client.Client, configurationChan chan<- types.ConfigMessage) {
	configurationChan <- types.ConfigMessage{
		ProviderName:  "kubernetes",
		Configuration: p.configuration(p.getBackends(c)),
	}
}

func (p *Kubernetes) getNodes(c *client.Client, configurationChan chan<- types.ConfigMessage) {
}

func (p *Kubernetes) getServices(c *client.Client) {
}

func (p *Kubernetes) getBackends(c *client.Client) {
}

func (p *Kubernetes) configuration(bs []KubernetesBackend) *types.Configuration {
	return nil
}
