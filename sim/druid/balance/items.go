package balance

import "github.com/wowsims/tbc/sim/core"

func init() {
	// Empty function to remove the warning from the UI
	// because this effect has been implemented in buffs.go
	core.NewItemEffect(27518, func(agent core.Agent) {})
}
