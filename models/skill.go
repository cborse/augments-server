package models

type Skill struct {
	ID    SkillID
	Name  string
	Core  bool
	Desc1 string
	Desc2 string
	Desc3 string
}

func GetSkill(id SkillID) Skill {
	return gSkills[id]
}

var gSkills = [...]Skill{
	{
		ID:    SKILL_NONE,
		Name:  "NONE",
		Core:  false,
		Desc1: "",
		Desc2: "",
	}, {
		ID:    SKILL_FULL_SIGHT,
		Name:  "FULL SIGHT",
		Core:  true,
		Desc1: "User's accuracy stat increases",
		Desc2: "by 10%.",
	}, {
		ID:    SKILL_FULL_VISION,
		Name:  "FULL VISION",
		Core:  true,
		Desc1: "User's evastion stat increases",
		Desc2: "by 10%.",
	}, {
		ID:    SKILL_FULL_CHOKE,
		Name:  "FULL CHOKE",
		Core:  true,
		Desc1: "Linear actions have increased",
		Desc2: "power; spread actions have ",
		Desc3: "decreased power.",
	}, {
		ID:    SKILL_CYLINDER,
		Name:  "CYLINDER",
		Core:  true,
		Desc1: "Spread actions have increased",
		Desc2: "power; linear actions have",
		Desc3: "decreased power.",
	}, {
		ID:    SKILL_DETECT,
		Name:  "DETECT",
		Core:  true,
		Desc1: "User and opponent become visible",
		Desc2: "if within a 1 ring radius from",
		Desc3: "eachother.",
	}, {
		ID:    SKILL_SCREEN_STRENGTH,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Strength action absored.",
	}, {
		ID:    SKILL_SCREEN_GROUND,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Ground action absored.",
	}, {
		ID:    SKILL_SCREEN_WATER,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Water action absored.",
	}, {
		ID:    SKILL_SCREEN_ICE,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Ice action absored.",
	}, {
		ID:    SKILL_SCREEN_CHEMICAL,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Chemical action absored.",
	}, {
		ID:    SKILL_SCREEN_METAL,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Metal action absored.",
	}, {
		ID:    SKILL_SCREEN_STONE,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Stone action absored.",
	}, {
		ID:    SKILL_SCREEN_SOLAR,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Solar action absored.",
	}, {
		ID:    SKILL_SCREEN_PSYCHE,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Psyche action absored.",
	}, {
		ID:    SKILL_SCREEN_WIND,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Wind action absored.",
	}, {
		ID:    SKILL_SCREEN_ELECTRIC,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Electric action absored.",
	}, {
		ID:    SKILL_SCREEN_SPIRIT,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Spirit action absorbed.",
	}, {
		ID:    SKILL_SCREEN_FIRE,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Fire action absored.",
	}, {
		ID:    SKILL_SCREEN_ILLUSION,
		Name:  "SCREEN",
		Core:  true,
		Desc1: "Creature takes 5% less damage",
		Desc2: "from any Illusion acction",
		Desc3: "absored.",
	}, {
		ID:    SKILL_REPLENISH,
		Name:  "REPLENISH",
		Core:  true,
		Desc1: "Increases users Energy by 5% of ",
		Desc2: "max at the end of each turn.",
	}, {
		ID:    SKILL_RENEW,
		Name:  "RENEW",
		Core:  true,
		Desc1: "Increases users Energy by 10% of",
		Desc2: "max at the end of each turn;",
		Desc3: "lowers Creatures Movement Speed.",
	}, {
		ID:    SKILL_STEADY,
		Name:  "STEADY",
		Core:  true,
		Desc1: "User can only move 1 tile per",
		Desc2: "turn but uses no Energy when",
		Desc3: "moving position.",
	}, {
		ID:    SKILL_RETAIN_IP,
		Name:  "RETAIN I.P.",
		Core:  true,
		Desc1: "Increases Inner Power stat by 50%",
		Desc2: "but user cannot use their last",
		Desc3: "available action.",
	}, {
		ID:    SKILL_RETAIN_ID,
		Name:  "RETAIN I.D.",
		Core:  true,
		Desc1: "Increases Inner Defence stat by",
		Desc2: "50% but user cannot use their",
		Desc3: "last available action.",
	}, {
		ID:    SKILL_RETAIN_OP,
		Name:  "RETAIN O.P.",
		Core:  true,
		Desc1: "Increases Outer Power stat by 50%",
		Desc2: "but user cannot use their last",
		Desc3: "available action.",
	}, {
		ID:    SKILL_RETAIN_OD,
		Name:  "RETAIN O.D.",
		Core:  true,
		Desc1: "Increases Outer Defence stat by",
		Desc2: "50% but user cannot use their",
		Desc3: "last available action.",
	}, {
		ID:    SKILL_RETAIN_MS,
		Name:  "RETAIN M.S.",
		Core:  true,
		Desc1: "Increases Movement Speed stat by",
		Desc2: "50% but user cannot use their",
		Desc3: "last available action.",
	}, {
		ID:    SKILL_RETAIN_AS,
		Name:  "RETAIN A.S.",
		Core:  true,
		Desc1: "Increases Action Speed stat by",
		Desc2: "50% but user cannot use their",
		Desc3: "last available action.",
	}, {
		ID:    SKILL_RETAIN_S,
		Name:  "RETAIN S",
		Core:  true,
		Desc1: "Increases Stamina stat by 50%",
		Desc2: "but user cannot use their last",
		Desc3: "available action.",
	}, {
		ID:    SKILL_RETAIN_A,
		Name:  "RETAIN A",
		Core:  true,
		Desc1: "Increases Accuracy stat by 50%",
		Desc2: "but user cannot use their last",
		Desc3: "available action.",
	}, {
		ID:    SKILL_RETAIN_E,
		Name:  "RETAIN E",
		Core:  true,
		Desc1: "Increases Evasion stat by 50%",
		Desc2: "but user cannot use their last",
		Desc3: "available action.",
	}, {
		ID:    SKILL_EXPERT,
		Name:  "EXPERT",
		Core:  true,
		Desc1: "Increases action power by 25% if",
		Desc2: "it is the same type as the user.",
	}, {
		ID:    SKILL_STEALTHY,
		Name:  "STEALTHY",
		Core:  true,
		Desc1: "Creature remains hidden if they",
		Desc2: "eliminate their foe.",
	}, {
		ID:    SKILL_LUCKY,
		Name:  "LUCKY",
		Core:  false,
		Desc1: "Increases Evasion stat by 100%",
		Desc2: "when the Creature is stationary.",
	}, {
		ID:    SKILL_EFFECTIVE,
		Name:  "EFFECTIVE",
		Core:  false,
		Desc1: "Action always hits at maximum",
		Desc2: "power.",
	}, {
		ID:    SKILL_PRECISE,
		Name:  "PRECISE",
		Core:  false,
		Desc1: "Action cannot miss.",
		Desc2: "",
	}, {
		ID:    SKILL_ORIGAMI,
		Name:  "ORIGAMI",
		Core:  false,
		Desc1: "Increases Evasion stat by 50%.",
		Desc2: "",
	}, {
		ID:    SKILL_ERRATIC,
		Name:  "ERRATIC",
		Core:  false,
		Desc1: "Increases Evasion stat by 100%",
		Desc2: "when the Creature is moving",
		Desc3: "position.",
	}, {
		ID:    SKILL_DECOY,
		Name:  "DECOY",
		Core:  false,
		Desc1: "Creature has a 25% chance of ",
		Desc2: "completely avoiding incoming",
		Desc3: "actions.",
	}, {
		ID:    SKILL_AGILITY,
		Name:  "AGILITY",
		Core:  false,
		Desc1: "Creature ignores any increased",
		Desc2: "change to the opponent's Evasion",
		Desc3: "stat.",
	}, {
		ID:    SKILL_POCKET,
		Name:  "POCKET",
		Core:  false,
		Desc1: "Accuracy and Evasion stats ",
		Desc2: "increase by 25% when opponent is",
		Desc3: "in adjacent tile.",
	}, {
		ID:    SKILL_EQUANIMITY,
		Name:  "EQUANIMITY",
		Core:  false,
		Desc1: "Creature's action cannot miss",
		Desc2: "when the opponent has low health.",
	}, {
		ID:    SKILL_RIFLED,
		Name:  "RIFLED",
		Core:  false,
		Desc1: "Creature's Accuracy stat ",
		Desc2: "increases by 200% when using  ",
		Desc3: "cannon actions.",
	}, {
		ID:    SKILL_BRILLIANT,
		Name:  "BRILLIANT",
		Core:  false,
		Desc1: "Increases Evasion stat by 200%",
		Desc2: "when the Creature is moving",
		Desc3: "position; Creature is detected.",
	}, {
		ID:    SKILL_SCARE_SPRING,
		Name:  "SCARE SPRING",
		Core:  false,
		Desc1: "Creature jumps high in the air",
		Desc2: "when frightened and avoids ",
		Desc3: "incoming action.",
	}, {
		ID:    SKILL_LONG_CLAW,
		Name:  "LONG CLAW",
		Core:  false,
		Desc1: "Power increases by 25% for ",
		Desc2: "slashing actions.",
	}, {
		ID:    SKILL_JAWS,
		Name:  "JAWS",
		Core:  false,
		Desc1: "Power increases by 25% for ",
		Desc2: "biting actions.",
	}, {
		ID:    SKILL_KNUCKLES,
		Name:  "KNUCKLES",
		Core:  false,
		Desc1: "Power increases by 25% for ",
		Desc2: "punching actions.",
	}, {
		ID:    SKILL_BIG_PAW,
		Name:  "BIG PAW",
		Core:  false,
		Desc1: "Power increases by 25% for ",
		Desc2: "crushing actions.",
	}, {
		ID:    SKILL_MAWS,
		Name:  "MAWS",
		Core:  false,
		Desc1: "Power increases by 25% for jet",
		Desc2: "actions.",
	}, {
		ID:    SKILL_ELASTIC,
		Name:  "ELASTIC",
		Core:  false,
		Desc1: "Offensive action covers double",
		Desc2: "in distance.",
	}, {
		ID:    SKILL_GLIDE_EJECT,
		Name:  "GLIDE EJECT",
		Core:  false,
		Desc1: "User moves 3 tiles backward when",
		Desc2: "hit by a super-effective action.",
	}, {
		ID:    SKILL_EQUALITY,
		Name:  "EQUALITY",
		Core:  false,
		Desc1: "User gets a primary action bonus",
		Desc2: "for its secondary type.",
	}, {
		ID:    SKILL_WIND_WISPER,
		Name:  "WIND WISPER",
		Core:  false,
		Desc1: "Power of Spirit-type actions is",
		Desc2: "increased by 25%.",
	}, {
		ID:    SKILL_QUICK_PAW,
		Name:  "QUICK PAW",
		Core:  false,
		Desc1: "Every action used will have ",
		Desc2: "maximum Action speed.",
	}, {
		ID:    SKILL_BUCK_TEETH,
		Name:  "BUCK TEETH",
		Core:  false,
		Desc1: "Biting actions can cause critical",
		Desc2: "damage.",
	}, {
		ID:    SKILL_BRAWLER,
		Name:  "BRAWLER",
		Core:  false,
		Desc1: "Power of Strength-type actions is",
		Desc2: "increased by 25%.",
	}, {
		ID:    SKILL_RUBBER,
		Name:  "RUBBER",
		Core:  false,
		Desc1: "Chance that any incoming action",
		Desc2: "will be redirected at the ",
		Desc3: "attacker.",
	}, {
		ID:    SKILL_FLUFFY,
		Name:  "FLUFFY",
		Core:  false,
		Desc1: "Creature takes 50% less damage ",
		Desc2: "from actions that make contact.",
	}, {
		ID:    SKILL_JELLY,
		Name:  "JELLY",
		Core:  false,
		Desc1: "Creature takes 75% less damage",
		Desc2: "from actions that make contact;",
		Desc3: "user cannot use contact actions.",
	}, {
		ID:    SKILL_POINTY,
		Name:  "POINTY",
		Core:  false,
		Desc1: "Creature's pointy body causes",
		Desc2: "damage to the opponent when there",
		Desc3: "is contact.",
	}, {
		ID:    SKILL_GOO,
		Name:  "GOO",
		Core:  false,
		Desc1: "Opponent's Speed stat lowers",
		Desc2: "when there is contact.",
	}, {
		ID:    SKILL_DENSITY,
		Name:  "DENSITY",
		Core:  false,
		Desc1: "Creature cannot be damaged by",
		Desc2: "contact actions.",
	}, {
		ID:    SKILL_IRON_ARMOUR,
		Name:  "IRON ARMOUR",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from contact actions; users Speed",
		Desc3: "& Stamina stats are lowered.",
	}, {
		ID:    SKILL_WHET_STONE,
		Name:  "WHET STONE",
		Core:  false,
		Desc1: "Creature's Outer Power increases",
		Desc2: "by 25% when hit by a Metal-type",
		Desc3: "action.",
	}, {
		ID:    SKILL_SUPERCOOLED,
		Name:  "SUPERCOOLED",
		Core:  false,
		Desc1: "Contact with this Creature may",
		Desc2: "cause freezing.",
	}, {
		ID:    SKILL_CHARGED,
		Name:  "CHARGED",
		Core:  false,
		Desc1: "Contact with this Creature may",
		Desc2: "cause paralysis.",
	}, {
		ID:    SKILL_STINGERS,
		Name:  "STINGERS",
		Core:  false,
		Desc1: "Contact with this Creature may",
		Desc2: "cause paralysis.",
	}, {
		ID:    SKILL_ALERT,
		Name:  "ALERT",
		Core:  false,
		Desc1: "User can detect any Creatures",
		Desc2: "within the immediate vicinity.",
	}, {
		ID:    SKILL_OBSERVANT,
		Name:  "OBSERVANT",
		Core:  false,
		Desc1: "User can detect any Creatures",
		Desc2: "within the vicinity.",
	}, {
		ID:    SKILL_VIGILANCE,
		Name:  "VIGILANCE",
		Core:  false,
		Desc1: "User can detect any Creatures ",
		Desc2: "within the area.",
	}, {
		ID:    SKILL_VANTAGE,
		Name:  "VANTAGE",
		Core:  false,
		Desc1: "User and opponents become easier",
		Desc2: "to detect.",
	}, {
		ID:    SKILL_SONAR,
		Name:  "SONAR",
		Core:  false,
		Desc1: "User can detect any large",
		Desc2: "Creatures within the vicinity.",
	}, {
		ID:    SKILL_ECHOLOCATE,
		Name:  "ECHOLOCATE",
		Core:  false,
		Desc1: "User can detect any small ",
		Desc2: "Creatures within the vicinity.",
	}, {
		ID:    SKILL_OPPORTUNIST,
		Name:  "OPPORTUNIST",
		Core:  false,
		Desc1: "User can detect any seriously",
		Desc2: "injured Creature.",
	}, {
		ID:    SKILL_CLOUDY,
		Name:  "CLOUDY",
		Core:  false,
		Desc1: "Creature's position can ",
		Desc2: "spontaneously return to hidden.",
	}, {
		ID:    SKILL_SNOW_COAT,
		Name:  "SNOW COAT",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Ice-type actions.",
	}, {
		ID:    SKILL_FURRY,
		Name:  "FURRY",
		Core:  false,
		Desc1: "Creature takes reduced damage ",
		Desc2: "from Ice-type actions.",
	}, {
		ID:    SKILL_BLUBBER,
		Name:  "BLUBBER",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Ice-type actions; Creatures",
		Desc3: "Movement Speed stat is lowered.",
	}, {
		ID:    SKILL_SUMMIT,
		Name:  "SUMMIT",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Ice-type actions.",
	}, {
		ID:    SKILL_HYDROPHOBIC,
		Name:  "HYDROPHOBIC",
		Core:  false,
		Desc1: "Creature takes reduced damage ",
		Desc2: "from Water-type actions.",
	}, {
		ID:    SKILL_IRRIGATE,
		Name:  "IRRIGATE",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Water-type actions and has",
		Desc3: "boosted stats.",
	}, {
		ID:    SKILL_SOAKED,
		Name:  "SOAKED",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Water-type actions and has",
		Desc3: "boosted stats.",
	}, {
		ID:    SKILL_SOILED,
		Name:  "SOILED",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Water-type actions and has a",
		Desc3: "boosted stat.",
	}, {
		ID:    SKILL_RETARDANT,
		Name:  "RETARDANT",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Fire-type actions.",
	}, {
		ID:    SKILL_FIREPROOF,
		Name:  "FIREPROOF",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Fire-type actions and their",
		Desc3: "effects.",
	}, {
		ID:    SKILL_RUBBERIZED,
		Name:  "RUBBERIZED",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Electric-type actions.",
	}, {
		ID:    SKILL_INSULATED,
		Name:  "INSULATED",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Electric-type actions.",
	}, {
		ID:    SKILL_GROUNDED,
		Name:  "GROUNDED",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Electric-type actions.",
	}, {
		ID:    SKILL_INSENSIBLE,
		Name:  "INSENSIBLE",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Psyche-type actions.",
	}, {
		ID:    SKILL_CRYSTAL_BODY,
		Name:  "CRYSTAL BODY",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Solar-type actions.",
	}, {
		ID:    SKILL_SOLAR_SHIELD,
		Name:  "SOLAR SHIELD",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Solar-type actions.",
	}, {
		ID:    SKILL_GRANITE,
		Name:  "GRANITE",
		Core:  false,
		Desc1: "Creature takes reduced damage",
		Desc2: "from Watr-type and Wind-type",
		Desc3: "actions.",
	}, {
		ID:    SKILL_BALLOON,
		Name:  "BALLOON",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Wind-type actions and is pushed",
		Desc3: "to the end of the action cover.",
	}, {
		ID:    SKILL_SAILS,
		Name:  "SAILS",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Wind-type actions and has a",
		Desc3: "boosted stat.",
	}, {
		ID:    SKILL_LEVITATE,
		Name:  "LEVITATE",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Ground-type actions.",
	}, {
		ID:    SKILL_MEMBRANE,
		Name:  "MEMBRANE",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Chemical-type actions.",
	}, {
		ID:    SKILL_FILTER,
		Name:  "FILTER",
		Core:  false,
		Desc1: "Creature is not affected by gas,",
		Desc2: "smoke, or dust actions.",
	}, {
		ID:    SKILL_LONG_NECK,
		Name:  "LONG NECK",
		Core:  false,
		Desc1: "Creature avoids being hit by the",
		Desc2: "full force of an action.",
	}, {
		ID:    SKILL_ARMOUR_SUIT,
		Name:  "ARMOUR SUIT",
		Core:  false,
		Desc1: "Creature only takes damage from",
		Desc2: "super-effective actions; Speed",
		Desc3: "stats are lowered.",
	}, {
		ID:    SKILL_BONELESS,
		Name:  "BONELESS",
		Core:  false,
		Desc1: "Creature is uneffected by any",
		Desc2: "crushing actions.",
	}, {
		ID:    SKILL_SCAVENGER,
		Name:  "SCAVENGER",
		Core:  false,
		Desc1: "Creature gets double the increase",
		Desc2: "from elimination perks.",
	}, {
		ID:    SKILL_HARVEST,
		Name:  "HARVEST",
		Core:  false,
		Desc1: "Any elimination perk attained ",
		Desc2: "will bring the Creature's health",
		Desc3: "back to maximum.",
	}, {
		ID:    SKILL_VITALITY,
		Name:  "VITALITY",
		Core:  false,
		Desc1: "Any elimination perk attained ",
		Desc2: "will bring the Creature's energy",
		Desc3: "back to maximum.",
	}, {
		ID:    SKILL_AFTERLIFE,
		Name:  "AFTERLIFE",
		Core:  false,
		Desc1: "The Creature's elimination perk",
		Desc2: "will cause damage to the",
		Desc3: "opponent.",
	}, {
		ID:    SKILL_EFFICIENT,
		Name:  "EFFICIENT",
		Core:  false,
		Desc1: "Creature uses less energy when",
		Desc2: "attacking or moving position.",
	}, {
		ID:    SKILL_ENDURANCE,
		Name:  "ENDURANCE",
		Core:  false,
		Desc1: "Creature uses less energy when",
		Desc2: "moving position.",
	}, {
		ID:    SKILL_PRESSURE,
		Name:  "PRESSURE",
		Core:  false,
		Desc1: "Opponent's use more energy when",
		Desc2: "attacking.",
	}, {
		ID:    SKILL_LOAD,
		Name:  "LOAD",
		Core:  false,
		Desc1: "Creature recovers energy when hit",
		Desc2: "by an Electric-type action.",
	}, {
		ID:    SKILL_WINDMILL,
		Name:  "WINDMILL",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Wind-type actions and recovers ",
		Desc3: "the equivalent in energy.",
	}, {
		ID:    SKILL_CHUBBYWEIGHT,
		Name:  "CHUBBYWEIGHT",
		Core:  false,
		Desc1: "Creature recovers energy more",
		Desc2: "quickly when energy is low.",
	}, {
		ID:    SKILL_ADORABLE,
		Name:  "ADORABLE",
		Core:  false,
		Desc1: "The Creature's adorableness",
		Desc2: "causes a chance that the opponent",
		Desc3: "will not attack.",
	}, {
		ID:    SKILL_HYDRATE,
		Name:  "HYDRATE",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Water-type actions and recovers",
		Desc3: "the equivalent in health.",
	}, {
		ID:    SKILL_SOAK,
		Name:  "SOAK",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Water-type actions and recovers",
		Desc3: "20% of max health.",
	}, {
		ID:    SKILL_DURABLE,
		Name:  "DURABLE",
		Core:  false,
		Desc1: "Creature cannot be eliminated",
		Desc2: "from full health with one attack.",
	}, {
		ID:    SKILL_PERSISTENT,
		Name:  "PERSISTENT",
		Core:  false,
		Desc1: "Creature cannot be eliminated",
		Desc2: "from full health in one turn.",
	}, {
		ID:    SKILL_REGENERATE,
		Name:  "REGENERATE",
		Core:  false,
		Desc1: "User recovers health each turn.",
		Desc2: "",
	}, {
		ID:    SKILL_OVERDRIVE,
		Name:  "OVERDRIVE",
		Core:  false,
		Desc1: "Creature acts with full effort",
		Desc2: "at the expense of their health.",
	}, {
		ID:    SKILL_RESOLVED,
		Name:  "RESOLVED",
		Core:  false,
		Desc1: "Creature has increased stats at",
		Desc2: "the expnse of their health and",
		Desc3: "energy.",
	}, {
		ID:    SKILL_BLOODSUCKER,
		Name:  "BLOODSUCKER",
		Core:  false,
		Desc1: "Creature converts 20% of the ",
		Desc2: "damage dealt to the opponent into",
		Desc3: "gained health.",
	}, {
		ID:    SKILL_ROOTED,
		Name:  "ROOTED",
		Core:  false,
		Desc1: "Creatures movement is limited to",
		Desc2: "1-tile position per turn but",
		Desc3: "increases Inner Defence stat.",
	}, {
		ID:    SKILL_TORTOISE,
		Name:  "TORTOISE",
		Core:  false,
		Desc1: "Creatures movement is limited to",
		Desc2: "1-tile position per turn but",
		Desc3: "increases Outer Defence stat.",
	}, {
		ID:    SKILL_SEDATE,
		Name:  "SEDATE",
		Core:  false,
		Desc1: "Creatures movement is limited to ",
		Desc2: "2-tile positions per turn but",
		Desc3: "increases both Defence stats.",
	}, {
		ID:    SKILL_KINGLY,
		Name:  "KINGLY",
		Core:  false,
		Desc1: "Opponents cannot move into a ",
		Desc2: "1-ring radius of the Creature.",
	}, {
		ID:    SKILL_RACER,
		Name:  "RACER",
		Core:  false,
		Desc1: "Creature will always succeed in",
		Desc2: "gaining position when competing",
		Desc3: "for a tile.",
	}, {
		ID:    SKILL_ROYAL_GUARD,
		Name:  "ROYAL GUARD",
		Core:  false,
		Desc1: "Nothing can prevent the Creature",
		Desc2: "from moving positions.",
	}, {
		ID:    SKILL_EKRANOPLAN,
		Name:  "EKRANOPLAN",
		Core:  false,
		Desc1: "Creature uses less energy and",
		Desc2: "acts faster when moving position.",
	}, {
		ID:    SKILL_RESILIENT,
		Name:  "RESILIENT",
		Core:  false,
		Desc1: "Creatures Defence stats double",
		Desc2: "when HP is at 25% of max.",
	}, {
		ID:    SKILL_ADAPTABLE,
		Name:  "ADAPTABLE",
		Core:  false,
		Desc1: "Creatures Inner Defnece stat",
		Desc2: "doubles when HP is at 50% of max.",
	}, {
		ID:    SKILL_ADAPTATION,
		Name:  "ADAPTATION",
		Core:  false,
		Desc1: "All attacks taken are defended ",
		Desc2: "using Outer Defence stat.",
	}, {
		ID:    SKILL_ABSOLUTE,
		Name:  "ABSOLUTE",
		Core:  false,
		Desc1: "Creature's stats cannot be",
		Desc2: "lowered by any means.",
	}, {
		ID:    SKILL_COMPETITVE,
		Name:  "COMPETITVE",
		Core:  false,
		Desc1: "Outer Power stat increases to ",
		Desc2: "match opponents.",
	}, {
		ID:    SKILL_FLOW,
		Name:  "FLOW",
		Core:  false,
		Desc1: "Eliminating an opponent increases",
		Desc2: "Inner Power stat by 25%.",
	}, {
		ID:    SKILL_ANTICIPATION,
		Name:  "ANTICIPATION",
		Core:  false,
		Desc1: "Increased Action Speed when",
		Desc2: "opponent and Creature are",
		Desc3: "attacking on the same turn.",
	}, {
		ID:    SKILL_HESITANT,
		Name:  "HESITANT",
		Core:  false,
		Desc1: "Opponents hesitate when attacking",
		Desc2: "this Creature in a 1-ring radius.",
	}, {
		ID:    SKILL_EXPLOSIVE,
		Name:  "EXPLOSIVE",
		Core:  false,
		Desc1: "Creature's Action Speed is ",
		Desc2: "doubled when attacking opponents",
		Desc3: "in close range.",
	}, {
		ID:    SKILL_CAPACITOR,
		Name:  "CAPACITOR",
		Core:  false,
		Desc1: "Creature's Speed stats increase",
		Desc2: "when hit with an Electric-type",
		Desc3: "action.",
	}, {
		ID:    SKILL_FRIGHTENING,
		Name:  "FRIGHTENING",
		Core:  false,
		Desc1: "Opponent's Inner Power stat is",
		Desc2: "lowered if attacked by this",
		Desc3: "Creature.",
	}, {
		ID:    SKILL_TERRIFYING,
		Name:  "TERRIFYING",
		Core:  false,
		Desc1: "Opponent's Outer Power stat is",
		Desc2: "lowered if attacked by this",
		Desc3: "Creature.",
	}, {
		ID:    SKILL_TEMPER,
		Name:  "TEMPER",
		Core:  false,
		Desc1: "Creature's Outer Defence stat ",
		Desc2: "increases whenever they are",
		Desc3: "attacked.",
	}, {
		ID:    SKILL_MOODY,
		Name:  "MOODY",
		Core:  false,
		Desc1: "Creature's Action Speed stat",
		Desc2: "increases whenever they are ",
		Desc3: "attacked.",
	}, {
		ID:    SKILL_RELATIVITY,
		Name:  "RELATIVITY",
		Core:  false,
		Desc1: "Opponent's Speed stats decrease",
		Desc2: "the closer they are to this ",
		Desc3: "Creature.",
	}, {
		ID:    SKILL_MICROPHOBIA,
		Name:  "MICROPHOBIA",
		Core:  false,
		Desc1: "Lowers Outer Power and Inner",
		Desc2: "Defence stats of much larger",
		Desc3: "opponents.",
	}, {
		ID:    SKILL_FRENZIED,
		Name:  "FRENZIED",
		Core:  false,
		Desc1: "Creature's Outer Power stat",
		Desc2: "increases but Evasion stat lowers",
		Desc3: "when the Creature has low health.",
	}, {
		ID:    SKILL_BURDEN,
		Name:  "BURDEN",
		Core:  false,
		Desc1: "If the Creature is the last",
		Desc2: "remaining team member their Inner",
		Desc3: "Power stat is maxed.",
	}, {
		ID:    SKILL_CHARCOAL,
		Name:  "CHARCOAL",
		Core:  false,
		Desc1: "All the Creature's stats increase",
		Desc2: "when hit by a Fire-type action.",
	}, {
		ID:    SKILL_GEAR_SHIFT,
		Name:  "GEAR SHIFT",
		Core:  false,
		Desc1: "User swaps Action Speed stat",
		Desc2: "with the opponent if opponent has",
		Desc3: "a higher stat.",
	}, {
		ID:    SKILL_SUPERCELL,
		Name:  "SUPERCELL",
		Core:  false,
		Desc1: "Creature's Inner Power and Outer",
		Desc2: "Power stats increase when hit by",
		Desc3: "a Water-type action.",
	}, {
		ID:    SKILL_LIFT,
		Name:  "LIFT",
		Core:  false,
		Desc1: "Creature's Inner Power and Outer",
		Desc2: "Power stats increase when hit by",
		Desc3: "a Wind-type action.",
	}, {
		ID:    SKILL_SUGAR_RUSH,
		Name:  "SUGAR RUSH",
		Core:  false,
		Desc1: "Creature's Speed stats increase",
		Desc2: "each turn of battle.",
	}, {
		ID:    SKILL_FORTUNATE,
		Name:  "FORTUNATE",
		Core:  false,
		Desc1: "If the Creature acts after an",
		Desc2: "opponent, the opponent's attack",
		Desc3: "will do 20% less damage.",
	}, {
		ID:    SKILL_DIVERT,
		Name:  "DIVERT",
		Core:  false,
		Desc1: "The Creature will intercept any",
		Desc2: "uneffective close-range opponent",
		Desc3: "actions directed at an ally.",
	}, {
		ID:    SKILL_RECOMBINE,
		Name:  "RECOMBINE",
		Core:  false,
		Desc1: "If the Creature eliminates an ",
		Desc2: "opponent it will absorb the ",
		Desc3: "opponent's highest stat.",
	}, {
		ID:    SKILL_FIRE_HAZARD,
		Name:  "FIRE HAZARD",
		Core:  false,
		Desc1: "Damage from Fire-type actions",
		Desc2: "will cause continuous damage;",
		Desc3: "Inner Power stat is maxed.",
	}, {
		ID:    SKILL_STALKER,
		Name:  "STALKER",
		Core:  false,
		Desc1: "Creature can use 3-tile movements",
		Desc2: "without being detected.",
	}, {
		ID:    SKILL_SHADOW,
		Name:  "SHADOW",
		Core:  false,
		Desc1: "Creature can act without being",
		Desc2: "detected.",
	}, {
		ID:    SKILL_CONCEALED,
		Name:  "CONCEALED",
		Core:  false,
		Desc1: "Creature remains hidden from",
		Desc2: "detection actions or skills.",
	}, {
		ID:    SKILL_CAMOUFLAGED,
		Name:  "CAMOUFLAGED",
		Core:  false,
		Desc1: "Creature returns to hidden if it",
		Desc2: "is stationary for one turn.",
	}, {
		ID:    SKILL_CHAMELEON,
		Name:  "CHAMELEON",
		Core:  false,
		Desc1: "Creature's position can never be",
		Desc2: "detected.",
	}, {
		ID:    SKILL_CURSED_CLOAK,
		Name:  "CURSED CLOAK",
		Core:  false,
		Desc1: "If the Creature is hit by a",
		Desc2: "Spiri-type action their position",
		Desc3: "will return to hidden.",
	}, {
		ID:    SKILL_DIRE_LUCK,
		Name:  "DIRE LUCK",
		Core:  false,
		Desc1: "If this Creature is attacked the",
		Desc2: "opponent will lose 10% of max ",
		Desc3: "Heatlh and 30% of max Energy.",
	}, {
		ID:    SKILL_GRIM_LUCK,
		Name:  "GRIM LUCK",
		Core:  false,
		Desc1: "If this Creature is attacked the",
		Desc2: "opponent will lose 10% of max ",
		Desc3: "Health.",
	}, {
		ID:    SKILL_AWFUL_LUCK,
		Name:  "AWFUL LUCK",
		Core:  false,
		Desc1: "If this Creature is attacked the",
		Desc2: "opponent will lose 30% of max ",
		Desc3: "Energy.",
	}, {
		ID:    SKILL_NESTING,
		Name:  "NESTING",
		Core:  false,
		Desc1: "Defensive stats are increased",
		Desc2: "when the Creature has higher",
		Desc3: "health.",
	}, {
		ID:    SKILL_SLYLY,
		Name:  "SLYLY",
		Core:  false,
		Desc1: "Creature cannot be blocked from",
		Desc2: "moving positions.",
	}, {
		ID:    SKILL_MAZE,
		Name:  "MAZE",
		Core:  false,
		Desc1: "If the Creature is detected there",
		Desc2: "will be multiple illusions",
		Desc3: "nearby.",
	}, {
		ID:    SKILL_DESPAIR,
		Name:  "DESPAIR",
		Core:  false,
		Desc1: "The opponent who detects this",
		Desc2: "Creature will have all stats",
		Desc3: "reduced.",
	}, {
		ID:    SKILL_CORK,
		Name:  "CORK",
		Core:  false,
		Desc1: "Creature is immune to damage from",
		Desc2: "Water-type actions and Solar-type",
		Desc3: "actions.",
	}, {
		ID:    SKILL_CRACKEDSHELL,
		Name:  "CRACKEDSHELL",
		Core:  false,
		Desc1: "If the user losses 50% of health",
		Desc2: "from a contact action, defences",
		Desc3: "reduce and speed stats max.",
	},
}
