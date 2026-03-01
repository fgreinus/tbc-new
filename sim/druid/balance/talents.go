package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/druid"
)

func (moonkin *BalanceDruid) ApplyBalanceTalents() {
	moonkin.registerStarlightWrath()
	moonkin.registerFocusedStarlight()
	moonkin.registerImprovedMoonfire()
	moonkin.registerVengenace()
	moonkin.registerNaturesGrace()
	moonkin.registerLunarGuidance()
	moonkin.registerMoonglow()
	moonkin.registerBalanceOfPower()
	moonkin.registerDreamstate()
	moonkin.registerWrathOfCenarius()
	moonkin.registerMoonfury()
}

func (moonkin *BalanceDruid) registerStarlightWrath() {
	if moonkin.Talents.StarlightWrath == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask: druid.DruidSpellStarfire,
		TimeValue: time.Millisecond * time.Duration(-100*float64(moonkin.Talents.StarlightWrath)),
		Kind:      core.SpellMod_CastTime_Flat,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask: druid.DruidSpellWrath,
		TimeValue: time.Millisecond * time.Duration(-100*float64(moonkin.Talents.StarlightWrath)),
		Kind:      core.SpellMod_CastTime_Flat,
	})
}

func (moonkin *BalanceDruid) registerFocusedStarlight() {
	if moonkin.Talents.FocusedStarlight == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellStarfire,
		FloatValue: 0.02 * float64(moonkin.Talents.FocusedStarlight),
		Kind:       core.SpellMod_BonusCrit_Percent,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellWrath,
		FloatValue: 0.02 * float64(moonkin.Talents.FocusedStarlight),
		Kind:       core.SpellMod_BonusCrit_Percent,
	})
}

func (moonkin *BalanceDruid) registerImprovedMoonfire() {
	if moonkin.Talents.ImprovedMoonfire == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfire,
		FloatValue: 0.05 * float64(moonkin.Talents.ImprovedMoonfire),
		Kind:       core.SpellMod_DamageDone_Pct,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfire,
		FloatValue: 0.05 * float64(moonkin.Talents.ImprovedMoonfire),
		Kind:       core.SpellMod_BonusCrit_Percent,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfireDoT,
		FloatValue: 0.05 * float64(moonkin.Talents.ImprovedMoonfire),
		Kind:       core.SpellMod_DamageDone_Pct,
	})
}

func (moonkin *BalanceDruid) registerVengenace() {
	if moonkin.Talents.Vengeance == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellStarfire,
		FloatValue: 0.20 * float64(moonkin.Talents.Vengeance),
		Kind:       core.SpellMod_CritMultiplier_Flat,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellWrath,
		FloatValue: 0.20 * float64(moonkin.Talents.Vengeance),
		Kind:       core.SpellMod_CritMultiplier_Flat,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfire,
		FloatValue: 0.20 * float64(moonkin.Talents.Vengeance),
		Kind:       core.SpellMod_CritMultiplier_Flat,
	})
}

func (moonkin *BalanceDruid) registerNaturesGrace() {
	if !moonkin.Talents.NaturesGrace {
		return
	}

	var proccedAt time.Duration
	var proccedSpell *core.Spell

	naturesGraceMod := moonkin.AddDynamicMod(core.SpellModConfig{
		ClassMask: druid.DruidSpellStarfire | druid.DruidSpellWrath,
		TimeValue: -500 * time.Millisecond,
		Kind:      core.SpellMod_CastTime_Flat,
	})
	naturesGraceMod.Deactivate()

	moonkin.NaturesGraceProcAura = moonkin.RegisterAura(core.Aura{
		Label:    "Natures Grace Proc",
		ActionID: core.ActionID{SpellID: 16886},
		Duration: core.NeverExpires,
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if !moonkin.Wrath.IsEqual(spell) && !moonkin.StarfireR8.IsEqual(spell) && !moonkin.StarfireR6.IsEqual(spell) {
				return
			}

			if proccedAt == sim.CurrentTime && proccedSpell == spell {
				// Means this is another hit from the same cast that procced Nature's Grace.
				return
			}

			aura.Deactivate(sim)
		},
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			naturesGraceMod.Activate()
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			naturesGraceMod.Deactivate()
		},
	})

	moonkin.RegisterAura(core.Aura{
		Label: "Natures Grace",
		// Commented out actionID so it does not show in the sim
		// ActionID: core.ActionID{SpellID: 16880},
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellResult *core.SpellResult) {
			if spellResult.Outcome.Matches(core.OutcomeCrit) {
				proccedAt = sim.CurrentTime
				proccedSpell = spell
				moonkin.NaturesGraceProcAura.Activate(sim)
			}
		},
	})
}

