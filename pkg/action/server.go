package action

import (
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log/level"
	"github.com/nilo72/owncloud_exporter/pkg/client"
	"github.com/nilo72/owncloud_exporter/pkg/config"
	"github.com/go-kit/kit/log"
	"github.com/webhippie/go-owncloud/owncloud"
)



// Server handles the server sub-command.
func Server(cfg *config.Config, logger log.Logger) error {
	level.Info(logger).Log(
		"msg", "Launching owncloud exporter",
	)

	// Initialize the clientset
	// TODO: Generate from swagger spec?
	ccf := client.ClientConfig{
		Address: cfg.Target.Address,
		Timeout: cfg.Target.Timeout,
	}

	_ = client.NewClient(&ccf)

	// run the exporter functions
	// TODO: not yet implemented
	return nil

}


func handler(cfg *config.Config, logger log.Logger, client *owncloud.Client) *chi.Mux {
	//
	return nil
}