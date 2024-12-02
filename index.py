import chess
import random

input_fen = "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"


def generate_random_moves(fen: str):
    board = chess.Board(fen)
    black_moves = []
    white_moves = []

    # Generiere 6 Züge (3 für Schwarz und 3 für Weiß)
    for _ in range(6):
        legal_moves = list(board.legal_moves)
        move = random.choice(legal_moves)
        move_notation = board.san(move)

        if board.turn == chess.BLACK:
            black_moves.append(move_notation)
        else:
            white_moves.append(move_notation)

        board.push(move)

    return black_moves, white_moves, board.is_checkmate()


while True:
    black_moves, white_moves, is_checkmate = generate_random_moves(input_fen)

    # Ausgabe der Züge in der Reihenfolge Schwarz - Weiß
    output = black_moves + white_moves
    print(" ".join(output))

    if is_checkmate:
        print("\nSchachmatt erreicht!")
        break
    else:
        print("\nKein Schachmatt. ...\n")
