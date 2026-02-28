package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (druid *Druid) registerFaerieFireSpell() {
	manaCostOptions := core.ManaCostOptions{
		FlatCost: 145,
	}
	gcd := core.GCDDefault
	ignoreHaste := false
	cd := core.Cooldown{}
	formMask := Humanoid | Moonkin

	if druid.InForm(Cat | Bear) {
		manaCostOptions = core.ManaCostOptions{}
		formMask = Cat | Bear
		cd = core.Cooldown{
			Timer:    druid.NewTimer(),
			Duration: time.Second * 6,
		}
	}

	if druid.InForm(Cat) {
		gcd = time.Second
		ignoreHaste = true
	}

	druid.FaerieFireAuras = druid.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.FaerieFireAura(target, druid.Talents.ImprovedFaerieFire > 0)
	})

	druid.FaerieFire = druid.RegisterSpell(formMask, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 26993},
		SpellSchool: core.SpellSchoolNature,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL,

		ManaCost: manaCostOptions,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: gcd,
			},
			IgnoreHaste: ignoreHaste,
			CD:          cd,
		},

		ThreatMultiplier: 1,
		DamageMultiplier: 1,
		CritMultiplier:   druid.DefaultSpellCritMultiplier(),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := 0.0
			outcome := spell.OutcomeMagicHit

			result := spell.CalcAndDealDamage(sim, target, baseDamage, outcome)

			if result.Landed() {
				aura := druid.FaerieFireAuras.Get(target)
				aura.Activate(sim)
			}
		},

		RelatedAuraArrays: druid.FaerieFireAuras.ToMap(),
	})
}
