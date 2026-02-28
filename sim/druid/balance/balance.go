package balance

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterBalanceDruid() {
	core.RegisterAgentFactory(
		proto.Player_BalanceDruid{},
		proto.Spec_SpecBalanceDruid,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewBalanceDruid(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_BalanceDruid)
			if !ok {
				panic("Invalid spec value for Balance Druid!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewBalanceDruid(character *core.Character, options *proto.Player) *BalanceDruid {
	balanceOptions := options.GetBalanceDruid()
	selfBuffs := druid.SelfBuffs{}

	moonkin := &BalanceDruid{
		Druid:   druid.New(character, druid.Moonkin, selfBuffs, options.TalentsString),
		Options: balanceOptions.Options,
	}

	moonkin.registerTreants()

	moonkin.SelfBuffs.InnervateTarget = &proto.UnitReference{}
	if balanceOptions.Options.ClassOptions.InnervateTarget != nil {
		moonkin.SelfBuffs.InnervateTarget = balanceOptions.Options.ClassOptions.InnervateTarget
	}

	moonkin.EnableManaBar()
	moonkin.RegisterMoonkinFormAura()

	return moonkin
}

type BalanceDruid struct {
	*druid.Druid
	Options *proto.BalanceDruid_Options

	ManaMetric *core.ResourceMetrics

	StarfireR8 *druid.DruidSpell
	StarfireR6 *druid.DruidSpell

	ForceOfNature *druid.DruidSpell
	Treants       TreantAgents
}

func (moonkin *BalanceDruid) GetDruid() *druid.Druid {
	return moonkin.Druid
}

func (moonkin *BalanceDruid) Initialize() {
	moonkin.Druid.Initialize()

	moonkin.RegisterBalanceSpells()
}

func (moonkin *BalanceDruid) ApplyTalents() {
	moonkin.Druid.ApplyTalents()
	moonkin.ApplyBalanceTalents()
}

func (moonkin *BalanceDruid) RegisterBalanceSpells() {
	moonkin.RegisterMoonkinFormSpell()

	moonkin.registerStarfireSpell()
	moonkin.registerForceOfNatureSpell()
}

func (moonkin *BalanceDruid) registerTreants() {
	for idx := range moonkin.Treants {
		treant := moonkin.NewDefaultTreant(TreantConfig{
			EnableAutos:             true,
			WeaponDamageCoefficient: 130,
			NonHitExpStatInheritance: func(ownerStats stats.Stats) stats.Stats {
				return stats.Stats{
					stats.AttackPower: ownerStats[stats.SpellDamage] * 0.5,
				}
			},
		})
		moonkin.Treants[idx] = treant
		moonkin.AddPet(treant)
	}
}

func (moonkin *BalanceDruid) Reset(sim *core.Simulation) {
	moonkin.Druid.Reset(sim)
}
