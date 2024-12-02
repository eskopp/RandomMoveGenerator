package main

import (
	"fmt"
	"github.com/notnil/chess"
	"math/rand"
	"time"
)

// Funktion, um zufällige Züge zu generieren und eine Folge von 4 Zügen (3 von Schwarz und 1 von Weiß) zu erstellen
func generateRandomMoveSequence(game *chess.Game) ([]string, error) {
	var moves []string

	// Schwarz zieht 3 Mal
	for i := 0; i < 3; i++ {
		// Generiere alle legalen Züge
		legalMoves := game.ValidMoves()

		if len(legalMoves) == 0 {
			return nil, fmt.Errorf("keine legalen Züge verfügbar")
		}

		// Zufälligen Zug wählen
		move := legalMoves[rand.Intn(len(legalMoves))]
		moves = append(moves, move.String()) // Zug als Notation speichern

		// Führe den Zug aus
		err := game.Move(move)
		if err != nil {
			return nil, fmt.Errorf("fehler beim Ausführen eines Zugs: %v", err)
		}
	}

	// Weiß zieht 1 Mal
	legalMoves := game.ValidMoves()
	if len(legalMoves) == 0 {
		return nil, fmt.Errorf("keine legalen Züge verfügbar")
	}
	move := legalMoves[rand.Intn(len(legalMoves))]
	moves = append(moves, move.String()) // Zug als Notation speichern

	// Führe den Zug aus
	err := game.Move(move)
	if err != nil {
		return nil, fmt.Errorf("fehler beim Ausführen eines Zugs: %v", err)
	}

	return moves, nil
}

// Hauptfunktion für das Spiel
func playGame(inputFEN string, maxAttempts int) {
	// Erstelle das Spiel aus der FEN-Zeichenkette
	position, err := chess.FEN(inputFEN)
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Position aus der FEN:", err)
		return
	}
	game := chess.NewGame(position)

	seenPositions := make(map[string]struct{})
	attemptCount := 0
	startTime := time.Now()

	// Schleife für die maximale Anzahl an Versuchen
	for {
		// Generiere Züge
		moves, err := generateRandomMoveSequence(game)
		if err != nil {
			fmt.Println("Fehler beim Generieren der Züge:", err)
			break
		}

		// Speichere die aktuelle Stellung in FEN
		positionFEN := game.Position().String()

		// Überprüfen, ob diese Stellung schon gesehen wurde
		if _, exists := seenPositions[positionFEN]; exists {
			game = chess.NewGame(position) // Setze das Spiel zurück
			continue
		}

		// Füge die aktuelle Stellung hinzu
		seenPositions[positionFEN] = struct{}{}

		// Gib alle 1000 geprüften Varianten aus
		if attemptCount%1000 == 0 {
			fmt.Printf("Variante %d überprüft...\n", attemptCount)
		}

		// Überprüfe, ob Schachmatt erreicht wurde
		if game.Outcome() == chess.WhiteWon || game.Outcome() == chess.BlackWon {
			elapsedTime := time.Since(startTime)
			fmt.Printf("\x1b[32mZüge: %s\x1b[0m\n", moves)
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

		// Spiel zurücksetzen
		game = chess.NewGame(position)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Zufallsgenerator initialisieren

	// Beispiel-FEN
	inputFEN := "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"
	maxAttempts := 500000 // Maximale Anzahl an Varianten

	// Spiel starten
	playGame(inputFEN, maxAttempts)
}
