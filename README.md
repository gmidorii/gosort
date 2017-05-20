# gosort

## Package sort
- Sortできるstructの条件  
→ 下記Intergaceを実装していること
```
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

- sort.Search(n int, f func(int) booli int
  - f()がtrue になる最初のindexを返却する
  - 一致を探すことも、ある数値以上の値を探すことも可能
  - 配列は、ソートしておく必要がある