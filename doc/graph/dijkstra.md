# 迪杰斯特拉算法-dijkstra
只能求出某点到其它点最短距离，并不能得出任意两点之间的最短距离。

## 算法步骤：
1. 将所有边初始化无穷大
2. 选择一个开始的顶点，添加到优先队列中
3. 对于该点的所有邻接顶点进行判断，如果该点的距离小于原先的值 ，则将该值进行更新
4. 将改点的所有邻接顶点添加到优先队列中
5. 从优先队列中挑选出一个路径值最小的顶点，将其弹出，作为新的顶点，重复3，4，5，直到所有点都被处理过一次

## 参考
- [漫画：图的 “最短路径” 问题](https://zhuanlan.zhihu.com/p/65340385)
- [Dijkstra算法和Floyd算法](http://www.cnblogs.com/biyeymyhjob/archive/2012/07/31/2615833.html)
- [Dijkstra算法](https://blog.csdn.net/qq_35644234/article/details/60870719)
- [github](https://github.com/RyanCarrier/dijkstra)