package ur;

const mastersNumDices = 3;
const mastersNumPieces = 7;

var mastersPaths = map[Player][]uint {
    White: { 19, 18, 17, 16, 8, 9, 10, 11, 12, 13, 14, 6, 7, 15, 23, 22, },
    Black: { 3, 2, 1, 0, 8, 9, 10, 11, 12, 13, 14, 22, 23, 15, 7, 6, },
}
