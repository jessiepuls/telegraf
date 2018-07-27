package nginx_plus

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

func init() {
	inputs.Add("nginx_plus_status", func() telegraf.Input {
		return &NginxPlusStatus{}
	})
}
