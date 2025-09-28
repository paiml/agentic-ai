import { simplePingPong } from "./simple.ts";

async function main(): Promise<void> {
  console.log("🦕 Simple Deno Actor Demo");

  const messages = await simplePingPong();

  console.log(`✅ Exchanged ${messages.length} messages`);
  messages.forEach((msg, i) => {
    console.log(`${i + 1}: ${JSON.stringify(msg)}`);
  });
}

if (import.meta.main) {
  main();
}
