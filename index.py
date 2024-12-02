import chess
import random

input_fen = "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"

def generate_random_moves(fen: str, num_moves: int = 6):
    board = chess.Board(fen)
    moves = []
    for _ in range(num_moves):
        if board.is_game_over():
            break
        move = random.choice(list(board.legal_moves))
        moves.append(move)
        board.push(move)
    return moves

move_sequence = generate_random_moves(input_fen, num_moves=6)
Specification of the color selection via the FEN
board = chess.Board(input_fen)
for i, move in enumerate(move_sequence, 1):
    print(f"Move {i}: {board.san(move)}")
    board.push(move)
