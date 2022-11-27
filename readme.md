# sokoban-go
sokoban-go is a CLI automated sokoban puzzles solver written in Go.

It provides easy to extend structures for implementing better searching algorithms
without having to design the whole program.

## Usage:

	sokoban-go [algorithm] [puzzlepath]

current algorithm options are:

	human
	dfs
	bfs
	astar-gh2
	astar-gh1
	astar-h1
	astar-h2
	hc-gh2
	hc-gh1
	hc-h2
	hc-h1

puzzlepath should be a suitable 2D JSON string array,
see "puzzles" folder for examples and [state/consts] for possible tiles characters.

## Example:
Input (I know the puzzle is not hard but I kept the name for some reason):

    sokoban-go astar-gh2 ./puzzles/hard.json

Output:
```
🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕🟩🟩⭕🟥
🟥⬜⬜⬜⬜🟥
🟥⬜🐈⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕🟩🟩⭕🟥
🟥⬜🐈⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕🟩🟩⭕🟥
🟥🐈⬜⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥😿🟩🟩⭕🟥
🟥⬜⬜⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥🐈⬜⬜⬜🟥
🟥⭕🟩🟩⭕🟥
🟥⬜⬜⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜🐈⬜⬜🟥
🟥⭕🟩🟩⭕🟥
🟥⬜⬜⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕🐈🟩⭕🟥
🟥⬜🟩⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕⬜🐈✅🟥
🟥⬜🟩⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕⬜⬜✅🟥
🟥⬜🟩🐈⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕⬜⬜✅🟥
🟥🟩🐈⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕⬜⬜✅🟥
🟥🟩⬜⬜⬜🟥
🟥⬜🐈⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥⭕⬜⬜✅🟥
🟥🟩⬜⬜⬜🟥
🟥🐈⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

🟥🟥🟥🟥🟥🟥
🟥⬜⬜⬜⬜🟥
🟥✅⬜⬜✅🟥
🟥🐈⬜⬜⬜🟥
🟥⬜⬜⬜⬜🟥
🟥🟥🟥🟥🟥🟥

YAY
Moves = 12
```

## License
[UNLICENSE](./UNLICENSE)