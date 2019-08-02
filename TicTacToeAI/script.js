var origBoard;
const huPlayer = 'O';
const aiPlayer = 'X';
const winCombos = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [6, 4, 2]
]


const cells = document.querySelectorAll('.cell');
startGame();


// Resets display at start of the game
function startGame() {
    // At the start of every game, restart the display
    document.querySelector(".endgame").style.display = 'none';
    // Create Array of length 9, with #s 0-8
    origBoard = Array.from(Array(9).keys());
    for (var i = 0; i < cells.length; i++) {
        cells[i].innerText = '';
        cells[i].style.removeProperty('background-color');
        cells[i].addEventListener('click', turnClick, false);
    }
}


// Passes clicked square and player to turn function
function turnClick(square) {
    // If id == number, human nor AI has occupied that spot
    if (typeof origBoard[square.target.id] == 'number') {
        turn(square.target.id, huPlayer);
        // After human players move, AI makes a move
        if (!checkTie()) turn(bestSpot(), aiPlayer);
    }
    // If clicked spot is occupied, do nothing
}


// Holds array indicating which player placed an X or an 0 and in what spot
function turn(squareID, player) {
    origBoard[squareID] = player;
    document.getElementById(squareID).innerText = player;
    let gameWon = checkWin(origBoard, player);
    if (gameWon) gameOver(gameWon)
}


// Checks if the board holds a winning player
// a = accumulator, value returned at end
// e = element in board array being traversed
// i = index
function checkWin(board, player) {
    let plays = board.reduce((a, e, i) =>
    (e === player) ? a.concat(i) : a, []);
    let gameWon = null;
    // Win allows access to the indexed array
    for (let [index, win] of winCombos.entries()) {
        // Checks if the player has all three spots occupied in a winning array
        if (win.every(elem => plays.indexOf(elem) > -1)) {
            gameWon = {index: index, player: player}
            break;
        }
    }
    return gameWon;
}


function gameOver(gameWon) {
    // If human player wins, set bg to blue, if AI wins set bg to red
    for (let index of winCombos[gameWon.index]) {
        document.getElementById(index).style.backgroundColor = gameWon.player == huPlayer ? "blue" : "red";
    }
    // Turn off eventListener so that further clicks do not trigger anything
    for (var i = 0; i < cells.length; i++) {
        cells[i].removeEventListener('click', turnClick, false);
    }
    declareWinner(gameWon.player == huPlayer ? "You win!" : "You lose.");
}


function declareWinner(who) {
    document.querySelector(".endgame").style.display = "block";
    document.querySelector(".endgame .text").innerText = who;
}


// Finds all spots in the cross that are not occupied
function emptySquares() {
    return origBoard.filter(s => typeof s == 'number')
}


// Calls minimax function to find best spot for AI 
function bestSpot() {
    return minimax(origBoard, aiPlayer).index;
}


// If length is 0, every square is occupied, therefore a tie
function checkTie() {
    if (emptySquares().length == 0) {
        for (var i = 0; i < cells.length; i++) {
            cells[i].style.backgroundColor = "green";
            cells[i].removeEventListener('click', turnClick, false);
        }
        declareWinner("Tie Game!");
        return true;
    }
    return false;
}


function minimax(newBoard, player) {
    var availSpots = emptySquares();

    // If human wins, return -10
    if (checkWin(newBoard, huPlayer)) {
        return {score: -10};
    // If AI wins, return 10
    } else if (checkWin(newBoard, aiPlayer)){
        return {score: 10};
    // If tie, return 0
    } else if (availSpots.length === 0) {
        return {score: 0}
    }

    // Collect available spots in move array
    var moves = [];
    for (var i = 0; i < availSpots.length; i++){
        var move = {};
        move.index = newBoard[availSpots[i]];
        // Set empty spot on newBoard to cur player
        newBoard[availSpots[i]] = player;

        // Call minimax function with other player and newBoard
        if (player == aiPlayer) {
            var result = minimax(newBoard, huPlayer);
            move.score = result.score
        } else {
            var result = minimax(newBoard, aiPlayer);
            move.score = result.score
        }

        // Reset board to init state
        newBoard[availSpots[i]] = move.index;

        // Pushes move object to moves array
        moves.push(move);
    }

    // Evaluates the best move in the move array
    var bestMove;
    // Algorithm goes through all possible moves for AI, while storing the best move
    if (player === aiPlayer) {
        var bestScore = -10000;
        for (var i = 0; i < moves.length; i++) {
            if (moves[i].score > bestScore) {
                bestScore = moves[i].score;
                bestMove = i;
            }
        }
    // Same evaluation for human player, however looks for move with lowest score to store
    } else {
        var bestScore = 10000;
        for (var i = 0; i < moves.length; i++) {
            if (moves[i].score < bestScore) {
                bestScore = moves[i].score;
                bestMove = i;
            }
        }
    }

    return moves[bestMove];
}