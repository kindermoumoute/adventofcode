package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
		`480`,
		``,
	},
}

var puzzle = `Button A: X+11, Y+73
Button B: X+65, Y+17
Prize: X=18133, Y=4639

Button A: X+49, Y+13
Button B: X+24, Y+79
Prize: X=6664, Y=948

Button A: X+37, Y+28
Button B: X+13, Y+90
Prize: X=2448, Y=2574

Button A: X+13, Y+91
Button B: X+77, Y+70
Prize: X=7241, Y=14105

Button A: X+14, Y+34
Button B: X+50, Y+29
Prize: X=8984, Y=5882

Button A: X+53, Y+96
Button B: X+47, Y+22
Prize: X=2753, Y=3282

Button A: X+80, Y+38
Button B: X+11, Y+56
Prize: X=3726, Y=18768

Button A: X+13, Y+47
Button B: X+37, Y+11
Prize: X=17682, Y=7654

Button A: X+28, Y+80
Button B: X+66, Y+17
Prize: X=18824, Y=1350

Button A: X+65, Y+95
Button B: X+88, Y+31
Prize: X=10004, Y=5543

Button A: X+69, Y+17
Button B: X+45, Y+94
Prize: X=9291, Y=8093

Button A: X+44, Y+21
Button B: X+17, Y+36
Prize: X=15790, Y=9971

Button A: X+41, Y+59
Button B: X+32, Y+15
Prize: X=12751, Y=8224

Button A: X+13, Y+17
Button B: X+94, Y+11
Prize: X=9937, Y=2138

Button A: X+73, Y+47
Button B: X+19, Y+41
Prize: X=12528, Y=1492

Button A: X+71, Y+14
Button B: X+14, Y+49
Prize: X=18670, Y=8677

Button A: X+13, Y+60
Button B: X+37, Y+21
Prize: X=3419, Y=2151

Button A: X+76, Y+57
Button B: X+14, Y+71
Prize: X=2138, Y=4931

Button A: X+29, Y+63
Button B: X+83, Y+16
Prize: X=8054, Y=2873

Button A: X+77, Y+87
Button B: X+67, Y+12
Prize: X=3828, Y=2223

Button A: X+59, Y+21
Button B: X+17, Y+34
Prize: X=15622, Y=4584

Button A: X+13, Y+42
Button B: X+65, Y+30
Prize: X=10533, Y=15782

Button A: X+34, Y+25
Button B: X+13, Y+73
Prize: X=1411, Y=4273

Button A: X+15, Y+44
Button B: X+40, Y+19
Prize: X=4535, Y=9716

Button A: X+54, Y+13
Button B: X+19, Y+72
Prize: X=3916, Y=15249

Button A: X+12, Y+46
Button B: X+73, Y+42
Prize: X=7764, Y=9918

Button A: X+38, Y+78
Button B: X+94, Y+38
Prize: X=3890, Y=4266

Button A: X+72, Y+24
Button B: X+14, Y+99
Prize: X=6092, Y=8634

Button A: X+62, Y+28
Button B: X+23, Y+40
Prize: X=2207, Y=3040

Button A: X+48, Y+12
Button B: X+37, Y+76
Prize: X=5981, Y=5180

Button A: X+19, Y+40
Button B: X+54, Y+30
Prize: X=11652, Y=690

Button A: X+44, Y+21
Button B: X+17, Y+46
Prize: X=4212, Y=10084

Button A: X+25, Y+71
Button B: X+48, Y+11
Prize: X=15978, Y=16005

Button A: X+15, Y+16
Button B: X+80, Y+17
Prize: X=5425, Y=1345

Button A: X+33, Y+60
Button B: X+41, Y+24
Prize: X=15377, Y=12548

Button A: X+70, Y+43
Button B: X+17, Y+47
Prize: X=14123, Y=8273

Button A: X+13, Y+37
Button B: X+26, Y+12
Prize: X=4527, Y=4989

Button A: X+55, Y+15
Button B: X+37, Y+70
Prize: X=9232, Y=15265

Button A: X+76, Y+54
Button B: X+22, Y+74
Prize: X=962, Y=2026

Button A: X+99, Y+79
Button B: X+11, Y+40
Prize: X=8360, Y=8201

Button A: X+34, Y+13
Button B: X+13, Y+44
Prize: X=4601, Y=9921

Button A: X+12, Y+24
Button B: X+57, Y+35
Prize: X=16634, Y=17390

Button A: X+46, Y+28
Button B: X+20, Y+44
Prize: X=996, Y=19284

Button A: X+79, Y+32
Button B: X+12, Y+54
Prize: X=635, Y=3410

Button A: X+63, Y+27
Button B: X+12, Y+45
Prize: X=473, Y=19763

Button A: X+64, Y+20
Button B: X+12, Y+42
Prize: X=9336, Y=5308

Button A: X+62, Y+32
Button B: X+22, Y+43
Prize: X=374, Y=18365

Button A: X+56, Y+75
Button B: X+63, Y+15
Prize: X=6573, Y=2490

Button A: X+23, Y+43
Button B: X+49, Y+22
Prize: X=18315, Y=3100

Button A: X+29, Y+44
Button B: X+37, Y+13
Prize: X=5059, Y=8953

Button A: X+19, Y+25
Button B: X+53, Y+22
Prize: X=5828, Y=3754

Button A: X+11, Y+47
Button B: X+87, Y+48
Prize: X=18116, Y=10748

Button A: X+59, Y+37
Button B: X+14, Y+83
Prize: X=3848, Y=4046

Button A: X+78, Y+40
Button B: X+49, Y+83
Prize: X=4919, Y=5937

Button A: X+23, Y+90
Button B: X+83, Y+49
Prize: X=8121, Y=9991

Button A: X+13, Y+75
Button B: X+87, Y+67
Prize: X=6647, Y=10513

Button A: X+30, Y+63
Button B: X+65, Y+30
Prize: X=7695, Y=8654

Button A: X+25, Y+46
Button B: X+51, Y+28
Prize: X=13547, Y=19260

Button A: X+47, Y+70
Button B: X+72, Y+29
Prize: X=9582, Y=7230

Button A: X+14, Y+29
Button B: X+53, Y+21
Prize: X=234, Y=18295

Button A: X+46, Y+16
Button B: X+19, Y+56
Prize: X=17380, Y=19424

Button A: X+15, Y+97
Button B: X+44, Y+22
Prize: X=3113, Y=10417

Button A: X+20, Y+61
Button B: X+71, Y+14
Prize: X=2487, Y=497

Button A: X+22, Y+60
Button B: X+37, Y+16
Prize: X=14487, Y=15396

Button A: X+68, Y+25
Button B: X+31, Y+94
Prize: X=3036, Y=7394

Button A: X+14, Y+42
Button B: X+76, Y+32
Prize: X=6348, Y=4644

Button A: X+38, Y+15
Button B: X+42, Y+77
Prize: X=16798, Y=16487

Button A: X+57, Y+37
Button B: X+15, Y+44
Prize: X=2681, Y=18959

Button A: X+18, Y+55
Button B: X+63, Y+33
Prize: X=3914, Y=6084

Button A: X+86, Y+21
Button B: X+59, Y+66
Prize: X=7005, Y=5580

Button A: X+89, Y+32
Button B: X+18, Y+74
Prize: X=9279, Y=5092

Button A: X+17, Y+30
Button B: X+59, Y+28
Prize: X=1009, Y=3092

Button A: X+79, Y+60
Button B: X+25, Y+73
Prize: X=8056, Y=8333

Button A: X+20, Y+43
Button B: X+30, Y+12
Prize: X=15470, Y=8933

Button A: X+93, Y+18
Button B: X+62, Y+66
Prize: X=12059, Y=5628

Button A: X+32, Y+12
Button B: X+21, Y+51
Prize: X=9997, Y=12767

Button A: X+35, Y+62
Button B: X+50, Y+22
Prize: X=8745, Y=9312

Button A: X+17, Y+65
Button B: X+91, Y+59
Prize: X=6219, Y=6731

Button A: X+37, Y+72
Button B: X+80, Y+22
Prize: X=6272, Y=7660

Button A: X+63, Y+40
Button B: X+39, Y+86
Prize: X=6918, Y=8924

Button A: X+11, Y+70
Button B: X+78, Y+23
Prize: X=6453, Y=16364

Button A: X+78, Y+11
Button B: X+19, Y+85
Prize: X=4506, Y=10108

Button A: X+43, Y+13
Button B: X+36, Y+56
Prize: X=1947, Y=8857

Button A: X+43, Y+25
Button B: X+27, Y+53
Prize: X=1393, Y=5899

Button A: X+66, Y+23
Button B: X+15, Y+37
Prize: X=17246, Y=2071

Button A: X+37, Y+38
Button B: X+92, Y+17
Prize: X=7758, Y=3241

Button A: X+29, Y+60
Button B: X+58, Y+16
Prize: X=14358, Y=11880

Button A: X+20, Y+47
Button B: X+39, Y+14
Prize: X=19545, Y=8215

Button A: X+19, Y+91
Button B: X+55, Y+11
Prize: X=3826, Y=2422

Button A: X+22, Y+64
Button B: X+55, Y+14
Prize: X=4610, Y=16272

Button A: X+15, Y+47
Button B: X+74, Y+42
Prize: X=902, Y=5446

Button A: X+29, Y+85
Button B: X+24, Y+17
Prize: X=3585, Y=6987

Button A: X+64, Y+29
Button B: X+22, Y+55
Prize: X=12848, Y=14552

Button A: X+96, Y+11
Button B: X+19, Y+20
Prize: X=8917, Y=2002

Button A: X+38, Y+15
Button B: X+11, Y+24
Prize: X=19628, Y=6146

Button A: X+62, Y+13
Button B: X+11, Y+65
Prize: X=16243, Y=6217

Button A: X+43, Y+69
Button B: X+47, Y+23
Prize: X=14766, Y=11968

Button A: X+83, Y+17
Button B: X+42, Y+72
Prize: X=7761, Y=6915

Button A: X+21, Y+56
Button B: X+41, Y+11
Prize: X=4697, Y=13187

Button A: X+12, Y+45
Button B: X+52, Y+16
Prize: X=9272, Y=617

Button A: X+98, Y+21
Button B: X+41, Y+85
Prize: X=8226, Y=6488

Button A: X+15, Y+53
Button B: X+71, Y+38
Prize: X=12980, Y=9260

Button A: X+22, Y+50
Button B: X+67, Y+39
Prize: X=5029, Y=19449

Button A: X+86, Y+56
Button B: X+36, Y+96
Prize: X=4298, Y=8168

Button A: X+12, Y+40
Button B: X+32, Y+21
Prize: X=4732, Y=15519

Button A: X+42, Y+18
Button B: X+24, Y+52
Prize: X=9584, Y=9244

Button A: X+35, Y+71
Button B: X+99, Y+33
Prize: X=6574, Y=7294

Button A: X+11, Y+22
Button B: X+49, Y+11
Prize: X=14913, Y=6304

Button A: X+28, Y+17
Button B: X+31, Y+62
Prize: X=4016, Y=6238

Button A: X+18, Y+68
Button B: X+71, Y+25
Prize: X=4567, Y=7281

Button A: X+14, Y+53
Button B: X+43, Y+11
Prize: X=496, Y=967

Button A: X+11, Y+27
Button B: X+42, Y+28
Prize: X=581, Y=17365

Button A: X+11, Y+61
Button B: X+74, Y+12
Prize: X=12762, Y=6204

Button A: X+50, Y+71
Button B: X+39, Y+17
Prize: X=7311, Y=19292

Button A: X+12, Y+37
Button B: X+69, Y+17
Prize: X=18530, Y=11291

Button A: X+15, Y+84
Button B: X+82, Y+81
Prize: X=6619, Y=11727

Button A: X+30, Y+52
Button B: X+50, Y+15
Prize: X=16680, Y=2062

Button A: X+12, Y+37
Button B: X+55, Y+29
Prize: X=3189, Y=2663

Button A: X+31, Y+92
Button B: X+44, Y+26
Prize: X=3844, Y=4924

Button A: X+72, Y+42
Button B: X+13, Y+33
Prize: X=15389, Y=5609

Button A: X+58, Y+25
Button B: X+40, Y+73
Prize: X=1290, Y=7329

Button A: X+97, Y+18
Button B: X+21, Y+63
Prize: X=8836, Y=4122

Button A: X+61, Y+30
Button B: X+15, Y+58
Prize: X=12291, Y=15042

Button A: X+37, Y+62
Button B: X+34, Y+11
Prize: X=6489, Y=17107

Button A: X+72, Y+22
Button B: X+11, Y+39
Prize: X=7940, Y=18284

Button A: X+56, Y+17
Button B: X+11, Y+42
Prize: X=1548, Y=18821

Button A: X+49, Y+32
Button B: X+11, Y+71
Prize: X=2070, Y=5117

Button A: X+90, Y+14
Button B: X+34, Y+72
Prize: X=3338, Y=5656

Button A: X+98, Y+41
Button B: X+69, Y+95
Prize: X=4630, Y=5905

Button A: X+35, Y+11
Button B: X+68, Y+78
Prize: X=9151, Y=8369

Button A: X+93, Y+66
Button B: X+18, Y+42
Prize: X=5181, Y=4992

Button A: X+18, Y+34
Button B: X+47, Y+14
Prize: X=4297, Y=14264

Button A: X+52, Y+18
Button B: X+20, Y+75
Prize: X=4476, Y=6519

Button A: X+89, Y+27
Button B: X+61, Y+86
Prize: X=12618, Y=9025

Button A: X+65, Y+20
Button B: X+44, Y+55
Prize: X=10135, Y=6850

Button A: X+89, Y+23
Button B: X+19, Y+88
Prize: X=1541, Y=4802

Button A: X+60, Y+30
Button B: X+28, Y+91
Prize: X=5960, Y=7985

Button A: X+54, Y+11
Button B: X+79, Y+94
Prize: X=8465, Y=5386

Button A: X+41, Y+75
Button B: X+75, Y+37
Prize: X=5261, Y=5115

Button A: X+49, Y+18
Button B: X+27, Y+61
Prize: X=7642, Y=19052

Button A: X+72, Y+30
Button B: X+15, Y+74
Prize: X=6009, Y=2978

Button A: X+56, Y+38
Button B: X+11, Y+27
Prize: X=19769, Y=5953

Button A: X+85, Y+33
Button B: X+35, Y+57
Prize: X=6515, Y=3441

Button A: X+31, Y+12
Button B: X+20, Y+36
Prize: X=6966, Y=6428

Button A: X+33, Y+64
Button B: X+81, Y+39
Prize: X=6198, Y=7651

Button A: X+88, Y+76
Button B: X+26, Y+76
Prize: X=7544, Y=9728

Button A: X+43, Y+20
Button B: X+19, Y+49
Prize: X=4446, Y=10629

Button A: X+76, Y+17
Button B: X+83, Y+93
Prize: X=9974, Y=5655

Button A: X+13, Y+48
Button B: X+80, Y+78
Prize: X=1000, Y=2388

Button A: X+89, Y+30
Button B: X+47, Y+80
Prize: X=3091, Y=2710

Button A: X+11, Y+36
Button B: X+74, Y+33
Prize: X=11533, Y=18533

Button A: X+97, Y+20
Button B: X+66, Y+62
Prize: X=5489, Y=4374

Button A: X+11, Y+78
Button B: X+79, Y+18
Prize: X=5011, Y=12902

Button A: X+16, Y+47
Button B: X+28, Y+11
Prize: X=4780, Y=3810

Button A: X+25, Y+15
Button B: X+51, Y+92
Prize: X=4136, Y=6227

Button A: X+99, Y+32
Button B: X+49, Y+80
Prize: X=7813, Y=3552

Button A: X+12, Y+59
Button B: X+68, Y+24
Prize: X=11240, Y=12151

Button A: X+46, Y+74
Button B: X+34, Y+15
Prize: X=1736, Y=1840

Button A: X+31, Y+57
Button B: X+47, Y+16
Prize: X=2670, Y=4167

Button A: X+31, Y+17
Button B: X+19, Y+35
Prize: X=13499, Y=17561

Button A: X+12, Y+56
Button B: X+51, Y+24
Prize: X=16196, Y=16288

Button A: X+90, Y+77
Button B: X+28, Y+98
Prize: X=8116, Y=7462

Button A: X+78, Y+30
Button B: X+12, Y+48
Prize: X=14534, Y=10310

Button A: X+19, Y+35
Button B: X+57, Y+23
Prize: X=6029, Y=3951

Button A: X+79, Y+22
Button B: X+24, Y+77
Prize: X=5110, Y=2970

Button A: X+85, Y+54
Button B: X+14, Y+92
Prize: X=4443, Y=5482

Button A: X+48, Y+93
Button B: X+94, Y+34
Prize: X=6446, Y=6416

Button A: X+25, Y+45
Button B: X+41, Y+11
Prize: X=1549, Y=14899

Button A: X+32, Y+75
Button B: X+63, Y+17
Prize: X=17499, Y=17482

Button A: X+64, Y+79
Button B: X+58, Y+16
Prize: X=4980, Y=2478

Button A: X+41, Y+49
Button B: X+41, Y+11
Prize: X=3567, Y=2439

Button A: X+23, Y+79
Button B: X+94, Y+60
Prize: X=5420, Y=7576

Button A: X+11, Y+65
Button B: X+76, Y+11
Prize: X=2316, Y=1194

Button A: X+15, Y+70
Button B: X+40, Y+12
Prize: X=7155, Y=3546

Button A: X+13, Y+34
Button B: X+67, Y+16
Prize: X=1274, Y=18032

Button A: X+55, Y+18
Button B: X+70, Y+82
Prize: X=5340, Y=5234

Button A: X+30, Y+25
Button B: X+22, Y+79
Prize: X=2232, Y=6774

Button A: X+15, Y+27
Button B: X+41, Y+24
Prize: X=4139, Y=19259

Button A: X+34, Y+13
Button B: X+32, Y+54
Prize: X=6670, Y=7075

Button A: X+26, Y+70
Button B: X+58, Y+12
Prize: X=12464, Y=8436

Button A: X+26, Y+12
Button B: X+30, Y+90
Prize: X=3242, Y=3324

Button A: X+12, Y+49
Button B: X+50, Y+25
Prize: X=18232, Y=19589

Button A: X+32, Y+93
Button B: X+96, Y+38
Prize: X=7360, Y=8135

Button A: X+27, Y+83
Button B: X+51, Y+28
Prize: X=3813, Y=5025

Button A: X+14, Y+36
Button B: X+23, Y+13
Prize: X=6999, Y=6683

Button A: X+23, Y+49
Button B: X+48, Y+19
Prize: X=7062, Y=4966

Button A: X+11, Y+20
Button B: X+43, Y+18
Prize: X=15413, Y=4578

Button A: X+22, Y+45
Button B: X+78, Y+40
Prize: X=5424, Y=5715

Button A: X+58, Y+14
Button B: X+21, Y+62
Prize: X=3942, Y=9636

Button A: X+14, Y+49
Button B: X+67, Y+32
Prize: X=13568, Y=19028

Button A: X+25, Y+53
Button B: X+48, Y+17
Prize: X=2534, Y=4648

Button A: X+75, Y+31
Button B: X+15, Y+43
Prize: X=12230, Y=7334

Button A: X+56, Y+18
Button B: X+34, Y+75
Prize: X=2428, Y=4916

Button A: X+33, Y+12
Button B: X+40, Y+62
Prize: X=5860, Y=6734

Button A: X+32, Y+19
Button B: X+18, Y+41
Prize: X=5468, Y=4256

Button A: X+19, Y+48
Button B: X+57, Y+11
Prize: X=3958, Y=6769

Button A: X+62, Y+33
Button B: X+35, Y+87
Prize: X=7144, Y=7221

Button A: X+22, Y+46
Button B: X+42, Y+13
Prize: X=12032, Y=7880

Button A: X+12, Y+98
Button B: X+82, Y+53
Prize: X=8934, Y=11911

Button A: X+99, Y+14
Button B: X+53, Y+43
Prize: X=7369, Y=3989

Button A: X+43, Y+11
Button B: X+29, Y+66
Prize: X=15386, Y=11441

Button A: X+16, Y+84
Button B: X+80, Y+70
Prize: X=4064, Y=6286

Button A: X+32, Y+70
Button B: X+92, Y+62
Prize: X=6700, Y=6162

Button A: X+81, Y+16
Button B: X+14, Y+76
Prize: X=18450, Y=13756

Button A: X+12, Y+45
Button B: X+99, Y+44
Prize: X=8805, Y=4548

Button A: X+15, Y+55
Button B: X+48, Y+13
Prize: X=2630, Y=530

Button A: X+46, Y+14
Button B: X+34, Y+79
Prize: X=7490, Y=19570

Button A: X+82, Y+20
Button B: X+28, Y+50
Prize: X=6026, Y=4060

Button A: X+25, Y+61
Button B: X+57, Y+14
Prize: X=7179, Y=12356

Button A: X+17, Y+55
Button B: X+77, Y+64
Prize: X=6710, Y=7825

Button A: X+18, Y+82
Button B: X+76, Y+15
Prize: X=7552, Y=2275

Button A: X+22, Y+76
Button B: X+99, Y+28
Prize: X=4587, Y=2344

Button A: X+23, Y+77
Button B: X+70, Y+15
Prize: X=7601, Y=7519

Button A: X+97, Y+49
Button B: X+16, Y+44
Prize: X=9439, Y=5271

Button A: X+17, Y+75
Button B: X+89, Y+47
Prize: X=7269, Y=4763

Button A: X+61, Y+23
Button B: X+20, Y+66
Prize: X=4990, Y=16918

Button A: X+37, Y+24
Button B: X+30, Y+71
Prize: X=4807, Y=5592

Button A: X+69, Y+81
Button B: X+97, Y+19
Prize: X=9123, Y=3879

Button A: X+87, Y+51
Button B: X+28, Y+82
Prize: X=8193, Y=9525

Button A: X+55, Y+26
Button B: X+44, Y+70
Prize: X=1441, Y=878

Button A: X+32, Y+81
Button B: X+77, Y+35
Prize: X=8587, Y=7824

Button A: X+60, Y+23
Button B: X+25, Y+59
Prize: X=5180, Y=15432

Button A: X+20, Y+44
Button B: X+82, Y+50
Prize: X=5328, Y=4680

Button A: X+99, Y+31
Button B: X+12, Y+35
Prize: X=8991, Y=4315

Button A: X+87, Y+84
Button B: X+21, Y+79
Prize: X=10389, Y=15727

Button A: X+21, Y+56
Button B: X+62, Y+33
Prize: X=4460, Y=14523

Button A: X+13, Y+47
Button B: X+61, Y+31
Prize: X=1748, Y=2908

Button A: X+38, Y+54
Button B: X+40, Y+15
Prize: X=12932, Y=12461

Button A: X+80, Y+81
Button B: X+53, Y+14
Prize: X=5085, Y=4157

Button A: X+59, Y+20
Button B: X+23, Y+66
Prize: X=19571, Y=12020

Button A: X+95, Y+29
Button B: X+43, Y+61
Prize: X=7584, Y=6528

Button A: X+58, Y+34
Button B: X+17, Y+32
Prize: X=19342, Y=5344

Button A: X+33, Y+69
Button B: X+93, Y+19
Prize: X=8796, Y=7338

Button A: X+11, Y+91
Button B: X+88, Y+73
Prize: X=8162, Y=9227

Button A: X+46, Y+17
Button B: X+13, Y+21
Prize: X=7121, Y=5232

Button A: X+94, Y+15
Button B: X+70, Y+84
Prize: X=6620, Y=5499

Button A: X+90, Y+12
Button B: X+32, Y+57
Prize: X=6794, Y=1275

Button A: X+77, Y+24
Button B: X+12, Y+49
Prize: X=1880, Y=13220

Button A: X+31, Y+16
Button B: X+36, Y+61
Prize: X=17747, Y=9472

Button A: X+39, Y+22
Button B: X+20, Y+50
Prize: X=12004, Y=6812

Button A: X+45, Y+15
Button B: X+23, Y+50
Prize: X=19415, Y=7715

Button A: X+61, Y+20
Button B: X+14, Y+46
Prize: X=5750, Y=11570

Button A: X+58, Y+21
Button B: X+31, Y+61
Prize: X=1375, Y=14929

Button A: X+41, Y+17
Button B: X+26, Y+50
Prize: X=6617, Y=19793

Button A: X+22, Y+47
Button B: X+60, Y+29
Prize: X=1578, Y=17740

Button A: X+98, Y+39
Button B: X+43, Y+74
Prize: X=4399, Y=5107

Button A: X+39, Y+13
Button B: X+11, Y+28
Prize: X=7736, Y=12927

Button A: X+66, Y+26
Button B: X+17, Y+41
Prize: X=4827, Y=13747

Button A: X+30, Y+70
Button B: X+38, Y+12
Prize: X=6488, Y=4962

Button A: X+60, Y+14
Button B: X+29, Y+75
Prize: X=9256, Y=6404

Button A: X+31, Y+71
Button B: X+67, Y+26
Prize: X=5436, Y=2683

Button A: X+21, Y+40
Button B: X+51, Y+23
Prize: X=1124, Y=16811

Button A: X+50, Y+22
Button B: X+14, Y+62
Prize: X=16126, Y=11714

Button A: X+61, Y+19
Button B: X+14, Y+49
Prize: X=6079, Y=11392

Button A: X+44, Y+13
Button B: X+19, Y+61
Prize: X=4014, Y=1744

Button A: X+19, Y+53
Button B: X+78, Y+42
Prize: X=7629, Y=3691

Button A: X+61, Y+87
Button B: X+71, Y+28
Prize: X=6881, Y=4539

Button A: X+28, Y+18
Button B: X+19, Y+37
Prize: X=7803, Y=4541

Button A: X+72, Y+17
Button B: X+13, Y+40
Prize: X=4995, Y=8730

Button A: X+57, Y+25
Button B: X+13, Y+30
Prize: X=17779, Y=5175

Button A: X+33, Y+94
Button B: X+66, Y+19
Prize: X=7854, Y=8852

Button A: X+41, Y+13
Button B: X+23, Y+40
Prize: X=1189, Y=10311

Button A: X+65, Y+12
Button B: X+11, Y+23
Prize: X=6352, Y=1634

Button A: X+38, Y+14
Button B: X+35, Y+64
Prize: X=3141, Y=14160

Button A: X+36, Y+57
Button B: X+43, Y+21
Prize: X=15949, Y=16718

Button A: X+59, Y+28
Button B: X+46, Y+80
Prize: X=2531, Y=1492

Button A: X+49, Y+11
Button B: X+12, Y+69
Prize: X=2917, Y=15210

Button A: X+79, Y+19
Button B: X+69, Y+97
Prize: X=6361, Y=6917

Button A: X+99, Y+36
Button B: X+31, Y+77
Prize: X=12300, Y=10191

Button A: X+72, Y+32
Button B: X+41, Y+71
Prize: X=5370, Y=5870

Button A: X+35, Y+76
Button B: X+58, Y+18
Prize: X=8521, Y=13664

Button A: X+72, Y+14
Button B: X+11, Y+65
Prize: X=16753, Y=18845

Button A: X+31, Y+63
Button B: X+25, Y+11
Prize: X=16191, Y=11049

Button A: X+78, Y+19
Button B: X+29, Y+93
Prize: X=3771, Y=1692

Button A: X+12, Y+70
Button B: X+62, Y+13
Prize: X=15090, Y=8127

Button A: X+60, Y+87
Button B: X+82, Y+34
Prize: X=7192, Y=6523

Button A: X+27, Y+65
Button B: X+50, Y+23
Prize: X=2386, Y=5530

Button A: X+38, Y+80
Button B: X+49, Y+14
Prize: X=4040, Y=12496

Button A: X+89, Y+20
Button B: X+16, Y+29
Prize: X=4183, Y=3201

Button A: X+31, Y+60
Button B: X+40, Y+13
Prize: X=8798, Y=19419

Button A: X+65, Y+32
Button B: X+18, Y+54
Prize: X=9936, Y=5532

Button A: X+22, Y+77
Button B: X+92, Y+44
Prize: X=6512, Y=4444

Button A: X+36, Y+15
Button B: X+48, Y+78
Prize: X=4640, Y=2996

Button A: X+26, Y+13
Button B: X+52, Y+99
Prize: X=2990, Y=2809

Button A: X+78, Y+27
Button B: X+30, Y+97
Prize: X=4590, Y=3581

Button A: X+33, Y+62
Button B: X+44, Y+20
Prize: X=19053, Y=7610

Button A: X+26, Y+50
Button B: X+48, Y+18
Prize: X=6822, Y=1488

Button A: X+88, Y+44
Button B: X+24, Y+77
Prize: X=7896, Y=6743

Button A: X+68, Y+13
Button B: X+23, Y+67
Prize: X=7264, Y=12390

Button A: X+84, Y+21
Button B: X+46, Y+47
Prize: X=10484, Y=5461

Button A: X+90, Y+60
Button B: X+26, Y+88
Prize: X=5278, Y=8324

Button A: X+11, Y+68
Button B: X+54, Y+12
Prize: X=16759, Y=11092

Button A: X+75, Y+37
Button B: X+25, Y+91
Prize: X=6550, Y=10626

Button A: X+31, Y+11
Button B: X+47, Y+79
Prize: X=14055, Y=13771

Button A: X+86, Y+32
Button B: X+48, Y+90
Prize: X=8214, Y=4932

Button A: X+32, Y+18
Button B: X+22, Y+55
Prize: X=232, Y=1508

Button A: X+69, Y+35
Button B: X+18, Y+54
Prize: X=4076, Y=9380

Button A: X+26, Y+11
Button B: X+15, Y+24
Prize: X=4055, Y=4577

Button A: X+26, Y+95
Button B: X+88, Y+55
Prize: X=8032, Y=7225

Button A: X+55, Y+37
Button B: X+11, Y+32
Prize: X=13608, Y=4473

Button A: X+16, Y+57
Button B: X+96, Y+40
Prize: X=8816, Y=7851

Button A: X+54, Y+28
Button B: X+28, Y+50
Prize: X=1332, Y=9554

Button A: X+88, Y+77
Button B: X+23, Y+81
Prize: X=1588, Y=1633

Button A: X+78, Y+51
Button B: X+17, Y+44
Prize: X=5052, Y=11829

Button A: X+16, Y+61
Button B: X+77, Y+45
Prize: X=4172, Y=4969

Button A: X+12, Y+46
Button B: X+66, Y+38
Prize: X=6554, Y=5652

Button A: X+26, Y+60
Button B: X+42, Y+18
Prize: X=9650, Y=8558

Button A: X+30, Y+64
Button B: X+92, Y+41
Prize: X=11254, Y=9724

Button A: X+25, Y+53
Button B: X+40, Y+17
Prize: X=12455, Y=17318

Button A: X+64, Y+37
Button B: X+15, Y+41
Prize: X=3855, Y=6003

Button A: X+12, Y+42
Button B: X+75, Y+21
Prize: X=13496, Y=4988

Button A: X+85, Y+29
Button B: X+33, Y+54
Prize: X=5598, Y=3876

Button A: X+86, Y+30
Button B: X+11, Y+65
Prize: X=15180, Y=11840

Button A: X+28, Y+63
Button B: X+58, Y+11
Prize: X=2432, Y=16018

Button A: X+28, Y+49
Button B: X+25, Y+14
Prize: X=11147, Y=18589

Button A: X+64, Y+25
Button B: X+29, Y+71
Prize: X=4242, Y=12936

Button A: X+14, Y+41
Button B: X+59, Y+24
Prize: X=7187, Y=3683

Button A: X+29, Y+11
Button B: X+39, Y+72
Prize: X=5051, Y=8159

Button A: X+26, Y+46
Button B: X+41, Y+21
Prize: X=1150, Y=2410

Button A: X+36, Y+77
Button B: X+95, Y+34
Prize: X=4955, Y=4338`
