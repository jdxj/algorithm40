package chapter2

// 位运算
// 判断奇偶: x&1
// 清除最低位的1: x = x&(x-1), 应用: 求数字中, 二进制位为1的个数
//   1010 (10)
// & 1001 (10-1)
//   1000
// 得到最低位的1: x = x&(-x)