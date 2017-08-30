package ur;

import "math/rand"

const finkelNumDices = 4
const finkelNumPieces = 5

var finkelPaths = map[*Board]map[Player][]uint {
    &UrBoard: {
        White: { 19, 18, 17, 16, 8, 9, 10, 11, 12, 13, 14, 15, 23, 22, },
        Black: { 3, 2, 1, 0, 8, 9, 10, 11, 12, 13, 14, 15, 7, 6, },
    },
}

type FinkelRuleset struct {
    paths map[Player][]uint
    numDices uint
    numPieces uint
}

func NewFinkelRuleset(board *Board) FinkelRuleset {
    paths := finkelPaths[board]
    if paths == nil {
        panic("FinkelRuleset has no paths for this board!")
    }
    return FinkelRuleset{paths: paths, numDices: finkelNumDices, numPieces: finkelNumPieces}
}

func (r *FinkelRuleset) NumDices() uint {
    return r.numDices
}

func (r *FinkelRuleset) NumPieces() uint {
    return r.numPieces
}

func (r *FinkelRuleset) ChooseMove(g *Game, idx uint) {
    if g.Phase != ChooseMove { panic("Not in ChooseMove phase!") }

    move := g.PossibleMoves[idx]
    g.PossibleMoves = g.PossibleMoves[:0]

    var piece Piece
    if move.From == OutOfGame {
        g.Stashes[g.CurrentPlayer] -= 1
        piece = Piece{ Player: g.CurrentPlayer }
    } else {
        piece = *g.Pieces[move.From].Pop()
    }

    extraTurn := false
    if move.To == OutOfGame {
        g.Scores[g.CurrentPlayer] += 1

        if g.Scores[g.CurrentPlayer] == r.numPieces {
            g.Phase = HasWon
            return
        }
    } else {
        oldPiece := g.Pieces[move.To].Pop()
        g.Pieces[move.To].Push(piece)
        if oldPiece != nil {
            g.Stashes[oldPiece.Player] += 1
        }

        extraTurn = g.Board.Tiles[move.To] == Rosetta
    }

    if !extraTurn {
        g.CurrentPlayer = g.CurrentPlayer.Next()
    }
    g.Phase = ThrowDices
}

func (r *FinkelRuleset) ThrowDices(g *Game) {
    if g.Phase != ThrowDices { panic("Not in ThrowDices phase!") }

    roll := 0
    for i := range g.Dices {
        v := rand.Int() % 2
        g.Dices[i] = v == 1
        if v == 1 { roll += 1 }
    }

    defer func() {
        if len(g.PossibleMoves) == 0 {
            g.CurrentPlayer = g.CurrentPlayer.Next()
            g.Phase = ThrowDices
        } else {
            g.Phase = ChooseMove
        }
    }()

    if roll == 0 { return }

    path := r.paths[g.CurrentPlayer]
    for i,pos := range path[:len(path) + 1 - roll] {
        piece := g.Pieces[pos].Top()

        if piece == nil || piece.Player != g.CurrentPlayer {
            if i + 1 == roll && g.Stashes[g.CurrentPlayer] > 0 {
                g.addPossibleMove(Move{From: OutOfGame, To: pos})
            }
            continue
        }

        j := i + roll
        if j < len(path) {
            targetPos := path[j]
            targetPiece := g.Pieces[targetPos].Top()

            canBeat := g.Board.Tiles[targetPos] != Rosetta
            if targetPiece == nil || (targetPiece.Player != g.CurrentPlayer && canBeat) {
                g.addPossibleMove(Move{From: pos, To: targetPos})
            }
        } else {
            g.addPossibleMove(Move{From: pos, To: OutOfGame})
        }
    }
}

