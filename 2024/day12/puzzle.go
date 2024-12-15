package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`AAAA
BBCD
BBCC
EEEC`,
		`140`,
		`80`,
	}, {
		`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
		`1930`,
		`1206`,
	},
}

var puzzle = `GGGGOOOOOOOOXXXXXXXXXXXXXXXXXSSSSSSSSSSSWWWWWWWWWWYYYYYYYYYYYYYYVVVVVVVVVVVVVVAGEEEEEEEEGGGGGGGGGGGGGEEBEQEYYYYYYYYBSSSBBBBBBBBBBBBBBBBBBBBB
GOOOOOOOOOOOOXXXXXXXXXXXXXXXXSSSSSSSSSSWWSWWWWWWYWYYYYYYYYYYYYYYYPVVVVVVVVVVVVVVEEEEEEEEGGGGGGGGGGEGGEEEEEEYYYYYYYYBBBBBBBBBBBBBBBBBBBBBBBBB
GGGGOOOOOOOOOXXXXXXXXXXXXXXXSSSSSSSSSSSWSSWWWWWWYYYYYYYYYYYYYYYYYYVVVVVVVVVVVVVVEEEEEEEEGGGGGGGGGEEEEEEEEEEEYYUYYYYBBBBBBBBBBBBBDBBBOOBBBBBI
GGGGGOOOOOOOXXXXXXXXXXXXXXXSSSSSSSSSSSSSSSQQWWRWYYYYYYYYYYYYYYYYYYVVVVVVVVVVVVVVEEEEEEEEGGGGGGGGGNEEEEEEEEEEYUUUUUUBBBBBBBBBDDDDDBBBOOBBBBII
GGGOOOOOOOOOXXXXXXXXXXXXXSSSSSSSSSSSSSSSSSQQQRRRYYYYYYYYYYYYYYYYYYYVVVVVVVVVVVVVEEEEEEEEGGGNGGGGGNTTEEEEEEEEUUUUSSSBBBBBBDDDDDDDDDDOOOBBBBBI
GGGGOOOOOOOXXXXXXXXXXXXXXSSSSSSSSSSSSSSSSRMMQRRRYYYYYYYYYYYYYYYYYYYVVVVVVVVVVVVVEEEEEEEEGGNNGGGGGNTTEEEEEEEEEUUSSQQQOBDDDDDDDDDDDDDDOBBBBBBB
GGGGGOOOOOOXXXXXXXXXXXXXXSSSSSSSSSSSSSSSRRRRRRRRRJQQYQYYYYYYYYYYYVVVVVVVVVVVVVVVEEEEEEEEGGNNNNNGNNNEEJUUUUUUUUUUQQQQQDDDDDDDDDDDDDDBBBBBBBBB
GGGYGYOOOOOOXXXXXXXXXXXXXXXSYYSSSSSSSSSRRRRRRRRRJJJQQQCYYYYYIVYBBVVVNNNNNVVVVVVVEEEEEEEEGGNNNNNNNNNEJJUUUUUUUUUUQQQQQQQDDDDDDDDDDDDBBBBBBSSK
GGGYYYOOOOOYYYXXXGGXXXGGGGGYYYYYSSSSOSSRRRRRRRRRRQQQQQQYYYEEIIIBBVVVNNNNNVVVVVEEEEEEEEEEGGGNNNNNNNNNNJGGUUUUUUUUNQQQQQYDDDDDDDDDDDDSSKSSSSSS
GGGGGYYOYOOYYXXXGGGGGGZGGGYYYYSSSSSSSSSSRRRRRRRRAVVQQQYYYYEIIIBBBVVVNNNNNNVVEEEEEEEGGGGGGGGNNNNNNNNNJJGGUUUUUUUUUQQQQQQDDDDDDDDDDDDDSSSSSSSS
GGGGGYYYYYYYYYXXXXGGGGGGGGGGYYYSSSYYYSSRRRRRRRRRVVVQQQQIYYIIIIIIBBVNNNNNNNVVEEEEEEEUUUGGGGGNNLGNNNNGGGGGGUUUUUHHQQQQTQQDDDDADDDDNDBBBSSSSSSS
GGYGYYYYYYYYYYXXXXGGGGGGGGGGYYYYYYYRRSRRRRRRRRRRRRVVVQQIIVIIIIIIIGNNNNNNNNMMEEEEEEEEEGGEEERRRGGGGGGGGGGUUUUUUUHGQSQTTTCDDDXXDDDDNDBSSSSSSSSS
YYYYYYYYYYYYYYXXXGGGGGGGGGGGYYYYYYYRRRRRRRRRRRRRRRVVVVQQQVXIIIIIQGNNNNNNNNUMEEEEEEEEEEEEEEEERGGGGGGGGGGUUUIUUHHHHTTTTTCDDDXXDDDDNNSSSSSSSSSS
YYYYYYYYYYYYYYYXXGGGGGGGGGGGGYYYYTRRRTRRRRRRRRRRRRRVVVVVVVIIIGIGGGGNNNNNNNNAEEEEEEEEEEEEEEEERGGGGGGGGGGIIUIIGGHHHTTTTTCDDDXXXXNDDNSSSSSSSSSS
YYYYYYYYYYYYYYYXGGGGGGGGGGGGGGGGYTKTTTTRRRRRRRRRRRVVVVVVVVVVIGGGGGGGNNNNNNZMMEEEEEEEEEEEEEEERRRGGGGGRRIIIIIIIIOHHTTTTTTTXXXXXXNNNNSSSSSSSSSS
YYYYYYYYYYYYYYYGGGGGGGGGGGGGGGTTTTTTTTGTRRRCRRRUUUUUUUUUVVGVGGGGGGGGGNNNNMMMMMMEEEMEEEEEEERRRRGGGGRRRRRIIIIIXITHHTTTTTSSSXXXXXNUUNSSSSSSSSSS
OOYYYYYYYYYYYSYGGGGGGGGGGGGGGGGGGTTTTTTTTCCCRRRUUUUUUUUUVVGGGGGGGGGGGGGGTTMMMMMEEEMMEEEEEEERRRGRRRRRRRRIIIIIXTTTTTTTTSSSXXXXXXUUUNNNSSSSSSSS
OYYYYYYYYYYYSSGGGSSSGGSGGGGGGGGTZTTTTTTTTCCCCCRUUUUUUUUUVVVTGGGGGGGGGGGGMMMMMMMMMMMEEEEEEDDTTRRRRRRRRRBIIIIITTTTTTTSSSSSXXXXXXUUNNNNSSSSSSSS
OYYYYYYYYYYYSSSSSSSSSSSSGGGGGHGTTTTTTTTTTTTCCCYUUUUUUUUUZVNOOGGGGGGGGGGGGMMMMMMMMMMEEEEEUTTTTRRRRRRRRRBIIIIIPTTTTTSSSSSSXXXXXUUUNNNNNNSSSSSS
OOYOYYYYYSSSSSSSSSSSSSSSGGGGGHHHTTTDTTTTTCTUUUUUUUUUUUUUODOOOGGGGGGGGGGMMMMMMMMMMMMZUUUUUTTTTTTRRRRRRRBBIIIBBBBTTSSSSSSSSXXNNUNNNNNNNNNSSSSN
OOYOOYYYGSSSSSSSSSSSSSSSGGGGGHHHHTTTTTZCCCTUUUUUUUUUUUUUOOOOOOOGGGGGGGGYMTMMMMMMMMMMUUUUATTTTTTRRRRRRRBBBBBBSBBSTTSSSSSSSSXXNNNNNNNNNNNSNNNN
OBOOOOYBGGSSSSSSSSSSSSSSTGUGHHHHHHHTTTCCCCCUUUUUUUUUBBBOOOOOOOOPOOGGGGGYYYMMMLMMMMMMUUUUAATTTTTTRRRRRRBBBBBSSDSSTSSSSSSSSSXNNNNNNNNNNNNNNNNN
BBBOOYYBGSSSSSSSSSSSSSSSTTIKIHRHHHIIECCCCCCUUUUUUUUUBBBOOOOOOOOOOOIOGGYYYYYLLLMMMMMMUUUUATTTTTTTRRRRRRRBBBBSSSSSSSSSSSSJJJJNNNNNJNNNNNNNNNNN
BBBBBBBBBBSZSSSSSSSSSSSSIIIIIRRRRHISCCCCCCCUUUUUUUUUBOBBBOOOOOOOOOOOOOYYYYYYLLMMMMMMMUUUGGTTTTTTGRRRRRRBBBBSSSSSSSSSSSSJJJJJJJJJJNNNNNNNBBNN
BBBBBBBBBBSZSSSSLSLSSSSSIIIIIRRRIIIICCYCCCCUUUUUUUUUBOBBSOOOOOOOOOOYYYYYYYYYLLUUUMZMMUUUUGKTTGTTGGRRRRRRBBSSSSSSSSSSSSSJJJJJJJJJJJNBBBBBBBNN
BBBBBBBBBVSSSSSSLLLLSSIIIIIIRRIRIIIAACCCCCTUUUUUUUUUVOOBSOOOOOOOOOTTYYYYYYLLLUUUUUUUUUUUGGGTGGGGGGGGRRRRBBSSSSSSSSSSSSSJJJJJJJJJJJJJBBBBBBNN
BBBBBBBBBSSSSSSSLLLLLIIIIIIIRIIRIIIJAAJJCCCUUUUUUUUUOOOOOOOOOOOOOOYYYYYYYYYLLUUUUUUUUUQUGGGGGGGGGGGCCRRRBBSSSSSSSSSSSSSJJJJJJJJJJJJJBBBBBBBN
BBBBBBBBBSBBSSSLLLLLAAAILIIIIIIIIJJJAJJJCJJUUUUUUUUUVVVOOOOOOOOOOYYYYYYYYYYYYYUUUUUUUUQQGGGGGGGGGGGGCRCCSSSSSSSSSSSSSJJJJJJJJJJJJJJBBBBBBBBN
BBBBBBBBBBBSSSLLLLLLLLLLLLIIIIIIJJJJJJJJCJJUUUUUUUUUVVVOOOOOOOOOOYJJYYYYYPYYYNNUUUUUQQQGGGGGGGGGGGGCCCCCSSSSSSSSSSSHSSJJJJJJJJJJJJJBBBBBBBNN
BBBBBBBBBBBSSSSLLLLLLLLLLLIIIIIIXXJJJJJJJJJUUUUUUUJJVVVLQQQOOOOOOOOJYYYYYPPPPPNNUUUUUQQQUGGGGGGGGGGGCCCCSSSSSSSSSSSHHSJJJJJJJJJJJJJBBBBBBBBB
BBBBBBBBBBBGSSLLLLLLLLLLLIIIIIIIXXJEEJJJJJJUUUUUUUJJJVVLQQXQOOOOOOUYYYYYYPNPPPNPPPUUUQQQQGGGGGGGGGGGCCSSSSSSSSSSHHHHHSOJJJJJJJJJJJJBBBBBBBBN
BBBBBBBBBBPWLLLLLLLLLLLLLIIIIIIIXXXXEDDJJJJUUUUUUUJJJJQQQQQQQQQQQOUUYYYYYQPPPLPPPPOOQQQQGGGGGGCCCGGCCCSSSSSSSSSSHHHHHOOOJJJJJJJJJJJBBBBBBNBN
BBBBBBBBWWWWWWALLLLLLLLLIIIIIIIXXXXXXDDJJJVUUUUUUUJJJQQQQQQQQAAUUUUUUUUYYQPPPPPPPPOSSSSSSSGGGGCCCCCCCCCCSSSSSHHHHHHHHHOOOJJJJJJJJJJJBBNNNNBN
BBBBBBBBBBBWWWALLLLLLLLLLIIIIIIIXDXXXDDDVVVVVVVJJJJJJQQQQQQQAAAUUUUUUUYYYYPPPPPPPPPSSSSSSSGGICCCCCCCCCSSSSSSSSSSHHHHHOOOOJJJJJKKJJJJLSNSNNNN
BBBBBBBBBBWWWWJJJLLLLLLLLLLIIIIIDDDXDDVIIVVVVVVVJJJJJTQQQQQQAUUUUUUUUUUUYUUPPPPPSSPSSSSSSSGGGCCCCCCCCCCSSSSSSSSHHHHGHOOJJJJJJKKJJLJLLSNSNNNN
BBBBBBBBBWWWWJJJJJLLWLRJSIIIIIDDDDDDDDVVVVVVVVVPJJJJQQQQQQQQQGUUUUUUUUUUUUPPPPPPPSSSSSSSSSGCCCCCCCCCCCSSSSSSSSSSSSSSMJJJJJJJJJJLLLLLLSSSNNNN
BBDDBBBDBWJJJJJJJJJJJJJJJIIIIIDDDDDDDDVVVVVVVVVVVJJQQQQQQQQQQQUUUUUUUUUUUUPPPPPPPSCSSSSSSSGCCCCCCCCCCCCSSCSSSSSSSSSSJJJJJJJJJJJLLLSSSSSSSSNS
BPDDDDDDWWWJJJJJJJJJJJJIIIIIIIWDDDDDDDVVVVVVVVVVJJJQQQQQQQQQQQJUUUUUUUUUUUUUPPPPPPCSSSSSSSCCCCCCCCCCCCCCCCCSSSSSSSSSPJJJJJJJJJJJLLLSSSSSSSSS
PPPNNDWWWWWJJJJJJJJJJJJIIIIIIWWWDDDDDDDDVVVVVVVVJJQQQQQQQQQJJJJUUUUUUUUUUUUPGPPPDDSSSSSSSSSSJJCCCCCCCCCCCSSSSSSSSSSSSJJJJJJJJJJJJSSSSSSSSSSS
PNNNNNNWWRWJJJJJJJJJJJJJIIIIIWWWWDDDDDDDCVVVVVJJJJJJQJQJQJJJJJJUUUUUUUUUUUPPPPPPDSSSSSSSSSSACCCCCCCCCCCCCSSSSSSSLSSSJJAJJJJJJJOOOOSSSSSSSSSS
PPNNNNNWRRWWJJJJJJJJJJJIIIIIIWWWWDDDDDDDCVVVVVVJJJJJQJQJJJJSJJJUUUUUUUUUUUPUFFPPDDSVSSSSSASAAACCCCCCCWWOSSSSSSSSSSSSSAAJJJAJJTSSSSSSSSSSSSSS
PPNNNNNNNRWRRJJJJJJJJXJIIIIIIWWWWWDDDDDDDVWVVVVVJJJJJJJJJJJSJJJLLUUUUHUUUUUUFFPPFSSFSSSAAAAAAACCIIIJCWWWWSSSSSSSSSSSAAAAAAAAATTSSSSSSSSSSSSS
PPNNNNNNRRRRJJJJJJJJXXXIIAIWWWWWWHDDADDIIVWWWWVJJJJJJJJJJJJSSSSSSUUUUUUUUUUUUFFJFFFFFFAAAAAAAAAIIIIIWWWWWSUUUUUUUUUAAAAAAAAAATTSSSSSSSSSSSSS
NNNNNNNNNRRRRJJGGJXXXXXIAAAAWWWWGDDDAADLLWWWLLMLLLJJJJJJJJSSSSSSSSUUUUUUUUUUUJJJJFFFFFFFFAAKAAAIIIIIWWWWLSUUUUUUUUUAAAAAAAAAATTSSSSSSSSSSSSS
NNNNNNNNNRRRRRGGGJXXXAAAAAAAWAWADDDDAADLLLWWLLLLLLLLJJJJJJSSSSSSSSSOPUUUUUKKCCJJJFFFFFFFFFAKKAIIIIIIWWWWWSUUUUUUUUUAAAAAAAAAATTTSSSSSSSSSSSS
NNNNNNNNNRRRRGGGGGXXAAAAAAAAWAWAADDAAAAALLLLLLLLLLLLJJJJJJJBSSSSSSSPPPPUUKBKCKFFFFFFFFFFKKKKKAAAAIIIIWWAAAUUUUUUUUUAAAAAAAAAATTTTSSSJJSSSSSS
NNNNNNNNNRRRRRRGGGGGAAAAAAAAAAAAAAAAYYYALLLLLLLLLLLJJJJJJJBBBSSSSSSSPPPKKKKKKKFFFFFFFFFFKKKKKKAAIIIIUUWAAAUUUUUUUUUAAAAAAAATATTTTTTSJJSSSSSS
NNNNNNRNRRRRRRRGGGGWAAAAAAAAAAAAAAAAYYYZLLLLLLLLLHHHJJJJJJBBBSSSSSPPPPPPPTKKKKKKKKKKKFKKKKKKKKKLLIIIIIWWWYUUUUUUUUUAAAAAAATTTTTTTTJJJJJSJSJS
NNNNNRRNRRRRRRRGGGGGGGGAAAAAAAAAAAAYYYYYLLLLLLLLLLHHHHJJJBBBFSXSSSSPPPPPPPCKKKKKKKKKKKKKKKKKKAALLIIIIILZYYUUUUUUUUUAAAAAAATTTTTTJJJJJJJSJJJS
SQQNUURRRRRRRRRGGGGGGGGGAAAAAAAAAAYYYYYLLLLLLLLLLLLHHHJJJJBBBBBBSSPPPPPPPPKKKKKKKKKKKKKKKKKKKAALLIIIWWLZYYYYYYYAAAAAAAAAAATTTTTTJEJJJJJJJJJS
QQQQQUUUURRRRRGGGGGGGGGGAAAAAAYYYYYYYYYLLLLLLLLLLLLHHHYJJJBBBBBBDPPPPSPPPAKKKKKKKKKKKKKKKKLKAALLLLLLLLLLYYYYYYYAAAAAAAAAATTTTTTJLJJJJJJJJJJE
VQQQQQUUURRRRGGGGGGGGGGGGAAAAAYYYYYYYYYLLLLLLLLLLLLHHYYYYNBBBBBBDPPPPSPPPAKKKKKKKKKKKKKKKKLAALLLLLLLLLLLLLYYYYYAAAAAAATTATTTTTTJJJJJJJJJJJJJ
VVVQQUUUUURRRRGGGGGGGGGGGAAAAAYYYYYYYYLLLLLLLLLLLLLHYYYYBBBBBBBBDPSSSSSPPAAKTKKKKKKKKKKLLLLLLLLLLLLLLLLLLYYYYYYYAAAAATTTTTTTTTTTJJJJJJJJJJCJ
VVVVUUUUTRRRRRRGGGGGGGGGAAAEAYYYYYYYYYLLLLLLLLLLLLLYYYYYBBBBBBEBDSSSSSSSAAATTKKKKKKKKKKLLLLLLLLLLLLLLLLLLLLYYYYYYAAAAAAATAZZZZTTTJJJJJJJJJJM
VVVUUUUUTTRRRRRRGJGGGUGXRRRERRRRYYYYYYYLLLLLLLLLLKYYYYYYYBBTWWEEWSSSSSSSSTTTTTTKKTKKKKLLLLLLLLLLLLLLLLLLLLFYYYYYYYYAAAAAAAZZZZZZZJJJJJJJJMMM
VVVUUUUUUTRRRRUUJJUUUUXXXRRRRRRRYYYYZZYYYLLLLKKKKKKYYYYYYYWWWWWWWSSSSSSSSWWTTTTTKTKKKKLLLLLLLLLLLLLLLLLLLLYYYYYYYYGGGGGGGGGGZZZZZZJJJJJJJMMM
VVVVUUUTTTRRRRUUUUUUUUXXXRRRRRRRVYZZZZZYYLLLLLKKKKYYYYYYYYWWWWWWWWSSSSSSSTTTTTTTTTTTKPPLLLLPZIZLLLLLLLLLLLYYYYYYYYGGGGGGGGGGZZZZZZZZZJJJJJMM
VVVVQUTTTTRTRRUUUUUUUUXXRRRRRRRVVZZZZZZZYYLLKKKKKKKYYYYWYYWWWWWWWWWSSSSSTTTTTTTTTTTTKPPLPPPPZZZZZLLLLLLLTTYYYYYYYPGGGGGGGGGGZZZZZZZZMMJMMMMM
VVVVVUTTTTTTRRUUUUUUUUUUURRRRRRRRRZZZZZZYGGLGKKKKKKWWWWWWWWWWWWWWWWWSSTTTTTTTTTTTTTTTPPPPPPPZZZZZZLLLLLLYYYYYYYYYPGGGGGGGGGGGGGGZZZZZMMMMMMM
VVVVVTTTTTTTRUUUUUUUUUUUURRRRRRRRBZZZZZZZGGGGKGKKKKWWWWWWWWWWWWWWWSSSVRTTTTTTTTTTTTPPPPPPPPZZZZZZNLLLLZZYYYYYYYPPPGGGGGGGGGGGGGGZZMMMMMMMMMM
YYYVTTTTTTTTRRUUUUUUUUUUURRRRRRRZBZZZZZZGGGGGGGGGKKKWWWWWWWWWWWWWVVVVVVTTTTTTTTTTTTTTPPPPZZZZZZZZNNNNLLPPYYYYYYPPPGGGGGGGGGGGGGGZZMMMMMMMMMM
YYYYYTTTTTTTTUUUUUUUUURURRRRRRRZZZZZZZZGGZGGGGGGKKKKWWWWWWWWWWWVVVVVVVVTXTTTTTTTTTTQPPPPPZZZZZZZZNNNNLNPPYYYYPYPPPGGGGGGGGGGGGGGZZZMMMMMMMMM
YYYYTTTTTTTTMMMUUUUHRRRRRRRRRRRRVZZZZZZZZZZGGGGBBBBKWWWWWWWWWWVVVVVVVVVVVVTTTTTTTTQQQPPPZZZZZZZZZZNNNNNPPYZYYPYPPPGGGGGGGGGGGGGGOOZZZZMMMMMM
YYVVIIETTTTTMMMUUUURRRRRRRRRRQCQQAAZZZZZZZGGGGGBBKKKKKMWWWWWWWWVVVVVVVVVTTTTTTTTTHQQZZZZZZZZZZZZZZNJNPPPPPYYYPPPPLLLLLLGGGGGGGGGOOZZZZMMMMMM
YYVIIIETTTTTTMMUUURRRRRRRRRRQQQQQQQQZZZZZZXGGGKKKKKKKMMMWWWWWWWWVVVVVVVVTTTTTTTTTTQQQZZZZZZZZZZJZZJJNJPPPPPYPPPPPLLLLLLGGGGGGGGGOOZZZMMMMMMM
VVVVVEETTTTTTTMMMMRRRRRRRRQQQQQQQMQXZZXXXXXGGKKKKKKKMMMMMMWWWWWWWVVVVVVVTVTTTTTTTTQQQQQZZZZZZZZJJZZJJJPPPPPYIIPPPPLLDLLGGGGGGGGGZZZZZMZMMMMM
VVVJVEEGTTTTTMMMMMRRRRRRRRRQQQQQQQXXXZXXXKKKGKKKKKKKMMMMMMWWWWWWUUVVVVVVTVTTTTTTTTQQQWQQQZZQZZZJJJJJPPPPPPIIIIIIIPPLLLLGGGGGGGGGVZZZZZZZMMMM
VVVVOEEEEETTTMMMMMMRRRRRRRRQPQQQQQXXXZZXXKKKKKKKKKKKMMMMMMMMWWWWWUBBVVVVTVTTTTJTWWWWQWWQQQZQZZZJJJJPPPPPPPPIIIIIPPPPPLLLLHHOOOOOZZZZZZZZMMMM
VVVVVVEEETTTTMMMMMMMRRRRRRPPPQQQQXXXXXXXXKKKKKKKKKKKMMMMMMBMBWBBBBBBVVVVVKKUUTJJWWWWWWWQQQQQQZJJJJJPPPPPPPPPPIPPPPPPPHHHHHOOOOOOOZZZZZZZZZMM
VVVVVVEEEETTTMMMMMMMRRRRRPPPPAQQXXAXXXXXXAAAKKKKKKKMMMMMMMBBBBBBBBBBBKKKKKKUUUWWWWWWWWWQQQQQQZZJJJJJJPPPPPPPPPPPPPPPPHHHHHOOOOOOZZZZZZZZZMMM
PVVVVVSEEEEEMMMMMMMMMRRRRPPPPAAAAAAAXXXXXAAAAAAAKARMMMMRRMHBBBBBBBBBKKKKKKUUUWWWWWWWWWWWQQQQQZZJJNJPPPPPPPPPPPEEPPLPHHHHHOOOOOOOZZZZZZZZZMMM
VVVVVVSEEEEEUUUUUMMMMMRRRPPPPAAAAAAAXXXXAAAAAAAAAARMRRRRRRRRBCBBBBBBKKKKKKKWWWWWWWWWWWWWQQQQQNNNNNJPPPPPPPPPPPEEZPPHHHSSHOOOOOOOZZZZZZAZAAMM
VVVVVVVUUHHHGGUUUUMMMMRPRPPPPAAAGKKAAXXAAAAAAAAAAARRRRRRRRRRVBBBBBBBKKKKKKKWWWWWWWWWWWWQQQQQQNNNNNJPPPPPPPPPPPEEEPPEEHSSOOOOOOOOOZZZZZAAAAMM
VVVVVVVUUHHHGHUUUUUUUUPPPPPPPGGGGKKAXXAAAAAAAAAAAARRRRRRRRRRRBBBBBBBBKKKKKKUTWWWWWWWWWQQQQQQQQNNNJJPPPPPPPPPPPREEEEESSSSSSOOOOOOOZZZZZAAAAAA
VVVVVJVUUHHHHHUUUUUPPPPPPPPPPGGGKKGXXXAAAAAAAAAAAARRRRRRRRRRREEBBEEEBBKKKLKTTTTWWWWWWWWWQQQQQQQQNJJJJJPPPPPPPPEEEEEEESSSSOOOOOOOOOOZZAAAAAAA
VVVVVJVUUHHHHHHUUUUPWWPPPPPPPPGGGGGGXXAAAAAAAAAAAAARRRRRRRRRREEWWWWWWWWKKKTTTTTWWWWWWQQQQQQQQQQNNNNJJJPPPPPPPEEEEEEEESSSSOOOOOOOOOOWAAAAAAAA
UUUUUUUUUUHHHUUUUUPPPWPPPPPPPPGGGGGGXXAAAAAAAAAAAYRRRRRRRREEEEFWWWWWWWWKKTTTTTTWWWWWWQQQQQQQQQQUUNNJJJJJJPPPEEEEEEEOSSSSSNNNNNNONOOOAAAAAAAA
UUUUUUUUUHHHHUUUUUPPPWWPPPPPPPGGGGZZAAAAAAAAAAAAAYYRRRRREEEEWWWWWWWWWWWKKKKKKKKKKWWWWQQQQQQQQQQUYNNYYYJJJPPPXEEEEEOOOSSONNNNNNNNNNNAAAAAAAAA
UUUUUUUUHHHHHHUUUUUUUWWPPPPPPPPGGGPPAAAAAAGZAAZZYYYRRRRRREEEWWWWWWWWWWWKOKKKKKKKKIIIIQQQQQQQWWUUYYYYYYYJJYXXXEEFOYOOOOOOONNNNNNNNNNAAAAAAAAA
UUUFEFBBBHWWUUUUUUWWWWWWWPPQPPPPPPPPEAAFAAZZVZZZYYYRRRRREEEFWWWWWWWWWWWWWWWWWWWWKIIIHQQQQQQQWWUUYYYYYYYYYYXXXEXOOOOOOOONNNNNNNNNANAAAAAAAAAA
FFUFFFBBBBWWWWUUUWWWWWWWWPPQQQPPPPPPPPYZZZZZZZZZYYYFREEEEEFFWWWWWWWWWWWWWWWWWWWWKIHIHQQQQQWWWWYYYYYYYYYYYYXXXXXOOOOOOONNNNNNNNNNAAAAAAAAAAAA
FFFOFFFBBWWWWUUUWWWWWWWKKQPQQQPPPPPPPPYYZZZZZZZZZYFFEEEEEFFFFWWWWWWWWWWWWWWWWWWWKIHHHQQQQQWWWWWWWYYYYYYYYYYYYYOOOOOOOONNNNNNNNNNNAAAAAAAAAAA
FFOOFFBBBWWWWWWWWWWWWKKKKQQQQQPPPPPPPPPYYZZZZZZZZZFFEEEEEFFFFWWWWWWWWWWWWWWWWWWWKHHHHQHHHWWWWWWWWWWWYYYYYYYYYIOOOOOOOOOONNNNNNNNNEAIIAAAAAAA
FFFFFFFBFFFWWWWWWWWWWKKKKQKQPQPPPPPPPPPPPZZZZZZWWWFFFFFFFOFFFWWWWWWWWWWMKKKKKKKKKHHHHHHHWWWWWWWWWWWWYYYYYYYIIIIOOOOOOONNNNNNNNNNEEEEEIAAAAAA
XFFFFFFFFFFWWWWWWWWWWWKKKKKQPPPPSSPPFPPZZZZZZZZWWWFFFFFFFFFFFWWWWWWWWWWKKKKKKKKKJJJHHHHHHWWWWWWWWRYYYYYYYYYIIIIIOOOOOONNNNNNNNNEEEEEEAAAAAAA
XFFFFFFFFFFFFWWWWWWWWKKKKKKKSPPSSSSSSSPPZZZZZWWWWWFFFFFFFFFFFWWWWWWWWWWKKKKKKKKKJJTTYHHHHHWWWWWWWWJYYYYYYYIIIIIIOOOOONNNNNNNNNNEEEEEENAAAAAA
FFFFFFFFFFFWWWWWWWWKKKKKKKKPSSSSSSSSSSPZZZZZZZZWWWWFFFFFFFFWWWWWWWWWMMKKKKKKKKKKJJJTYHHHHWWWWWWWWWYYYYYYYYIIIIIIOOOOOONNNNNNNNNEEREEEEAAAAAB
FFFFFFFFFRRRRRRRWKKKKKKKKKPPSSSSSSSSSZZZZZZZZZZZWWWFFFFFFFFWWWWWWWFFJJKKKKKKKKKKJJJJYYYYHWWWWWWWWWWWYYYYYYYYIIIIOOOOOOONVVNNNEEEEEEEEEAATABB
FFFFFFFFFFRRYRRRKKKKKKKKKKKKDSSSSSSSSZZZZZZZZVWWWVVFFFFFFUFWWWWWWWVFFJKKKJJJJJJJJJJJYYYYYFWWWWWWWWWWYYYYYYYYIIIIOOOOOOOOAVVNNNEEEEEEEEEEEAEB
FFFFFFFFFFRRRRRRRVKKKKKKKKKSSSSSSSSXSSZZZZZZZVVVVVVFFFFFFUFWWWWWWWVVVJJJJJJJJJJJJJJFYYYYWWWWWWWWWWYYYYYYYYYIIIIIIOOOBBVOVVVVVEEEEEEEEEEEEEEE
FFFFFFFFFFFRRRRRRVRKKKKKKKKKKSSSSSSSSZZZZZZZRRVVVVZZFFVVVFFWWWWWWWVVVJJJJJJJJJJJJJNNNNNWWWWWLWWWWYYYYYYYYIIIIIIIIIOBBVVVVVVVVVEEEEEEEEEEEEEE
FFFFFFFPPPRRRRRRRRRRKKKKKKKKKSSSSSSSSSZZZZZZRRVVVVVZZZVVVVVWWWWWWVVVJJJJJJJJJJJJJJJJNNNNWWWWWLLWLLLLLLYIIIIIIIIIIMMBBVVVVVJJJJJEEEEEEEEEEEEE
FFFFFFFPPPRRRRRRRRKKKKKKKKKKKKSSSSSSSSHZZZZZRVVVVVVVVVVVVVVVVVVVVVVVVVJJJJJJJJJJJJJJJNNNNNWWWWLLLLLLLLIIIIBIIIIIMMMMBVVVJJJJJJJJEEEEEEEEEQED
FFFFFFRPPPRRRRRRRRRKKKKKKKKKKHHSHSHHLLBXZBRRRRRVVVVVVVVVVVVVVVVVUVVVVJJJJJJJJJJJJJJJNNNNNWWWWWLLLLLLLLIIIBBBIIIIMMMMBVVJJJJJJJJJJJEEEEEEEEED
FEFFFFRPPPRRRRRRRRRKKKKKKKKAHHLHHHHHBLBBBBRRRRVVVVVVVVUUVUUUUUUUUUBBBBBJJJJJJJJJJNNNNNNNNWWWLLLLLLBLBBBSSBBBBBIIMMMMMVFJJJJJJJJJJJJJEEEEEEEE
EEBBBFRPPPRRRRRRRRRRHKKKJKKJJHHHHHHHBBBBBBRRVRRVVVVVVUUUVVUUUUUUUUUEBBBBBBJJBBBNNNNNNNNZNWWWLLLLLLBLBBBBBBBBBIIIMMMMMMEAJJJJJJJJJHJJEEEEEEEE
EEBBBBRPPPRRRRRRRRRRHHHHJKKJCHHHHHHVBBBBBBBRVVVVVVVVVUUUUUUUUUUUUUUBBBBBBBBBBBBNNNNNNNZZZWWWWLZLLLBBBBBBBBBBOOIIIIMDDMEEJJJJJJJJJJJJOJQQEQSE
EEBLBMMPPPRRRRRRRRRRHHJJJJJJCHHHHHHVBBBBBBBVVVVVVVVVVUUUUUUUUUUUUUUBBBBBBBBBBBBNBBNNNNNZZWWWWWZZZBBBBBBBBBBOOOIIIIDDDMEDJJJJJJJJJJJJJJQQQQQH
EEBBBBBPPPRRRRRRRRRRRHHJJJJJJHHHHHVVBBBBBBYYVVVVVVVVVUUUUUUUUUUUUUUBBBBBBBBBBBBBBBNNNZZZZZZZZZZZZBBBBBBBBOOOOOIIIODDDDEDJJJJJJJJJJJJJJQQQQQH
EEEBBBBPPPRRRRRRRRPPPJJJJJJJJJHHHHVBBBBBBBVVVVVVVVVVVVUUUUUUUUUUUUBBBBBBBBBBBSBBBZZZZZZZZZZZZZZZZZBBBBBOGOOOOOOOOOODDDDDJJJJJJJJJJJSSQQQQQQQ
BBBBBBBBJBRPPRRRRRPPPJPJJJJJJJJKHVVBBGGBGBGGVVVVVVVVVVVUUUUUUUUUUUBBBBBBBBBBBSBXXZZZZZZZZZZZZZZZZBBBBBBOOOOOOOOOOODDDDDDJJJJJJJJJJJQQQQQQQQQ
KBBBBBBBBBDPPPRRRPPPPPPPPPPJJJJKHHVBBHGGGGGVVLVVVVVVVVUUUUUUUUUKKUKBBBBBBBBBSSSSSSSZZZZZZZZZZZZZZBBBBOOOOOOOOOOOOODDDDDDDDDDJJJJJJJJQQQQQQQQ
KBBBBBBBBBDPDPRRPPPPPPPPPPPJJWJJHWBBBHHGGMGVVVVVVVVVVVVUUUUUUUUKKKKBBBBBBBBSSSSSSSSSZZZZZZZZZZZZBBBBBBOOOOOOOOOOOIOODDDDDDDDDDDDDQQQQQQQQQQQ
BBBBBBBBBBDDDPRPPPPPPPPPPPPJPWWWWWWWHHHGMZZZZZZZZVVVVVVHKKKKKKKKKKKUEBBBBBBSSSSSSSSSZZZZZZZZZZZZBBBBBBOOOOOOOOOOOOOODDDVDDDDDDDDDDQQQQQQQQQQ
BBBBBBBBBBDDDDDDPPPPPPPPPPPPPWWWWHHHHHHHMZZZZZZZZVVBVVVHCKKKKKKKKUUUUGGBBBSSSSSSSSSSZZZZZZZZZZZZZZOOOBOOOOOOOOOOOOOOVVVVDDDDDDDDDDDDDQQQQQQQ
BBBBBBBBBDDDDDDDPPPPPPPPPPPPPWWWHHHHHHHMMZZZZZZZZVVVQQCCCKKKKKKKKUUUUUBBBBYYSSSSSSNZZZZZZZTZZZZZOOOOOOOOOOOOOOOOOOOVVVVZZZZDDDDDDQQDDQQQQQQQ
BBBBBBBBBBDDDDDDDPPPPPPPPPPPWWWWHHHHHHHHHZZZZZZZZWWMMQCCCCCKKKWKUUUUUUUBBBBYYSYSSSSSZZZZZZTTTTTTTSOOOOMOXOOOOOOOOOOVVVVVZZZZDDDDDQQDQQQQQQQQ
MBBBBBBBBJDDDDDDDDLPPPPPPPPPPPPWHHHHHHHHHZZZZZZZZMMMMCCCCKKKKKKKUUUUUUUYYYYYYYYSSSSSZZZZZTTTTTTTTOOOOOOOXOOOOOOOOOOVVVVVZZZZZZDDQQQQQQQQQQQQ
MMJJJJJJJJJJDJJDDDLLLLLPPAQQPQQHHHHHHHHHHZZZZZZZZMMMMCCCCJKKDDDDUIUUUUUYYYYYYYYSSUUZZZZZZTTTTTTTAAOOOOOOOOOOOOOOOOBVVVSSSZZZZDDDDDDQQQQQQQQR
MMJJJJJJJJJJJJJJJLLLLLLLPAQQPQQQQQHHHHHHHHGGMMMMMMMMMNCCCCCCDDDDFFUUUUUYFFYYYYYYSZUZZZZZTTXTTJJJJJOOOOOOOOOOOOOOOBBVVVSSSSZZDDDEEDEQQQQQQQQR
MMJJJJJJJJJJJJJJJJJJLLLLPAQQQQQQQQHHHHVVVHGGMMMMMMMMMMMCCCCCDDDDDFFFUUFYFFYYYYYZYZZZZZZZZZJJJJJJJOOOOOOODLOOLLROOBBBVVSVVSSSEEEEEEEQQQQRQRRR
MMJJJJJJYJJJJJJJJJLLLLLLQAQQQQQQQQHVVHVVVGGGGMMMMMMMMMMMNCCCWDDWDFFFUFFYFFYYYYYZZZZZZZZZZZJJJJJJJOOOOOOOLLLLLLLOBBBVVVVVSSSSEEEEEEEQQQQRRRRR
MMMJJJJJJJJJJJJJJLLLLLLLQQQQQQQQVVHVVHVVVGGMMMMMMMMMMMMMMNCSWDDWWFFFFFFFFFFYYYYZZZZZZZZZZZZYYYYJJOOOOOOLLLLLLLLLLVVVVVVYESSSSEEEEEEEFERRRRRR
MMJJJJJJJJJJJJJJJLLLLLQQQQQQQQQQVVVVVVVVVVVMMMMMMMMMMMMMMMCSWWWWWFFFFFFFFFFYYYZZZZZZZZZZZJAYYYYJJJOOOOALQLLLLLLLKVVVVVVYEEEEEEEEEEEEEERRRRRR
MMGGJJJJSSJJJLJLJLLLLLLQQQQQQQQQQQVVVVVVVVVVMMMMMMMMMMMMMWSSWWWWWFFFFFFFFFFYYYZZZZZZZZZZZJJYYYYJJBOOOOOOSSLLLLLKKVVVVVVEEEEEEEEEEEEEEAERRRRR
MMGGJJSSSSSJJLLLLLLLLLLQQQQQQQQQQVVVVVVVVVVVMMMMMMMMMMMMMWWWWWWWWFFFFFFFFFFFYYYZZZZZZZZZZZYYYYYJJJWOOOCCSSLLLELKKKVVVVVEEEEEEEEEEEEEEEERRRRR
MMGGJSSLLSLLJLLLLLLLLLLQQQQQQQVVVVVVVVVVVVVVVJJMMMMMMMMWWWWWWWWWWFFFFFFFFFYFYYYYZZZZZZZZZZYYYYYJJJJSSSSCSSSLLEEVVVVVVVVVEEEEEEEEEEEEEEERRRRR
MMMMSSSLLLLLJLLBLLLLLQLQQQQQQVVVVVVVVVIVVVVZZJJMMMMMMMWWWWWWWWWWFFFFFFFFFYYYYYYYZZZZZZZZZZYYYYYJJJSSSSSSSSSSEEEVVVVVVVVVEEEEEEEEEEEEEEEERRRR
SSSMSLLLLLLLLLBBBLLLLQFFFFFFFFFVVVVVVVIVVVZZJJJJJMMMMMMMWWWWWWWWFFFFFFFFFFFYYYYYZZZZZZZZZZYYYYYJRJSSSSSSSSSSSRRRRRRRRVVVVWEEEEEEEEEEEEERRRRR
SSSSSSLLLLLLBBBBBBBBBBFFFFFFFFFVVVVVVVIIIIZZJJJJMMMMMMMMWWWWWWWWFFFFFFFFFFFYYYYYYYFFWZZZZLYYYYYJRRSSSSSSSSSSSRRRRRRRRVVVWWEETEEEEEEEGEEERRRR
SSSSSLLLLLLLLBBBBBBBBBFFFFFFFFFVVVVVIVIIIZZZJJJJXXXXWMMMWWWWWWWWAFFFFFFFFFFFVVYYVVFFFFFLLLYYYYYJRRSSSSSSSSSSSRRRRRRRRVVVVEEIEEEEIEEEEERRRRRR
SYYSSLLLLLLLBBBBBBBBBBFFFFFFFFFFFFVIIIIISZZJJJJJXXXWWMMMMWWWWWEFFFFFFFFFDFVVVVVVVAFFFFFFFLYYYYYJJSSSSSSSSSSSSRRRRRRRRVVVVIIIEEEIIEEEXERRRRRR
SYYYSSSSLLSLBBBBBBBBBBBFFFFFFFFFFFVIIIISSSSSJJJXXXXWWWWWWWWWWXXFFFFFFFFDDDDVVVVVAAFFFFFFFFYYYYJJJSSSSSSSSSSSSRRRRRRRRNNVVVIEEEIIIIRRRRRRRRRR
SYSSSSSSSSSSBBBBBBBBBBTFFFFFFFFFFFVIIIIIISSSSXXXXXXWWWWWWWWMWXXFFFFFFFFDDDDVVVVVVAFFFFFFFFYYYYJJJSSSSSSSSSSSSRRRRRRRRNIIIIIIIIIIIIRRRRRRRRRR
SYSSSSSSSSSBBBBBBBBBBBTFFFFFFFFFFFIIIIIISSSXXXXXXXXXWWWWWWWMXXXMMFFFFFDDDDDVDVVVVFFFFFFFFYYYYYYJYSSSOOSSSSSSSRRRRRRRRNNIIIIIIIIIIRRRRRRRRRRR
SSXXXSSXSSSSSUBBBBBBBTTFFFFFFFFFFFIIPPPPXSSSXXXXIXXXWWWWWMMMMMMMMFFFFDDDDDDDDVVVFFFFFFFFFYFFYYYYYYYOOOSSSSOSURRRRRRRRRRIIIIIIIIIIIRRRRRRRWRW
SSSXXXXXSSSSUUBBBBBBBTTTTTFFFFFFFFAPPPPLLLLLLLLLIIWWWWWWMMMMMMMMMMFMMDDDDDDDDDDQQFFFFFFFFFFYYYYYOOOOOOOOOOOUUUGGGGRRRRRNLIIIILIIIIRGWRRRWWWW
SSSXXXXXSSSUUUUBBBBBBTTTTTTTBBBBBAAAPPPLLLLLLLLLIIIWWWWWIIIMMMMMMMMMMDDDDDDDDDDQZFFFFFFFFFFYYYYYOOOOOOOOOOOUUGGGGDRRRRRNLLLLLLIIIIGGWWWWWWWW
SSSXXXXXSSSUUUUBBBBTTTTTTTTTBBBBBAAFFFFLLLLLLLLLIIIIIIIIIIIMMMMMMMMMMDDDDDDDQQQQQQFFFFFFQFFYYYBBBOOOOOOOOOUUGGGGGDRRRRRNNLLLLLLLIIGGGWWWWGWW
SSXXXXXXSSSUUUUUBBBTTTTTTTTTTRRBBBABFFFLLLLLLLLLIIIIIIIIIIIMLMMMMMMMMDDDDDDDQQQQQQFFFFQQQFQYQBBBBOOOOOOOOOOOOGGGGDRRRRROLLLLLLLLLIGGGWWWGGWW
XXXXXXXXSSUUUUUUUUUGGTTTTTTTTRRRRBBBFFFLLLLLLLLLIIIIIIIIILLLLLMMMMMMMDDDDQDDDDQQQQFQFFQQQQQQQQBBBOOOOOOOOTTYGGGGOORRRRROOLLLLLLLLGGGGGGGGGGG
XXXXXXXXSXUUUUUUUUUGGTTTTTTTTRRRBBBBFFFLLLLLLLLLIIIIIIIIIILLLLMMMMMMMDDVDQQQDQQQQQQQFQQQQQQQQQBBOOOOOPOJJYYYGGGGGORRRRROOLLLLLLLLLGGGGGGGGGG
XXXXXXXXXXUUUUUUGGGGGTTTTTTTRRRRRBBBBFFFFFFFFFFFIIIIIIILLLLLVLLLLMLMVVVVVFQQQQQQQQQQQQQQQQQQQQBBOOOOOOOJJJYYYYGYGOOOOOOOOOOLLLLLLIGGGGGGGGGG
XXXXXXXXXUUUUUUUUGNNGGTTTTTTTTKRBBBBBBFFFFFFFFFIIIIIIIIIVLLLVVLLLLLLVVVVVFFQQQQQQQQQQQQQQQQBQBBBOOOOOOJJJJYYYYYYGOOOOOOOOOOLLGGGGGGGGGGGGGGG
XXXXXXXXXUUUUUUUGGNNGGTTTTTTTTTRBBBBBBFFFFFFFFFIFFIIIRRRVVVVVVVLVVLLVVVVFFFQSQQQQQQQQQQQQQQBBBBBBBOOOJJJJJYYYYYYGGGOOOOOOOLLLLGGGGGGGGGGGGGG
XXXXXXXXXXXUUUUUUNNNNTTTTTTTTTTBBBBBBBBRFFFFFFFFFFFIIRRRVAVVVVVLVVVVVVVFFFFFQQQQQQQQQQQNQZQBBBBBBBOOOJJJJJYYYYYYYGOOOOOOOOOLLLGGGGGGGGGGGGGG
XXXXXXXXXXXUUUUNNNNNNNTTTTTTTTTTTBRRRRRRFFFFFFFFFRRRRRRRVVVVVZZVVVVVVVVVVFFFFFQQQQQQQQQQQQBBBBBBBBOOOJJJJJYYSYYYYGGOOOOOWOOLLLGGKGGGGGGGGGGG
XXXXXXXXXXXXUUUNNNNNNTTTWTTHHHTTHRRRRRRRFFFFFFFFFRRRRRRRVVVVVVZZVVVVVVVVVFFFFFFQQQQQUQQBBBBBBBBBBBORBQQQQQQQYYGGYGGOOOOOWOOLLLLGGGGGGGGGGGGG
XXXXXXXXXXXXUUNNNNNNNNTWWTTHHHHHHHRRRRRRRRRRRRRRRRRRRRRRRVVVOVVVVVVVVVVVVVFFFFFFFFFQBBBBBBBBBBBBBBBBBWWQQQQQYGGGGGOOOOOOOOOLLLLGGGGGGGGGGGGG
XXXXXXXXXXXUUUUNNNNNNNNNWWTTHHHHHHRRRRRRRRRRRRRRRRRRRRRRRVOOOOVVVVVVVVBBVVFFFFFFFFFFBBBBBBBBBBBBBBBBBWQQQQQQYGGGGGOOOPPPOLLLLLLLLGGGGGGGGGGG`