# go-tour
Tutorial presentation to introduce the Go programming language.

This is a work in progress.

Eventually I intend to write a set of lessons to teach Go programming.
For now, there is a “whirlwind tour” presentation intended to be given in about an hour
to provide a basic understanding of what Go is and why it would be worthwhile to continue
studying it. If needed, some sections can be removed or glossed over to reduce the time.

## Important Note on Content
These materials are *not* designed to be stand-alone training materials.
They are intended to be presented by a knowledgable instructor. As such, there is a lot that
the teacher should be expanding on in addition to what is written on the slides themselves.
Just reading the slides alone only conveys a fraction of the lesson.

## Go Version
The slides in the presentations assume access to the latest Go compiler. If you are using
an older version of Go, you may need to make some adjustments. For example,
  * The use of `range` to iterate over an integer interval was introduced in Go 1.22.
  * The `any` type as an alias for `interface{}` was introduced in Go 1.18.
  * The introduction of unique local index variables for each iteration of a loop was introduced in Go 1.22. Previous versions just changed the index variable values which could lead to unexpected behavior when closures inside the loop body were bound to the index variables.
