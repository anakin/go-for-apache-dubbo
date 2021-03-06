package selector

import (
	"fmt"
)

import (
	"github.com/dubbo/go-for-apache-dubbo/client"
	"github.com/dubbo/go-for-apache-dubbo/registry"
)

var (
	ServiceArrayEmpty = fmt.Errorf("emtpy service array")
)

type Selector interface {
	Select(ID int64, array client.ServiceArrayIf) (registry.ServiceURL, error)
}
