[toc]
# 图

## 说明
1. 邻接矩阵表达图的数据结构
2. 邻接表（矩阵比较稀疏的时候） V(顶点) E(边)

## 图的遍历
### 广度优先
BFS: 从图的某一节点出发， 首先依次访问该节点的所有相邻节点，再按照这些节点被访问的先后顺序，
依次访问与他们相邻的所有未被访问到的节点，重复此过程，直至所有节点均被访问。

### 深度优先
DFS: 假设初始状态是图中所有顶点均未被访问，则从某个顶点v出发，首先访问该顶点，然后依次从它的各个未被访问的邻接点出发深度优先搜索遍历图，直至图中所有和v有路径相通的顶点都被访问到。 若此时尚有其他顶点未被访问到，则另选一个未被访问的顶点作起始点，重复上述过程，直至图中所有顶点都被访问到为止。

# 参考
- [Golang 数据结构：图](https://wuyin.io/2018/06/22/golang-data-structure-graph/)
- [graph](https://flaviocopes.com/golang-data-structure-graph/)
- [图的遍历之 深度优先搜索和广度优先搜索](https://www.cnblogs.com/skywang12345/p/3711483.html)
- [Dijkstra, Floyd, Kruskal, Prim](https://github.com/muzixing/graph_algorithm)