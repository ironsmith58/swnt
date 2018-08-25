package name

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	//"github.com/jmcvetta/randutil"
	"github.com/nboughton/rollt"
)

var badPrefix = regexp.MustCompile(`[cflmnr][^aeiouyh]`)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name generator. This is a primitive first effort that will be refined over time

// Generate creates a random name by combining alternating vowels and consonants
func Generate(ln int) string {
	name := ""
	for i := rand.Intn(2); len(name) <= ln; i++ {
		if i%2 != 0 {
			//c, _ := randutil.WeightedChoice(cnWeighted)
			c := con.Roll()

			if name == "" {
				//for badPrefix.MatchString(c.Item.(string)) { // I don't like starting a name with these
				for badPrefix.MatchString(c) { // I don't like starting a name with these
					//c, _ = randutil.WeightedChoice(cnWeighted)
					c = con.Roll()
				}
			}

			//name += c.Item.(string)
			name += c
		} else {
			//c, _ := randutil.WeightedChoice(vlWeighted)
			//name += c.Item.(string)
			name += vl.Roll()
		}
	}

	return strings.ToUpper(string(name[0])) + string(name[1:])
}

var vl = rollt.Table{
	Name: "vowels",
	Dice: "10d5",
	Items: []rollt.Item{
		{Match: []int{10}, Text: "ii"},
		{Match: []int{11}, Text: "yu"},
		{Match: []int{12}, Text: "uy"},
		{Match: []int{13}, Text: "oy"},
		{Match: []int{14}, Text: "ao"},
		{Match: []int{15}, Text: "ye"},
		{Match: []int{16}, Text: "ae"},
		{Match: []int{17}, Text: "oe"},
		{Match: []int{18}, Text: "eo"},
		{Match: []int{19}, Text: "oi"},
		{Match: []int{20}, Text: "ua"},
		{Match: []int{21}, Text: "au"},
		{Match: []int{22}, Text: "ia"},
		{Match: []int{23}, Text: "ey"},
		{Match: []int{24}, Text: "oo"},
		{Match: []int{25}, Text: "io"},
		{Match: []int{26}, Text: "ea"},
		{Match: []int{27}, Text: "y"},
		{Match: []int{28}, Text: "o"},
		{Match: []int{29}, Text: "a"},
		{Match: []int{30}, Text: "e"},
		{Match: []int{31}, Text: "i"},
		{Match: []int{32}, Text: "u"},
		{Match: []int{33}, Text: "ou"},
		{Match: []int{34}, Text: "ee"},
		{Match: []int{35}, Text: "ai"},
		{Match: []int{36}, Text: "ie"},
		{Match: []int{37}, Text: "ei"},
		{Match: []int{38}, Text: "ue"},
		{Match: []int{39}, Text: "ay"},
		{Match: []int{40}, Text: "ui"},
		{Match: []int{41}, Text: "oa"},
		{Match: []int{42}, Text: "yi"},
		{Match: []int{43}, Text: "ya"},
		{Match: []int{44}, Text: "eu"},
		{Match: []int{45}, Text: "iu"},
		{Match: []int{46}, Text: "yo"},
		{Match: []int{47}, Text: "aa"},
		{Match: []int{48}, Text: "uo"},
		{Match: []int{49}, Text: "uu"},
		{Match: []int{50}, Text: "'"},
	},
}

