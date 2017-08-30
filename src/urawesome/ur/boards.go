package ur

var UrBoard = Board{
    Width: 8,
    Tiles: []Tile{
        Rosetta, FourEyes, FiveDots, FourEyes, Invalid, Invalid, Rosetta, Pyramid,
        Finish, FiveDots, FourCinques, Rosetta, FiveDots, FourCinques, FourEyes, FiveDots,
        Rosetta, FourEyes, FiveDots, FourEyes, Invalid, Invalid, Rosetta, Pyramid,
    },
}

/// also called "Tau", pretty much identical to long jiroft board
var AsebBoard = Board{
    Width: 12,
    Tiles: []Tile{
        Rosetta, Blank, Blank, Blank, Invalid, Invalid, Invalid, Invalid, Invalid, Invalid, Invalid, Invalid,
        Blank, Blank, Blank, Rosetta, Blank, Blank, Blank, Rosetta, Blank, Blank, Blank, Rosetta,
        Rosetta, Blank, Blank, Blank, Invalid, Invalid, Invalid, Invalid, Invalid, Invalid, Invalid, Invalid,
    },
}

// http://www.thehindu.com/features/metroplus/society/tradtional-board-games-from-kochi-to-iraq/article7711918.ece
// var AshaBoard

// https://boardgamegeek.com/boardgame/39628/jiroft
//  https://sites.google.com/site/boardandpieces/list-of-games/the-jiroft-games
// var LongJiroftBoard
// var ShortJiroftBoard

