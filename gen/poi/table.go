package poi

import (
	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/gen/table"
)

// Table represents the linked roll tables for Point of Interest generation as described in
// Stars Without Number (Revised Edition) pg 171
var Table = table.ThreePart{
	Headers: [3]string{"A Point", "Occupied By", "With This Situation"},
	Tables: []table.ThreePartSubTable{
		{
			Name: "Deep-space station",
			SubTable1: rollt.List{
				"Dangerously odd transhumans",
				"Freeze-dried ancient corpses",
				"Secretive military observers",
				"Eccentric oligarch and minions",
				"Deranged but brilliant scientist",
			},
			SubTable2: rollt.List{
				"Systems breaking down",
				"Foreign sabotage attempt",
				"Black market for the elite",
				"Vault for dangerous pretech",
				"Supply base for pirates",
			},
		},
		{
			Name: "Asteroid base",
			SubTable1: rollt.List{
				"Zealous religious sectarians",
				"Failed rebels from another world",
				"Wage-slave corporate miners",
				"Independent asteroid prospectors",
				"Pirates masquerading as otherwise",
			},
			SubTable2: rollt.List{
				"Life support is threatened",
				"Base needs a new asteroid",
				"Dug out something nasty",
				"Fighting another asteroid",
				"Hit a priceless vein of ore",
			},
		},
		{
			Name: "Remote moon base",
			SubTable1: rollt.List{
				"Unlucky corporate researchers",
				"Reclusive hermit genius",
				"Remnants of a failed colony",
				"Military listening post",
				"Lonely overseers and robot miners",
			},
			SubTable2: rollt.List{
				"Something dark has awoken",
				"Criminals trying to take over",
				"Moon plague breaking out",
				"Desperate for vital supplies",
				"Rich but badly-protected",
			},
		},
		{
			Name: "Ancient orbital ruin",
			SubTable1: rollt.List{
				"Robots of dubious sentience",
				"Trigger-happy scavengers",
				"Government researchers",
				"Military quarantine enforcers",
				"Heirs of the original alien builders",
			},
			SubTable2: rollt.List{
				"Trying to stop it awakening",
				"Meddling with strange tech",
				"Impending tech calamity",
				"A terrible secret is unearthed",
				"Fighting outside interlopers",
			},
		},
		{
			Name: "Research base",
			SubTable1: rollt.List{
				"Experiments that have gotten loose",
				"Scientists from a major local corp",
				"Black-ops governmental researchers",
				"Secret employees of a foreign power",
				"Aliens studying the human locals",
			},
			SubTable2: rollt.List{
				"Perilous research underway",
				"Hideously immoral research",
				"Held hostage by outsiders",
				"Science monsters run amok",
				"Selling black-market tech",
			},
		},
		{
			Name: "Asteroid belt",
			SubTable1: rollt.List{
				"Grizzled belter mine laborers",
				"Ancient automated guardian drones",
				"Survivors of destroyed asteroid base",
				"Pirates hiding out among the rocks",
				"Lonely military patrol base staff",
			},
			SubTable2: rollt.List{
				"Ruptured rock released a peril",
				"Foreign spy ships hide there",
				"Gold rush for new minerals",
				"Ancient ruins dot the rocks",
				"War between rival rocks",
			},
		},
		{
			Name: "Gas giant mine",
			SubTable1: rollt.List{
				"Miserable gas-miner slaves or serfs",
				"Strange robots and their overseers",
				"Scientists studying the alien life",
				"Scrappers in the ruined old mine",
				"Impoverished separatist group",
			},
			SubTable2: rollt.List{
				"Things are emerging below",
				"They need vital supplies",
				"The workers are in revolt",
				"Pirates secretly fuel there",
				"Alien remnants were found",
			},
		},
		{
			Name: "Refueling station",
			SubTable1: rollt.List{
				"Half-crazed hermit caretaker",
				"Sordid purveyors of decadent fun",
				"Extortionate corporate minions",
				"Religious missionaries to travelers",
				"Brainless automated vendors",
			},
			SubTable2: rollt.List{
				"A ship is in severe distress",
				"Pirates have taken over",
				"Has corrupt customs agents",
				"Foreign saboteurs are active",
				"Deep-space alien signal",
			},
		},
	},
}