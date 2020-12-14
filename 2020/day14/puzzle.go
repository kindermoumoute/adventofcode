package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`,
		``,
		208,
	},
}

var puzzle = `mask = 0101XX01X00X1X1011X1X000000101X10001
mem[7625] = 333450
mem[24015] = 49860
mem[42377] = 14966
mem[17961] = 3341
mem[37065] = 2066541
mask = 111101XX101X1110110101X01X101X100X0X
mem[17330] = 176272
mem[20250] = 11696927
mem[28122] = 103
mem[23322] = 1224
mem[20521] = 306265426
mem[56993] = 5315703
mask = 0101110101111X0010100000X10X000X1X1X
mem[52304] = 107295284
mem[48546] = 12756871
mem[26246] = 794803013
mem[32041] = 796
mem[20164] = 6770841
mem[48951] = 2607
mask = 110101110110X11X110X100110X110010110
mem[33686] = 1985
mem[36236] = 652423
mem[1772] = 614
mem[48552] = 61799
mem[51199] = 1768
mem[38041] = 13211
mask = 11XX01011X111110X100X1000X100X0X1001
mem[10184] = 61765
mem[17636] = 883
mem[63317] = 137
mem[4641] = 602692
mem[31656] = 7910
mem[40895] = 3100935
mask = 1101111100X0X110X0X00X00111000010X01
mem[12544] = 1947
mem[10149] = 1246276
mem[2653] = 116139
mem[41147] = 5353
mem[49305] = 186429425
mem[27914] = 675
mem[30788] = 2718
mask = 0101110101110100X010XX11011XX0011101
mem[45140] = 129899
mem[54670] = 2700003
mem[21662] = 5681585
mem[57364] = 44013091
mem[11437] = 212
mem[53267] = 45116
mask = 11100101101X11101100X010000X10X110X1
mem[4630] = 919
mem[14230] = 228044604
mem[25029] = 189185
mem[41789] = 497795355
mem[29298] = 1740608
mem[55321] = 603
mask = 00X1111110X010101X0101011X0110100X11
mem[55427] = 6760330
mem[16391] = 8497685
mask = 01110111X010111011001X1101X1001X0011
mem[59831] = 11449720
mem[24494] = 21722136
mask = 1111X011XXX010X0X100010000001001XX10
mem[10149] = 2044502
mem[26810] = 1824407
mem[24894] = 495492533
mem[13252] = 160075
mem[23314] = 96205288
mem[31410] = 123057882
mask = 1X111011XX1001X011001101101X0X11100X
mem[17518] = 2310
mem[33686] = 2407271
mem[55864] = 15074402
mask = 1101X1X1X011111110001X011010010100X1
mem[6096] = 54975
mem[10934] = 358
mem[57154] = 1490
mask = 111X0XX10X101110110001X101X0000101X1
mem[64796] = 210910
mem[4752] = 3068
mem[5629] = 1532924
mem[53057] = 8797918
mem[16901] = 3392
mem[971] = 297
mem[48019] = 80129224
mask = 11100111001011101100011111X1X0X001X0
mem[4689] = 6580
mem[61719] = 506
mem[3114] = 231484
mem[26961] = 157879585
mem[10668] = 8177
mem[5634] = 2471
mask = 0101X1011011111X11100001X000010X1X00
mem[36080] = 46408099
mem[22073] = 419830102
mem[35679] = 1702
mem[53576] = 277386
mask = 111110010X1001101100X10011XX00000100
mem[21538] = 12730624
mem[38041] = 5930570
mem[41831] = 6150636
mem[38958] = 152898386
mem[41147] = 1057472
mem[13265] = 1434313
mem[6405] = 841
mask = 1111X0110110X0101X00X100100101000011
mem[15676] = 940
mem[20515] = 97
mem[23882] = 903041426
mask = 111110110110011011XX00X11110001X0110
mem[35745] = 981
mem[24128] = 11108
mem[4977] = 1279634
mask = X1110101011X1110110001010X10X00X0110
mem[23442] = 1781008
mem[54225] = 13834
mem[48910] = 39611
mem[59114] = 1939
mem[5139] = 922190
mask = 1X11X1X1X011X11111011111X011X00010X0
mem[16307] = 28883
mem[35549] = 1434246
mem[63781] = 101884
mem[56134] = 113928
mem[24099] = 20077761
mem[7516] = 5573416
mem[59069] = 14919
mask = 11110X010110X11011X011011101X0010101
mem[43391] = 791798
mem[2652] = 12659
mem[4799] = 17430
mem[62099] = 10386985
mem[2180] = 55661090
mask = 1110010XX101XX1011X000100101XX000010
mem[20835] = 105254603
mem[39463] = 940293735
mem[50577] = 2537
mem[14690] = 130220
mem[23057] = 121136841
mask = 1110010101X01110110X01X10110010111X0
mem[18309] = 297206536
mem[17182] = 45095
mem[61572] = 826
mem[44910] = 2788180
mem[37318] = 25645
mem[49915] = 233629
mask = 1101X010011X1011010001000X000100001X
mem[35857] = 19517506
mem[5684] = 8268861
mem[43961] = 777
mem[65296] = 4313
mem[11847] = 639775290
mem[3832] = 5094416
mask = X101X11100101XXX00100X1X0101X0110X01
mem[11009] = 156925
mem[6286] = 290862
mem[63205] = 985
mem[26579] = 32557999
mem[57238] = 5759116
mem[40172] = 16031
mask = 011X01X110111110110X1X1XX11000XX10X0
mem[12977] = 53152
mem[34861] = 12118
mem[58473] = 13657
mem[58199] = 5196811
mask = 111X0101001X11001X000100X0X11XX1X000
mem[7162] = 58695
mem[35819] = 264552401
mem[41789] = 30314
mem[49377] = 191785
mask = 1111001101101XX0X1X00X010001X00XXX10
mem[2291] = 30592
mem[63205] = 491496
mem[38504] = 1049100484
mem[34919] = 178729
mem[25424] = 156
mem[65520] = 58112582
mem[52035] = 323124091
mask = X1011101X111X1X01X1001XX01X0000X1111
mem[33094] = 31397
mem[51348] = 56754
mem[14230] = 41266
mem[20522] = 11648
mask = 11X101111X1011101X0101X100X1X111X01X
mem[11136] = 3560556
mem[5587] = 921
mem[18309] = 67938806
mask = 11100101X011X1X001001XX000001XX01001
mem[4833] = 3160
mem[25408] = 123351536
mem[25049] = 109348838
mem[63781] = 390880058
mem[55179] = 25336
mem[6968] = 758
mask = 1X1X010110111X1001001000X0X011001001
mem[53455] = 9337510
mem[14234] = 60613906
mem[27428] = 11014923
mem[50788] = 476612064
mem[35743] = 54469469
mem[1614] = 123865983
mask = X11110010110011011000X01110000X0X001
mem[47883] = 31228
mem[31214] = 64574504
mem[16405] = 23878
mem[60239] = 101342
mem[16684] = 146627857
mem[16663] = 650
mask = 1111001100XXX100X10X10001000100X0XX0
mem[48927] = 5082359
mem[2881] = 1779
mem[24736] = 2016486
mem[39479] = 2018954
mem[18818] = 225
mem[31119] = 12745
mem[63944] = 1033694
mask = 1110X101101X0X0001000101X01000001X00
mem[37512] = 43119
mem[33449] = 14364
mask = XX01010X1111X1101100X1010011X0111011
mem[44289] = 10155
mem[36861] = 484156
mem[5649] = 6231
mem[44588] = 13058
mem[36130] = 117
mask = 110111X11X111110X11X1010111001X00000
mem[52268] = 501592648
mem[5815] = 66091
mem[41650] = 95926
mem[59871] = 6758
mem[51464] = 5743
mem[60497] = 62281782
mem[20522] = 93843022
mask = X0111011001001001100X00110X01X111101
mem[53275] = 17387
mem[60055] = 118119
mem[8006] = 136
mem[31684] = 438174
mask = 111100010X1X1110110X01101001X0000000
mem[7140] = 3017
mem[29078] = 38418489
mem[63673] = 171084
mem[22294] = 3837
mem[63317] = 313
mem[27563] = 52466
mask = 1111X011011X11000100X000011011X10101
mem[46366] = 3934
mem[39437] = 100546
mask = 1111X0X101100110X100X100000XX0X10001
mem[39576] = 21930920
mem[65424] = 462
mem[56340] = 353590342
mask = 111001X1001111001100XX000XX010010X11
mem[55316] = 706671322
mem[25591] = 518584
mem[50729] = 2900209
mem[33498] = 11876
mem[31561] = 259813
mem[35113] = 147880795
mask = 110101011X1XX1X0X1001110X0X1110XX000
mem[29355] = 625
mem[37409] = 220996783
mem[37711] = 135727391
mem[37047] = 3660
mem[46341] = 426315
mem[15120] = 1754833
mask = X110010XXXXX11X01100010X00X100010001
mem[53073] = 1381843
mem[13260] = 93739
mem[37872] = 52
mem[32740] = 3704961
mem[59346] = 1171
mem[29949] = 16961
mask = 1111X01X01X01X00111X000101XX10100110
mem[40895] = 8972
mem[41295] = 54333696
mem[20850] = 2653
mem[19282] = 19093
mem[38608] = 11892
mem[48832] = 206906
mem[20432] = 871
mask = X111001101101010X10X1X0101100X1010X0
mem[29949] = 23166498
mem[1906] = 8837
mem[49790] = 1479345
mask = 11X10110011X1X110100X0X000100X0100X0
mem[53057] = 6958067
mem[56592] = 937
mem[36056] = 511488563
mem[59203] = 1751
mem[37047] = 20461
mem[46044] = 11961
mask = 0X011111101X1X10X101X10XX001001X0011
mem[2180] = 290
mem[51448] = 1927
mem[33875] = 617047
mem[62545] = 38166
mem[24159] = 769
mem[56134] = 763869
mask = 11100X0XXX111110X10011100111000X10X0
mem[46115] = 74662
mem[31913] = 7995
mem[1663] = 3351
mem[15242] = 7255
mem[38592] = 1282670
mem[52969] = 54696
mem[6096] = 1276
mask = 1XX101010X111X111X0X0001X01000X100X1
mem[22073] = 58445318
mem[24069] = 44324
mask = 11110011X1101010110X000110X00000011X
mem[16663] = 24384890
mem[2762] = 37315460
mem[52499] = 215572
mem[19346] = 51671
mask = 1111001100X1010X01X1000X10101X00X0X0
mem[35139] = 73247
mem[8006] = 8172
mem[1578] = 11860073
mem[45896] = 947482
mem[58045] = 28856
mem[8032] = 225912010
mask = 1X11X0110XXX11X0X1001000X0X111010010
mem[34167] = 7412
mem[33617] = 11192
mem[37247] = 13040665
mem[64827] = 57758603
mem[5649] = 520014659
mem[7105] = 1558987
mask = 110101XX1010111111000X101111X0000001
mem[3800] = 328348
mem[60536] = 66209620
mem[5063] = 2271
mem[32395] = 133940
mem[11737] = 5453
mem[48512] = 94680
mask = XX010X10X11110110100X100XX1101010X11
mem[26620] = 2476136
mem[33872] = 930
mem[21275] = 17463
mem[13265] = 3033
mem[13058] = 1658898
mem[40725] = 430
mask = 110111111X0111101X00100100001001X0XX
mem[18233] = 8421
mem[17436] = 31387
mem[53222] = 67556
mem[20521] = 1141367
mem[54918] = 2596
mem[40403] = 1004410657
mask = 0X1X010110X111X0X11100X001000111101X
mem[17673] = 118390439
mem[27333] = 7879
mem[29078] = 15837725
mask = 01011XX1XX1X10100101010000X101X0000X
mem[59493] = 6025045
mem[16555] = 21344536
mask = 110111111X1XXX101010110X001000000101
mem[39637] = 233
mem[36506] = 939031780
mem[47242] = 33211645
mem[6561] = 217613
mask = 111101010111X111100X10011X111000X000
mem[63673] = 3846623
mem[56658] = 329423
mem[38481] = 547857
mask = X1X10X110XX011101100000111011X01000X
mem[48927] = 6384
mem[43391] = 5514
mem[21936] = 131847
mem[43030] = 810434897
mem[57830] = 35743
mem[7640] = 305
mask = 01X11111101110101101001010XXX11XXX1X
mem[60143] = 90582
mem[32081] = 7673
mem[14349] = 25972726
mem[23840] = 86305185
mem[48353] = 608866275
mem[62103] = 22452
mask = 1110XX01X0110100X1001101X01X0110010X
mem[28115] = 27314978
mem[19040] = 136523686
mem[62545] = 322784
mem[35413] = 6136384
mem[35902] = 77354799
mem[9741] = 79081
mem[63898] = 1698818
mask = XX0X11111X1X10101101XX10111001X0X110
mem[33891] = 16545
mem[40774] = 92261
mem[59203] = 15655
mem[45590] = 36588183
mask = 1111010100101110X100X10X1010010X0X00
mem[48951] = 11761726
mem[40080] = 61805
mask = 11111X0100100X1XX10011X101X1X0X11000
mem[63205] = 2654080
mem[61251] = 6112
mem[38917] = 3928
mem[43961] = 887252
mask = 1111X0X10X1001101100XX0011000X1X0100
mem[44267] = 991
mem[22054] = 129208
mem[24736] = 2063224
mask = 000111111011101011X10XX00X0100100111
mem[35194] = 9210267
mem[50934] = 1766913
mem[10934] = 21495328
mem[610] = 412419
mem[36953] = 11943
mask = 1X1X01010010111X11000000X11X00010000
mem[28679] = 637124
mem[37318] = 881
mem[57422] = 2767596
mask = 1X10X100X1111X1011000X11X1X11110X1X1
mem[27363] = 25462469
mem[31504] = 28338
mem[44316] = 1809290
mem[15190] = 7443
mem[50765] = 15365
mem[40403] = 544499158
mask = 0110011110XX1110110111X00110X0X1110X
mem[213] = 2945928
mem[14529] = 34489340
mem[24633] = 2265142
mem[14164] = 42798
mask = X101X1X110X1111011XX10011XX0110100X0
mem[2881] = 253851450
mem[38504] = 998965
mem[49657] = 570454
mem[971] = 41
mask = X11101011011X11X110XX1011111X000XX11
mem[29349] = 1319294
mem[14166] = 116618287
mem[53112] = 289785
mem[55957] = 293394
mem[4582] = 254866
mem[18343] = 57127617
mem[10149] = 15993396
mask = 1011010X011111X11X01000101100001X001
mem[21561] = 112597
mem[13051] = 29108
mem[12802] = 175602144
mem[478] = 15086994
mem[53889] = 55877423
mem[46492] = 62745
mem[48353] = 1279694
mask = 1X1X010X1X110100X1X01101000000X0100X
mem[42565] = 233433
mem[33409] = 387972295
mask = 0110010X0110110X1X0X0X10100XX1110001
mem[20392] = 423739412
mem[17500] = 362907
mem[36030] = 53419432
mem[19860] = 15602203
mask = 11X10101XX1X111X1X000X00X11000000000
mem[64796] = 198944
mem[21525] = 3974377
mem[11737] = 136209878
mem[54225] = 444434
mask = X10111X1XX1XX1101X1010X01000100100X1
mem[26125] = 490282156
mem[7498] = 516
mask = 1111X10X001X11101X00X101X10X1101000X
mem[59368] = 1003333
mem[5649] = 730126
mask = 1X0111X1111X01X011100X000X11X1X11X11
mem[17709] = 6746
mem[32839] = 3345828
mask = 11XX0110011X1111010011101101010X0011
mem[28541] = 60604
mem[42817] = 1758
mem[25224] = 902483
mem[36265] = 120914
mem[32519] = 5399
mem[36622] = 834098
mask = X1X1110101110XX011101110010010011101
mem[32388] = 1033759874
mem[63976] = 4606
mem[64485] = 25543800
mem[61692] = 64754077
mem[6489] = 851
mem[57647] = 8936
mem[39558] = 10983494
mask = XXX111111000X010100X0101X100100X0010
mem[16499] = 450156060
mem[28412] = 430144217
mem[6561] = 786352
mask = 100X0100111111101X001000001XX01X0X01
mem[32046] = 24306
mem[48960] = 26990757
mem[39206] = 4884
mask = 111X01110X101X101X001100000100011100
mem[1131] = 1269
mem[23819] = 14028
mem[53267] = 104931
mask = X1010010X1111X1101000X1010010X10XX11
mem[23882] = 3644
mem[51931] = 238578
mem[4752] = 912
mem[23830] = 1956722
mem[27068] = 823297352
mask = 0X0X1101111X0110X1100111010100001100
mem[37047] = 3669952
mem[9844] = 63716241
mem[34635] = 897860
mem[25683] = 1353
mem[29342] = 1769
mem[55576] = 400091397
mask = X10101011X01XX101XX111011XX011110001
mem[51710] = 821128099
mem[33947] = 237866140
mem[23402] = 1725292
mem[51603] = 3384
mem[610] = 21940
mask = 1111001X01101110010X100101X111X0X000
mem[62557] = 7498
mem[1254] = 599
mem[19803] = 2675
mask = 110111X11X11111011101X10X1X000X1X0XX
mem[2762] = 258673
mem[60042] = 20606346
mem[3212] = 13178507
mem[8887] = 412915
mem[18462] = 619
mem[10311] = 7584
mem[63816] = 24680261
mask = 1111011X01101X1XX100000000000X01X000
mem[29821] = 883
mem[57154] = 5765093
mem[35745] = 73
mask = 101XXX1101011100010010010001010XX111
mem[4607] = 8269571
mem[14234] = 6849
mem[64727] = 3522
mask = 0X011XX1X01111X0010101X011000110X011
mem[50853] = 6689303
mem[26913] = 129493
mem[8439] = 8784
mem[42377] = 5663270
mem[64796] = 175
mask = 111X001101X01X1011X01111110110110110
mem[7758] = 28452
mem[34374] = 559
mem[3328] = 5312
mem[46467] = 235850961
mask = 111101110X1011X0110000010X0XX000100X
mem[17500] = 6530
mem[16930] = 11847064
mem[7140] = 188556307
mem[6232] = 102869047
mem[39269] = 51308
mem[28924] = 4828
mask = 11X10111001011101100X0100000011XX110
mem[32674] = 684
mem[12892] = 156288
mem[21192] = 46944
mem[48554] = 998041
mem[18453] = 136061
mem[22294] = 39009055
mem[12802] = 63780
mask = 11110XX1011011101X0X0000001000100X00
mem[61692] = 1273
mem[56716] = 12434142
mem[31684] = 13133132
mem[16545] = 2489415
mem[39483] = 1048235210
mask = 01X1X1X1101X11X011X1010001X101110010
mem[53092] = 3823
mem[47384] = 4573
mem[50275] = 284570347
mem[1129] = 28366135
mask = 1X110101X01110111X010X0000000X0XX001
mem[22024] = 945785
mem[55321] = 56363
mem[28412] = 3465`
