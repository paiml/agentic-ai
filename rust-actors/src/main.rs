use rust_actors::simple::{ping_pong_demo, Message};
use std::sync::mpsc;

fn main() {
    println!("🦀 Simple Rust Actor Demo");

    let (tx1, rx1) = mpsc::channel::<Message>();
    let (tx2, rx2) = mpsc::channel::<Message>();

    let messages = ping_pong_demo(tx1, rx1, tx2, rx2);

    println!("✅ Exchanged {} messages", messages.len());
    for (i, msg) in messages.iter().enumerate() {
        println!("{}: {:?}", i + 1, msg);
    }
}