var con = rollt.Table{
	Name: "consonants",
	Dice: "10d5",
	Items: []rollt.Item{
		{Match: []int{10}, Text: "tt"},
		{Match: []int{11}, Text: "rr"},
		{Match: []int{12}, Text: "ct"},
		{Match: []int{13}, Text: "pr"},
		{Match: []int{14}, Text: "ns"},
		{Match: []int{15}, Text: "bl"},
		{Match: []int{16}, Text: "sh"},
		{Match: []int{17}, Text: "ld"},
		{Match: []int{18}, Text: "k"},
		{Match: []int{19}, Text: "nd"},
		{Match: []int{20}, Text: "ll"},
		{Match: []int{21}, Text: "nt"},
		{Match: []int{22}, Text: "st"},
		{Match: []int{23}, Text: "f"},
		{Match: []int{24}, Text: "ng"},
		{Match: []int{25}, Text: "w"},
		{Match: []int{26}, Text: "th"},
		{Match: []int{27}, Text: "m"},
		{Match: []int{28}, Text: "n"},
		{Match: []int{29}, Text: "d"},
		{Match: []int{30}, Text: "r"},
		{Match: []int{31}, Text: "s"},
		{Match: []int{32}, Text: "t"},
		{Match: []int{33}, Text: "l"},
		{Match: []int{34}, Text: "c"},
		{Match: []int{35}, Text: "v"},
		{Match: []int{36}, Text: "b"},
		{Match: []int{37}, Text: "p"},
		{Match: []int{38}, Text: "h"},
		{Match: []int{39}, Text: "g"},
		{Match: []int{40}, Text: "wh"},
		{Match: []int{41}, Text: "ch"},
		{Match: []int{42}, Text: "ss"},
		{Match: []int{43}, Text: "rs"},
		{Match: []int{44}, Text: "nc"},
		{Match: []int{45}, Text: "fr"},
		{Match: []int{46}, Text: "rt"},
		{Match: []int{47}, Text: "gr"},
		{Match: []int{48}, Text: "rd"},
		{Match: []int{49}, Text: "sp"},
		{Match: []int{50}, Text: "ck"},
	},
}

