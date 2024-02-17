const BlankIcon = 0,
    FishLeft = 1,
    FishRight = 2,
    FishCircle = 3,
    FishUp = 4,
    SnakeInfinity = 5,
    SnakeEight = 6,
    SnakeTwoHead = 7,
    SnakeCloud = 8,
    DragonLeft = 9,
    DragonRight = 10,
    DragonUp = 11,
    DragonDown = 12,
    BirdDown = 13,
    BirdPerched = 14,
    BirdLeft = 15,
    BirdBehind = 16;

class Wish {
    id= 0
    name = ''
    location = ''
    effect = ''

    /**
     * @type {Number[]}
     */
    values = []

    constructor(id, name, location, effect) {
        this.id = id;
        this.name = name;
        this.location = location;
        this.effect = effect;
    }
}

const Symbols = new Wish(-1, "Wish Symbols", "None", "None");
Symbols.values = [
    FishLeft,   FishRight,    FishCircle, FishUp,      SnakeInfinity,
    SnakeEight, SnakeTwoHead, SnakeCloud, DragonLeft,  DragonRight,
    DragonUp,   DragonDown,   BirdDown,   BirdPerched, BirdLeft,
    BirdBehind, BlankIcon,    FishLeft,   FishRight,   FishCircle,
]

const BlankWall = new Wish(0, "Blank", "None", "None");
BlankWall.values = Array.from(Array(20), ()=> BlankIcon)

const Wish1 = new Wish(1, "To Feed An Addiction", "Shattered throne, after ogre fight", "Grants an etherial key, used in raid treasure room")
Wish1.values = [
    FishLeft,   FishLeft,   BirdPerched, BirdPerched, BirdPerched,
    FishLeft,   FishLeft,   BlankIcon,   BirdPerched, BirdPerched,
    SnakeEight, SnakeEight, BlankIcon,   FishUp,      FishUp,
    SnakeEight, SnakeEight, DragonRight, FishUp,      FishUp,
]

const Wish2 = new Wish(2, "Material Validation", "Raid, on ceiling before riven", "Spawns chest between 3rd & 4th encounters")
Wish2.values = [
    FishUp,    BlankIcon, DragonUp,    BlankIcon, FishCircle,
    FishUp,    BlankIcon, DragonUp,    BlankIcon, FishCircle,
    FishRight, BlankIcon, DragonRight, BlankIcon, FishCircle,
    FishRight, BlankIcon, DragonRight, BlankIcon, SnakeInfinity,
]

const Wish3 = new Wish(3, "Others To Celebrate Your Success", "New Dreaming City Cutscene (titan shield)", "Awards 'Numbers of Power' Emblem")
Wish3.values = [
    FishLeft,   BlankIcon,     FishLeft,   BlankIcon,     FishLeft,
    BlankIcon,  SnakeInfinity, BlankIcon,  SnakeInfinity, BlankIcon,
    DragonLeft, BlankIcon,     DragonLeft, BlankIcon,     DragonLeft,
    BlankIcon,  BirdDown,      BlankIcon,  BirdDown,      BlankIcon,
]

const Wish4 = new Wish(4, "To Look Athletic And Elegant", "Raid, after the large bridge before 2nd encounter", "Warp to Shuro Chi")
Wish4.values = [
    BirdLeft,     BirdBehind, BirdBehind, BirdBehind, BirdDown,
    SnakeTwoHead, BirdLeft,   SnakeCloud, BirdDown,   SnakeTwoHead,
    SnakeTwoHead, BirdDown,   BlankIcon,  BirdLeft,   SnakeTwoHead,
    BirdDown,     SnakeEight, SnakeEight, SnakeEight, BirdLeft,
]

const Wish5 = new Wish(5, "For A Promising Future", "Raid, ascendant realm after Shuro Chi", "Warp to Morgeth")
Wish5.values = [
    SnakeInfinity, DragonDown, BirdLeft, DragonDown, SnakeInfinity,
    BlankIcon, DragonLeft, BirdBehind, DragonLeft, BlankIcon,
    BlankIcon, DragonLeft, BirdLeft, DragonLeft, BlankIcon,
    SnakeInfinity, DragonDown, BirdLeft, DragonDown, SnakeInfinity,
]

const Wish6 = new Wish(6, "To Move The Hands Of Time", "Raid, before elevator to the vault", "Warp to Vault")
Wish6.values = [
    DragonRight, BirdDown,    BirdBehind, BirdDown,    DragonRight,
    FishLeft,    DragonRight, BlankIcon,  DragonRight, FishLeft,
    FishLeft,    BlankIcon,   BlankIcon,  BlankIcon,   FishLeft,
    DragonUp,    DragonUp,    DragonUp,   DragonUp,    DragonUp
]

