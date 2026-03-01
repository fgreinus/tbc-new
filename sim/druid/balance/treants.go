package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Testing results
// 30 sek mit full gear
// (799 spelldmg, 799 nature, 79 str, 81 agi, 316 sta, 339 int, 160 spirit, 120 spellhit (9,51%), 15,06% spellcrit)
// 51 hits, average normal 187, 203 dps, 5,9% crit, min:132, max:214, crit-min:358, crit-max:374

// 30sek ohne spelldmg trinket
// 756 spell, 756 nature, 79 str, 81 agi, 316 sta, 339 int,
// 51 hits, average normal 184, dps 188, 7,8% crit

// noch weniger gear
// 464 spelldmg, 295 int,
// 50 hits, average normal 163, avg dps 129, min 123, max 185, 12% crit?!

// Extension of PetAgent interface, for treants.
type TreantAgent interface {
	core.PetAgent

	Enable(sim *core.Simulation)
}

// Embed this in spec-specific treant structs.
type DefaultTreantImpl struct {
	core.Pet
}

// Overwrite these for spec variants that register spells.
func (treant *DefaultTreantImpl) Initialize()                              {}
func (treant *DefaultTreantImpl) ExecuteCustomRotation(_ *core.Simulation) {}

func (treant *DefaultTreantImpl) Reset(sim *core.Simulation) {
	treant.Disable(sim)
}

func (treant *DefaultTreantImpl) OnEncounterStart(_ *core.Simulation) {
}

func (treant *DefaultTreantImpl) GetPet() *core.Pet {
	return &treant.Pet
}

func (treant *DefaultTreantImpl) Enable(sim *core.Simulation) {
	treant.EnableWithTimeout(sim, treant, time.Second*30)
}

type TreantConfig struct {
	NonHitExpStatInheritance core.PetStatInheritance
	EnableAutos              bool
	WeaponDamageCoefficient  float64
}

func (balance *BalanceDruid) NewDefaultTreant(config TreantConfig) *DefaultTreantImpl {
	treant := &DefaultTreantImpl{
		Pet: core.NewPet(core.PetConfig{
			Name:                            "Treant",
			Owner:                           &balance.Character,
			NonHitExpStatInheritance:        config.NonHitExpStatInheritance,
			EnabledOnStart:                  false,
			IsGuardian:                      true,
			HasDynamicMeleeSpeedInheritance: false,
			HasDynamicCastSpeedInheritance:  false,
			HasResourceRegenInheritance:     false,
		}),
	}

	if !config.EnableAutos {
		return treant
	}

	baseWeaponDamage := config.WeaponDamageCoefficient

	treant.EnableAutoAttacks(treant, core.AutoAttackOptions{
		MainHand: core.Weapon{
			BaseDamageMin:        baseWeaponDamage,
			BaseDamageMax:        baseWeaponDamage,
			SwingSpeed:           2,
			NormalizedSwingSpeed: 2,
			CritMultiplier:       balance.DefaultMeleeCritMultiplier(),
			SpellSchool:          core.SpellSchoolPhysical,
		},

		AutoSwingMelee: true,
	})

	treant.OnPetEnable = func(sim *core.Simulation) {
		// Treant spawns in front of boss but moves behind after first swing.
		treant.PseudoStats.InFrontOfTarget = true
		pa := sim.GetConsumedPendingActionFromPool()
		pa.NextActionAt = sim.CurrentTime + time.Millisecond*500

		pa.OnAction = func(_ *core.Simulation) {
			treant.PseudoStats.InFrontOfTarget = false
		}

		sim.AddPendingAction(pa)
	}

	return treant
}

type TreantAgents [3]TreantAgent
