# 图相关知识点与问题汇总

## 判断图中是否有环
- 拓扑排序
    + 有向图直接使用拓扑排序
    + 无向图，稍微变形
        * 求出图中所有点的度
        * 删除图中所有度<=1的顶点；并移除与之关联的边；且将关联边的另一个顶点的度减一
        * 如果图中仍有度<=1的顶点；重复步骤2
        * 最后图中，仍有未被删除的顶点，则存在环
- 深度遍历
    + 深度遍历过程中，某点的一条边指向已访问过的顶点
    + 此访问过的顶点不是该点dfs中的父节点
    
- 参考
    + [LeetCode-684](https://leetcode.com/problems/redundant-connection/description/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=transfer2china)
    + [判断一个图是否有环](https://www.jianshu.com/p/acbd585f5c60)