# Thread Mob Storage Vault Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 177 | ship |
| stress | sync drift | 193 | ship |
| edge | local state | 149 | ship |
| recovery | conflict cost | 223 | ship |
| stale | form pressure | 197 | ship |

Start with `recovery` and `edge`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

`recovery` is the optimistic case; use it to make sure the scoring path still rewards strong signal.
