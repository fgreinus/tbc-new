package balance

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get caster sets included. (we use spellfire here)
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterBalanceDruid()
}

func TestBalance(t *testing.T) {

	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassDruid,
			Race:       proto.Race_RaceNightElf,
			OtherRaces: []proto.Race{proto.Race_RaceTauren},
			SpecOptions: core.SpecOptionsCombo{Label: "Balance", SpecOptions: &proto.Player_BalanceDruid{
				BalanceDruid: &proto.BalanceDruid{
					Options: &proto.BalanceDruid_Options{
						ClassOptions: &proto.DruidOptions{},
					},
				},
			}},
			GearSet:  core.GetGearSet("../../../ui/druid/balance/gear_sets", "preraid"),
			Talents:  "510022312503135231351--500233",
			Rotation: core.GetAplRotation("../../../ui/druid/balance/apls", "standard"),
			ItemFilter: core.ItemFilter{
				WeaponTypes: []proto.WeaponType{
					proto.WeaponType_WeaponTypeDagger,
					proto.WeaponType_WeaponTypeMace,
					proto.WeaponType_WeaponTypeOffHand,
					proto.WeaponType_WeaponTypeStaff,
				},
				ArmorType: proto.ArmorType_ArmorTypeLeather,
				RangedWeaponTypes: []proto.RangedWeaponType{
					proto.RangedWeaponType_RangedWeaponTypeIdol,
				},
				EnchantBlacklist: []int32{2673, 3225, 3273},
			},
		},
	}))
}
