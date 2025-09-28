// TDD: Test written FIRST according to specification
use rust_actors::simple::{ping_pong_demo, Message};
use std::sync::mpsc;
use std::time::Duration;

#[test]
fn test_three_round_ping_pong() {
    // GIVEN: Two actors for ping-pong exchange
    let (tx1, rx1) = mpsc::channel::<Message>();
    let (tx2, rx2) = mpsc::channel::<Message>();

    // WHEN: Running the ping-pong demo
    let messages = ping_pong_demo(tx1, rx1, tx2, rx2);

    // THEN: Exactly 6 messages exchanged (3 pings, 3 pongs)
    assert_eq!(messages.len(), 6);
    assert_eq!(
        messages
            .iter()
            .filter(|m| matches!(m, Message::Ping(_)))
            .count(),
        3
    );
    assert_eq!(
        messages
            .iter()
            .filter(|m| matches!(m, Message::Pong(_)))
            .count(),
        3
    );
}

#[test]
fn test_message_ordering() {
    // GIVEN: Two actors
    let (tx1, rx1) = mpsc::channel::<Message>();
    let (tx2, rx2) = mpsc::channel::<Message>();

    // WHEN: Running the demo
    let messages = ping_pong_demo(tx1, rx1, tx2, rx2);

    // THEN: Messages alternate ping-pong
    for (i, msg) in messages.iter().enumerate() {
        if i % 2 == 0 {
            assert!(
                matches!(msg, Message::Ping(_)),
                "Expected Ping at position {}",
                i
            );
        } else {
            assert!(
                matches!(msg, Message::Pong(_)),
                "Expected Pong at position {}",
                i
            );
        }
    }
}

#[test]
fn test_deterministic_behavior() {
    // GIVEN: Multiple runs with same setup
    let run1 = {
        let (tx1, rx1) = mpsc::channel();
        let (tx2, rx2) = mpsc::channel();
        ping_pong_demo(tx1, rx1, tx2, rx2)
    };

    let run2 = {
        let (tx1, rx1) = mpsc::channel();
        let (tx2, rx2) = mpsc::channel();
        ping_pong_demo(tx1, rx1, tx2, rx2)
    };

    // THEN: Identical behavior
    assert_eq!(run1.len(), run2.len());
    for (msg1, msg2) in run1.iter().zip(run2.iter()) {
        match (msg1, msg2) {
            (Message::Ping(n1), Message::Ping(n2)) => assert_eq!(n1, n2),
            (Message::Pong(n1), Message::Pong(n2)) => assert_eq!(n1, n2),
            _ => panic!("Message types don't match"),
        }
    }
}

#[test]
fn test_performance() {
    use std::time::Instant;

    // GIVEN: Performance requirement of <10ms
    let (tx1, rx1) = mpsc::channel();
    let (tx2, rx2) = mpsc::channel();

    // WHEN: Measuring execution time
    let start = Instant::now();
    ping_pong_demo(tx1, rx1, tx2, rx2);
    let elapsed = start.elapsed();

    // THEN: Complete within 10ms
    assert!(
        elapsed < Duration::from_millis(10),
        "Took {:?}, expected <10ms",
        elapsed
    );
}
