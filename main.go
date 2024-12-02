package main

import (
	"fmt"
	"github.com/notnil/chess"
	"math/rand"
	"time"
)

// Funktion, um zufällige Züge zu generieren und eine Folge von 4 Zügen (3 von Schwarz und 1 von Weiß) zu geben
func generateRandomMoveSequence(game *chess.Game) ([]string, error) {
	var moves []string

	// Schwarz zieht 3 Mal
	for i := 0; i < 3; i++ {
		// Generiere alle legalen Züge
		legalMoves := game.LegalMoves()
		if len(legalMoves) == 0 {
			return nil, fmt.Errorf("keine legalen Züge verfügbar")
		}

		// Zufälligen Zug wählen
		move := legalMoves[rand.Intn(len(legalMoves))]
		moves = append(moves, move.String()) // Zug als Notation speichern

		// Führe den Zug aus
		game.Move(move)
	}

	// Weiß zieht 1 Mal
	legalMoves := game.LegalMoves()
	if len(legalMoves) == 0 {
		return nil, fmt.Errorf("keine legalen Züge verfügbar")
	}
	move := legalMoves[rand.Intn(len(legalMoves))]
	moves = append(moves, move.String()) // Zug als Notation speichern

	// Führe den Zug aus
	game.Move(move)

	return moves, nil
}

func main() {
	// Zufallsgenerator initialisieren
	rand.Seed(time.Now().UnixNano())

	// Beispiel-FEN für das Schachbrett (eine beliebige Stellung)
	inputFEN := "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"

	// Spiel aus der FEN-Zeichenkette erstellen
	game, err := chess.NewGame(chess.FEN(inputFEN))
	if err != nil {
		fmt.Println("Fehler beim Erstellen des Spiels:", err)
		return
	}

	// Zufällige Zugfolge generieren
	moves, err := generateRandomMoveSequence(game)
	if err != nil {
		fmt.Println("Fehler beim Generieren der Züge:", err)
		return
	}

	// Ausgabe der Züge
	fmt.Println("Zugfolge:", moves)
}
