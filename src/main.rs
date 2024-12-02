use chess::{Board, Game, MoveGen};
use rand::seq::SliceRandom;
use std::collections::HashSet;
use std::time::Instant;

fn generate_random_moves(game: &mut Game, max_moves: usize) -> (Vec<String>, bool) {
    let mut black_moves = Vec::new();
    let mut white_moves = Vec::new();

    for _ in 0..max_moves {
        // Verwende eine unveränderliche Referenz auf das Schachbrett
        let legal_moves: Vec<_> = MoveGen::new_legal(&game.board()).collect();

        if let Some(random_move) = legal_moves.choose(&mut rand::thread_rng()) {
            let move_notation = game.board().san(*random_move);
            if game.board().turn().is_black() {
                black_moves.push(move_notation);
            } else {
                white_moves.push(move_notation);
            }
            game.play(*random_move).unwrap();

            if game.board().is_checkmate() {
                return (black_moves, true);
            }
        }
    }
    (black_moves, game.board().is_checkmate())
}

fn play_game(input_fen: &str) {
    let mut seen_positions = HashSet::new(); // Zum Speichern der bereits geprüften Schachstellungen
    let mut game = Game::new_from_fen(input_fen).unwrap();
    let mut attempt_count = 0;

    let start_time = Instant::now();

    loop {
        let (black_moves, is_checkmate) = generate_random_moves(&mut game, 6); // Maximal 6 Züge

        let position_fen = game.board().to_string();

        // Prüfen, ob diese FEN-Position bereits in der HashMap aufgetreten ist
        if seen_positions.contains(&position_fen) {
            game = Game::new_from_fen(input_fen).unwrap(); // Zurücksetzen des Spiels
            continue;
        }

        // Die Stellung zur HashSet hinzufügen
        seen_positions.insert(position_fen);

        // Alle 1000 geprüften Varianten eine Nachricht ausgeben
        if attempt_count % 1000 == 0 {
            println!("Variante {} überprüft...", attempt_count);
        }

        // Prüfen, ob Schachmatt erreicht wurde
        if is_checkmate {
            let elapsed_time = start_time.elapsed();
            // Ausgabe der Züge und des Ergebnisses in grün
            println!("\x1b[32m{}\x1b[0m", black_moves.join(" "));
            println!("\x1b[32mSchachmatt erreicht!\x1b[0m");
            println!("\x1b[32mLaufzeit: {:.2} Sekunden\x1b[0m", elapsed_time.as_secs_f32());
            println!("\x1b[32mAnzahl der geprüften Varianten: {}\x1b[0m", attempt_count);
            break; // Beende das Spiel, wenn Schachmatt erreicht wurde
        }

        attempt_count += 1;
        game = Game::new_from_fen(input_fen).unwrap(); // Zurücksetzen des Spiels
    }
}

fn main() {
    let input_fen = "8/4N3/3k4/8/2R1n3/8/N5K1/8 b - - 0 1"; // Beispiel FEN
    play_game(input_fen);
}
