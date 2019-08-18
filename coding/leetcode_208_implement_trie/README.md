# Trie

Trie is a special kind of tree which facilitate string searching where each node points to all the next characters in the provided words and root node represents empty string. For example the below trie representation shows how `khaled`, `kha`, `khalid`, `karam`, `bob` is stored

```info
                                         ""
                                        /  \
                                       /    b
                                      /      \
                                     /        o
                                    /          \
                                    k           b (word)
                                   / \
                                  h   a
                                /      \
                        (word) a        r
                              /          \
                             l            a
                            /  \           \
                           e    i            m (word)
                          /      \
                  (word) d        d (word)
```

## Implementation

Two things we need to find out

* Data structure to hold all children nodes and fetch the child node quickly
* how to identify if the character resembles the end of a word

In this implementation I will use a `map` to with the character as a key and a reference to node as a value and that will provide us with checking if a character is found and fetching the related node in `O(1)`

In the structure representing the node a flag will be `true` if the current character represents end of a word otherwise this flag will be set to `false`

## Complexity

`Insert`, `Search`, `StartsWith` are O(K) where K is the number of characters in the word
