# ipmitool

This ipmitool library is an `ipmitool` command wrapper to be used as an ipmi client.

This library is made for the purposes of Bancadati and is therefor limited in functionality and supported commands.  
For now only the (chassis) power command is supported.

## Usage

```go
package main

import (
	"fmt"
    "log"

	"github.com/bancadati/ipmitool"
)

func main() {
    cl, err := ipmitool.NewClient("192.198.1.1", 0, "IPMIUSER", "Password")
	if err != nil {
		log.Fatal(err)
	}

	status, err := cl.Power.Status()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(status)

    if status == ipmitool.PowerStateOff {
		err := cl.Power.On()
		if err != nil {
			log.Fatal(err)
		}
	}
}
```