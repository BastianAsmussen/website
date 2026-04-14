# Projects

Things I've built. Roughly ordered by how interesting I find them.

---

## [dotfiles](https://github.com/BastianAsmussen/dotfiles)

**Nix.** My entire system configuration, declaratively managed with NixOS flakes.
Spans multiple machines, each with its own hardware profile. Full-disk encryption
via LUKS with FIDO2 token auth, secrets managed through sops-nix with age keys,
and a custom Neovim build. The kind of setup where reinstalling the OS takes ten
minutes and everything is exactly how you left it.

---

## [ROS](https://github.com/BastianAsmussen/ROS)

**Rust.** A hobby x86_64 kernel. `#![no_std]`, `#![no_main]`, custom target
spec, the works. The kernel initialises a full IDT covering every x86_64
exception (divide error through security exception) plus hardware interrupts for
the PIT timer, keyboard, and RTC, all using Rust's `extern "x86-interrupt"` ABI.

Memory is managed with a fixed-size block allocator (9 size classes, 8 to 2048
bytes) that falls back to a linked-list allocator for larger requests. Two
earlier designs, bump and plain linked-list, are still in the codebase.
Heap lives at 16 TiB virtual and is mapped at boot using the bootloader's
page table info.

The scheduler is a cooperative async executor built directly on Rust's
`Future`/`Poll`. Tasks are stored in a `BTreeMap`, woken by pushing IDs into a
lock-free `ArrayQueue`. `sleep_if_idle` does `enable_and_hlt` atomically to
avoid the interrupt-check race.

Also has an ATA disk driver, a VGA text buffer, CMOS RTC and PIT timekeeping,
and a FAT filesystem implementation with boot sector parsing, cluster chain
traversal, and directory entry reads. Tests run in QEMU, exiting via port `0xF4`
with results over serial.

---

## [fft-rs](https://github.com/BastianAsmussen/fft-rs)

**Rust.** An iterative Cooley-Tukey radix-2 FFT with a hand-rolled complex
number type. The core is a bit-reversal permutation followed by in-place
butterfly passes, no recursion, no heap allocation per call. Twiddle factors
are precomputed by incrementally multiplying a step complex number rather than
calling `cos`/`sin` in a loop, so the cache costs one trig evaluation regardless
of transform size. The `Complex` type uses `mul_add` for fused multiply-add in
the butterfly, and the butterfly loop is structured for cache locality.

While working on this I independently arrived at something I was calling "weird
half-numbers": numbers with two parts, like coordinates, where multiplying them
together rotates you instead of just scaling. Multiply one by itself and you end
up pointing in a completely different direction. A mathematician friend informed
me I had accidentally rediscovered the complex plane.

---

## [Lithium](https://github.com/BastianAsmussen/Lithium)

**Rust.** A programming language with Rust-inspired syntax (`fn`, `let`, `:` for
type annotations, `->` for return types). The pipeline runs lexer to recursive
descent parser to semantic analyzer. The parser produces a typed AST covering
all the usual statement forms: variable and function declarations, if/else,
while, C-style for, break/continue/return, and blocks. The semantic analyzer
walks the AST with a visitor, maintains a scope stack, and catches undefined
symbols, uninitialized variable reads, and assignments to function names.
Writing a language front-end forces you to actually understand how scoping rules
work rather than just using them.

---

## [RSE](https://github.com/BastianAsmussen/RSE)

**Rust.** A search engine from scratch: async crawler, Snowball stemmer, and a
PageRank-inspired ranker, backed by Postgres. Two services sharing a common crate.

The crawler runs three concurrent Tokio pipelines behind a control loop: one
channel feeds URLs to scrapers, scrapers emit `Website` items into a second
channel for processors, and a third channel carries discovered URLs back to the
control loop. A `Barrier` synchronises shutdown. Each domain's `robots.txt` is
fetched once and cached in a `RwLock<HashMap>`. The processor strips `<script>`
and `<style>` nodes from the DOM, collects body text, lowercases and de-punctuates
it, then stems every word with a Snowball stemmer selected by the page's
`<html lang="...">` attribute (15 languages). Stemmed word frequencies, forward
links with counts, and page metadata are written to Postgres via Diesel.

The search server stems the query the same way as indexing, looks up pages
with matching keyword stems, then scores them in two passes: a TF-style relevance
score (sum of `keyword_freq * query_term_freq`) and a PageRank pass that
accumulates weighted backlink scores. Results are sorted by the combined rank.
