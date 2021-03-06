package chapter2

// 121, 122, 123, 309, 188, 714 股票买卖系列

// 121: 买一次, 卖一次
// 122: 买无数次, 卖无数次

// DP 应该是多维的
//   - [i]: 到第 i 天, 最大收益
//   - [j]: 0, 不拥有股票; 1, 拥有一股
//   - [k]: 交易次数

//                     --
//                    | MP[i-1, k, 0]          不动
// MP[i, k, 0] = Max -|
//                    | MP[i-1, k-1, 1] + a[i] 卖出
//                     --
//
//                     --
//                    | MP[i-1, k, 1]          不动
// MP[i, k, 1] = Max -|
//                    | MP[i-1, k-1, 0] - a[i] 买入
//                     --
// 最后遍历 MP[n-1, {0...k}, 0] 中的最大值
