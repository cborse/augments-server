package models

type ActionID uint32
type SkillID uint32
type SpeciesID uint32
type CoverID uint8
type EffectID uint8
type PerkID uint8
type RarityID uint8
type StyleID uint8
type TypeID uint8

const (
	ACTION_NONE = ActionID(iota)
	ACTION_SMACK_STRENGTH
	ACTION_SMACK_GROUND
	ACTION_SMACK_WATER
	ACTION_SMACK_ICE
	ACTION_SMACK_CHEMICAL
	ACTION_SMACK_METAL
	ACTION_SMACK_STONE
	ACTION_SMACK_SOLAR
	ACTION_SMACK_PSYCHE
	ACTION_SMACK_WIND
	ACTION_SMACK_ELECTRIC
	ACTION_SMACK_SPIRIT
	ACTION_SMACK_FIRE
	ACTION_SMACK_ILLUSION
	ACTION_JAB_STRENGTH
	ACTION_JAB_GROUND
	ACTION_JAB_WATER
	ACTION_JAB_ICE
	ACTION_JAB_CHEMICAL
	ACTION_JAB_METAL
	ACTION_JAB_STONE
	ACTION_HOOK_STRENGTH
	ACTION_HOOK_GROUND
	ACTION_HOOK_WATER
	ACTION_HOOK_ICE
	ACTION_HOOK_CHEMICAL
	ACTION_HOOK_METAL
	ACTION_HOOK_STONE
	ACTION_BOOM_SOLAR
	ACTION_BOOM_PSYCHE
	ACTION_BOOM_WIND
	ACTION_BOOM_ELECTRIC
	ACTION_BOOM_SPIRIT
	ACTION_BOOM_FIRE
	ACTION_BOOM_ILLUSION
	ACTION_SPURT_SOLAR
	ACTION_SPURT_PSYCHE
	ACTION_SPURT_WIND
	ACTION_SPURT_ELECTRIC
	ACTION_SPURT_SPIRIT
	ACTION_SPURT_FIRE
	ACTION_SPURT_ILLUSION
	ACTION_MEDITATE
	ACTION_BARRICADE
	ACTION_HONE
	ACTION_FORTIFY
	ACTION_HASTEN
	ACTION_BRISKEN
	ACTION_CONDITION
	ACTION_MARK
	ACTION_NIMBLE
	ACTION_BLOCK
	ACTION_CHECK
	ACTION_SLIP
	ACTION_DODGE
	ACTION_REVERSAL
	ACTION_ADVANCE
	ACTION_BOUND
	ACTION_DASH
	ACTION_RECOVER
	ACTION_MAUL
	ACTION_CRUSH
	ACTION_LASH
	ACTION_CHARGE
	ACTION_TAIL_HAMMER
	ACTION_SUFFOCATION
	ACTION_SLAM
	ACTION_SNARE_CRUSH
	ACTION_IMPALEMENT
	ACTION_TONGUE_SMACK
	ACTION_HEAD_SMASH
	ACTION_ENERGY_DRAIN
	ACTION_CANNON_SHOT
	ACTION_THROW_TACKLE
	ACTION_SPIN_KICK
	ACTION_DRAIN_BITE
	ACTION_ROAR
	ACTION_SILK_STRANDS
	ACTION_DEFLECTION
	ACTION_GOO_STRIKE
	ACTION_BUNKER
	ACTION_SUMO_TOSS
	ACTION_TEEP
	ACTION_SLUMBER
	ACTION_EARTHQUAKE
	ACTION_METEOR_STORM
	ACTION_TUNNELTHRASH
	ACTION_TSUNAMI
	ACTION_TIMBER
	ACTION_PRIMAL_ROAR
	ACTION_EARTH_BLAST
	ACTION_CRYSTAL_HEAL
	ACTION_HYDRA_CHOMP
	ACTION_CASCADE
	ACTION_SALT_BITE
	ACTION_SUPRCRITICAL
	ACTION_DROWN
	ACTION_HURRICANE
	ACTION_TYPHOON
	ACTION_AQUA_WALLOP
	ACTION_HYDRA_BLADE
	ACTION_HYDRA_JET
	ACTION_TAIL_WHACK
	ACTION_FLASH_FLOOD
	ACTION_RISE_CURRENT
	ACTION_EBB_CURRENT
	ACTION_BUBBLE_BURST
	ACTION_ICE_JET
	ACTION_FROST_FANG
	ACTION_AVALANCHE
	ACTION_FROST_WAVE
	ACTION_ICE_DRILL
	ACTION_SNOWBALL
	ACTION_ICEQUAKE
	ACTION_FLASH_FREEZE
	ACTION_INJECTION
	ACTION_TOXIC_SPRAY
	ACTION_LIQUIFY
	ACTION_HONEY_SHOT
	ACTION_STICKY_SAP
	ACTION_PARA_DUST
	ACTION_PARA_POWDER
	ACTION_STEEL_CHOMP
	ACTION_DISCHARGE
	ACTION_MERCURY_JET
	ACTION_LEAD_HAMMER
	ACTION_DETONATION
	ACTION_MACE_STRIKE
	ACTION_REND
	ACTION_CRUSHER
	ACTION_MAGNET_SNIPE
	ACTION_METEOR_FALL
	ACTION_DRILL_CHARGE
	ACTION_METAL_WAVE
	ACTION_GYRO_SMASH
	ACTION_TORPEDO
	ACTION_SNAP_TRAP
	ACTION_GEAR_SPIN
	ACTION_STONE_DAGGER
	ACTION_STONE_SHARD
	ACTION_ROCK_CHOMP
	ACTION_ROLL_CHARGE
	ACTION_GRANITE_FIST
	ACTION_STONE_BLAST
	ACTION_SOLAR_JET
	ACTION_PHOTON_BEAM
	ACTION_RADIATION
	ACTION_LEAF_BURST
	ACTION_ROOT_LASH
	ACTION_SUN_STRIKE
	ACTION_SOLAR_STORM
	ACTION_LASER_BEAM
	ACTION_GULP
	ACTION_DAZZLE
	ACTION_SOLAR_SHIELD
	ACTION_SYNTHESIS
	ACTION_COTTONBURROW
	ACTION_GLEAM_BURST
	ACTION_SOLAR_RUSH
	ACTION_AWE_STRIKE
	ACTION_DREAD_AURA
	ACTION_SUPERFRIGHT
	ACTION_BARRIER
	ACTION_SUSPENSION
	ACTION_HYPNOTIZE
	ACTION_IMPAIR
	ACTION_IMPEDIMENT
	ACTION_DESPAIRING
	ACTION_WIND_SURGE
	ACTION_WIND_TERROR
	ACTION_GALE_CRUSH
	ACTION_TORNADO
	ACTION_CYCLONE
	ACTION_WIND_ARROW
	ACTION_SKYDIVE
	ACTION_WIND_SHIELD
	ACTION_TURBULENCE
	ACTION_LULLING_COO
	ACTION_WINDED
	ACTION_WIND_SPRINT
	ACTION_WIND_RUSH
	ACTION_CLOUD_SHIELD
	ACTION_VOLT_CHOMP
	ACTION_VOLT_SURGE
	ACTION_SHOCK_CHARGE
	ACTION_BOLT_SPHERE
	ACTION_LIGHTNING
	ACTION_THUNDERBOLT
	ACTION_SPARK
	ACTION_THUNDER_ROAR
	ACTION_STATIC
	ACTION_ELECTRO_PUSH
	ACTION_SHADOWSTRIKE
	ACTION_VENGEANCE
	ACTION_SEVERANCE
	ACTION_SHARED_FATE
	ACTION_SPIRIT_REND
	ACTION_SPIRIT_BRAY
	ACTION_DEADLY_CURSE
	ACTION_BLAZE_BITE
	ACTION_FIRE_JET
	ACTION_TORCH
	ACTION_FLAMETHROWER
	ACTION_INFERNO
	ACTION_ERUPTION
	ACTION_PYROCLASTIC
	ACTION_FLAME_BURST
	ACTION_BOMB_DROP
	ACTION_LAVA_BLAST
	ACTION_COMBUSTION
	ACTION_COVER_SMOKE
	ACTION_BLACK_POWDER
	ACTION_FEINT_SNAP
	ACTION_SHADOW_SCARE
	ACTION_TELEPORT
	ACTION_FOLD_DODGE
	ACTION_COTTONTEMPER
	ACTION_MOONLIGHT
	ACTION_WAR_DANCE
	ACTION_TRUMPET_ROAR
	ACTION_CHEST_DRUM
	ACTION_WOOL_STIFFEN
	ACTION_SPRINT
	ACTION_RUSH
	ACTION_SPRING_HOP
	ACTION_VOLT_SPRINT
	ACTION_CHEM_CRUNCH
	ACTION_VOLT_RUSH
	ACTION_SLINGSTONE
	ACTION_SANDBLAST
	ACTION_INK_SCREEN
	ACTION_WEB_SHOT
	ACTION_POWER_PUNCH
	ACTION_PLOD
	ACTION_MIRAGE
	ACTION_PEEK_A_BOO
	ACTION_COTTON_BURROW
	ACTION_COTTON_TEMPER
	ACTION_TUNNEL_THRASH
	ACTION_SHADOW_STRIKE
	ACTION_SUPER_FRIGHT
	ACTION_BLIZZARD
	ACTION_FLOOD_CURRENT
	ACTION_SOLAR_SPRINT
	ACTION_HIBERNATE
	ACTION_RAINBOW_BEAM
	ACTION_BLOCKADE
	ACTION_LIGHTNING_BOLT
	ACTION_SOUL_FEAST
	ACTION_WEB_NET
	ACTION_SLASHER

	ACTION_COUNT
)

