package stomp

import (
	"github.com/prmuthu/xk6-stomp-tls"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/stomp", stomp.New())
}
