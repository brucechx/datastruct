# 拓扑排序(Kahn算法)

## 算法思想
- 维护一个入度为0的顶点集合(V)
- 每次从顶点集合中取一个点放到队列list(如果需要排序，可以使用优先队列)
- 然后循环遍历该点引出的所有边；从图中移除这条边，并获得边的另一个点；
- 如果该点的入度在移除这条边后，为0，则放入顶点集合(V)；
- 如果顶点集合V不为0;则重复2，3，4步骤；直至集合为空；
- 当集合为空时；如果图中还存在边，则图中有环；
  否则，队列list 即为图的拓扑排序结果；

## 参考
- [拓扑排序-Kahn算法](http://ddrv.cn/a/49505)
- [拓扑排序的实现方法以及环路检测](https://zhuanlan.zhihu.com/p/34871092)
- [Golang实现拓扑排序-DFS算法版](https://studygolang.com/articles/24427)
- [拓扑排序模板（Kahn算法和DFS实现）](https://blog.csdn.net/baodream/article/details/80368764?depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-24&utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-24)