const (
	COVER_NONE = CoverID(iota)
	COVER_RING1
	COVER_RING2
	COVER_ARC2
	COVER_ARC3
	COVER_ARC4
	COVER_ARC5
	COVER_TILE1
	COVER_TILE2
	COVER_TILE3
	COVER_TILE4
	COVER_TILE5
	COVER_LINE2
	COVER_LINE3
	COVER_LINE4
	COVER_LINE5
)

const (
	EFFECT_NONE = EffectID(iota)
)

const (
	PERK_NONE = PerkID(iota)
	PERK_HEALTH
	PERK_ENERGY
	PERK_INNER_POWER
	PERK_INNER_DEFENSE
	PERK_OUTER_POWER
	PERK_OUTER_DEFENSE
	PERK_MOVEMENT_SPEED
	PERK_ACTION_SPEED
	PERK_STAMINA
	PERK_ACCURACY
	PERK_EVASION
)

const (
	RARITY_COMMON = RarityID(iota)
	RARITY_UNCOMMON
	RARITY_RARE
)

const (
	SKILL_NONE = SkillID(iota)
	SKILL_FULL_SIGHT
	SKILL_FULL_VISION
	SKILL_FULL_CHOKE
	SKILL_CYLINDER
	SKILL_DETECT
	SKILL_SCREEN_STRENGTH
	SKILL_SCREEN_GROUND
	SKILL_SCREEN_WATER
	SKILL_SCREEN_ICE
	SKILL_SCREEN_CHEMICAL
	SKILL_SCREEN_METAL
	SKILL_SCREEN_STONE
	SKILL_SCREEN_SOLAR
	SKILL_SCREEN_PSYCHE
	SKILL_SCREEN_WIND
	SKILL_SCREEN_ELECTRIC
	SKILL_SCREEN_SPIRIT
	SKILL_SCREEN_FIRE
	SKILL_SCREEN_ILLUSION
	SKILL_REPLENISH
	SKILL_RENEW
	SKILL_STEADY
	SKILL_RETAIN_IP
	SKILL_RETAIN_ID
	SKILL_RETAIN_OP
	SKILL_RETAIN_OD
	SKILL_RETAIN_MS
	SKILL_RETAIN_AS
	SKILL_RETAIN_S
	SKILL_RETAIN_A
	SKILL_RETAIN_E
	SKILL_EXPERT
	SKILL_STEALTHY
	SKILL_LUCKY
	SKILL_EFFECTIVE
	SKILL_PRECISE
	SKILL_ORIGAMI
	SKILL_ERRATIC
	SKILL_DECOY
	SKILL_AGILITY
	SKILL_POCKET
	SKILL_EQUANIMITY
	SKILL_RIFLED
	SKILL_BRILLIANT
	SKILL_SCARE_SPRING
	SKILL_LONG_CLAW
	SKILL_JAWS
	SKILL_KNUCKLES
	SKILL_BIG_PAW
	SKILL_MAWS
	SKILL_ELASTIC
	SKILL_GLIDE_EJECT
	SKILL_EQUALITY
	SKILL_WIND_WISPER
	SKILL_QUICK_PAW
	SKILL_BUCK_TEETH
	SKILL_BRAWLER
	SKILL_RUBBER
	SKILL_FLUFFY
	SKILL_JELLY
	SKILL_POINTY
	SKILL_GOO
	SKILL_DENSITY
	SKILL_IRON_ARMOUR
	SKILL_WHET_STONE
	SKILL_SUPERCOOLED
	SKILL_CHARGED
	SKILL_STINGERS
	SKILL_ALERT
	SKILL_OBSERVANT
	SKILL_VIGILANCE
	SKILL_VANTAGE
	SKILL_SONAR
	SKILL_ECHOLOCATE
	SKILL_OPPORTUNIST
	SKILL_CLOUDY
	SKILL_SNOW_COAT
	SKILL_FURRY
	SKILL_BLUBBER
	SKILL_SUMMIT
	SKILL_HYDROPHOBIC
	SKILL_IRRIGATE
	SKILL_SOAKED
	SKILL_SOILED
	SKILL_RETARDANT
	SKILL_FIREPROOF
	SKILL_RUBBERIZED
	SKILL_INSULATED
	SKILL_GROUNDED
	SKILL_INSENSIBLE
	SKILL_CRYSTAL_BODY
	SKILL_SOLAR_SHIELD
	SKILL_GRANITE
	SKILL_BALLOON
	SKILL_SAILS
	SKILL_LEVITATE
	SKILL_MEMBRANE
	SKILL_FILTER
	SKILL_LONG_NECK
	SKILL_ARMOUR_SUIT
	SKILL_BONELESS
	SKILL_SCAVENGER
	SKILL_HARVEST
	SKILL_VITALITY
	SKILL_AFTERLIFE
	SKILL_EFFICIENT
	SKILL_ENDURANCE
	SKILL_PRESSURE
	SKILL_LOAD
	SKILL_WINDMILL
	SKILL_CHUBBYWEIGHT
	SKILL_ADORABLE
	SKILL_HYDRATE
	SKILL_SOAK
	SKILL_DURABLE
	SKILL_PERSISTENT
	SKILL_REGENERATE
	SKILL_OVERDRIVE
	SKILL_RESOLVED
	SKILL_BLOODSUCKER
	SKILL_ROOTED
	SKILL_TORTOISE
	SKILL_SEDATE
	SKILL_KINGLY
	SKILL_RACER
	SKILL_ROYAL_GUARD
	SKILL_EKRANOPLAN
	SKILL_RESILIENT
	SKILL_ADAPTABLE
	SKILL_ADAPTATION
	SKILL_ABSOLUTE
	SKILL_COMPETITVE
	SKILL_FLOW
	SKILL_ANTICIPATION
	SKILL_HESITANT
	SKILL_EXPLOSIVE
	SKILL_CAPACITOR
	SKILL_FRIGHTENING
	SKILL_TERRIFYING
	SKILL_TEMPER
	SKILL_MOODY
	SKILL_RELATIVITY
	SKILL_MICROPHOBIA
	SKILL_FRENZIED
	SKILL_BURDEN
	SKILL_CHARCOAL
	SKILL_GEAR_SHIFT
	SKILL_SUPERCELL
	SKILL_LIFT
	SKILL_SUGAR_RUSH
	SKILL_FORTUNATE
	SKILL_DIVERT
	SKILL_RECOMBINE
	SKILL_FIRE_HAZARD
	SKILL_STALKER
	SKILL_SHADOW
	SKILL_CONCEALED
	SKILL_CAMOUFLAGED
	SKILL_CHAMELEON
	SKILL_CURSED_CLOAK
	SKILL_DIRE_LUCK
	SKILL_GRIM_LUCK
	SKILL_AWFUL_LUCK
	SKILL_NESTING
	SKILL_SLYLY
	SKILL_MAZE
	SKILL_DESPAIR
	SKILL_CORK
	SKILL_CRACKEDSHELL
	SKILL_INFECT
	SKILL_TAG
	SKILL_PITCHER
	SKILL_WIND_WHISPER
	SKILL_SPEED_TRIGGER
	SKILL_SPRINGY
	SKILL_FORAGER
	SKILL_COOLIE
	SKILL_REINFORCED
	SKILL_SYMBIOSIS
	SKILL_TRACKLESS
	SKILL_COMPETITIVE
	SKILL_STUNNING_PEARL
	SKILL_EXOSKELETON
	SKILL_CUDDLES
	SKILL_TANK
	SKILL_ENTICING_EYES
	SKILL_CRACKED_SHELL
	SKILL_SPOOKY_SOUL
	SKILL_SOLAR_POWER
	SKILL_LUCKY_CLOVER
	SKILL_POWERFUL
	SKILL_STICK_GUARD
	SKILL_FIRST_STRIKE

	SKILL_COUNT
)

