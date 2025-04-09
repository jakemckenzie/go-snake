# Snake in Go using Bubbletea

My go-to "hello world" in a language is [implementing Snake](https://en.wikipedia.org/wiki/Snake_(video_game_genre)). I've been playing around with a lot of functional languages off github. I think I may do Snake in F# or Ocaml next.

![First Attempt](https://github.com/jakemckenzie/go-snake/blob/main/docs/attempt-1.gif?raw=true)

## Notes 1:

I reached the channels section of the golang section of [boot.dev](https://www.boot.dev/courses/learn-golang). I thought I'd take a break from barreling through the coursework to stop and read documentation and actually build something. The go garbage collector is nice, I found the use of [`net/http/pprof`](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html) to be helpful. I don't use the differences in this project but it's [been exciting to read how pointers in Go differ from C](https://pkg.go.dev/unsafe).

>A uintptr is an integer, not a reference. Converting a Pointer to a uintptr creates an integer value with no pointer semantics. Even if a uintptr holds the address of some object, the garbage collector will not update that uintptr's value if the object moves, nor will that uintptr keep the object from being reclaimed.

Unrelated but it might as well go here. I was in an F# channel and [someone posted morphdom](https://www.npmjs.com/package/morphdom) which seems like a really useful npm package. I wish I had known about it years ago. Its benefits for me, deciding about the correctness of the DOM node tree seems fairly obvious. It was a problem that came about fairly regularly. That led me down a rabbit hole of learning about Idiomorph which has a similar approach. Here is a [good discussion between Micah Geisel and Delaney discussing Idiomorph](https://www.youtube.com/watch?v=IrtBBqyDrJU), it's cool to see very different types of engineers coming to grips with the same problem. "manipulating a dom that is stable is actually fast" is obvious to me coming from the react world, I hope they keep pushing this.
## TODOs:

- ~~**look into bubbletea to make the game look better**: emojis or shaders(?)~~
- ~~**ignore user input if already going a direction**: if snake is already going -> prevent <-, etc~~
- ~~**separate out from main**: main got fairly long, best to separate this out into different go files~~
- **add menu?**: I like [cli of life's menu](https://github.com/gabe565/cli-of-life/tree/main), probably [worth reading](https://medium.com/@nexidian/writing-an-interactive-cli-menu-in-golang-d6438b175fb6).
- ~~**TRS-80 CRT Effect**: never implemented this for snake, might be fun, might be outside the scope~~