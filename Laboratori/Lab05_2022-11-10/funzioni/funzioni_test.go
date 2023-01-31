/*
In questo esercizio prendiamo in considerazione solo caratteri ASCII (byte).
In un file funzioni.go definire le seguenti funzioni:

- hasUpper(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa contiene almeno una lettera latina maiuscola (tra 'A' e 'Z'), false altrimenti.

- firstUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione della prima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.

- lastUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione dell'ultima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.

- countDigitsLettersOthers(s string) int int int
La funzione riceve una stringa come parametro, conta quante cifre, quante lettere e quanti altri caratteri (né cifre né lettere) contiene e restituisce i tre risultati in questo ordine.

- isPalindrome(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa è palindroma, e false altrimenti. Una parola è palindroma se può essere letta  sia da sinistra a destra che da destra a sinistra. Ad esempio "osso" e "ingegni" sono palindrome, ma anche "" (la stringa vuota) e "t" (le stringhe di un solo carattere).

Scrivete infine un main che legge una stringa da standard input, usa le funzioni qui sopra per determinare se la stringa letta contiene maiuscole, in che posizione è la prima maiuscola, in che posizione è l'ultima maiuscola, quante cifre, lettere e altri caratteri contiene, se è palindroma, e stampa i risultati ("ha maiuscole" / "non ha maiuscole", "prima maiuscola in posizione ...", “palindroma" / "non palindroma", ecc.).

nomefile: funzioni.go

*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

// palindromi
func TestPalindromeSingleChar(t *testing.T) {
	if !isPalindrome("A") {
		t.Fail()
	}
}

func TestPalindromeSameChar(t *testing.T) {
	if !isPalindrome("bbbbbbbbbbbbbbbbbb") {
		t.Fail()
	}
}

func TestNonPalindromo(t *testing.T) {
	if isPalindrome("ciaoajlajdlas") {
		t.Fail()
	}
}

// hasupper
func TestHasUpperSingle(t *testing.T) {
	if !hasUpper("A") {
		t.Fail()
	}
	if hasUpper("b") {
		t.Fail()
	}
}

func TestHasUpperNormal(t *testing.T) {
	if !hasUpper("askjhdaksjsIlaskdjaslkdj") {
		t.Fail()
	}
}

// count
func TestCountNormal(t *testing.T) {
	d, l, o := countDigitsLettersOthers("sakdja87213612kjhkasjhdkas8q912as k askdhaksdhkasdkjas")
	//fmt.Println(d, " ", l, " ", o)
	if d != 12 && l != 40 && o != 2 {
		t.Fail()
	}
}

func TestCountEmpty(t *testing.T) {
	d, l, o := countDigitsLettersOthers("")
	if d != 0 && l != 0 && o != 0 {
		t.Fail()
	}
}

func TestCountOneOfEach(t *testing.T) {
	d, l, o := countDigitsLettersOthers("1a,")
	//fmt.Println(d, " ", l, " ", o)
	if d != 1 && l != 1 && o != 1 {
		t.Fail()
	}
}

// firstupper
func TestFirstUpper(t *testing.T) {
	if firstUpper("udtqABsso") != 4 {
		t.Fail()
	}
}
func TestFirstUpperNo(t *testing.T) {
	if firstUpper("udtqddsso") != -1 {
		t.Fail()
	}
}

// lastupper
func TestLastUpper(t *testing.T) {
	if lastUpper("udtqAvBsso") != 6 {
		t.Fail()
	}
}

/* non funziona causa punteggiatura e maiuscole
func TestPalindromoRecord(t *testing.T) {
	if !isPalindrome("Ai lati, a esordir, dama e re, Pertini trepida, tira lieti moccoli, dialoga – vocina, pipa… -, ricorre alle battute. È durata!… ne patì Trap: allena – mèritasi lodi testé – Juvitalia, mai amata. Il boato n’eruppe su filato, mero atto d’ira: assorga da gai palati, ingoi l’arena! Si rise, noi: gara azzurra – felicità, reti – e ricca! Né tacerò pose, ire, rapidi miti; citerò paure… però meritan oro. Ci sono rari tiri? Sia! ma i latini eroi goderono di rigore – c’è fallo -; “Fatale far tale rete”: lassa prosopopea nei peani dona aìre facile. Ma “fatale” malessere globi dilata, rene, vene ci necrotizza: ratto, vago, da finir al còre (l’oblierà? Dall’idea – l’Erinni! – trepiderà: tic e tac…)… Lapsus saliente (idra! sillabo!): non amai Cabrini; flusso acre – pus era? sudore? -, bile d’ittero ci assalì: risa brutali, amaro icore… Fiore italo, cari miei, secca, alidirà vizzito là, se sol – a foci nuove diretti, fisi – a metà recedete: l’itala idea di vis (i redivivi, noti, ilari miti!) trapasserà, inerte e vana, in italianità lisa, banal. Attutite relativa ira, correte: eterni onori n’avrete! Sibili – tre “fi” – di arbitro: finita lì metà partita; reca loro l’animo di lotta, fidata ripresa! mira, birra rida’! attuta ire, bile! La si disse “eterea”, la Catalogna: alla pari terrò cotali favolose ore… Notte molle, da re! Poeti m’illusero “Va’!”, “Fa’!”, “Osa!”) colla fusione – esile, serica, viva -, rime lepide, tra anelito d’età d’oro e rudezze d’orpello; così cederò all’eros, ai sensi rei; amai -l’amavo… – una grata città, la gag, la vita; nutro famosa cara sete, relativa a Lalo, Varese, De Falla, Petrassi, e Ravel, e Adam, e Nono… Sor… bene, totale opaca arte; né pago fui per attori, dive, divi (lo sarò?)… Là ogni avuto, mai sopito piacere s’evaporò, leggera falena era: se con amor, lì, alla cara – cotale! – virile sera – coi gaudi sereni, grevi da dare angine, beati – lo paragono, decàde a ludo, mollica, vile cineseria, onere. Sì! Taccola barocca allora rimane, meno mi tange: solo apatia apporterà, goffa noia… Paride, Ettore e soci trovarono sì dure sorti – riverberare di pira desueta! – coi gelosi re dei Dori (trono era d’ira, Era, Muse); a Ilio nati e no, di elato tono, di rango, là tacitati – re… mogi -, videro Elleni libare, simil a Titani, su al Pergamo: idem i Renani e noi… “… caparbi”, vaticinò – tono trepido -, ed ora tange là tale causale trofeo (coppa di rito è la meta della partita), trainer fisso; mìralo come l’anemone: fisso, raro, da elogi… D’animo nobile, divo mai, mai tetro, fatale varò la tattica. Cito Gay, ognor abile devo dir: da Maracanà sono tacco, battuta… Ai lati issò vela l’ala latina Bruno: cerca la rete, si batte assai, opera lì, fora, rimargina… Bergomi, nauta ragazzo, riserra giù sì care fila: è l’età… Coi gradi vedo – troppa la soavità… – capitano Dino, razza ladina. Rete vigila! dilàtati…!: la turba, l’arena, ti venera. Ad ogni rado, torpido e no, tirabile tiro, trapelò rapidità sua: parò (la tivù, lì, diè nitidi casi). Di tutto – fiero, mai di fatica, vivace – raccatta: e, se tarpate, le ali loro – è la verità – paion logore. Zoff (ùtinam !) è dei.. Parà: para… Piede, mani, tuffo: zero gol, noi a patire. Vale oro: lì, là… è l’età… “Pratese, attacca! reca vivacità!”, “Fidiamo!”, “Rei fottuti disaciditi!”… Nei diluvi, talora pausati, di parole partorite lì, baritone o di proto, da ring o da arene (“Vita nera là, brutalità tali da ligi veterani, da… lazzaroni!”, “Dònati! pàcati! va’! osa!: l’apporto devi dar!”, “Giocate leali, feraci!”, “Su i garresi!”, “Rozza gara!”, “Tu, animo!”, “Grèbani! Grami!”, “Raro filare!”; poi: “Assaetta!”, “Bis!” e “Ter!”), alacre, con urbanità, l’alalà levossi: “Italia!”, a tutta bocca, tonò. Sana cara Madrid, ove delibaron Goya… gotica città talora velata: forte ti amiamo! Vi delibo nomina di goleador a Rossi – fenomenale! -: mo’, colà, rimossi freni artati (tra palle date male o tiri dappoco è forte la sua celata legnata), rode, o d’ipertono, tonicità, vibra. Pacione inane, rimediò magre, plausi – nati tali – miserabili nelle ore di Vigo (meritàti!); Catalogna ridonò totale idoneità – noi lì a esumare, a ridare onor -, tiro diede, riso; le giocate use – da ripide, rare, brevi, ritrose, rudi – son ora vorticose e rotte, e d’ira paion affogare (troppa?). Aìta, Paolo!: segna, timone mena, mira, rolla, accora, balòccati sereno, aìre – se Nice lì vacillò – modula e da’ (cedono…): gara polita e benigna – e rada, di vergine residua… – gioca. Re s’è lì rivelato (Caracalla? Il romano Cesare!): anela, fa, regge loro, pavese reca…: ipotiposi amo. Tu va’ in goal, ora! Sol, ivi, devi dirottare più foga: penetra a capo elato – tenebroso non è… -, ma da elevare, issar te, palla, fede, sera (vola, là) a vitale rete! Sarà caso… Ma Fortuna ti valga galattica targa, nuova malìa: mai Eris ne sia sorella! Or è deciso; colle prodezze, dure o rodate doti – lena, arte di Pelé, mira -, vivaci rese lì sé e noi: su fallo (caso a favore sul limite, opera dell’ometto nero) è solo, va filato, corre, tira, palla angolata cala… è rete! Essi di sale, l’Iberia tutta a dir “Arriba!”, rimaser. Pirata? Di fatto li domina… Loro lacerati tra patemi; Latini forti, braidi, fertili: bis e ter van, ìrono in rete… E terrò cari a vita: le reti; tutta l’anabasi latina; i Latini, a nave e treni, a ressa partiti (mìrali!); i toni vivi, derisivi, d’aedi alati; le tede cerate (“Mai sì fitte” ridevo: unico falò s’esalò, tizzi vari di là accesi); e i miracolati eroi, feroci… Oramai la turba si rilassa: i coretti deliberò d’usare. Supercaos sul finir! Baciamano? No: balli sardi, etnei lassù (spalcate!); citaredi per tinnire, là, ed il “la” dare; il Bolero, clarini, fado, gavotta, razzi, torce (Nice n’è venerata) lì. Di bolge, resse, la melata famelica “feria” anodina è piena, e po’ po’ sorpassa l’etere la trafelata folla. Fecero giri d’onore: dogi o re, in Italia, mai si ritirarono sì coronati. Remore, Perù, aporetici timidi pareri… e sopore, catenacci reiterati, Cile, far ruzza: a ragione si risanerà lì ogni itala piaga; da grossa a ridotta, o remota, lì fu, seppure nota, obliata. Mai amai la tivù: jet-set, idoli, satire…; ma nella partita – penata, rude e tutta bella: erro? – ci rapì: panico vago, lai di locco, mite ilarità di Pertini… tre pere a Madrid, rosea Italia") {
		t.Fail()
	}
}
*/
