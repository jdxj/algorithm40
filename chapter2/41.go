package chapter2

// 231 2的幂
// 1. mod
// 2. math log
// 3. 位运算
// 规律: 只要是2的幂, 就是1000... -> x&(x-1) == 0

// 338 比特位计数
// 1. 利用 x&(x-1)
// 2. count[i] = count[i&(i-1)] + 1
