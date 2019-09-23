# All base digits

This code generates a sequence with the following method:

* For an integer n, count the number of digits required to represent n in all bases from 1 to n.

## Output

| n   | f(n) | Digits of n in all bases, for reference  |
| --- | ---  | ---                                      |
| 1   | 1    | 1                                        |
| 2   | 4    | 11 10                                    |
| 3   | 7    | 111 11 10                                |
| 4   | 11   | 1111 100 11 10                           |
| 5   | 14   | 11111 101 12 11 10                       |
| 6   | 17   | 111111 110 20 12 11 10                   |
| 7   | 20   | 1111111 111 21 13 12 11 10               |
| 8   | 24   | 11111111 1000 22 20 13 12 11 10          |
| 9   | 28   | 111111111 1001 100 21 14 13 12 11 10     |
| 10  | 31   | 1111111111 1010 101 22 20 14 13 12 11 10 |