const Wish7 = new Wish(7, "To Help A Friend In Need", "Inside Riven's jumping puzzle", "Warp to Riven")
Wish7.values = [
    BirdBehind, FishUp,    DragonDown,  BlankIcon,    SnakeEight,
    DragonLeft, BlankIcon, BlankIcon,   SnakeTwoHead, BirdLeft,
    DragonLeft, BlankIcon, BirdPerched, SnakeTwoHead, BirdLeft,
    BirdBehind, FishUp,    DragonDown,  BlankIcon,    SnakeEight
]

const Wish8 = new Wish(8, "To Stay Here Forever", "Ledge near Shuro Chi", "Plays song 'Hope for the Future'")
Wish8.values = [
    BlankIcon, DragonDown, BirdPerched, DragonUp, BlankIcon,
    BlankIcon, DragonDown, BirdPerched, DragonUp, BlankIcon,
    BlankIcon, DragonDown, BirdBehind,  DragonUp, BlankIcon,
    BlankIcon, DragonDown, BirdPerched, DragonUp, BlankIcon,
]

const Wish9 = new Wish(9, "To Stay Here Forever", "Building between 2nd and 3rd encounters", "Adds Failsafe voice lines")
Wish9.values = [
    BlankIcon, BlankIcon, BlankIcon, BlankIcon,   BlankIcon,
    BlankIcon, FishRight, FishRight, FishRight,   BlankIcon,
    BlankIcon, FishRight, FishRight, BirdPerched, BlankIcon,
    BlankIcon, BlankIcon, BlankIcon, BlankIcon,   BlankIcon,
]

const Wish10 = new Wish(10, "To Stay Here Forever", "Cathedral of Scars, 3rd round on cliff in the distance", "Adds Drifter voice lines")
Wish10.values = [
    FishRight, FishRight,  BlankIcon,    BlankIcon,    FishRight,
    FishRight, DragonDown, DragonDown,   BlankIcon,    FishUp,
    FishRight, BlankIcon,  SnakeTwoHead, SnakeTwoHead, FishUp,
    BlankIcon, BlankIcon,  BlankIcon,    FishUp,       FishUp,
]

const Wish11 = new Wish(11, "To Stay Here Forever", "Nessus, Sunken Cavern near Vex portal", "Grunt birthday party effect")
Wish11.values = [
    BirdLeft,   BirdLeft,   BlankIcon, BirdBehind,    BirdBehind,
    BirdLeft,   BlankIcon,  FishRight, SnakeInfinity, BirdBehind,
    DragonLeft, BlankIcon,  FishRight, BlankIcon,     SnakeEight,
    DragonLeft, DragonLeft, BlankIcon, SnakeEight,    SnakeEight,
]

const Wish12 = new Wish(12, "To Open Your Mind To New Ideas", "Titan, Siren's Watch on the back of a door", "Adds effects to heads of fireteam")
Wish12.values = [
    FishLeft, BlankIcon, BlankIcon,  BlankIcon,  DragonLeft,
    FishLeft, FishLeft,  DragonLeft, DragonLeft, DragonLeft,
    FishUp,   FishUp,    DragonDown, BirdLeft,   BirdLeft,
    FishUp,   BlankIcon, BlankIcon,  BlankIcon,  BirdLeft,
]

const Wish13 = new Wish(13, "For The Means To Feed An Addiction", "Treasure room after opening all chests", "Petra's Run (flawless mode)")
Wish13.values = [
    BirdDown,   FishRight, BirdPerched, BirdLeft,  SnakeCloud,
    BirdDown,   FishRight, BlankIcon,   BirdLeft,  SnakeCloud,
    SnakeCloud, BirdLeft,  BlankIcon,   FishRight, BirdDown,
    SnakeCloud, BirdLeft,  BlankIcon,   FishRight, BirdDown,
]

const Wish14 = new Wish(14, "For Love And Support", "Mara Sov's throne world", "Spawns Ahamkara eggs")
Wish14.values = [
    SnakeInfinity, DragonRight, BirdPerched, DragonLeft, SnakeInfinity,
    BlankIcon,     BlankIcon,   BirdBehind,  BlankIcon,  BlankIcon,
    FishLeft,      BlankIcon,   SnakeEight,  SnakeCloud, BlankIcon,
    SnakeInfinity, DragonDown,  BlankIcon,   BlankIcon,  DragonUp,
]

export const WishConfigurations = {
    BlankWall,
    Symbols,
    Walls: [
        BlankWall,
        Wish1,
        Wish2,
        Wish3,
        Wish4,
        Wish5,
        Wish6,
        Wish7,
        Wish8,
        Wish9,
        Wish10,
        Wish11,
        Wish12,
        Wish13,
        Wish14
    ],
}
