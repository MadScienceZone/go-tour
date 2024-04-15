# Presentation #0: Whirlwind Tour of Go
This 30-40 minute lesson is intended to introduce Go to an audience of experienced developers. It very quickly shows what Go syntax is like, quickly glossing over how it's already similar to other C-derived languages but slowing down to cover where it extends or diverts from that base.
It then goes on to cover in more detail Go's more unique aspects, highlighting the concepts of
  * The Go ecosystem, including automatic documentation generation and importing 3rd party packages 
  * Factored notation
  * Control flow, including compound `if`, the various forms of `for`,  and `switch`
  * Arrays and slices
  * Object orientation, including composition, methods, and polymorphism via interfaces
  * Concurrency with goroutines and channels
  * Error handling
  * Cross-thread memory access and the difference between Go-idiomatic approaches and more traditional ones
  * Contexts
  * JSON support

Code samples are provided in the `examples` directory. They are also loaded into the Go playground. Links to the playground code appear at the top of the corresponding slides in case students want to try playing with the examples live during the class.

## Timing
During a practice run of the material for this presentation, trying to keep the pace brisk (especially for the early bits where we're laying the foundation of basic Go syntax and such), the timing for when we started each major section turned out thusly:
  00:00 Intro
  03:52 Basics, Intrinsic Types
  12:13 Ecosystem
  14:38 Factored Notation
  18:49 Arrays
  25:14 Structs
  38:58 Goroutines
  45:09 Error Handling
  46:00 ID Generator Example
  51:00 Contexts
  55:40 JSON
  58:15 End
