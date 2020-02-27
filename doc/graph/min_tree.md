[toc]

# 最小生成树
## 含义
连通无向图G=(V, E), V是点集合，E是边集合，对于每条边(u,v)∈E, 权重w(u,v),存在一无环子集，其权重最小，此即为最小生成树。

## 普里姆算法 [prim]
1. 输入：一个加权连通图，其中顶点集合为V，边集合为E；
2. 初始化：Vnew = {x}，其中x为集合V中的任一节点（起始点），Enew = {},为空；
3. 重复下列操作，直到Vnew = V：
    * 在集合E中选取权值最小的边<u, v>，其中u为集合Vnew中的元素，而v不在Vnew集合当中，并且v∈V（如果存在有多条满足前述条件即具有相同权值的边，则可任意选取其中之一）；
    * 将v加入集合Vnew中，将<u, v>边加入集合Enew中；
4. 输出：使用集合Vnew和Enew来描述所得到的最小生成树。


## Kruskal算法
1. 记Graph中有v个顶点，e个边
2. 新建图Graphnew，Graphnew中拥有原图中相同的e个顶点，但没有边
3. 将原图Graph中所有e个边按权值从小到大排序
4. 循环：从权值最小的边开始遍历每条边 直至图Graph中所有的节点都在同一个连通分量中
    if 这条边连接的两个节点于图Graphnew中不在同一个连通分量中 添加这条边到图Graphnew中

## 参考
- [最小生成树-百度百科](https://baike.baidu.com/item/%E6%9C%80%E5%B0%8F%E7%94%9F%E6%88%90%E6%A0%91/5223845)
- [小生成树---Prim算法和Kruskal算法](https://www.cnblogs.com/zhangming-blog/p/5414514.html)
- [最小生成树-Prim算法和Kruskal算法](https://www.cnblogs.com/biyeymyhjob/archive/2012/07/30/2615542.html)