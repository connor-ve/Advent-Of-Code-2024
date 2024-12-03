## Advent of Go (code-2024)

This year I will be using AoC to help me improve my skills with golang.

I will be posting and updating my codebase every day to showcase a way to solve AoC using Golang.

Link -> https://adventofcode.com/2024

## Progression

| Day | Stars | Difficulty |
| --- | ----- | ---------- |
| 1   | ⭐⭐  | 🧩         |
| 2   | ⭐⭐  | 🧩🧩       |
| 3   | ⭐⭐  | 🧩\*       |
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

\* Solution was simple, but was unaware of the technology or library

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

## Day 2 - Learned

How to remove individual values in Golang

```go
// I previous was using this
// If i wanted to remove i
v = append(v[:i], v[i+1:]...)
// This works but when iterating over each other I was getting repeated values

// Creating a 1 instance smaller array and then appending values worked and was a much better solution.
temp := make([]int, 0, len(v)-1)
temp = append(temp, v[:i]...)
temp = append(temp, v[i+1:]...)
```

## Day 3 - Learned

That I need to learn Regexp. Through googling, I was able to piece together this regexp

```go
// Creating the Regexp
r, _ := regexp.Compile(`mul\((\d*)?,(\d*)?\)`)
// PartTwo with ors
r, _ := regexp.Compile(`mul\((\d*)?,(\d*)?\)|do\(\)|don't\(\)`)

// running it to get matches
matches := r.FindAllStringSubmatch(s, -1)

for _, match := range matches {
  match[0]
  match[1]
  ...
  // It appears out of index doesnt matter it will be null by default
  // Returns as a string
}
```
