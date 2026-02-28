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
	// TODO
}

func (moonkin *BalanceDruid) registerLunarGuidance() {
	if moonkin.Talents.LunarGuidance == 0 {
		return
	}

	moonkin.AddStatDependency(stats.Intellect, stats.SpellDamage, 0.08*float64(moonkin.Talents.LunarGuidance))
	moonkin.AddStatDependency(stats.Intellect, stats.HealingPower, 0.08*float64(moonkin.Talents.LunarGuidance))
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

	moonkin.AddStatDependency(stats.Intellect, stats.MP5, float64(moonkin.Talents.Dreamstate)*0.04)
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
