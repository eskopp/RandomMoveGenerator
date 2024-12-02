package main

import (
	"fmt"
	"github.com/notnil/chess"
	"math/rand"
	"time"
)

func generateRandomMoves(game *chess.Game, maxMoves int) ([]string, bool) {
	var blackMoves []string
	var whiteMoves []string

	// Schleife über maxMoves
	for i := 0; i < maxMoves; i++ {
		// Generiere alle legalen Züge
		legalMoves := game.LegalMoves()

		// Zufälligen Zug wählen
		move := legalMoves[rand.Intn(len(legalMoves))]

		// Züge als Notation aufzeichnen
		moveNotation := move.String()

		// Bestimme die Farbe des Spielers
		if game.Position().Turn() == chess.Black {
			blackMoves = append(blackMoves, moveNotation)
		} else {
			whiteMoves = append(whiteMoves, moveNotation)
		}

		// Führe den Zug aus
		game.Move(move)

		// Prüfen, ob Schachmatt erreicht wurde
		if game.IsCheckmate() {
			return append(blackMoves, whiteMoves...), true
		}
	}

	return append(blackMoves, whiteMoves...), false
}

func playGame(inputFEN string, maxAttempts int) {
	game, err := chess.NewGameFromFEN(inputFEN)
	if err != nil {
		fmt.Println("Fehler beim Erstellen des Spiels:", err)
		return
	}

	seenPositions := make(map[string]struct{})
	attemptCount := 0
	startTime := time.Now()

	// Maximal erlaubte Versuche, die Variante zu finden
	for {
		// Züge generieren
		moves, isCheckmate := generateRandomMoves(game, 6)

		// FEN der aktuellen Stellung
		positionFEN := game.Position().FEN()

		// Überprüfen, ob diese Stellung schon gesehen wurde
		if _, exists := seenPositions[positionFEN]; exists {
			game, _ = chess.NewGameFromFEN(inputFEN) // Zurücksetzen des Spiels
			continue
		}

		// Die Stellung hinzufügen
		seenPositions[positionFEN] = struct{}{}

		// Alle 1000 geprüften Varianten eine Nachricht ausgeben
		if attemptCount%1000 == 0 {
			fmt.Printf("Variante %d überprüft...\n", attemptCount)
		}

		// Schachmatt erreichen
		if isCheckmate {
			elapsedTime := time.Since(startTime)
			fmt.Printf("\x1b[32m%s\x1b[0m\n", moves)
			fmt.Println("\x1b[32mSchachmatt erreicht!\x1b[0m")
			fmt.Printf("\x1b[32mLaufzeit: %.2f Sekunden\x1b[0m\n", elapsedTime.Seconds())
			fmt.Printf("\x1b[32mAnzahl der geprüften Varianten: %d\x1b[0m\n", attemptCount)
			break
		}

		attemptCount++

		// Wenn die maximale Anzahl an Versuchen überschritten wurde
		if attemptCount > maxAttempts {
			fmt.Println("Maximale Anzahl an Varianten erreicht, keine Lösung gefunden.")
			break
		}

		game, _ = chess.NewGameFromFEN(inputFEN) // Zurücksetzen des Spiels
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Zufallsgenerator initialisieren

	// Beispiel-FEN für das Schachbrett
	inputFEN := "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"
	maxAttempts := 500000 // Maximale Anzahl an Varianten

	// Spiel starten
	playGame(inputFEN, maxAttempts)
}
