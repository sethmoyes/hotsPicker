package heroes

type role string

const (
	rangedAssassin role = "Ranged Assassin"
	meleeAssassin       = "Melee Assassin"
	bruiser             = "Bruiser"
	tank                = "Tank"
	support             = "Support"
	healer              = "Healer"
)

type Comp int

const (
	teamFight Comp = iota
	pick
	splitPush
	dive
	safe
	poke
)

// Hero represents each hero in HotS
type Hero struct {
	Name string
	// Difficulty string
	// Universe   string
	Role role
	Comp []Comp
}

var abathur = Hero{Name: "Abathur", Role: support, Comp: []Comp{splitPush, teamFight}}
var alarak = Hero{Name: "Alarak", Role: meleeAssassin, Comp: []Comp{teamFight}}
var alexstrasza = Hero{Name: "Alexstraza", Role: healer, Comp: []Comp{teamFight, poke}}
var ana = Hero{Name: "Ana", Role: healer, Comp: []Comp{safe, teamFight}}
var anduin = Hero{Name: "Anduin", Role: healer, Comp: []Comp{teamFight, safe}}
var anubarak = Hero{Name: "Anub'arak", Role: tank, Comp: []Comp{dive}}
var artanis = Hero{Name: "Artanis", Role: bruiser, Comp: []Comp{pick, teamFight}}
var arthas = Hero{Name: "Arthas", Role: tank, Comp: []Comp{teamFight, safe}}
var auriel = Hero{Name: "Auriel", Role: healer, Comp: []Comp{safe}}
var azmodan = Hero{Name: "Azmodan", Role: rangedAssassin, Comp: []Comp{poke, splitPush}}
var blaze = Hero{Name: "Blaze", Role: tank, Comp: []Comp{teamFight, dive, safe}}
var brightwing = Hero{Name: "Brightwing", Role: healer, Comp: []Comp{safe, splitPush}}
var cassia = Hero{Name: "Cassia", Role: meleeAssassin, Comp: []Comp{dive, teamFight}}
var chen = Hero{Name: "Chen", Role: tank, Comp: []Comp{teamFight, pick, dive}}
var cho = Hero{Name: "Cho", Role: tank, Comp: []Comp{teamFight, poke}}
var chromie = Hero{Name: "Chromie", Role: rangedAssassin, Comp: []Comp{poke}}
var dva = Hero{Name: "D.Va", Role: tank, Comp: []Comp{dive, teamFight}}
var deathwing = Hero{Name: "Deathwing", Role: tank, Comp: []Comp{teamFight, splitPush}}
var deckard = Hero{Name: "Deckard", Role: healer, Comp: []Comp{teamFight, safe, pick}}
var dehaka = Hero{Name: "Dehaka", Role: tank, Comp: []Comp{splitPush, pick}}
var diablo = Hero{Name: "Dibalo", Role: tank, Comp: []Comp{teamFight, pick, safe}}
var etc = Hero{Name: "E.T.C", Role: tank, Comp: []Comp{teamFight, safe}}
var falstad = Hero{Name: "Falstad", Role: rangedAssassin, Comp: []Comp{splitPush, teamFight}}
var fenix = Hero{Name: "Fenix", Role: rangedAssassin, Comp: []Comp{teamFight, dive}}
var gall = Hero{Name: "Gall", Role: rangedAssassin, Comp: []Comp{teamFight, poke}}
var garrosh = Hero{Name: "Garrosh", Role: tank, Comp: []Comp{pick, teamFight, safe}}
var gazlowe = Hero{Name: "Gazlowe", Role: meleeAssassin, Comp: []Comp{splitPush, poke}}
var genji = Hero{Name: "Genji", Role: meleeAssassin, Comp: []Comp{dive, teamFight}}
var greymane = Hero{Name: "Graymane", Role: rangedAssassin, Comp: []Comp{dive, safe, splitPush}}
var guldan = Hero{Name: "Gul'dan", Role: rangedAssassin, Comp: []Comp{teamFight, safe, poke}}
var hanzo = Hero{Name: "Hanzo", Role: rangedAssassin, Comp: []Comp{poke, teamFight}}
var illidan = Hero{Name: "Illidan", Role: meleeAssassin, Comp: []Comp{dive, splitPush}}
var imperius = Hero{Name: "Imperius", Role: bruiser, Comp: []Comp{teamFight, pick}}
var jaina = Hero{Name: "Jaina", Role: rangedAssassin, Comp: []Comp{teamFight, poke}}
var johanna = Hero{Name: "Johanna", Role: tank, Comp: []Comp{teamFight, safe}}
var junkrat = Hero{Name: "Junkrat", Role: rangedAssassin, Comp: []Comp{pick, splitPush, poke}}
var kaelthas = Hero{Name: "Kael'thas", Role: rangedAssassin, Comp: []Comp{teamFight, poke, pick}}
var kelthuzad = Hero{Name: "Kel'Thuzad", Role: rangedAssassin, Comp: []Comp{poke, teamFight}}
var kerrigan = Hero{Name: "Kerrigan", Role: meleeAssassin, Comp: []Comp{dive, teamFight}}
var kharazim = Hero{Name: "Kharazim", Role: healer, Comp: []Comp{dive}}
var leoric = Hero{Name: "Leoric", Role: bruiser, Comp: []Comp{teamFight, splitPush}}
var lili = Hero{Name: "Li li", Role: healer, Comp: []Comp{teamFight, safe}}
var liming = Hero{Name: "Li-Ming", Role: rangedAssassin, Comp: []Comp{poke, teamFight}}
var ltmorales = Hero{Name: "Lt. Morales", Role: healer, Comp: []Comp{safe}}
var lucio = Hero{Name: "Lucio", Role: healer, Comp: []Comp{teamFight, safe, dive}}
var lunara = Hero{Name: "Lunara", Role: rangedAssassin, Comp: []Comp{teamFight, poke}}
var maiev = Hero{Name: "Maiev", Role: meleeAssassin, Comp: []Comp{teamFight, dive}}
var malganis = Hero{Name: "Mal'Ganis", Role: tank, Comp: []Comp{dive, teamFight, safe}}
var malfurion = Hero{Name: "Malfurion", Role: healer, Comp: []Comp{teamFight, pick, poke}}
var malthael = Hero{Name: "Malthael", Role: meleeAssassin, Comp: []Comp{dive, splitPush}}
var medivh = Hero{Name: "Medivh", Role: support, Comp: []Comp{poke, safe, teamFight}}
var mephisto = Hero{Name: "Mephisto", Role: rangedAssassin, Comp: []Comp{poke, teamFight}}
var muradin = Hero{Name: "Muradin", Role: tank, Comp: []Comp{teamFight, safe, dive, pick}}
var murky = Hero{Name: "Murky", Role: meleeAssassin, Comp: []Comp{splitPush, pick}}
var nazeebo = Hero{Name: "Nazeebo", Role: rangedAssassin, Comp: []Comp{splitPush, teamFight, poke}}
var nova = Hero{Name: "Nova", Role: rangedAssassin, Comp: []Comp{pick}}
var orphea = Hero{Name: "Orphea", Role: rangedAssassin, Comp: []Comp{teamFight, safe, poke}}
var probius = Hero{Name: "Probius", Role: rangedAssassin, Comp: []Comp{splitPush, poke}}
var qhira = Hero{Name: "Qhira", Role: meleeAssassin, Comp: []Comp{dive}}
var ragnaros = Hero{Name: "Ragnaros", Role: meleeAssassin, Comp: []Comp{splitPush, teamFight}}
var raynor = Hero{Name: "Raynor", Role: rangedAssassin, Comp: []Comp{safe, teamFight}}
var rehgar = Hero{Name: "Rehgar", Role: healer, Comp: []Comp{teamFight, dive, safe}}
var rexxar = Hero{Name: "Rexxar", Role: tank, Comp: []Comp{splitPush}}
var samuro = Hero{Name: "Samuro", Role: meleeAssassin, Comp: []Comp{splitPush, dive}}
var sgthammer = Hero{Name: "Sgt. Hammer", Role: rangedAssassin, Comp: []Comp{splitPush, teamFight, safe, poke}}
var sonya = Hero{Name: "Sonya", Role: bruiser, Comp: []Comp{teamFight, splitPush}}
var stitches = Hero{Name: "Stitches", Role: tank, Comp: []Comp{pick, teamFight}}
var stukov = Hero{Name: "Stukov", Role: healer, Comp: []Comp{teamFight, safe, dive}}
var sylvanas = Hero{Name: "Sylvanas", Role: rangedAssassin, Comp: []Comp{splitPush, teamFight}}
var tassadar = Hero{Name: "Tassadar", Role: rangedAssassin, Comp: []Comp{poke, teamFight}}
var thebutcher = Hero{Name: "The Butcher", Role: meleeAssassin, Comp: []Comp{pick, dive}}
var thelostvikings = Hero{Name: "The Lost Vikings", Role: support, Comp: []Comp{splitPush}}
var thrall = Hero{Name: "Thrall", Role: meleeAssassin, Comp: []Comp{dive, teamFight}}
var tracer = Hero{Name: "Tracer", Role: rangedAssassin, Comp: []Comp{dive, teamFight}}
var tychus = Hero{Name: "Tychus", Role: rangedAssassin, Comp: []Comp{teamFight, safe}}
var tyrael = Hero{Name: "Tyrael", Role: tank, Comp: []Comp{teamFight, dive, safe}}
var tyrande = Hero{Name: "Tyrande", Role: healer, Comp: []Comp{teamFight, poke}}
var uther = Hero{Name: "Uther", Role: healer, Comp: []Comp{safe}}
var valeera = Hero{Name: "Valeera", Role: meleeAssassin, Comp: []Comp{dive, pick}}
var valla = Hero{Name: "Valla", Role: rangedAssassin, Comp: []Comp{teamFight, pick}}
var varian = Hero{Name: "Varian", Role: tank, Comp: []Comp{dive, teamFight}}
var whitemane = Hero{Name: "Whitemane", Role: healer, Comp: []Comp{teamFight, safe}}
var xul = Hero{Name: "Xul", Role: meleeAssassin, Comp: []Comp{splitPush, teamFight}}
var yrel = Hero{Name: "Yrel", Role: bruiser, Comp: []Comp{teamFight, dive, safe}}
var zagara = Hero{Name: "Zagara", Role: rangedAssassin, Comp: []Comp{poke, splitPush}}
var zarya = Hero{Name: "Zarya", Role: support, Comp: []Comp{teamFight, safe}}
var zeratul = Hero{Name: "Zeratul", Role: meleeAssassin, Comp: []Comp{dive, pick}}
var zuljin = Hero{Name: "Zul'jin", Role: rangedAssassin, Comp: []Comp{safe, teamFight}}

// AllHeroes is a var holding all heroes in the game
var AllHeroes = []Hero{
	abathur, alarak, alexstrasza, ana, anduin, anubarak, artanis, arthas, auriel, azmodan, blaze, brightwing, cassia, chen, cho, chromie, dva, deathwing, deckard, dehaka, diablo, etc, falstad, fenix, gall, garrosh, gazlowe, genji, greymane, guldan, hanzo, illidan, imperius, jaina, johanna, junkrat, kaelthas, kelthuzad, kerrigan, kharazim, leoric, lili, liming, ltmorales, lucio, lunara, maiev, malganis, malfurion, malthael, medivh, mephisto, muradin, murky, nazeebo, nova, orphea, probius, qhira, ragnaros, raynor, rehgar, rexxar, samuro, sgthammer, sonya, stitches, stukov, sylvanas, tassadar, thebutcher, thelostvikings, thrall, tracer, tychus, tyrael, tyrande, uther, valeera, valla, varian, whitemane, xul, yrel, zagara, zagara, zeratul, zuljin,
}
