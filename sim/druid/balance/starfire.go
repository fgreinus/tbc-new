package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/druid"
)

const (
	StarfireBonusCoeff = 1
	StarfireR8MinDmg   = 550
	StarfireR8MaxDmg   = 647

	IvoryMoongoddess int32 = 27518
)

func (moonkin *BalanceDruid) registerStarfireSpell() {
	bonusSpellDamage := 0.0
	if moonkin.Druid.Equipment.Ranged().ID == IvoryMoongoddess {
		bonusSpellDamage = 55
	}

	moonkin.StarfireR8 = moonkin.RegisterSpell(druid.Humanoid|druid.Moonkin, core.SpellConfig{
		ActionID:       core.ActionID{SpellID: 26986},
		SpellSchool:    core.SpellSchoolArcane,
		ProcMask:       core.ProcMaskSpellDamage,
		ClassSpellMask: druid.DruidSpellStarfire,
		Flags:          core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			FlatCost: 370,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 3500,
			},
		},

		BonusCoefficient: StarfireBonusCoeff,

		BonusSpellDamage: bonusSpellDamage,

		DamageMultiplier: 1,

		CritMultiplier: moonkin.DefaultSpellCritMultiplier(),

		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := moonkin.CalcAndRollDamageRange(sim, StarfireR8MinDmg, StarfireR8MaxDmg)
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
		},
	})
}
