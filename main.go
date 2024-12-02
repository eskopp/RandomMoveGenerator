package main

import (
	"fmt"
	"github.com/notnil/chess"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Funktion, um zufällige Züge zu generieren und eine Folge von 6 Zügen (3 Schwarz, 3 Weiß) zu prüfen
func generateRandomMoveSequence(game *chess.Game) ([]string, error) {
	var moves []string

	// Schwarz zieht 3 Mal
	for i := 0; i < 3; i++ {
		legalMoves := game.ValidMoves()
		if len(legalMoves) == 0 {
			return nil, fmt.Errorf("keine legalen Züge verfügbar für Schwarz")
		}

		move := legalMoves[rand.Intn(len(legalMoves))]
		moves = append(moves, move.String()) // Zug als Notation speichern
		if err := game.Move(move); err != nil {
			return nil, fmt.Errorf("fehler beim Ausführen eines Zugs für Schwarz: %v", err)
		}
	}

	// Weiß zieht 3 Mal
	for i := 0; i < 3; i++ {
		legalMoves := game.ValidMoves()
		if len(legalMoves) == 0 {
			return nil, fmt.Errorf("keine legalen Züge verfügbar für Weiß")
		}

		move := legalMoves[rand.Intn(len(legalMoves))]
		moves = append(moves, move.String()) // Zug als Notation speichern
		if err := game.Move(move); err != nil {
			return nil, fmt.Errorf("fehler beim Ausführen eines Zugs für Weiß: %v", err)
		}

		// Nach dem dritten Zug von Weiß prüfen, ob Schachmatt erreicht ist
		if /*i == 2 && */ game.Outcome() == chess.WhiteWon {
			fmt.Printf("Hallo Welt")
			os.Exit(0)
			return moves, nil
		}
	}

	return moves, fmt.Errorf("kein Schachmatt nach 6 Zügen gefunden")
}

// Hauptfunktion für das Spiel
func playGame(inputFEN string, maxAttempts int) {
	position, err := chess.FEN(inputFEN)
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Position aus der FEN:", err)
		return
	}
	game := chess.NewGame(position)

	seenVariants := make(map[string]struct{}) // Speichert bereits geprüfte Varianten
	attemptCount := 0
	uniqueVariants := 0
	startTime := time.Now()

	// Schleife für die maximale Anzahl an Versuchen
	for {
		moves, err := generateRandomMoveSequence(game)
		attemptCount++

		// Konvertiere die Zugfolge in einen eindeutigen Schlüssel
		variantKey := strings.Join(moves, " ")

		// Prüfen, ob diese Variante bereits geprüft wurde
		if _, exists := seenVariants[variantKey]; exists {
			// Variante wurde bereits geprüft, überspringen
			game = chess.NewGame(position)
			continue
		}

		// Variante ist neu, speichern und zählen
		seenVariants[variantKey] = struct{}{}
		uniqueVariants++
		if uniqueVariants%10000 == 0 {
			fmt.Printf("Lebenszeichen: %d Varianten geprüft...\n", uniqueVariants)
		}
		// Ausgabe der Variante
		// fmt.Printf("Variante %d: %v\n", uniqueVariants, moves)
		if err == nil {
			// Erfolgreiche Variante gefunden
			elapsedTime := time.Since(startTime)
			fmt.Printf("\x1b[32mSchachmatt erreicht!\x1b[0m\n")
			fmt.Printf("\x1b[32mZüge: %v\x1b[0m\n", moves)
			fmt.Printf("\x1b[32mLaufzeit: %.2f Sekunden\x1b[0m\n", elapsedTime.Seconds())
			fmt.Printf("\x1b[32mAnzahl der geprüften Varianten: %d\x1b[0m\n", uniqueVariants)
			break
		}
		// Wenn die maximale Anzahl an Versuchen überschritten wurde
		if uniqueVariants >= maxAttempts {
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
	maxAttempts := 9999999999 // Maximale Anzahl an Varianten

	// Spiel starten
	playGame(inputFEN, maxAttempts)
}
