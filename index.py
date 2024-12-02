import chess
import random
import time

# Eingabe der Startstellung (FEN)
input_fen = "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"
# input_fen = "6k1/4r3/8/8/8/8/4qPPP/2R3K1 b - - 0 1"

def generate_random_moves(board: chess.Board, max_moves: int = 6):
    """
    Generiert eine Liste von zufälligen, legalen Zügen für das gegebene Schachbrett.

    Args:
        board (chess.Board): Das aktuelle Schachbrett.
        max_moves (int): Die maximale Anzahl der Züge, die generiert werden sollen (standardmäßig 6).

    Returns:
        move_sequence (list): Liste der Züge in SAN-Notation.
        is_checkmate (bool): Ob Schachmatt erreicht wurde.
    """
    move_sequence = []

    for _ in range(max_moves):
        if board.is_game_over():
            break
        legal_moves = list(board.legal_moves)
        move = random.choice(legal_moves)
        move_notation = board.san(move)
        move_sequence.append(move_notation)
        board.push(move)

        if board.is_checkmate():
            return move_sequence, True

    return move_sequence, board.is_checkmate()


def play_game(input_fen: str):
    """
    Spielt Varianten von Zügen, bis Schachmatt erreicht wird, ohne dass eine Zugsequenz doppelt vorkommt.

    Args:
        input_fen (str): Die Startstellung in FEN-Notation.
    """
    board_initial = chess.Board(input_fen)
    seen_sequences = set()  # Set (HashMap) zum Speichern der bereits gesehenen FEN-Strings (Schachstellungen)
    attempt_count = 0  # Zähler für die Anzahl der Versuche

    start_time = time.time()  # Starte die Zeitmessung

    while True:
        # Erstelle eine Kopie des initialen Schachbretts
        board = chess.Board(input_fen)

        # Generiere zufällige Zugsequenz
        move_sequence, is_checkmate = generate_random_moves(board)

        attempt_count += 1  # Erhöhe die Anzahl der Versuche

        # Konvertiere die Zugsequenz zu einem FEN-String
        position_fen = board.fen()

        # Prüfen, ob diese FEN-Position bereits in der HashMap aufgetreten ist
        if position_fen in seen_sequences:
            continue  # Diese Variante überspringen

        # Füge die neue FEN-Position zur HashMap (Set) hinzu
        seen_sequences.add(position_fen)

        # Alle 1000 geprüften Varianten eine Nachricht ausgeben
        if attempt_count % 1000 == 0:
            print(f"Variante {attempt_count} überprüft...")

        # Prüfen, ob Schachmatt erreicht wurde
        if is_checkmate:
            elapsed_time = time.time() - start_time  # Berechne die Zeit, die seit dem Start vergangen ist
            # Ausgabe der Züge und des Ergebnisses in grün
            print(f"\033[32m{' '.join(move_sequence)}\033[0m")  # Grüne Ausgabe der Züge
            print("\033[32mSchachmatt erreicht!\033[0m")  # Grüne Ausgabe für Schachmatt
            print(f"\033[32mLaufzeit: {elapsed_time:.2f} Sekunden\033[0m")  # Grüne Ausgabe der Laufzeit
            print(f"\033[32mAnzahl der geprüften Varianten: {attempt_count}\033[0m")  # Grüne Ausgabe der Varianten
            break  # Beende das Spiel, wenn Schachmatt erreicht wurde


# Das Spiel starten
play_game(input_fen)