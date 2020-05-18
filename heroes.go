package heroes

type Role string

const (
	RangedAssassin Role = "Ranged Assassin"
	MeleeAssassin       = "Melee Assassin"
	Bruiser             = "Bruiser"
	Tank                = "Tank"
	Support             = "Support"
	Healer              = "Healer"
)

type Comp int

const (
	TeamFight Comp = iota
	Pick
	SplitPush
	Dive
	Safe
	Poke
)

type CompCountName []compCountName

type compCountName struct {
	Name  Comp
	Count int
}

// CompCount holds the number of comp types across the heroes on the team
type CompCount struct {
	TeamFightCount int
	PickCount      int
	SplitPushCount int
	DiveCount      int
	SafeCount      int
	PokeCount      int
}

// Hero represents each hero in HotS
type Hero struct {
	Name string
	// Difficulty string
	// Universe   string
	Role Role
	Comp []Comp
}

var abathur = Hero{Name: "Abathur", Role: Support, Comp: []Comp{SplitPush, TeamFight}}
var alarak = Hero{Name: "Alarak", Role: MeleeAssassin, Comp: []Comp{TeamFight}}
var alexstrasza = Hero{Name: "Alexstraza", Role: Healer, Comp: []Comp{TeamFight, Poke}}
var ana = Hero{Name: "Ana", Role: Healer, Comp: []Comp{Safe, TeamFight}}
var anduin = Hero{Name: "Anduin", Role: Healer, Comp: []Comp{TeamFight, Safe}}
var anubarak = Hero{Name: "Anub'arak", Role: Tank, Comp: []Comp{Dive}}
var artanis = Hero{Name: "Artanis", Role: Bruiser, Comp: []Comp{Pick, TeamFight}}
var arthas = Hero{Name: "Arthas", Role: Tank, Comp: []Comp{TeamFight, Safe}}
var auriel = Hero{Name: "Auriel", Role: Healer, Comp: []Comp{Safe}}
var azmodan = Hero{Name: "Azmodan", Role: RangedAssassin, Comp: []Comp{Poke, SplitPush}}
var blaze = Hero{Name: "Blaze", Role: Tank, Comp: []Comp{TeamFight, Dive, Safe}}
var brightwing = Hero{Name: "Brightwing", Role: Healer, Comp: []Comp{Safe, SplitPush}}
var cassia = Hero{Name: "Cassia", Role: MeleeAssassin, Comp: []Comp{Dive, TeamFight}}
var chen = Hero{Name: "Chen", Role: Tank, Comp: []Comp{TeamFight, Pick, Dive}}
var cho = Hero{Name: "Cho", Role: Tank, Comp: []Comp{TeamFight, Poke}}
var chromie = Hero{Name: "Chromie", Role: RangedAssassin, Comp: []Comp{Poke}}
var dva = Hero{Name: "D.Va", Role: Tank, Comp: []Comp{Dive, TeamFight}}
var deathwing = Hero{Name: "Deathwing", Role: Tank, Comp: []Comp{TeamFight, SplitPush}}
var deckard = Hero{Name: "Deckard", Role: Healer, Comp: []Comp{TeamFight, Safe, Pick}}
var dehaka = Hero{Name: "Dehaka", Role: Tank, Comp: []Comp{SplitPush, Pick}}
var diablo = Hero{Name: "Diablo", Role: Tank, Comp: []Comp{TeamFight, Pick, Safe}}
var etc = Hero{Name: "E.T.C", Role: Tank, Comp: []Comp{TeamFight, Safe}}
var falstad = Hero{Name: "Falstad", Role: RangedAssassin, Comp: []Comp{SplitPush, TeamFight}}
var fenix = Hero{Name: "Fenix", Role: RangedAssassin, Comp: []Comp{TeamFight, Dive}}
var gall = Hero{Name: "Gall", Role: RangedAssassin, Comp: []Comp{TeamFight, Poke}}
var garrosh = Hero{Name: "Garrosh", Role: Tank, Comp: []Comp{Pick, TeamFight, Safe}}
var gazlowe = Hero{Name: "Gazlowe", Role: MeleeAssassin, Comp: []Comp{SplitPush, Poke}}
var genji = Hero{Name: "Genji", Role: RangedAssassin, Comp: []Comp{Dive, TeamFight}}
var greymane = Hero{Name: "Graymane", Role: RangedAssassin, Comp: []Comp{Dive, Safe, SplitPush}}
var guldan = Hero{Name: "Gul'dan", Role: RangedAssassin, Comp: []Comp{TeamFight, Safe, Poke}}
var hanzo = Hero{Name: "Hanzo", Role: RangedAssassin, Comp: []Comp{Poke, TeamFight}}
var illidan = Hero{Name: "Illidan", Role: MeleeAssassin, Comp: []Comp{Dive, SplitPush}}
var imperius = Hero{Name: "Imperius", Role: Bruiser, Comp: []Comp{TeamFight, Pick}}
var jaina = Hero{Name: "Jaina", Role: RangedAssassin, Comp: []Comp{TeamFight, Poke}}
var johanna = Hero{Name: "Johanna", Role: Tank, Comp: []Comp{TeamFight, Safe}}
var junkrat = Hero{Name: "Junkrat", Role: RangedAssassin, Comp: []Comp{Pick, SplitPush, Poke}}
var kaelthas = Hero{Name: "Kael'thas", Role: RangedAssassin, Comp: []Comp{TeamFight, Poke, Pick}}
var kelthuzad = Hero{Name: "Kel'Thuzad", Role: RangedAssassin, Comp: []Comp{Poke, TeamFight}}
var kerrigan = Hero{Name: "Kerrigan", Role: MeleeAssassin, Comp: []Comp{Dive, TeamFight}}
var kharazim = Hero{Name: "Kharazim", Role: Healer, Comp: []Comp{Dive}}
var leoric = Hero{Name: "Leoric", Role: Bruiser, Comp: []Comp{TeamFight, SplitPush}}
var lili = Hero{Name: "Li li", Role: Healer, Comp: []Comp{TeamFight, Safe}}
var liming = Hero{Name: "Li-Ming", Role: RangedAssassin, Comp: []Comp{Poke, TeamFight}}
var ltmorales = Hero{Name: "Lt. Morales", Role: Healer, Comp: []Comp{Safe}}
var lucio = Hero{Name: "Lucio", Role: Healer, Comp: []Comp{TeamFight, Safe, Dive}}
var lunara = Hero{Name: "Lunara", Role: RangedAssassin, Comp: []Comp{TeamFight, Poke}}
var maiev = Hero{Name: "Maiev", Role: MeleeAssassin, Comp: []Comp{TeamFight, Dive}}
var malganis = Hero{Name: "Mal'Ganis", Role: Tank, Comp: []Comp{Dive, TeamFight, Safe}}
var malfurion = Hero{Name: "Malfurion", Role: Healer, Comp: []Comp{TeamFight, Pick, Poke}}
var malthael = Hero{Name: "Malthael", Role: MeleeAssassin, Comp: []Comp{Dive, SplitPush}}
var medivh = Hero{Name: "Medivh", Role: Support, Comp: []Comp{Poke, Safe, TeamFight}}
var mephisto = Hero{Name: "Mephisto", Role: RangedAssassin, Comp: []Comp{Poke, TeamFight}}
var muradin = Hero{Name: "Muradin", Role: Tank, Comp: []Comp{TeamFight, Safe, Dive, Pick}}
var murky = Hero{Name: "Murky", Role: MeleeAssassin, Comp: []Comp{SplitPush, Pick}}
var nazeebo = Hero{Name: "Nazeebo", Role: RangedAssassin, Comp: []Comp{SplitPush, TeamFight, Poke}}
var nova = Hero{Name: "Nova", Role: RangedAssassin, Comp: []Comp{Pick}}
var orphea = Hero{Name: "Orphea", Role: RangedAssassin, Comp: []Comp{TeamFight, Safe, Poke}}
var probius = Hero{Name: "Probius", Role: RangedAssassin, Comp: []Comp{SplitPush, Poke}}
var qhira = Hero{Name: "Qhira", Role: MeleeAssassin, Comp: []Comp{Dive}}
var ragnaros = Hero{Name: "Ragnaros", Role: MeleeAssassin, Comp: []Comp{SplitPush, TeamFight}}
var raynor = Hero{Name: "Raynor", Role: RangedAssassin, Comp: []Comp{Safe, TeamFight}}
var rehgar = Hero{Name: "Rehgar", Role: Healer, Comp: []Comp{TeamFight, Dive, Safe}}
var rexxar = Hero{Name: "Rexxar", Role: Tank, Comp: []Comp{SplitPush}}
var samuro = Hero{Name: "Samuro", Role: MeleeAssassin, Comp: []Comp{SplitPush, Dive}}
var sgthammer = Hero{Name: "Sgt. Hammer", Role: RangedAssassin, Comp: []Comp{SplitPush, TeamFight, Safe, Poke}}
var sonya = Hero{Name: "Sonya", Role: Bruiser, Comp: []Comp{TeamFight, SplitPush}}
var stitches = Hero{Name: "Stitches", Role: Tank, Comp: []Comp{Pick, TeamFight}}
var stukov = Hero{Name: "Stukov", Role: Healer, Comp: []Comp{TeamFight, Safe, Dive}}
var sylvanas = Hero{Name: "Sylvanas", Role: RangedAssassin, Comp: []Comp{SplitPush, TeamFight}}
var tassadar = Hero{Name: "Tassadar", Role: RangedAssassin, Comp: []Comp{Poke, TeamFight}}
var thebutcher = Hero{Name: "The Butcher", Role: MeleeAssassin, Comp: []Comp{Pick, Dive}}
var thelostvikings = Hero{Name: "The Lost Vikings", Role: Support, Comp: []Comp{SplitPush}}
var thrall = Hero{Name: "Thrall", Role: MeleeAssassin, Comp: []Comp{Dive, TeamFight}}
var tracer = Hero{Name: "Tracer", Role: RangedAssassin, Comp: []Comp{Dive, TeamFight}}
var tychus = Hero{Name: "Tychus", Role: RangedAssassin, Comp: []Comp{TeamFight, Safe}}
var tyrael = Hero{Name: "Tyrael", Role: Tank, Comp: []Comp{TeamFight, Dive, Safe}}
var tyrande = Hero{Name: "Tyrande", Role: Healer, Comp: []Comp{TeamFight, Poke}}
var uther = Hero{Name: "Uther", Role: Healer, Comp: []Comp{Safe}}
var valeera = Hero{Name: "Valeera", Role: MeleeAssassin, Comp: []Comp{Dive, Pick}}
var valla = Hero{Name: "Valla", Role: RangedAssassin, Comp: []Comp{TeamFight, Pick}}
var varian = Hero{Name: "Varian", Role: Tank, Comp: []Comp{Dive, TeamFight}}
var whitemane = Hero{Name: "Whitemane", Role: Healer, Comp: []Comp{TeamFight, Safe}}
var xul = Hero{Name: "Xul", Role: MeleeAssassin, Comp: []Comp{SplitPush, TeamFight}}
var yrel = Hero{Name: "Yrel", Role: Bruiser, Comp: []Comp{TeamFight, Dive, Safe}}
var zagara = Hero{Name: "Zagara", Role: RangedAssassin, Comp: []Comp{Poke, SplitPush}}
var zarya = Hero{Name: "Zarya", Role: Support, Comp: []Comp{TeamFight, Safe}}
var zeratul = Hero{Name: "Zeratul", Role: MeleeAssassin, Comp: []Comp{Dive, Pick}}
var zuljin = Hero{Name: "Zul'jin", Role: RangedAssassin, Comp: []Comp{Safe, TeamFight}}

// AllHeroes is a var holding all heroes in the game
var AllHeroes = []Hero{
	abathur, alarak, alexstrasza, ana, anduin, anubarak, artanis, arthas, auriel, azmodan, blaze, brightwing, cassia, chen, cho, chromie, dva, deathwing, deckard, dehaka, diablo, etc, falstad, fenix, gall, garrosh, gazlowe, genji, greymane, guldan, hanzo, illidan, imperius, jaina, johanna, junkrat, kaelthas, kelthuzad, kerrigan, kharazim, leoric, lili, liming, ltmorales, lucio, lunara, maiev, malganis, malfurion, malthael, medivh, mephisto, muradin, murky, nazeebo, nova, orphea, probius, qhira, ragnaros, raynor, rehgar, rexxar, samuro, sgthammer, sonya, stitches, stukov, sylvanas, tassadar, thebutcher, thelostvikings, thrall, tracer, tychus, tyrael, tyrande, uther, valeera, valla, varian, whitemane, xul, yrel, zagara, zagara, zeratul, zuljin,
}

func (candidate Comp) IsInList(list []Comp) bool {
	for _, comp := range list {
		if comp == candidate {
			return true
		}
	}
	return false
}
