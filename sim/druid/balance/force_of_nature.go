package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/druid"
)

func (balance *BalanceDruid) registerForceOfNatureSpell() {
	if !balance.Talents.ForceOfNature {
		return
	}

	balance.ForceOfNature = balance.RegisterSpell(druid.Humanoid|druid.Moonkin, core.SpellConfig{
		ActionID: core.ActionID{SpellID: 33831},
		Flags:    core.SpellFlagAPL,

		SpellSchool:    core.SpellSchoolNature,
		ProcMask:       core.ProcMaskSpellDamage,
		ClassSpellMask: druid.DruidSpellStarfire,

		ManaCost: core.ManaCostOptions{
			BaseCostPercent: 12,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    balance.NewTimer(),
				Duration: 3 * time.Minute,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			for _, treant := range balance.Treants {
				treant.Enable(sim)
			}
		},
	})

	balance.AddMajorCooldown(core.MajorCooldown{
		Spell: balance.ForceOfNature.Spell,
		Type:  core.CooldownTypeDPS,
	})
}
