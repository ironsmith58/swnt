// Package name provides tables and a generator for NPC and place names
package name

import (
	"github.com/nboughton/go-roll"
	"github.com/nboughton/swnt/content/culture"
)

// table represents the collection of roll lists keyed by cultural background
type table struct {
	Culture culture.Culture
	Male    roll.List
	Female  roll.List
	Surname roll.List
	Place   roll.List
}

// Tables represents a set of table structs
type tables []table

// ByCulture returns a name table that matches the given culture
func (t tables) ByCulture(c culture.Culture) table {
	if c == culture.Any {
		c = culture.Random()
	}

	for _, tbl := range t {
		if tbl.Culture == c {
			return tbl
		}
	}

	return table{}
}

// Table represents rollable tables for individuals names
var Table = tables{
	table{
		Culture: culture.Arabic,
		Male:    roll.List{Items: []string{"Aamir", "Ayub", "Binyamin", "Efraim", "Ibrahim", "Ilyas", "Ismail", "Jibril", "Jumanah", "Kazi", "Lut", "Matta", "Mohammed", "Mubarak", "Mustafa", "Nazir", "Rahim", "Reza", "Sharif", "Taimur", "Usman", "Yakub", "Yusuf", "Zakariya", "Zubair"}},
		Female:  roll.List{Items: []string{"Aisha", "Alimah", "Badia", "Bisharah", "Chanda", "Daliya", "Fatimah", "Ghania", "Halah", "Kaylah", "Khayrah", "Layla", "Mina", "Munisa", "Mysha", "Naimah", "Nissa", "Nura", "Parveen", "Rana", "Shalha", "Suhira", "Tahirah", "Yasmin", "Zulehka"}},
		Surname: roll.List{Items: []string{"Abdel", "Awad", "Dahhak", "Essa", "Hanna", "Harbi", "Hassan", "Isa", "Kasim", "Katib", "Khalil", "Malik", "Mansoor", "Mazin", "Musa", "Najeeb", "Namari", "Naser", "Rahman", "Rasheed", "Saleh", "Salim", "Shadi", "Sulaiman", "Tabari"}},
		Place:   roll.List{Items: []string{"Adan", "Magrit", "Ahsa", "Masqat", "Andalus", "Misr", "Asmara", "Muruni", "Asqlan", "Qabis", "Baqubah", "Qina", "Basit", "Rabat", "Baysan", "Ramlah", "Baytlahm", "Riyadh", "Bursaid", "Sabtah", "Dahilah", "Salalah", "Darasalam", "Sana", "Dawhah", "Sinqit", "Ganin", "Suqutrah", "Gebal", "Sur", "Gibuti", "Tabuk", "Giddah", "Tangah", "Harmah", "Tarifah", "Hartum", "Tarrakunah", "Hibah", "Tisit", "Hims", "Uman", "Hubar", "Urdunn", "Karbala", "Wasqah", "Kut", "Yaburah", "Lacant", "Yaman"}},
	},
	table{
		Culture: culture.Chinese,
		Male:    roll.List{Items: []string{"Aiguo", "Bai", "Bohai", "Chao", "Dai", "Dawei", "Duyi", "Fa", "Fu", "Gui", "Hong", "Jianyu", "Kang", "Li", "Niu", "Peng", "Quan", "Ru", "Shen", "Shi", "Song", "Tao", "Xue", "Yi", "Yuan", "Zian"}},
		Female:  roll.List{Items: []string{"Biyu", "Changying", "Daiyu", "Huidai", "Huiliang", "Jia", "Jingfei", "Lan", "Liling", "Liu", "Meili", "Niu", "Peizhi", "Qiao", "Qing", "Ruolan", "Shu", "Suyin", "Ting", "Xia", "Xiaowen", "Xiulan", "Ya", "Ying", "Zhilan"}},
		Surname: roll.List{Items: []string{"Bai", "Cao", "Chen", "Cui", "Ding", "Du", "Fang", "Fu", "Guo", "Han", "Hao", "Huang", "Lei", "Li", "Liang", "Liu", "Long", "Song", "Tan", "Tang", "Wang", "Wu", "Xing", "Yang", "Zhang"}},
		Place:   roll.List{Items: []string{"Andong", "Luzhou", "Anqing", "Ningxia", "Anshan", "Pingxiang", "Chaoyang", "Pizhou", "Chaozhou", "Qidong", "Chifeng", "Qingdao", "Dalian", "Qinghai", "Dunhuang", "Rehe", "Fengjia", "Shanxi", "Fengtian", "Taiyuan", "Fuliang", "Tengzhou", "Fushun", "Urumqi", "Gansu", "Weifang", "Ganzhou", "Wugang", "Guizhou", "Wuxi", "Hotan", "Xiamen", "Hunan", "Xian", "Jinan", "Xikang", "Jingdezhen", "Xining", "Jinxi", "Xinjiang", "Jinzhou", "Yidu", "Kunming", "Yingkou", "Liaoning", "Yuxi", "Linyi", "Zigong", "Lushun", "Zoige"}},
	},
	table{
		Culture: culture.English,
		Male:    roll.List{Items: []string{ "Adam", "Albert", "Alexander", "Alfred", "Allan", "Andrew", "Anthony", "Archibald", "Arthur", "Basil", "Benjamin", "Brandon", "Brian", "Charles", "Christopher", "Colin", "Daniel", "David", "Donald", "Douglas", "Edgar", "Edmund", "Edward", "Eric", "Frank", "Gary", "George", "Gregory", "Harold", "Henry", "Ian", "Jack", "Jacob", "James", "Jason", "Jeffrey", "John", "Jonathan", "Joseph", "Joshua", "Justin", "Kenneth", "Kevin", "Larry", "Lewis", "Mark", "Matthew", "Michael", "Nicholas", "Oliver", "Patrick", "Paul", "Philip", "Raymond", "Richard", "Robert", "Ronald", "Ryan", "Samuel", "Scott", "Stephen", "Steven", "Thomas", "Timothy", "Walter", "William" }},
		Female:  roll.List{Items: []string{ "Abigail", "Alice", "Amanda", "Amy", "Angela", "Anna", "Anne", "Ashley", "Barbara", "Beatrice", "Betty", "Blanche", "Brenda", "Carol", "Catherine", "Charlotte", "Christine", "Claire", "Cynthia", "Deborah", "Donna", "Dorothy", "Eleanor", "Elizabeth", "Emily", "Emma", "Evelyn", "Frances", "Georgia", "Harriet", "Helen", "Jennifer", "Jessica", "Joan", "Judy", "Julia", "Karen", "Katherine", "Kathleen", "Kimberly", "Laura", "Linda", "Lisa", "Lucy", "Lydia", "Margaret", "Mary", "Melissa", "Michelle", "Molly", "Nancy", "Nicole", "Nora", "Pamela", "Patricia", "Rachel", "Rebecca", "Rosie", "Ruth", "Samantha", "Sandra", "Sarah", "Sharon", "Shirley", "Stephanie", "Susan", "Victoria", "Virginia" }},
		Surname: roll.List{Items: []string{ "Acy", "Adams", "Adolphson", "Ajax", "Allen", "Amery", "Amilia", "Anderson", "Anthonyson", "Bailey", "Baker", "Barbrow ", "Barker", "Beaufoy", "Bell", "Bennett", "Berrycloth", "Birdwhistle", "Black", "Bowman ", "Brazil", "Bread", "Brown", "Brumby", "Butler", "Bythesea", "Callen", "Carter", "Cass", "Chadburn", "Chamberlain", "Champion", "Chandler", "Chapman", "Choules", "Clark", "Clarke", "Coleman", "Collier", "Collins", "Cook", "Cooper", "Corbyn", "Cox", "Crusoe", "Cullimore", "Culpepper", "Daft", "Dagger", "Dankworth", "Darcy", "Davies", "Davis", "Decksheimer", "Denson", "Edevane", "Evenson", "Fernsby", "Fisher", "Fletcher", "Foster", "Gastrell", "Glass", "Gray", "Green", "Halfenaked", "Hall", "Harris", "Harrison", "Hill", "Jackson", "Johnson", "Jones", "Kellogg", "Kemp", "King", "Lee", "Lloyd", "Lorimer", "Loughty", "MacCaa", "Marshall", "Martin", "Miller", "Miracle", "Mitchell", "Moore", "Parker", "Pussett", "Relish", "Richardson", "Roberts", "Robinson", "Rymer", "Sallow", "Sexton", "Shaw", "Slora", "Smith", "Spinster", "Stoddard", "Taylor", "Thomas", "Thompson", "Tucker", "Tumbler", "Turner", "Villin", "Walker", "Waller", "Ward", "Watson", "Webster", "White", "Wilkinson", "Williams", "Wood", "Wright", "Young" }},
		Place:   roll.List{Items: []string{"Aldington", "Kedington", "Appleton", "Latchford", "Ashdon", "Leigh", "Berwick", "Leighton", "Bramford", "Maresfield", "Brimstage", "Markshall", "Carden", "Netherpool", "Churchill", "Newton", "Clifton", "Oxton", "Colby", "Preston", "Copford", "Ridley", "Cromer", "Rochford", "Davenham", "Seaford", "Dersingham", "Selsey", "Shackleton", "Doverdale", "Stanton", "Eastgate", "Elsted", "Stockham", "Ferring", "Stoke", "Gissing", "Sutton", "Heydon", "Thakeham", "Holt", "Thetford", "Trenton", "Hunston", "Thorndon", "Hutton", "Ulting", "Inkberrow", "Upton", "Inworth", "Westdale", "Westgate", "Westhorpe", "Wilmington", "Isfield", "Worcester"}},
	},
	table{
		Culture: culture.Gang,
		Male:    roll.List{Items: []string{ "Starport Steve", "4 Stroke", "AB, Almighty Blood Messiah", "Ace of Spades", "Adonis", "Alik the Greek", "All City", "Animal", "Ant", "Artichoke King", "Baby Face", "Baby Gangster", "Baby-Shanks", "Bakha-Bakha", "Baldo", "Baldy Dom", "Bananas", "Barber", "Bats", "Benny Squint", "Biba", "Big Al", "Big Frank Nitti", "Big G", "Big Head", "Blaze", "Bloodhound", "Blue Eyes", "Bonafide", "Bondar", "Books", "Boot", "Bootsie", "Bouncer", "Bow Wow", "Boxcars", "Brains", "Bronx", "Bugs", "Bugsy", "Bull", "Bullet", "Bumpy", "Burnout", "Butterass", "Butterfingers", "Camel", "Caveman", "Channel 4", "Charles Heung", "Chee-Chee", "Chicken Man", "Chicken man", "Chin", "Chongo", "Clouds", "Clown", "Cool Daddy Cake", "Cranky", "Creeper", "Cruiser", "Cruzito", "Cuts", "Da Boss", "Dapper Don", "Dasher", "Demon", "Dentist", "Digger", "Dogs", "Dreads", "Droopy", "Easy Z", "El Chuey", "Endo", "Enforcer", "FUBAR", "Fangs", "Fast Trigger", "Fat Boy", "Fisherman", "Flames", "Fox", "Freckles", "Free Style", "Frosty", "Gangsta C", "Georgian", "Ghost", "Gizmo", "Goldie", "Grandpa Hussan", "Greasy Thumb", "Gremlin", "Grim Reaper", "Ha-Ha", "Happy", "Hatchet", "Head Busta", "Henry Fok", "Horseface", "Hoza", "Hsu Hai-ching", "Ice Box", "Ice Cube", "Ice T", "Iceman", "Icepick", "Jelly", "Joe Banana", "Johnny Sausage", "Joker", "Kermit", "Kid Blast", "Kid Twist", "Killing Machine", "Knots", "Kolya", "Lai Changxing", "Larry Fab", "Lasha Rustavsky", "Lil Loco", "Lil Man", "Lil’ Loco", "Lips", "Little Insane", "Little Nicky", "Little Red", "Liu Yong", "Lord High Executioner", "Lord Lenin", "Love Candy", "Lucky", "Machine Gun", "Mad Dog", "Mad Hatter", "Manic Maniac", "Master Blaster", "Matches", "Maverick", "Max", "Mega Flex", "Messiah", "Mickey the Cobra", "Midnight", "Miklo", "Minor", "Mom", "Monster", "Montana", "Mr Cisco", "Mr Scrappy", "Murda King", "Muscles", "Music Man", "Mussolini", "Nails", "Numbers", "O Dog", "O.G.", "Old Creepy", "Oso", "Otarik", "Paco Magic", "Paiyaso", "Pat the Cat", "Peanuts", "Pee Wee", "Pelon", "Phones", "Pistol Pete", "Pit Bull", "Playboy", "Plumber", "Popcorn", "Popoye", "Pretty Boy", "Princess", "Pzo", "Qiao Si", "Quack-Quack", "Quiet Godfather", "Rascal", "Razor", "Red Face", "Red Ryder", "Repo Girl", "Rifleman", "Roshan Lankaransky", "Sam Sings in the Night", "Scarface", "Scars", "Scratch", "Shadow", "Shaggy", "Shagro Junior", "Sharpie", "Shellackhead", "Shorty", "Silver Dollar", "Slingshot", "Smallville", "Smiley", "Smoke", "Snake", "Sneakers", "Sniper", "Socks", "Soft Hammer", "Sonny Black", "Spider", "Spoon", "Spyder", "Stacks", "Sticks", "Swifty", "T-Bone", "Taco", "Tank", "Taro", "Tarzan", "Teflon Don", "That Guy", "The Owl", "The Prophet", "The Reluctant Prine", "Three Fingers", "Tick-Tock", "Tonksku-Tom", "Toto", "Trigger Mike", "Two Holes", "Vixen", "Wadd", "Wan Kuok-koi", "Whack-Whack", "Wheels", "Wicked", "Wild Child", "Wild Man", "Witch", "Wolf", "Young G", "Yu Zuomin", "Zhang Ziqiang", "Zhou Guanglong", "Zookeeper",
        }},
		Female:  roll.List{Items: []string{ "Baby One Punch", "Bad bitch", "Bear Trap", "Black Widow", "Chao", "Clumsy", "Cougar", "Dopey", "Flaco", "Foxy", "Ginger", "Ice Queen", "Joker", "Kiki", "Kinky", "Kitty", "La Giggles", "La Shy Girl", "Lil Tramposo", "Lil’ Loco", "Little Clown", "Little Puppet", "Mousey", "Murda Queen", "Pussy", "Red Stiletto", "Sharpies", "Shy Girl", "Smiler", "Snow Queen", "Spooky", "Tigress", "Travieso", }},
		Surname: roll.List{Items: []string{"Andreas", "Argyros", "Dimitriou", "Floros", "Gavras", "Ioannidis", "Katsaros", "Kyrkos", "Leventis", "Makris", "Metaxas", "Nikolaidis", "Pallis", "Pappas", "Petrou", "Raptis", "Simonides", "Spiros", "Stavros", "Stephanidis", "Stratigos", "Terzis", "Theodorou", "Vasiliadis", "Yannakakis"}},
		Place:   roll.List{Items: []string{
            "The Warehouse", "The Diner", "The Parking Lot", "The Abandoned House", "Tenement 39", "The Old School", "Scrap Yard", 
        }},
	},
	table{
		Culture: culture.Greek,
		Male:    roll.List{Items: []string{"Alexander", "Alexius", "Anastasius", "Christodoulos", "Christos", "Damian", "Dimitris", "Dysmas", "Elias", "Giorgos", "Ioannis", "Konstantinos", "Lambros", "Leonidas", "Marcos", "Miltiades", "Nestor", "Nikos", "Orestes", "Petros", "Simon", "Stavros", "Theodore", "Vassilios", "Yannis"}},
		Female:  roll.List{Items: []string{"Alexandra", "Amalia", "Callisto", "Charis", "Chloe", "Dorothea", "Elena", "Eudoxia", "Giada", "Helena", "Ioanna", "Lydia", "Melania", "Melissa", "Nika", "Nikolina", "Olympias", "Philippa", "Phoebe", "Sophia", "Theodora", "Valentina", "Valeria", "Yianna", "Zoe"}},
		Surname: roll.List{Items: []string{"Andreas", "Argyros", "Dimitriou", "Floros", "Gavras", "Ioannidis", "Katsaros", "Kyrkos", "Leventis", "Makris", "Metaxas", "Nikolaidis", "Pallis", "Pappas", "Petrou", "Raptis", "Simonides", "Spiros", "Stavros", "Stephanidis", "Stratigos", "Terzis", "Theodorou", "Vasiliadis", "Yannakakis"}},
		Place:   roll.List{Items: []string{"Adramyttion", "Kallisto", "Ainos", "Katerini", "Alikarnassos", "Kithairon", "Avydos", "Kydonia", "Dakia", "Lakonia", "Dardanos", "Leros", "Dekapoli", "Lesvos", "Dodoni", "Limnos", "Efesos", "Lykia", "Efstratios", "Megara", "Elefsina", "Messene", "Ellada", "Milos", "Epidavros", "Nikaia", "Erymanthos", "Orontis", "Evripos", "Parnasos", "Gavdos", "Petro", "Gytheio", "Samos", "Ikaria", "Syros", "Ilios", "Thapsos", "Illyria", "Thessalia", "Iraia", "Thira", "Irakleio", "Thiva", "Isminos", "Varvara", "Ithaki", "Voiotia", "Kadmeia", "Vyvlos"}},
	},
	table{
		Culture: culture.Indian,
		Male:    roll.List{Items: []string{"Amrit", "Ashok", "Chand", "Dinesh", "Gobind", "Harinder", "Jagdish", "Johar", "Kurien", "Lakshman", "Madhav", "Mahinder", "Mohal", "Narinder", "Nikhil", "Omrao", "Prasad", "Pratap", "Ranjit", "Sanjay", "Shankar", "Thakur", "Vijay", "Vipul", "Yash"}},
		Female:  roll.List{Items: []string{"Amala", "Asha", "Chandra", "Devika", "Esha", "Gita", "Indira", "Indrani", "Jaya", "Jayanti", "Kiri", "Lalita", "Malati", "Mira", "Mohana", "Neela", "Nita", "Rajani", "Sarala", "Sarika", "Sheela", "Sunita", "Trishna", "Usha", "Vasanta"}},
		Surname: roll.List{Items: []string{"Achari", "Banerjee", "Bhatnagar", "Bose", "Chauhan", "Chopra", "Das", "Dutta", "Gupta", "Johar", "Kapoor", "Mahajan", "Malhotra", "Mehra", "Nehru", "Patil", "Rao", "Saxena", "Shah", "Sharma", "Singh", "Trivedi", "Venkatesan", "Verma", "Yadav"}},
		Place:   roll.List{Items: []string{"Ahmedabad", "Jaisalmer", "Alipurduar", "Jharonda", "Alubari", "Kadambur", "Anjanadri", "Kalasipalyam", "Ankleshwar", "Karnataka", "Balarika", "Kutchuhery", "Bhanuja", "Lalgola", "Bhilwada", "Mainaguri", "Brahmaghosa", "Nainital", "Bulandshahar", "Nandidurg", "Candrama", "Narayanadri", "Chalisgaon", "Panipat", "Chandragiri", "Panjagutta", "Charbagh", "Pathankot", "Chayanka", "Pathardih", "Chittorgarh", "Porbandar", "Dayabasti", "Rajasthan", "Dikpala", "Renigunta", "Ekanga", "Sewagram", "Gandhidham", "Shakurbasti", "Gollaprolu", "Siliguri", "Grahisa", "Sonepat", "Guwahati", "Teliwara", "Haridasva", "Tinpahar", "Indraprastha", "Villivakkam"}},
	},
	table{
		Culture: culture.Japanese,
		Male:    roll.List{Items: []string{"Akira", "Daiki", "Daisuke", "Fukashi", "Goro", "Haruto", "Hinata", "Hiro", "Hiroya", "Hotaka", "Itsuki", "Katsu", "Katsuto", "Keishuu", "Kenta", "Kyuuto", "Mikiya", "Minato", "Mitsunobu", "Mitsuru", "Naruhiko", "Nobu", "Ren", "Shigeo", "Shigeto", "Shou", "Sou", "Shuji", "Takaharu", "Teruaki", "Tetsushi", "Tsukasa", "Yamato", "Yasuharu", "Yuuma"}},
		Female:  roll.List{Items: []string{"Aemi", "Airi", "Akari", "Ako", "Aoi", "Ayu", "Chikaze", "Eriko", "Hina", "Honoka", "Ichika", "Kaede", "Kaori", "Keiko", "Kyouka", "Mayumi", "Miho", "Mio", "Mei", "Namiko", "Natsu", "Nobuko", "Rei", "Riko", "Rin", "Ririsa", "Sakimi", "Sakura", "Shihoko", "Shika", "Shiori", "Tsukiko", "Tsuzune", "Yoriko", "Yorimi", "Yoshiko", "Yui", "Yuna"}},
		Surname: roll.List{Items: []string{ "Abe", "Abiko", "Agawa", "Ageda", "Ahane", "Aikawa", "Akihito", "Ando", "Aoki", "Arakaki", "Asuka", "Baba", "Bando", "Bushida", "Chikafuji", "Chinen", "Chisaka", "Daigo", "Eguchi", "Ejiri", "Endo", "Fuji", "Fujiki", "Fujiwara", "Fukuda", "Fukushima", "Furukawa", "Furuya", "Gato", "Genji", "Goda", "Goto", "Gushiken", "Hada", "Haga", "Hagino", "Haku", "Hamada", "Handa", "Hanyu", "Hara", "Haruki", "Hasegawa", "Hasimoto", "Hirohito", "Honda", "Honjo", "Hori", "Ieyasu", "Iida", "Ikeda", "Imada", "Imai", "Ine", "Ishida", "Ishii", "Ishikawa", "Ishizaki", "Ito", "Iwata", "Kagawa", "Kamada", "Kamei", "Kamida", "Kamiyama", "Kaneko", "Kasai", "Kasama", "Kase", "Katayama", "Kato", "Kibe", "Kido", "Kikuchi", "Kinjo", "Kita", "Kobayashi", "Koga", "Kojima", "Kokawa", "Komatsu", "Kondo", "Kouda", "Kubo", "Kudo", "Kurosawa", "Machida", "Maeda", "Maekawa", "Makino", "Masuda", "Matsubara", "Matsumara", "Matsushita", "Miura", "Miyamoto", "Morita", "Murakami", "Musashi", "Nagasawa", "Nagata", "Nakada", "Nakamura", "Nakano", "Nakayama", "Narita", "Naruhito", "Nishikawa", "Nishio", "Nobunaga", "Noguchi", "Nomura", "Oba", "Ochi", "Ogawa", "Ohara", "Ohtani", "Okamoto", "Omori", "Ono", "Onoda", "Osaka", "Oshiro", "Otake", "Ozawa", "Saito", "Sakamoto", "Sanada", "Sato", "Sone", "Suzuki", "Takahashi", "Tanaka", "Toriyama", "Watanabe", "Yamamoto", "Yamasaki" }},
		Place:   roll.List{Items: []string{ "Akita", "Aomori", "Asahikawa", "Bando", "Chiba", "Chigasaki", "Chikuma", "Chikusei", "Chino", "Enoshima", "Fujisawa", "Fukuoka", "Fukushima", "Fukuyama", "Hakodate", "Hammamatsu", "Himeji", "Hirosaki", "Hiroshima", "Hitachi", "Hitachinaka", "Hitachiomiya", "Hitachiota", "Ichihara", "Iida", "Iiyama", "Ina", "Inashiki", "Ishioka", "Itako", "Iwaki", "Kagaoshima", "Kamisu", "Kanazawa", "Kasama", "Kashima", "Kashiwa", "Kasumigaura", "Kawagoe", "Kawaguchi", "Kitaibaraki", "Kitakyushu", "Kiyose", "Kobe", "Kochi", "Koga", "Komagane", "Komoro", "Kumamoto", "Kurashiki", "Kure", "Kurume", "Kyoto", "Kyushu", "Maebashi", "Matsudo", "Matsue", "Matsumoto", "Matsuyama", "Mito", "Mitsukaido", "Miyazaki", "Morioka", "Moriya", "Nagano", "Nagaoka", "Nagasaki", "Nagoya", "Naha", "Naka", "Nakano", "Nara", "Niigata", "Nishinomiya", "Numazu", "Odawara", "Ogi", "Okaya", "Okayama", "Okazaki", "Omachi", "Osaka", "Otsu", "Ryugasaki", "Saku", "Sapporo", "Sasebo", "Sendai", "Settsu", "Shimonoseki", "Shimotsuma", "Shiojiri", "Shizuoka", "Shonan", "Suita", "Suwa", "Suzaka", "Takahagi", "Takamatsu", "Takarazuka", "Takasaki", "Takeo", "Tokushima", "Tokyo", "Tomi", "Toride", "Tottori", "Toyama", "Toyohashi", "Tsuchiura", "Tsukuba", "Ueda", "Ushiku", "Wakayama", "Yokohama", "Yoshikawa", "Yuki" }},
	},
	table{
		Culture: culture.Latin,
		Male:    roll.List{Items: []string{"Agrippa", "Appius", "Aulus", "Caeso", "Decimus", "Faustus", "Gaius", "Gnaeus", "Hostus", "Lucius", "Mamercus", "Manius", "Marcus", "Mettius", "Nonus", "Numerius", "Opiter", "Paulus", "Proculus", "Publius", "Quintus", "Servius", "Tiberius", "Titus", "Volescus"}},
		Female:  roll.List{Items: []string{"Appia", "Aula", "Caesula", "Decima", "Fausta", "Gaia", "Gnaea", "Hosta", "Lucia", "Maio", "Marcia", "Maxima", "Mettia", "Nona", "Numeria", "Octavia", "Postuma", "Prima", "Procula", "Septima", "Servia", "Tertia", "Tiberia", "Titia", "Vibia"}},
		Surname: roll.List{Items: []string{"Antius", "Aurius", "Barbatius", "Calidius", "Cornelius", "Decius", "Fabius", "Flavius", "Galerius", "Horatius", "Julius", "Juventius", "Licinius", "Marius", "Minicius", "Nerius", "Octavius", "Pompeius", "Quinctius", "Rutilius", "Sextius", "Titius", "Ulpius", "Valerius", "Vitellius"}},
		Place:   roll.List{Items: []string{"Abilia", "Lucus", "Alsium", "Lugdunum", "Aquileia", "Mediolanum", "Argentoratum", "Novaesium", "Ascrivium", "Patavium", "Asculum", "Pistoria", "Attalia", "Pompeii", "Barium", "Raurica", "Batavorum", "Rigomagus", "Belum", "Roma", "Bobbium", "Salernum", "Brigantium", "Salona", "Burgodunum", "Segovia", "Camulodunum", "Sirmium", "Clausentum", "Spalatum", "Corduba", "Tarraco", "Coriovallum", "Treverorum", "Durucobrivis", "Verulamium", "Eboracum", "Vesontio", "Emona", "Vetera", "Florentia", "Vindelicorum", "Lactodurum", "Vindobona", "Lentia", "Vinovia", "Lindum", "Viroconium", "Londinium", "Volubilis"}},
	},
	table{
		Culture: culture.Nigerian,
		Male:    roll.List{Items: []string{
            "Adesegun", "Abayomi", "Adama", "Ade", "Adebayo", "Adedamola", "Ademola", "Adewale", "Adeyemi", "Akin", "Akintola", "Amabere", "Amadi", "Arikawe", "Arinze", "Asagwara", "Atiba", "Ayinde", "Ayo", "Ayodeji", "Ayodele", "Ayomide", "Babatunde", "Chibueze", "Chibuike", "Chibuikem", "Chidera", "Chidi", "Chidubem", "Chigozie", "Chijioke", "Chike", "Chima", "Chinedu", "Chisom", "Chiwetei", "Chukwudi", "Chukwuebuka", "Chukwuemeka", "Chukwuma", "Damilola", "Ebubechukwu", "Efe", "Emeka", "Esangbedo", "Ezenwoye", "Folarin", "Genechi", "Geofrey", "Idowu", "Ifeanyi", "Ifeanyichukwu", "Ifeoluwa", "Ike", "Ikechukwu", "Ikenna", "Jidenna", "Juwon", "Kamsiyochukwu", "Kayin", "Kayode", "Kehinde", "Kelechi", "Kenechukwu", "Ketanndu", "Mazi", "Melubari", "Nkanta", "Nnamdi", "Obafemi", "Obi", "Obie", "Obinna", "Okey", "Olajuwon", "Olamide", "Olaoluwa", "Olatunde", "Olawale", "Olufemi", "Olumide", "Oluwadamilare", "Oluwadamilola", "Oluwadarasimi", "Oluwademilade", "Oluwafemi", "Oluwanifemi", "Oluwaseun", "Oluwaseyi", "Oluwatimilehin", "Oluwatobi", "Oluwatobiloba", "Oluwatosin", "Oreoluwa", "Osaze", "Oshay", "Ovie", "Praise", "Rapheal", "Saheed", "Shade", "Somtochukwu", "Taiwo", "Taye", "Tayo", "Temiloluwa", "Temitope", "Tobe", "Tobechukwu", "Tobenna", "Tobi", "Tombari", "Uchechukwu", "Uchenna", "Udofia", "Ugo", "Ugochukwu", "Ugonna", "Uyoata", "Uzochi", "Uzoma", }},
		Female:  roll.List{Items: []string{
            "Abike", "Abeni", "Abiola", "Adaeze", "Adah", "Adama", "Adanna", "Adaora", "Adenike", "Adeola", "Adesuwa", "Adunola", "Ajah", "Amadi", "Amara", "Amarachi", "Amma", "Anguli", "Anjolaoluwa", "Arewa", "Asari", "Ashani", "Ayomide", "Bisola", "Chiamaka", "Chidera", "Chidinma", "Chika", "Chimamanda", "Chimere", "Chinenye", "Chinyere", "Chioma", "Chisom", "Chizara", "Chizaram", "Daisi", "Damilola", "Deola", "Dera", "Dola", "Dosha", "Eduwa", "Emilohi", "Eniola", "Ezinne", "Favor", "Favour", "Fehintola", "Folasade", "Folashade", "Fumi", "Honour", "Ife", "Ifeoluwa", "Ifeoma", "Ijeoma", "Ima", "Inioluwa", "Ivie", "Iyanla", "Kelechi", "Kemi", "Mahparah", "Mayana", "Minika", "Nene", "Ngozi", "Nkechi", "Nkolika", "Nkoyo", "Nneka", "Nnenna", "Nneoma", "Nuanae", "Obie", "Obioma", "Ogechi", "Ola", "Olafemi", "Olamae", "Olamide", "Oluchi", "Oluwadamilola", "Oluwadarasimi", "Oluwakemi", "Oluwanifemi", "Oluwatamilore", "Oluwatobi", "Oni", "Oreoluwa", "Osa", "Ovie", "Praise", "Sade", "Shacara", "Shade", "Shakara", "Shanumi", "Simi", "Sola", "Sominabo", "Stephenia", "Suliat", "Tami", "Tari", "Tariere", "Temedire", "Temiloluwa", "Temitope", "Teniola", "Timaya", "Timi", "Toba", "Tobi", "Tomi", "Tonna", "Toya", "Toye", "Uchenna", "Yemisi", "Yetunde", "Zelie"}},
		Surname: roll.List{Items: []string{"Adegboye", "Ade", "Adebayo", "Adeniyi", "Adeniyi", "Adeyeku", "Adeyemi", "Adunola", "Afolabi", "Agbaje", "Ajayi", "Akande", "Akin", "Akpan", "Akpan", "Akpehi", "Alabi", "Aliki", "Alonge", "Amadi", "Amara", "Asuni", "Awe", "Ayeni", "Ayo", "Babalola", "Babangida", "Bada", "Bakare", "Balogun", "Banke", "Bolla", "Bose", "Chima", "Chisom", "Chuba", "Chukwu", "Coker", "Dada", "Dare", "Duro", "Eke", "Ekim", "Essien", "Eze", "Ezeiruaku", "Fabiola", "Fasola", "Favor", "Fife", "Fruit", "Funke", "Ibe", "Idowu", "Ige", "Igwe", "Ike", "Ivie", "Juba", "Kalu", "Kola", "Lade", "Lawal", "Longe", "Luper", "Machala", "Mosco", "Nale", "Njoku", "Nwachukwu", "Nwankwo", "Nwokolo", "Nwosu", "Nzeocha", "Obara", "Obey", "Obi", "Obie", "Odum", "Ogbonna", "Ojo", "Ojo", "Okafor", "Oke", "Okeke", "Okereke", "Okey", "Okoli", "Okolo", "Okon", "Okonkwo", "Okonkwo", "Okoro", "Okoye", "Ola", "Olaniyan", "Olawale", "Olumese", "Onajobi", "Oni", "Onuoha", "Opara", "Ore", "Osuji", "Pere", "Riches", "Sade", "Shade", "Shima", "Simi", "Sola", "Soyinka", "Taiwo", "Taye", "Tobe", "Tola", "Tope", "Toya", "Toye", "Ude", "Umana", "Ure", "Wike", "Yamusa", "Zamani",}},
		Place:   roll.List{Items: []string{"Abadan", "Abuja", "Ador", "Agatu", "Agbokim", "Akamkpa", "Akpabuyo", "Ala", "Askira", "Bakassi", "Bama", "Bayo", "Bekwara", "Benin", "Biase", "Boki", "Buruku", "Calabar", "Chibok", "Damboa", "Dikwa", "Etung", "Gboko", "Gubio", "Guzamala", "Gwoza", "Hawul", "Ibadan", "Ikom", "Jere", "Kalabalge", "Kano", "Katsina", "Knoduga", "Konshishatse", "Kukawa", "Kwande", "Kwayakusar", "Lagos", "Logo", "Mafa", "Makurdi", "Nganzai", "Obanliku", "Obi", "Obubra", "Obudu", "Odukpani", "Ogbadibo", "Ohimini", "Okpokwu", "Onitsha", "Otukpo", "Owerri", "Shani", "Ugep", "Uyo", "Vandeikya", "Warri", "Yala" }},
	},
	table{
		Culture: culture.Russian,
		Male:    roll.List{Items: []string{"Aleksandr", "Andrei", "Arkady", "Boris", "Dmitri", "Dominik", "Grigory", "Igor", "Ilya", "Ivan", "Kiril", "Konstantin", "Leonid", "Nikolai", "Oleg", "Pavel", "Petr", "Sergei", "Stepan", "Valentin", "Vasily", "Viktor", "Yakov", "Yegor", "Yuri"}},
		Female:  roll.List{Items: []string{"Aleksandra", "Anastasia", "Anja", "Catarina", "Devora", "Dima", "Ekaterina", "Eva", "Irina", "Karolina", "Katlina", "Kira", "Ludmilla", "Mara", "Nadezdha", "Nastassia", "Natalya", "Oksana", "Olena", "Olga", "Sofia", "Svetlana", "Tatyana", "Vilma", "Yelena"}},
		Surname: roll.List{Items: []string{"Abelev", "Bobrikov", "Chemerkin", "Gogunov", "Gurov", "Iltchenko", "Kavelin", "Komarov", "Korovin", "Kurnikov", "Lebedev", "Litvak", "Mekhdiev", "Muraviov", "Nikitin", "Ortov", "Peshkov", "Romasko", "Shvedov", "Sikorski", "Stolypin", "Turov", "Volokh", "Zaitsev", "Zhukov"}},
		Place:   roll.List{Items: []string{"Amur", "Omsk", "Arkhangelsk", "Orenburg", "Astrakhan", "Oryol", "Belgorod", "Penza", "Bryansk", "Perm", "Chelyabinsk", "Pskov", "Chita", "Rostov", "Gorki", "Ryazan", "Irkutsk", "Sakhalin", "Ivanovo", "Samara", "Kaliningrad", "Saratov", "Kaluga", "Smolensk", "Kamchatka", "Sverdlovsk", "Kemerovo", "Tambov", "Kirov", "Tomsk", "Kostroma", "Tula", "Kurgan", "Tver", "Kursk", "Tyumen", "Leningrad", "Ulyanovsk", "Lipetsk", "Vladimir", "Magadan", "Volgograd", "Moscow", "Vologda", "Murmansk", "Voronezh", "Novgorod", "Vyborg", "Novosibirsk", "Yaroslavl"}},
	},
	table{
		Culture: culture.Spanish,
		Male:    roll.List{Items: []string{"Alejandro", "Alonso", "Amelio", "Armando", "Bernardo", "Carlos", "Cesar", "Diego", "Emilio", "Estevan", "Felipe", "Francisco", "Guillermo", "Javier", "Jose", "Juan", "Julio", "Luis", "Pedro", "Raul", "Ricardo", "Salvador", "Santiago", "Valeriano", "Vicente"}},
		Female:  roll.List{Items: []string{"Adalina", "Aleta", "Ana", "Ascencion", "Beatriz", "Carmela", "Celia", "Dolores", "Elena", "Emelina", "Felipa", "Inez", "Isabel", "Jacinta", "Lucia", "Lupe", "Maria", "Marta", "Nina", "Paloma", "Rafaela", "Soledad", "Teresa", "Valencia", "Zenaida"}},
		Surname: roll.List{Items: []string{"Arellano", "Arispana", "Borrego", "Carderas", "Carranzo", "Cordova", "Enciso", "Espejo", "Gavilan", "Guerra", "Guillen", "Huertas", "Illan", "Jurado", "Moretta", "Motolinia", "Pancorbo", "Paredes", "Quesada", "Roma", "Rubiera", "Santoro", "Torrillas", "Vera", "Vivero"}},
		Place:   roll.List{Items: []string{"Aguascebas", "Loreto", "Alcazar", "Lujar", "Barranquete", "Marbela", "Bravatas", "Matagorda", "Cabezudos", "Nacimiento", "Calderon", "Niguelas", "Cantera", "Ogijares", "Castillo", "Ortegicar", "Delgadas", "Pampanico", "Donablanca", "Pelado", "Encinetas", "Quesada", "Estrella", "Quintera", "Faustino", "Riguelo", "Fuentebravia", "Ruescas", "Gafarillos", "Salteras", "Gironda", "Santopitar", "Higueros", "Taberno", "Huelago", "Torres", "Humilladero", "Umbrete", "Illora", "Valdecazorla", "Isabela", "Velez", "Izbor", "Vistahermosa", "Jandilla", "Yeguas", "Jinetes", "Zahora", "Limones", "Zumeta"}},
	},
}

