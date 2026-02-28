package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const (
	MoonfireBonusCoeff    = 0.15
	MoonfireDotBonusCoeff = 0.12999999523
)

func (druid *Druid) registerMoonfireSpell() {
	druid.registerMoonfireImpactSpell()
	druid.registerMoonfireDoTSpell()
}

func (druid *Druid) registerMoonfireDoTSpell() {
	druid.Moonfire.RelatedDotSpell = druid.Unit.RegisterSpell(core.SpellConfig{
		ActionID:       core.ActionID{SpellID: 26988}.WithTag(1),
		SpellSchool:    core.SpellSchoolArcane,
		ProcMask:       core.ProcMaskSpellDamage,
		ClassSpellMask: DruidSpellMoonfireDoT,
		Flags:          core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

		DamageMultiplier: 1,
		CritMultiplier:   druid.DefaultSpellCritMultiplier(),
		ThreatMultiplier: 1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Moonfire (DoT)",
			},
			NumberOfTicks:       4,
			TickLength:          time.Second * 3,
			BonusCoefficient:    MoonfireDotBonusCoeff,
			AffectedByCastSpeed: false,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.Snapshot(target, float64(600/dot.BaseTickCount))
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.Dot(target).Apply(sim)
		},
	})
}

func (druid *Druid) registerMoonfireImpactSpell() {
	druid.Moonfire = druid.RegisterSpell(Humanoid|Moonkin, core.SpellConfig{
		ActionID:       core.ActionID{SpellID: 26988},
		SpellSchool:    core.SpellSchoolArcane,
		ProcMask:       core.ProcMaskSpellDamage,
		ClassSpellMask: DruidSpellMoonfire,
		Flags:          core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			FlatCost: 495,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},

		DamageMultiplier: 1,

		CritMultiplier:   druid.DefaultSpellCritMultiplier(),
		ThreatMultiplier: 1,
		BonusCoefficient: MoonfireBonusCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := druid.CalcAndRollDamageRange(sim, 305, 357)
			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			spell.DealDamage(sim, result)
			if result.Landed() {
				druid.Moonfire.RelatedDotSpell.Cast(sim, target)
			}
		},
	})
}
