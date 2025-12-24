let playerChoice = document.getElementById("player-choice");
let opponentChoice = document.getElementById("opponent-choice");
let result = document.getElementById("result");
let wins = document.getElementById("wins");
let losses = document.getElementById("losses");
let ties = document.getElementById("ties");

wins.innerHTML = 0;
losses.innerHTML = 0;
ties.innerHTML = 0;

const choices = ["rock", "paper", "scissors"];
function matchup(userChoice) {
    const computerChoice = choices[Math.floor(Math.random() * choices.length)];
    opponentChoice.innerHTML = computerChoice;
    playerChoice.innerHTML = userChoice;
    if (userChoice === computerChoice) {
        result.innerHTML = "It's a tie!";
        ties.innerHTML++;
    }
    else if ((userChoice === "rock" && computerChoice === "scissors") ||
             (userChoice === "paper" && computerChoice === "rock") ||
             (userChoice === "scissors" && computerChoice === "paper")) {
        result.innerHTML = "You win!";
        wins.innerHTML++;
    }
    else {
        result.innerHTML = "You lose!";
        losses.innerHTML++;
    }

}

function reset() {
    playerChoice.innerHTML = "";
    computerChoice.innerHTML = "";
    result.innerHTML = "";
    wins.innerHTML = 0;
    losses.innerHTML = 0;
    ties.innerHTML = 0;
}