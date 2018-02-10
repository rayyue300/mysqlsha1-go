# Project Title

mysqlsha1-go

## Project Requirement

Given the MySQL password hashes, find the original passwords.

## Prerequisites

Using Go 1.9.4

## In-code Parameters

hashes, set, i

in the main function

## Output

```
security 1FDB0D828172183735F1ED9E45E6AF3CE04DE9D1 24.047ms
computer 81101DED975D54BD76A3C8EAD293597AE9BB143F 24.172ms
password 2470C0C06DEE42FD1618BB99005ADCA2EC9D1E19 24.953ms
abcdef C2D24DCA38E9E862098B85BF0AB35CAA52803797 3.079203s
iisct D5D82EF0B0A344A314E3B3E8D1E78A4697B97A53 19.621969s
```

## Credits

Answer to "Generate all possible n-character passwords"

by peterSO

https://stackoverflow.com/a/22741715

Answer to "Read text file into string array (and write)"

by Kyle Lemons

https://stackoverflow.com/a/18479916