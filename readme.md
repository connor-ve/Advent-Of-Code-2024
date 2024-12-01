## Advent of Code 2024

This year I will be using AoC to help me improve my skills with golang.

## Progression

| Day | Stars | Difficulty |
| --- | ----- | ---------- |
| 1   | ⭐⭐  | 🧩         |
| 2   |       |            |
| 3   |       |            |
| 4   |       |            |
| 5   |       |            |
| 6   |       |            |
| 7   |       |            |
| 8   |       |            |
| 9   |       |            |
| 10  |       |            |
| 11  |       |            |
| 12  |       |            |
| 13  |       |            |
| 14  |       |            |
| 15  |       |            |
| 16  |       |            |
| 17  |       |            |
| 18  |       |            |
| 19  |       |            |
| 20  |       |            |
| 21  |       |            |
| 22  |       |            |
| 23  |       |            |
| 24  |       |            |
| 25  |       |            |

## Day 1 - Learned

How to read file line by line

```go
// open the file
file, err := os.Open("input.txt")
if err != nil {
  os.Exit(0)
}
defer file.Close()

// Use a scanner to go line by line
scanner := bufio.NewScanner(file)
for scanner.Scan() {
  // inner logic by line
  line := scanner.Text() // Grab the line
  words := strings.Fields(line) // break down the words
}

if err := scanner.Err(); err != nil {
  fmt.Println(err)
}
```