const (
	SPECIES_NONE = SpeciesID(iota)
	SPECIES_OWLARK
	SPECIES_CONILASH
	SPECIES_PELISHIELD
	SPECIES_PUMKEY
	SPECIES_GHOSOTH
	SPECIES_HEFFERRA
	SPECIES_TELEGER
	SPECIES_RUBBARK
	SPECIES_DRALMON
	SPECIES_BIRTROCUTE
	SPECIES_DAIBLE
	SPECIES_HOPTON
	SPECIES_ORIKOI
	SPECIES_TRAWKADEE
	SPECIES_BEARBLE
	SPECIES_GALLEON
	SPECIES_ESKOGUIN
	SPECIES_VOLTSOG
	SPECIES_DREPIR
	SPECIES_DETTOCRAB
	SPECIES_GRAPPUTAN
	SPECIES_SERPITCHER
	SPECIES_HYPNOLLY
	SPECIES_SIMIALAY
	SPECIES_INDOGON
	SPECIES_ALANTOM
	SPECIES_ELECLASH
	SPECIES_WILLOSPER
	SPECIES_HIPPOCRUNCH
	SPECIES_ARMARAI
	SPECIES_SPROGGY
	SPECIES_TERRARAM
	SPECIES_GORAITH
	SPECIES_DAGGBAY
	SPECIES_VOLCENIX
	SPECIES_EKRANOPUFF
	SPECIES_TWIGILE
	SPECIES_KOABLAZE
	SPECIES_PUFFLY
	SPECIES_FERRESOL
	SPECIES_VERTEBLADE
	SPECIES_SILKAMB
	SPECIES_GLIGGY
	SPECIES_ICERATE
	SPECIES_SPAROCEROUS
	SPECIES_ELECTRAH
	SPECIES_MOUSECOY
	SPECIES_HOWLOON
	SPECIES_AUSTEROGON
	SPECIES_CAVYCLAY
	SPECIES_FOXFERNO
	SPECIES_TYTANICONE
	SPECIES_OTTECRUSH
	SPECIES_KATAHUM
	SPECIES_PANDAKHAN
	SPECIES_MUSHIFU
	SPECIES_JOLTALOTL
	SPECIES_GIRAFFRIGHT
	SPECIES_BLIZZBEAR
	SPECIES_CAMBRIOPOD
	SPECIES_KNIGHTOISE
	SPECIES_STATCHICK
	SPECIES_LASELLADY
	SPECIES_SORCEROTH
	SPECIES_TURFOUL
	SPECIES_SIANEMA
	SPECIES_OCTOGOO
	SPECIES_SIMORCHIS
	SPECIES_DANDEMOTH
	SPECIES_FERROCROC
	SPECIES_BISTONE
	SPECIES_MORSHELL
	SPECIES_STARSNARL
	SPECIES_LEOSQUALL
	SPECIES_DOVECOO
	SPECIES_EMPERORCA
	SPECIES_BEAVADDLE
	SPECIES_CATIMIUM
	SPECIES_OAKANNON
	SPECIES_WALLOPRUS
	SPECIES_LEVETEASL
	SPECIES_MOSSADOW
	SPECIES_RAYLING
	SPECIES_CORALGO
	SPECIES_ANEMOSTING
	SPECIES_LUMBEAR
	SPECIES_SMOULDTUS
	SPECIES_BOMBELLY
	SPECIES_REAPRAVE
	SPECIES_PELEPOM
	SPECIES_KELPAGI
	SPECIES_TELASTIER
	SPECIES_WOLFROST
	SPECIES_PHIRUSE
	SPECIES_CRABEOR
	SPECIES_DRILLNOME
	SPECIES_LEMSNOW
	SPECIES_WRECKO
	SPECIES_COSMOFLY
	SPECIES_CHESGNAW
	SPECIES_SABESEAL
	SPECIES_MAJEAGLE
	SPECIES_INVISLEON
	SPECIES_SPIRIBLOSSOM
	SPECIES_TOXICHAR
	SPECIES_SAILPHIN
	SPECIES_GRANOGRE
	SPECIES_CLAMDIBLE
	SPECIES_FIGGULP
	SPECIES_MANDOPPOSE
	SPECIES_PANDFLAME
	SPECIES_BUCKEETLE
	SPECIES_ORATURTLE
	SPECIES_TUMBLEGEIST
	SPECIES_SUMOCROAK
	SPECIES_THUNDLEO
	SPECIES_HIPPUDGE
	SPECIES_FRUBBLE
	SPECIES_HAYBRAY
	SPECIES_PLATYCURSE
	SPECIES_CARDICOAL
	SPECIES_THERAQUARTZ
	SPECIES_BAMMUAY
	SPECIES_GEARCOON
	SPECIES_BASHURRE
	SPECIES_FLANIMBUS
	SPECIES_TORPTUNA
	SPECIES_WALLASPRING
	SPECIES_BOKUREX
	SPECIES_MESQITO
	SPECIES_KOYOSKA
	SPECIES_SLYOTE
	SPECIES_IMPACTUS
	SPECIES_PEBBSLING
	SPECIES_MAPALLIC
	SPECIES_POSSESSUM
	SPECIES_DROMEDUNE
	SPECIES_CUDDLESQUID
	SPECIES_HEDGEAZE
	SPECIES_BALLEURA
	SPECIES_WYVEROK
	SPECIES_TULOPS
	SPECIES_BINDBAT
	SPECIES_CONFUSEBRA
	SPECIES_RUMBLEBEE
	SPECIES_PIGGYTANK
	SPECIES_LITHICHIEF
	SPECIES_DROWZSHEEP
	SPECIES_MAGESTER
	SPECIES_TRICERASHANK
	SPECIES_CORKCUB
	SPECIES_PSEUTLE
	SPECIES_SPECTOWL
	SPECIES_SPINOE
	SPECIES_CORACAVE
	SPECIES_CLOVESLASH
	SPECIES_TUNDROTH
	SPECIES_CATERSTICK
	SPECIES_SINISTING
	SPECIES_SKOOKY

	SPECIES_COUNT
)

const (
	STYLE_NONE = StyleID(iota)
	STYLE_INNER
	STYLE_OUTER
)

const (
	TYPE_NONE = TypeID(iota)
	TYPE_STRENGTH
	TYPE_GROUND
	TYPE_WATER
	TYPE_ICE
	TYPE_CHEMICAL
	TYPE_METAL
	TYPE_STONE
	TYPE_SOLAR
	TYPE_PSYCHE
	TYPE_WIND
	TYPE_ELECTRIC
	TYPE_SPIRIT
	TYPE_FIRE
	TYPE_ILLUSION
)