func (moonkin *BalanceDruid) registerLunarGuidance() {
	if moonkin.Talents.LunarGuidance == 0 {
		return
	}

	conversions := []float64{0.08, 0.16, 0.25}
	moonkin.AddStatDependency(stats.Intellect, stats.SpellDamage, conversions[moonkin.Talents.LunarGuidance-1])
	moonkin.AddStatDependency(stats.Intellect, stats.HealingPower, conversions[moonkin.Talents.LunarGuidance-1])
}

func (moonkin *BalanceDruid) registerMoonglow() {
	if moonkin.Talents.Moonglow == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellStarfire,
		FloatValue: -0.03 * float64(moonkin.Talents.Moonglow),
		Kind:       core.SpellMod_PowerCost_Pct,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellWrath,
		FloatValue: -0.03 * float64(moonkin.Talents.Moonglow),
		Kind:       core.SpellMod_PowerCost_Pct,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfire,
		FloatValue: -0.03 * float64(moonkin.Talents.Moonglow),
		Kind:       core.SpellMod_PowerCost_Pct,
	})
}

func (moonkin *BalanceDruid) registerMoonfury() {
	if moonkin.Talents.Moonfury == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellStarfire,
		FloatValue: 0.02 * float64(moonkin.Talents.Moonfury),
		Kind:       core.SpellMod_DamageDone_Pct,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfire,
		FloatValue: 0.02 * float64(moonkin.Talents.Moonfury),
		Kind:       core.SpellMod_DamageDone_Pct,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellMoonfireDoT,
		FloatValue: 0.02 * float64(moonkin.Talents.Moonfury),
		Kind:       core.SpellMod_DamageDone_Pct,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellWrath,
		FloatValue: 0.02 * float64(moonkin.Talents.Moonfury),
		Kind:       core.SpellMod_DamageDone_Pct,
	})
}

func (moonkin *BalanceDruid) registerBalanceOfPower() {
	if moonkin.Talents.BalanceOfPower == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidDamagingSpells,
		FloatValue: 2 * float64(moonkin.Talents.BalanceOfPower),
		Kind:       core.SpellMod_BonusHit_Percent,
	})
}

func (moonkin *BalanceDruid) registerDreamstate() {
	if moonkin.Talents.Dreamstate == 0 {
		return
	}

	conversions := []float64{0.04, 0.07, 0.10}
	moonkin.AddStatDependency(stats.Intellect, stats.MP5, conversions[moonkin.Talents.Dreamstate-1])
}

func (moonkin *BalanceDruid) registerWrathOfCenarius() {
	if moonkin.Talents.WrathOfCenarius == 0 {
		return
	}

	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellStarfire,
		FloatValue: 0.04 * float64(moonkin.Talents.BalanceOfPower),
		Kind:       core.SpellMod_BonusCoeffecient_Flat,
	})
	moonkin.AddStaticMod(core.SpellModConfig{
		ClassMask:  druid.DruidSpellWrath,
		FloatValue: 0.02 * float64(moonkin.Talents.BalanceOfPower),
		Kind:       core.SpellMod_BonusCoeffecient_Flat,
	})
}
