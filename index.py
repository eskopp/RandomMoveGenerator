import chess
import random

input_fen = "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"
start_player = "black"


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


if start_player.lower() == "black" and input_fen.split(" ")[1] == "w":
    parts = input_fen.split(" ")
    parts[1] = "b"
    input_fen = " ".join(parts)
elif start_player.lower() == "white" and input_fen.split(" ")[1] == "b":
    parts = input_fen.split(" ")
    parts[1] = "w"
    input_fen = " ".join(parts)


move_sequence = generate_random_moves(input_fen, num_moves=6)

board = chess.Board(input_fen)
for move in move_sequence:
    print(board.san(move))
    board.push(move)
