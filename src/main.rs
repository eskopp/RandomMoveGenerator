use chess::{Board, ChessMove, Game, MoveGen};
use rand::seq::SliceRandom;
use std::collections::HashSet;
use std::time::Instant;

fn generate_random_moves(game: &mut Game, max_moves: usize) -> (Vec<String>, bool) {
    let mut black_moves = Vec::new();
    let mut white_moves = Vec::new();

    // Wir können die Anzahl der geprüften Züge begrenzen (max_moves), um die Laufzeit zu steuern.
    let board = game.current_position();
    let legal_moves: Vec<ChessMove> = MoveGen::new_legal(&board).collect();

    for _ in 0..max_moves {
        if let Some(random_move) = legal_moves.choose(&mut rand::thread_rng()) {
            let move_notation = random_move.to_string();  // Züge als Notation
            if board.side_to_move() == chess::Color::Black {
                black_moves.push(move_notation);
            } else {
                white_moves.push(move_notation);
            }
            game.make_move(*random_move);
            if board.status() == chess::BoardStatus::Checkmate {
                return (black_moves, true);  // Schachmatt erreicht
            }
        }
    }

    (black_moves, game.current_position().status() == chess::BoardStatus::Checkmate)
}

fn play_game(input_fen: &str, max_attempts: usize) {
    let mut seen_positions = HashSet::new();
    let mut game = Game::new_from_fen(input_fen).unwrap();
    let mut attempt_count = 0;

    let start_time = Instant::now();

    loop {
        let (black_moves, is_checkmate) = generate_random_moves(&mut game, 6);

        let position_fen = game.current_position().to_string();

        if seen_positions.contains(&position_fen) {
            game = Game::new_from_fen(input_fen).unwrap();
            continue;
        }

        seen_positions.insert(position_fen);

        if attempt_count % 1000 == 0 {
            println!("Variante {} überprüft...", attempt_count);
        }

        if is_checkmate {
            let elapsed_time = start_time.elapsed();
            println!("\x1b[32m{}\x1b[0m", black_moves.join(" "));
            println!("\x1b[32mSchachmatt erreicht!\x1b[0m");
            println!("\x1b[32mLaufzeit: {:.2} Sekunden\x1b[0m", elapsed_time.as_secs_f32());
            println!("\x1b[32mAnzahl der geprüften Varianten: {}\x1b[0m", attempt_count);
            break;
        }

        attempt_count += 1;
        if attempt_count > max_attempts {
            println!("Maximale Anzahl an Varianten erreicht, keine Lösung gefunden.");
            break;  // Max. Versuche erreicht
        }

        game = Game::new_from_fen(input_fen).unwrap();
    }
}

fn main() {
    let input_fen = "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"; // Beispiel FEN
    let max_attempts = 500000;  // Setze ein realistisches Limit für die Anzahl der Varianten
    play_game(input_fen, max_attempts);
}
