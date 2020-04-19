# go-fucking-algorithm

解决没一类问题，都有一些固定的方式、方法或者说套路。

程序设计中的算法锻炼的是思维，练习算法的目的是建立计算机思维模式,针对算法的思维模式或者模型也就算法框架，统俗的说就是算法套路。

本仓库的目的就是强化开发中计算机思维模式，自我巩固语言基础，重新的常见的算法问题。

时时刻刻都要把自己当成小学生，经常回头看自己，经常会用意想不到的收获。

## 基本数据存储结构

内存中，最最最基本的数据结构(或者说数据存储方式)有且只有两种：

+ 数组（空间连续存储）
+ 链表（空间链接的存储）

其他的散列表、栈、队列、堆、树、图都是由基本的数组和链表进一步构建和表达的高级数据结构。

## 特点

链表主要靠遍历，访问数据。

数组可以靠索引，直接随机访问目标数据。


## 一些高级数据存储结构

如：「树」、「哈希表」、「图」，树有二叉和N叉，常用的是二叉，哈希质是列表(List)可用数组存储，也可以用链表存储，常常用作唯一性的Key，但是这时候要解决碰撞(出现了相同Key)问题，重点说一下图(如下)：

+ 「图」有两种常见的存储方式：
    + 1）基于链表的邻接表，邻使用链表相对节省空间，但没法随机取，
    + 2）基于(连续空间的)二维数组构成的邻接矩阵，可以随机存取，可以进行矩阵运算，相对消耗存储空间。


## 一些高级数据结构概念和模型


说到这个，脑子里最好形成一些空间结构图，辅助自己理解，一些重要的、常用于算法的高级数据结构概念和函数样例示例[Go]：


+ 非线性的「二叉树」遍历，根据访问节点值的次序会划分为前序、后序、中序三种勤劳的

```go
type TreeNode struct{
    Val interface{}
    Left, Right *TreeNode
}

void traverse(tree TreeNode) {
    //pre-order
    traverse(tree.Left)
    //center-order
    traverse(tree.Right)
    //after-order
}

```

+ 非线性N叉树遍历：

```go

type TreeNode struct{
    Val interface{}
    Children []TreeNode
}

func traverse(root TreeNode) {
    for child := range root.Children{
        traverse(child)
    }
}

```
+ 非线性森林叉树遍历，所谓森林就是有由个N叉树构成的数据集

根据N叉树，可以进一步引出森林的概念，也没什么特别的，就是在树上面在放一层数组，基本访问如下：

```go
type TreeNode struct{
    Val interface{}
    Children []TreeNode
}

func traverse(root TreeNode) {
    for child := range root.Children{
        traverse(child)
    }
}

func visitForest(forest []TreeNode) {
    for tree := range forest{
        traverse(tree)
    }
}

```

+ 链表遍历,链表有两种遍历方式，顺序和递归

```go

type LinkedListNode {
    Val interface{}
    Next *LinkedListNode
}

//顺序的线性遍历
func traverse(head LinkedListNode) {
    for (p := head; p != nil; p = p.Next) {
        // visit bt steps p.Val
    }
}

//递归遍历
func traverse(head LinkedListNode) {
 
    // visit head.Val
    // recursive
    traverse(head.Next)
    
}
```

+ 线性的数组遍历：

```go
func traverse(array int[]) {
    for (i := 0; i < len(array); i++) {
        // visit array[i]
    }
}
```

## 好了

自古牛角不可钻，唯有套路得码心，想到别的，再说，先开始撸码了

[零零散散记录一些问题点](./misc/NOTES.md)
[先温习一下Go语言基础](./language/README.md)