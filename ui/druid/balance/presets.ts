import { Encounter } from '../../core/encounter';
import * as PresetUtils from '../../core/preset_utils.js';
import { ConsumesSpec, Debuffs, IndividualBuffs, PartyBuffs, Profession, RaidBuffs, Stat, TristateEffect, UnitReference } from '../../core/proto/common.js';
import { BalanceDruid_Options as BalanceDruidOptions } from '../../core/proto/druid.js';
import { SavedTalents } from '../../core/proto/ui.js';
import { Stats } from '../../core/proto_utils/stats';
import { defaultRaidBuffMajorDamageCooldowns } from '../../core/proto_utils/utils';
import StandardApl from './apls/standard.apl.json';
import PreraidGear from './gear_sets/preraid.gear.json';
import { Race } from '../../core/proto/common';

export const PreraidPresetGear = PresetUtils.makePresetGear('Pre-raid', PreraidGear);

export const StandardRotation = PresetUtils.makePresetAPLRotation('Standard', StandardApl);

export const StandardEPWeights = PresetUtils.makePresetEpWeights(
	'Standard',
	Stats.fromMap({
		[Stat.StatIntellect]: 0.54,
		[Stat.StatSpirit]: 0.1,
		[Stat.StatSpellDamage]: 1,
		[Stat.StatArcaneDamage]: 1,
		[Stat.StatNatureDamage]: 0,
		[Stat.StatSpellCritRating]: 0.84,
		[Stat.StatSpellHasteRating]: 1.29,
		[Stat.StatSpellPenetration]: 1,
		[Stat.StatMana]: 1,
	}),
);

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/tbc/talent-calc and copy the numbers in the url.
export const StandardTalents = {
	name: 'Standard',
	data: SavedTalents.create({
		talentsString: '510022312503135231351--500233',
	}),
};

export const DefaultOptions = BalanceDruidOptions.create({
	classOptions: {
		innervateTarget: UnitReference.create(),
		faerieFire: true,
		insectSwarm: false,
	},
});

export const DefaultConsumables = ConsumesSpec.create({
	flaskId: 22861, // Flask of Blinding Light
	foodId: 27657, // Blackened Basilisk
	potId: 22832, // Super Mana Potion
	mhImbueId: 25122, // Brilliant Wizard Oil
	conjuredId: 12662, // Demonic Rune
	drumsId: 351355, // Greater Drums of Battle (Haste)
});

export const DefaultRaidBuffs = RaidBuffs.create({
	...defaultRaidBuffMajorDamageCooldowns(),
	arcaneBrilliance: true,
	divineSpirit: TristateEffect.TristateEffectImproved,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectRegular,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfSalvation: true,
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
});

export const DefaultPartyBuffs = PartyBuffs.create({
	totemOfWrath: TristateEffect.TristateEffectRegular,
	wrathOfAirTotem: TristateEffect.TristateEffectRegular,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: TristateEffect.TristateEffectRegular,
});

export const DefaultDebuffs = Debuffs.create({
	judgementOfWisdom: true,
	misery: true,
	curseOfElements: TristateEffect.TristateEffectRegular,
	improvedSealOfTheCrusader: true,
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Engineering,
	profession2: Profession.Tailoring,
	race: Race.RaceNightElf,
};

const encounter = Encounter.defaultEncounterProto();
encounter.duration = 120;

const ENCOUNTER_SINGLE_TARGET = PresetUtils.makePresetEncounter('Single Target Dummy', encounter);

export const PresetPreraidBuild = PresetUtils.makePresetBuild('Pre-raid', {
	gear: PreraidPresetGear,
	talents: StandardTalents,
	rotation: StandardRotation,
	epWeights: StandardEPWeights,
	encounter: ENCOUNTER_SINGLE_TARGET,
});
