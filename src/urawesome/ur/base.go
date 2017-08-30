package ur

type Tile string
const (
    Invalid Tile = "invalid"
    FourEyes Tile = "four-eyes"
    FiveDots Tile = "five-dots"
    Rosetta Tile = "rosetta"
    Finish Tile = "finish"
    FourCinques Tile = "four-cinques"
    Pyramid Tile = "pyramid"
    Blank Tile = "blank"
)

// A matrix of tiles representing the game board. This structure is supposed to be 
// immutable.
type Board struct {
    Width int
    Tiles []Tile
}

type Player string
const (
    White Player = "white"
    Black Player = "black"
)

func (p Player) Next() Player {
    if p == White { return Black } else { return White }
}

// A piece has a player who owns it, and it may be flipped. When pieces are
// flipped, and what it means for them to be flipped, is up to the rulset.
type Piece struct {
    Player Player
    Flipped bool
}

// Pieces are organized in stacks. The topmost piece is the active one and only
// it may be moved. How pieces are stacked, what kind of pieces are allowed to be
// stacked, is managed by the ruleset.
type Stack []Piece

func (s *Stack) Top() *Piece {
    l := len(*s)
    if l != 0 {
        return &(*s)[l - 1]
    } else {
        return nil
    }
}

func (s *Stack) Push(p Piece) {
    *s = append(*s, p)
}

func (s *Stack) Pop() *Piece {
    l := len(*s)
    if l != 0 {
        r := &(*s)[l - 1]
        *s = (*s)[:l - 1]
        return r
    } else {
        return nil
    }
}

// Piece is to be moved back to the players stash or to his score.
const OutOfGame = 1000

// Move a piece from `From` to `To`. `From` and `To` may denote stacks,
// or one of them may be `OutOfGame`.
type Move struct {
    From uint
    To uint
}

func NewMove(From uint, To uint) Move {
    return Move{ From, To }
}

type GamePhase string
const (
    ThrowDices GamePhase = "throw-dices"
    ChooseMove GamePhase = "choose-move"
    HasWon GamePhase = "has-won"
    HasConceded GamePhase = "has-conceded"
)

type Ruleset interface {
    NumDices() uint
    NumPieces() uint
    ChooseMove(*Game, uint)
    ThrowDices(*Game)
}

type Game struct {
    ruleset Ruleset
    Board *Board
    Pieces []Stack
    CurrentPlayer Player
    Dices []bool
    Phase GamePhase
    PossibleMoves []Move
    Scores map[Player]uint
    Stashes map[Player]uint
}

func NewGame(board *Board, ruleset Ruleset) Game {
    return Game {
        ruleset: ruleset,
        Board: board,
        Pieces: make([]Stack, len(board.Tiles)),
        CurrentPlayer: White,
        Dices: make([]bool, ruleset.NumDices()),
        Phase: ThrowDices,
        PossibleMoves: make([]Move, 0),
        Scores: map[Player]uint{
            White: 0,
            Black: 0,
        },
        Stashes: map[Player]uint{
            White: ruleset.NumPieces(),
            Black: ruleset.NumPieces(),
        },
    }
}

func (g *Game) Concede(p Player) {
    if g.Phase == HasWon || g.Phase == HasConceded { panic("Can't concede, game already ended!") }

    g.CurrentPlayer = p
    g.Phase = HasConceded
}
func (g *Game) ChooseMove(idx uint) {
    if g.Phase != ChooseMove { panic("Not in ChooseMove phase!") }

    g.ruleset.ChooseMove(g, idx)
}

func (g *Game) ThrowDices() {
    if g.Phase != ThrowDices { panic("Not in ThrowDices phase!") }

    g.ruleset.ThrowDices(g)
}

func (g *Game) addPossibleMove(m Move) {
    g.PossibleMoves = append(g.PossibleMoves, m)
}
