package chapter2

// 169 求众数
// count(x) > n/2

// 1. 暴力
//  两层循环
//    第一层, 指定当前值
//    第二层, 计算当前值出现的次数, 并比较

// 2. hash
//   key: x, value: x_count

// 3. sort(array_x)
//    之后再计算中值的个数
// O(N*logN)

// 4. divide & conquer
