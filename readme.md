# Project Title

mysqlsha1-go

## Project Requirement

Given the MySQL password hashes, find the original passwords.

## Method

Brute Force + Simple Dictionary Attack

## Prerequisites

Using Go 1.9.4

## In-code Parameters

hashes, set, i

in the main function

## Output

```
password 2470C0C06DEE42FD1618BB99005ADCA2EC9D1E19 127.305ms
qwertyuiop 6063C78456BB048BAF36BE1104D12D547834DFEA 127.466ms
computer 81101DED975D54BD76A3C8EAD293597AE9BB143F 127.606ms
abcdef C2D24DCA38E9E862098B85BF0AB35CAA52803797 128.793ms
security 1FDB0D828172183735F1ED9E45E6AF3CE04DE9D1 129.397ms
helloworld D35DB127DB631E6E27C6B75E8D376B04F64FAF83 537.625ms
abcdef C2D24DCA38E9E862098B85BF0AB35CAA52803797 3.264912s
iisct D5D82EF0B0A344A314E3B3E8D1E78A4697B97A53 19.896724s
```

## Credits

Answer to "Generate all possible n-character passwords"

by peterSO

https://stackoverflow.com/a/22741715

Answer to "Read text file into string array (and write)"

by Kyle Lemons

https://stackoverflow.com/a/18479916