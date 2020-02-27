# B(B+)树
## B树
# 定义
## 特性
那么B-Tree是满足下列条件的数据结构：<br>
d为大于1的一个正整数，称为B-Tree的度。
h为一个正整数，称为B-Tree的高度。
1. 每个非叶子节点由n-1个key和n个指针组成，其中d<=n<=2d。
2. 每个叶子节点最少包含一个key和两个指针，最多包含2d-1个key和2d个指针，叶节点的指针均为null 。
3. 所有叶节点具有相同的深度，等于树高h。
4. key和指针互相间隔，节点两端是指针。
5. 一个节点中的key从左到右非递减排列。
6. 所有节点组成树结构。
7. 每个指针要么为null，要么指向另外一个节点。

### 说明
* 如果某个指针在节点node最左边且不为null，则其指向节点的所有key小于v(key1)，其中v(key1)为node的第一个key的值。
* 如果某个指针在节点node最右边且不为null，则其指向节点的所有key大于v(keym)，其中v(keym)为node的最后一个key的值。
* 如果某个指针在节点node的左右相邻key分别是keyi和keyi+1且不为null，则其指向节点的所有key小于v(keyi+1)且大于v(keyi)。

## B+树
###与B-Tree相比，B+Tree有以下不同点：
* 每个节点的指针上限为2d而不是2d+1。
* 内节点不存储data，只存储key；叶子节点不存储指针。

## 参考
- [google-btree](https://godoc.org/github.com/google/btree)
- [google-btree](https://github.com/google/btree)
- [为什么 B-tree 在不同著作中度的定义有一定差别？](https://www.zhihu.com/question/19836260)
- [漫画算法：什么是 B 树？](http://blog.jobbole.com/111757/)
- [从B树、B+树、B*树谈到R 树](https://blog.csdn.net/v_JULY_v/article/details/6530142)
- [MySQL索引背后的数据结构及算法原理](http://blog.codinglabs.org/articles/theory-of-mysql-index.html)
- [july B树](https://github.com/julycoding/The-Art-Of-Programming-By-July/blob/master/ebook/zh/03.02.md)