// System is a list of possible star system names taken from a wikipedia page on fictional planet names.
var System = roll.List{Items: []string{"Abyormen", "Acheron", "Aegus", "Aether", "Ahnooie-4", "Aiur", "Aka", "Alcarinque", "Aldeian", "Alderaan", "Alkarinque", "Alpha", "Altair IV", "Alwas", "Amanga", "Amazo", "Ambar", "Anarres", "Anubis", "Aquarius", "Aquas", "Arcadia", "Arda", "Arisia", "Ark", "Arlia", "Armaghast", "Arrakis", "Astra", "Athos", "Athshe", "Atlantis", "Aurelia", "Auron", "Axturias", "Azeroth", "Baab", "Bajor", "Balaho", "Baloris", "Baltan", "Barathrum", "Barrayar", "Barsoom", "Bas-Lag", "Bazoik", "Beezee", "Belzagor", "Beowulf", "Bespin", "Beulah", "Bismoll", "Black Star", "Blue Sands", "Bog", "Bop", "Boskone", "Braal", "Brontitall", "Bryyo", "Cadwal", "Caladan", "Calafia", "Camazotz", "Caprica", "Carnil", "Carnivalia", "Centauri", "Cetaganda", "Chel", "Chthon", "Clin", "Corneria", "Coruscant", "Covenant", "Crematoria", "Crete", "Cybertron", "Cygnus Alpha", "Cyteen", "Dada", "Dagobah", "Darkover", "Dar Sai", "Daxam", "Dyan", "Deemi", "Demeter", "Denzi", "Dezoris", "Dhrawn", "Diso", "Doisac", "Dorsai", "Dosadi", "Dragon's Egg", "Dragon", "Dres", "Druidia", "Dryad", "Dump", "Duna", "Durdane", "Ea", "Earendil", "Eayn", "Echronedal", "Eeloo", "Elemmire", "Ellicoore 2", "Elysia", "Emerald", "Empyrrean", "Endor", "Epsilon 3", "Erna", "Eve", "Expel", "Exxilon", "Eylor", "Famille", "Fanbelt", "Far Away", "Fargett", "Fhloston", "Fichina", "Fiorina", "Flash", "Fortuna", "Freeza 79", "Fribbulus Xax", "Friedland", "Frystaat", "Galaxian 3", "Gallifrey", "Gamilon", "Gamilus", "Garissa", "Garrota", "Garth", "Gauda Prime", "Gaul", "Gelidus", "Genesis", "Gethen", "Giedi Prime", "Giganda", "Girath", "Gloob", "Gnarlach", "Gnosticus IV", "Gobotron", "God's Grove", "Golgota", "Gor", "Gorgona", "Gork", "Grayson", "Groth", "Gurun", "Gyodai", "Hades", "Hain", "Halvmork", "Harvest", "Hazard", "Heath", "Hebron", "Hegira", "Hekla", "Helicon", "Helion Prime", "Helliconia", "Hiigara", "Hikari", "Home", "Hope", "Horizon", "Hoth", "Houston", "Htrae", "Hummard", "Hyaita 4", "Hydros", "Hydross", "Hyperion", "Iga", "Ilu", "Imbar", "Incandescent", "Interchange", "Ireta", "Irk", "Ivy", "Ix", "Ixchel", "Janjur Qom", "Jijo", "Jinx", "Jobis", "Jool", "Jophekka", "Jurai", "Kaitain", "Kalimdor", "Kamino", "Kanassa", "Karn", "Kashyyyk", "Katina", "Kelewan", "Kerbin", "Kharak", "King", "Kithrup", "Klaus", "Klendathu", "Kobaia", "Kobol", "Komarr", "Koozebane", "Korath III", "Kosmos", "Krelar", "Krikkit", "Krishna", "Kronos", "Krypton", "Kukulkan", "Lagash", "La Maetelle", "Lamarckia", "Lave", "Laythe", "Leesti", "Legis XV", "Leonida", "Leslie", "Londinium", "Luinil", "Lumbar", "Lumen", "Lumiere", "Lusch", "Lusitania", "LV-426", "MacBeth", "Maetel", "Magma", "Magrathea", "Majipoor", "Mare Infinitus", "Marshmalia", "Marune", "Maske", "Maui-Covenant", "Medea", "Meiji", "Mejerr", "Melmac", "Mer", "Meridian", "Merle", "Mesklin", "Metaluna", "Midkemia", "Milokeenia", "Minbar", "Miranda", "Mogo", "Moho", "Mok", "Mondas", "Monea", "Mongo", "Mons", "Mor-Tax", "Morthrai", "Motavia", "Mote Prime", "Mount", "Mustafar", "Naboo", "Nackle", "Nacre", "Namek", "Narn", "Navi", "Nebula 71", "Nebula L-77", "Nebula Z95", "Nede", "Nemesis", "Nenar", "New Amazonia", "New Chicago", "New Earth", "New Namek", "New Terra", "New Vegeta", "Nihil", "Nirn", "Nopalgarth", "Norfolk", "Norion", "Nova Kong", "Nuliajuk", "Nyvan", "Oa", "Oban", "Omicron", "Omnivarium", "Onyx", "Optera", "Ork", "Ormazd", "Orthe", "Osiris", "Pacem", "Palain IX", "Palamok", "Palma", "Palshife", "Pandarve", "Pandora", "Pant", "Parvati", "Peladon", "Pern", "Petaybee", "Phaaze", "Pittsburgh", "Arb", "Plateau", "Plootarg", "Pluto", "Pol", "Pyrrus", "Q-13", "Qar'To", "Qom-Riyadh", "Qo'noS", "Quintessa", "Rain", "Rainbow", "Raxicor", "Reach", "Red Star", "Regis III", "Remulak", "Remus", "Rentus", "Requiem", "Resurgam", "Reverie", "Riedquat", "Rigel", "Rigel 7", "Rime", "River", "Roak", "Roche", "Romulus", "Rougpelt", "Rubanis", "Ruzhena", "Rykros", "Rylos", "Salusa Secundus", "Sanctuary", "Sanghelios", "Sangre", "Saraksh", "Sarris", "Saula", "Sauria", "Sauron", "Second Miltia", "Sedon", "Sedra", "Sergyar", "Serpo", "Sesharrim", "Shadow", "Shaggai", "Shikasta", "Shora", "Sigma Octanus IV", "Signo", "Sihnon", "Skaro", "Sky's Edge", "Solaria", "Solaris", "Solbrecht", "Sol Draconi", "Soror", "Sparta", "Spherus Magna", "Spider", "Spira", "SR-388", "Star One", "St. Ekaterina", "Stroggos", "Styx", "Synnax", "Tagora", "Talark", "Tallon IV", "Tamaran", "Tanis", "Tanith", "Tatooine", "Taurus", "Te", "Tek", "Tellus Secundus", "Tellus Tertius", "Temblor", "Tenebra", "Tergiverse IV", "Terminal", "Terminus", "Thalassa", "Thalassean", "Thallon", "Thel", "Thermia", "The Smoke Ring", "Thra", "Tiamat", "T'ien Shan", "Timbl", "Tirol", "Tissa", "Titan", "Titania", "Tleilax", "Tokyo", "Torto", "Traal", "Tralfamadore", "Trantor", "Trenco", "Tribute", "Trullion", "Tschai", "T'vao", "Twinsun", "Tylo", "U40", "Ummo", "Undomiel", "Unicron", "Uriel", "Urth", "Urtraghus", "Vall", "Vanguard 3", "Vega", "Vegandon", "Vegeta", "Velantia", "Velux", "Venom", "Vindine", "Vladislava", "Vorticon VI", "Vortis", "Vulcan", "Vusstra", "Wallach IX", "Waterloo", "Wegthor", "Wormwood", "Wyst", "X", "X-13", "Xenex", "Xenon", "Xindus", "Yaila", "Yavin 4", "Yellowstone", "Yugopotamia", "Zahir", "Zark", "Zarkon", "Zartron-9", "Zebes", "Zedelbrock 473", "Zedelbrock 475", "Zeelich", "Z'ha'dum", "Zog", "Zok", "Zokk", "Zoness", "Zorg", "Zutinma", "Zyrgon"}}
