package druid

import "github.com/wowsims/tbc/sim/core"

func (druid *Druid) ApplyTalents() {
	druid.registerIntensity()
}

func (druid *Druid) registerIntensity() {
	// druid.PseudoStats.SpiritRegenRateCasting = float64(druid.Talents.Intensity) * 0.1
}

func (druid *Druid) RegisterSharedFeralHotwMods() (*core.SpellMod, *core.SpellMod, *core.SpellMod) {
	return nil, nil, nil
}
