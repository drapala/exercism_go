# Pov

Welcome to Pov on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Reparent a graph on a selected node.

This exercise is all about re-orientating a graph to see things from a different
point of view. For example family trees are usually presented from the
ancestor's perspective:

```text
    +------0------+
    |      |      |
  +-1-+  +-2-+  +-3-+
  |   |  |   |  |   |
  4   5  6   7  8   9
```

But the same information can be presented from the perspective of any other node
in the graph, by pulling it up to the root and dragging its relationships along
with it. So the same graph from 6's perspective would look like:

```text
        6
        |
  +-----2-----+
  |           |
  7     +-----0-----+
        |           |
      +-1-+       +-3-+
      |   |       |   |
      4   5       8   9
```

This lets us more simply describe the paths between two nodes. So for example
the path from 6-9 (which in the first graph goes up to the root and then down to
a different leaf node) can be seen to follow the path 6-2-0-3-9

This exercise involves taking an input graph and re-orientating it from the point
of view of one of the nodes.

## Implementation Notes

POV / reparent / change root of a tree

The type name is Graph because you'll probably be implementing a general
directed graph representation, although the test program will only use
it to create a tree.  The term "arc" is used here to mean a directed edge.

The test program will create a graph with New, then use AddNode to add
leaf nodes.  After that it will use AddArc to construct the rest of the tree
from the bottom up.  That is, the `to` argument will always specify a node
that has already been added.

ArcList is a dump method to let the test program see your graph.  It must
return a list of all arcs in the graph.  Format each arc as a single string
like "from -> to". The test program can then easily sort the list and
compare it to an expected result.  You do not need to bother with sorting
the list yourself.

All this graph construction and dumping must be working before you start
on the interesting part of the exercise, so it is tested separately as
a first test.

API function ChangeRoot does the interesting part of the exercise.
OldRoot is passed (as a convenience) and you must return a graph with
newRoot as the root.  You can modify the original graph in place and
return it or create a new graph and return that.  If you return a new
graph you are free to consume or destroy the original graph.  Of course
it's nice to leave it unmodified.

## Source

### Created by

- @skeys

### Contributed to by

- @alebaffa
- @bitfield
- @ekingery
- @ferhatelmas
- @hilary
- @kytrinyx
- @leenipper
- @petertseng
- @robphoenix
- @sebito91
- @soniakeys
- @tleen

### Based on

Adaptation of exercise from 4clojure - https://www.4clojure.com/