/*
var vlWeighted = []randutil.Choice{
	{Weight: 257955, Item: "e"},
	{Weight: 134060, Item: "a"},
	{Weight: 129280, Item: "i"},
	{Weight: 98383, Item: "o"},
	{Weight: 39972, Item: "u"},
	{Weight: 30850, Item: "y"},
	{Weight: 23883, Item: "ou"},
	{Weight: 19815, Item: "ea"},
	{Weight: 11073, Item: "ee"},
	{Weight: 8982, Item: "io"},
	{Weight: 8363, Item: "ai"},
	{Weight: 7664, Item: "oo"},
	{Weight: 7365, Item: "ie"},
	{Weight: 3994, Item: "ey"},
	{Weight: 3894, Item: "ei"},
	{Weight: 3752, Item: "ia"},
	{Weight: 2762, Item: "ue"},
	{Weight: 2389, Item: "au"},
	{Weight: 2344, Item: "ay"},
	{Weight: 2121, Item: "ua"},
	{Weight: 2082, Item: "ui"},
	{Weight: 1949, Item: "oi"},
	{Weight: 1706, Item: "oa"},
	{Weight: 977, Item: "eo"},
	{Weight: 578, Item: "yi"},
	{Weight: 466, Item: "oe"},
	{Weight: 387, Item: "ya"},
	{Weight: 369, Item: "ae"},
	{Weight: 352, Item: "eu"},
	{Weight: 311, Item: "ye"},
	{Weight: 281, Item: "iu"},
	{Weight: 242, Item: "ao"},
	{Weight: 202, Item: "yo"},
	{Weight: 190, Item: "oy"},
	{Weight: 64, Item: "aa"},
	{Weight: 63, Item: "uy"},
	{Weight: 57, Item: "uo"},
	{Weight: 42, Item: "yu"},
	{Weight: 13, Item: "uu"},
	{Weight: 11, Item: "ii"},
}

var cnWeighted = []randutil.Choice{
	{Weight: 79825, Item: "r"},
	{Weight: 68762, Item: "d"},
	{Weight: 61735, Item: "s"},
	{Weight: 60203, Item: "n"},
	{Weight: 54142, Item: "t"},
	{Weight: 50410, Item: "m"},
	{Weight: 44578, Item: "l"},
	{Weight: 44510, Item: "th"},
	{Weight: 37663, Item: "c"},
	{Weight: 31015, Item: "w"},
	{Weight: 26880, Item: "v"},
	{Weight: 26288, Item: "ng"},
	{Weight: 21656, Item: "b"},
	{Weight: 20508, Item: "f"},
	{Weight: 20292, Item: "p"},
	{Weight: 19717, Item: "st"},
	{Weight: 19650, Item: "h"},
	{Weight: 15535, Item: "nt"},
	{Weight: 14047, Item: "g"},
	{Weight: 13307, Item: "ll"},
	{Weight: 13036, Item: "wh"},
	{Weight: 12552, Item: "nd"},
	{Weight: 10593, Item: "ch"},
	{Weight: 9304, Item: "k"},
	{Weight: 9151, Item: "ss"},
	{Weight: 7452, Item: "ld"},
	{Weight: 7385, Item: "rs"},
	{Weight: 7220, Item: "sh"},
	{Weight: 7204, Item: "nc"},
	{Weight: 6946, Item: "bl"},
	{Weight: 6308, Item: "fr"},
	{Weight: 6299, Item: "ns"},
	{Weight: 5862, Item: "rt"},
	{Weight: 5065, Item: "pr"},
	{Weight: 5047, Item: "gr"},
	{Weight: 4969, Item: "ct"},
	{Weight: 4702, Item: "rd"},
	{Weight: 4454, Item: "rr"},
	{Weight: 4288, Item: "sp"},
	{Weight: 4192, Item: "tt"},
	{Weight: 3984, Item: "ck"},
	{Weight: 3844, Item: "tr"},
	{Weight: 3744, Item: "rn"},
	{Weight: 3716, Item: "gh"},
	{Weight: 3412, Item: "cl"},
	{Weight: 3281, Item: "pl"},
	{Weight: 3104, Item: "cr"},
	{Weight: 2964, Item: "wn"},
	{Weight: 2899, Item: "br"},
	{Weight: 2846, Item: "rm"},
	{Weight: 2706, Item: "ft"},
	{Weight: 2587, Item: "sl"},
	{Weight: 2518, Item: "sc"},
	{Weight: 2380, Item: "q"},
	{Weight: 2358, Item: "dr"},
	{Weight: 2315, Item: "nl"},
	{Weight: 2265, Item: "kn"},
	{Weight: 2264, Item: "lt"},
	{Weight: 2199, Item: "pp"},
	{Weight: 2189, Item: "j"},
	{Weight: 2057, Item: "fl"},
	{Weight: 1935, Item: "ts"},
	{Weight: 1810, Item: "lf"},
	{Weight: 1756, Item: "z"},
	{Weight: 1736, Item: "nn"},
	{Weight: 1717, Item: "rv"},
	{Weight: 1692, Item: "rk"},
	{Weight: 1692, Item: "pt"},
	{Weight: 1681, Item: "x"},
	{Weight: 1666, Item: "rg"},
	{Weight: 1663, Item: "ls"},
	{Weight: 1622, Item: "mp"},
	{Weight: 1615, Item: "rl"},
	{Weight: 1612, Item: "ff"},
	{Weight: 1547, Item: "gl"},
	{Weight: 1452, Item: "ph"},
	{Weight: 1425, Item: "cc"},
	{Weight: 1414, Item: "mb"},
	{Weight: 1393, Item: "ms"},
	{Weight: 1385, Item: "sm"},
	{Weight: 1380, Item: "mm"},
	{Weight: 1370, Item: "nk"},
	{Weight: 1263, Item: "dd"},
	{Weight: 1211, Item: "ps"},
	{Weight: 1207, Item: "rc"},
	{Weight: 1107, Item: "gn"},
	{Weight: 1015, Item: "lm"},
	{Weight: 962, Item: "ds"},
	{Weight: 921, Item: "nf"},
	{Weight: 914, Item: "lk"},
	{Weight: 913, Item: "rb"},
	{Weight: 880, Item: "ws"},
	{Weight: 847, Item: "nv"},
	{Weight: 780, Item: "gg"},
	{Weight: 762, Item: "sk"},
	{Weight: 716, Item: "tw"},
	{Weight: 713, Item: "xp"},
	{Weight: 712, Item: "wl"},
	{Weight: 701, Item: "xt"},
	{Weight: 677, Item: "lv"},
	{Weight: 662, Item: "wr"},
		{Weight: 661, Item: "sn"},
		{Weight: 658, Item: "dg"},
		{Weight: 634, Item: "dl"},
		{Weight: 628, Item: "sw"},
		{Weight: 625, Item: "rh"},
		{Weight: 612, Item: "sq"},
		{Weight: 594, Item: "rw"},
		{Weight: 587, Item: "rf"},
		{Weight: 578, Item: "bs"},
		{Weight: 482, Item: "xc"},
		{Weight: 470, Item: "tl"},
		{Weight: 467, Item: "dn"},
		{Weight: 444, Item: "lw"},
		{Weight: 430, Item: "ks"},
		{Weight: 422, Item: "bj"},
		{Weight: 418, Item: "mn"},
		{Weight: 416, Item: "rp"},
		{Weight: 392, Item: "bt"},
		{Weight: 345, Item: "bb"},
		{Weight: 323, Item: "gs"},
		{Weight: 317, Item: "dw"},
		{Weight: 315, Item: "dv"},
		{Weight: 311, Item: "lp"},
		{Weight: 301, Item: "nh"},
		{Weight: 296, Item: "nr"},
		{Weight: 220, Item: "lc"},
		{Weight: 208, Item: "dm"},
		{Weight: 208, Item: "gm"},
		{Weight: 203, Item: "nw"},
		{Weight: 198, Item: "nm"},
		{Weight: 185, Item: "nj"},
		{Weight: 170, Item: "ml"},
		{Weight: 170, Item: "nq"},
		{Weight: 160, Item: "bn"},
		{Weight: 154, Item: "lb"},
		{Weight: 153, Item: "cs"},
		{Weight: 142, Item: "wd"},
		{Weight: 137, Item: "xh"},
		{Weight: 134, Item: "fs"},
		{Weight: 126, Item: "bv"},
		{Weight: 114, Item: "wf"},
		{Weight: 113, Item: "nx"},
		{Weight: 113, Item: "zz"},
		{Weight: 113, Item: "nz"},
		{Weight: 111, Item: "lh"},
		{Weight: 106, Item: "lg"},
		{Weight: 104, Item: "nb"},
		{Weight: 104, Item: "tm"},
		{Weight: 94, Item: "mf"},
		{Weight: 94, Item: "ln"},
		{Weight: 92, Item: "lr"},
		{Weight: 85, Item: "cq"},
		{Weight: 84, Item: "hn"},
		{Weight: 80, Item: "tf"},
		{Weight: 79, Item: "pn"},
		{Weight: 74, Item: "kl"},
		{Weight: 74, Item: "pm"},
		{Weight: 68, Item: "tn"},
		{Weight: 65, Item: "sb"},
		{Weight: 63, Item: "sg"},
		{Weight: 63, Item: "pw"},
		{Weight: 59, Item: "wt"},
		{Weight: 59, Item: "dj"},
		{Weight: 57, Item: "bd"},
		{Weight: 55, Item: "np"},
		{Weight: 54, Item: "bh"},
		{Weight: 51, Item: "sf"},
		{Weight: 48, Item: "bw"},
		{Weight: 40, Item: "bm"},
		{Weight: 39, Item: "sd"},
		{Weight: 37, Item: "zr"},
		{Weight: 34, Item: "df"},
		{Weight: 34, Item: "kh"},
		{Weight: 30, Item: "cz"},
		{Weight: 29, Item: "bc"},
		{Weight: 28, Item: "rz"},
		{Weight: 28, Item: "wb"},
		{Weight: 28, Item: "dh"},
		{Weight: 24, Item: "xq"},
		{Weight: 23, Item: "kr"},
		{Weight: 22, Item: "pk"},
		{Weight: 21, Item: "sr"},
		{Weight: 21, Item: "wc"},
		{Weight: 20, Item: "mt"},
		{Weight: 20, Item: "wm"},
		{Weight: 20, Item: "sj"},
		{Weight: 18, Item: "tb"},
		{Weight: 16, Item: "wk"},
		{Weight: 15, Item: "tz"},
		{Weight: 15, Item: "wp"},
		{Weight: 15, Item: "td"},
		{Weight: 15, Item: "hs"},
		{Weight: 14, Item: "tp"},
		{Weight: 14, Item: "zk"},
		{Weight: 14, Item: "pb"},
		{Weight: 13, Item: "kf"},
		{Weight: 12, Item: "tc"},
		{Weight: 11, Item: "zt"},
		{Weight: 11, Item: "rq"},
		{Weight: 11, Item: "kw"},
		{Weight: 10, Item: "dq"},
		{Weight: 9, Item: "kd"},
		{Weight: 9, Item: "cn"},
		{Weight: 9, Item: "dp"},
		{Weight: 8, Item: "mw"},
		{Weight: 7, Item: "wg"},
		{Weight: 7, Item: "kk"},
		{Weight: 7, Item: "cd"},
		{Weight: 7, Item: "db"},
		{Weight: 7, Item: "mr"},
		{Weight: 5, Item: "dt"},
		{Weight: 5, Item: "vr"},
		{Weight: 5, Item: "zv"},
		{Weight: 5, Item: "fj"},
		{Weight: 5, Item: "pf"},
		{Weight: 5, Item: "km"},
		{Weight: 5, Item: "hl"},
		{Weight: 4, Item: "xf"},
		{Weight: 4, Item: "fb"},
		{Weight: 3, Item: "vs"},
		{Weight: 3, Item: "hm"},
		{Weight: 3, Item: "gt"},
		{Weight: 3, Item: "fh"},
		{Weight: 3, Item: "gz"},
		{Weight: 3, Item: "kc"},
		{Weight: 3, Item: "jk"},
		{Weight: 2, Item: "mh"},
		{Weight: 2, Item: "mq"},
		{Weight: 2, Item: "dc"},
		{Weight: 2, Item: "md"},
		{Weight: 2, Item: "fk"},
		{Weight: 2, Item: "dz"},
		{Weight: 2, Item: "hr"},
		{Weight: 2, Item: "sv"},
		{Weight: 2, Item: "mc"},
		{Weight: 2, Item: "tg"},
		{Weight: 2, Item: "zb"},
		{Weight: 2, Item: "gw"},
		{Weight: 2, Item: "lz"},
		{Weight: 1, Item: "xl"},
		{Weight: 1, Item: "hd"},
		{Weight: 1, Item: "vl"},
		{Weight: 1, Item: "xd"},
		{Weight: 1, Item: "ht"},
		{Weight: 1, Item: "kt"},
		{Weight: 1, Item: "rx"},
		{Weight: 1, Item: "pz"},
		{Weight: 1, Item: "fn"},
		{Weight: 1, Item: "tk"},
		{Weight: 1, Item: "hf"},
		{Weight: 1, Item: "pj"},
		{Weight: 1, Item: "xb"},
		{Weight: 1, Item: "pd"},
		{Weight: 1, Item: "gd"},
		{Weight: 1, Item: "gb"},
		{Weight: 1, Item: "cm"},
		{Weight: 1, Item: "zm"},
		{Weight: 1, Item: "fc"},
		{Weight: 1, Item: "hh"},
		{Weight: 1, Item: "hw"},
		{Weight: 1, Item: "kb"},

}
*/
