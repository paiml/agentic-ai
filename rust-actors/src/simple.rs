// Simple 3-round ping-pong implementation per specification
use std::sync::mpsc::{Receiver, Sender};
use std::thread;
use std::time::Duration;

#[derive(Debug, Clone)]
pub enum Message {
    Ping(u32),
    Pong(u32),
}

// Run ping actor - sends 3 pings, receives 3 pongs
fn run_ping_actor(pong_tx: Sender<Message>, ping_rx: Receiver<Message>) {
    for round in 1..=3 {
        println!("Ping: Sending round {}", round);
        pong_tx.send(Message::Ping(round)).unwrap();

        if let Ok(Message::Pong(n)) = ping_rx.recv_timeout(Duration::from_millis(5)) {
            println!("Ping: Received pong {}", n);
        }
    }
}

// Run pong actor - receives 3 pings, sends 3 pongs
fn run_pong_actor(ping_tx: Sender<Message>, pong_rx: Receiver<Message>) -> Vec<Message> {
    let mut pongs = Vec::new();
    for _ in 1..=3 {
        if let Ok(Message::Ping(n)) = pong_rx.recv_timeout(Duration::from_millis(5)) {
            println!("Pong: Received ping {}", n);

            let pong_msg = Message::Pong(n);
            pongs.push(pong_msg.clone());
            ping_tx.send(pong_msg).unwrap();
            println!("Pong: Sent pong {}", n);
        }
    }
    pongs
}

// Collect messages in alternating order
fn collect_messages(pongs: Vec<Message>) -> Vec<Message> {
    let mut messages = Vec::new();
    for i in 1..=3u32 {
        messages.push(Message::Ping(i));
        if let Some(pong) = pongs.get((i - 1) as usize) {
            messages.push(pong.clone());
        }
    }
    messages
}

// Main demo function - orchestrates the 3-round exchange
pub fn ping_pong_demo(
    ping_tx: Sender<Message>,
    ping_rx: Receiver<Message>,
    pong_tx: Sender<Message>,
    pong_rx: Receiver<Message>,
) -> Vec<Message> {
    // Spawn ping actor
    let ping_handle = thread::spawn(move || {
        run_ping_actor(pong_tx, ping_rx);
    });

    // Spawn pong actor
    let pong_handle = thread::spawn(move || run_pong_actor(ping_tx, pong_rx));

    // Wait for completion and collect results
    ping_handle.join().unwrap();
    let pongs = pong_handle.join().unwrap();

    collect_messages(pongs)
}
