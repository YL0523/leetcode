package main
import (
	"fmt"
)


// 一种特殊的树，满足下面的两个条件：
// 1. 堆总是一颗完全二叉树(除了最后一层，其他都是满节点，最右一层先排左节点)
// 2. 堆中某个节点的值总是大于等于（小于等于）其所有子节点的值，如果是大于等于的情况就称为大顶堆，小于等于情况就是小顶堆
// 完全二叉树适合用数组存储，因为下标为i的元素，他的左子树下标为2i，右子树下标为2i+1，父节点就是i/2向上取整

//步骤：
// 	  建堆：	堆是用数组存储的，而且0下标不存，从1开始存储，建堆就是在原地通过交换位置达到建立堆的目的。完全二叉树我们知道，
//			如果最后一个元素的下标为n，则1到n/2是非叶子节点，需要自上而下的堆化(和叶子结点比较)，n/2到n是叶子节点，不需要堆化
//	  插入一个元素：	先插入的元素放在堆的最后，然后和父节点比较，如果大于父节点，就交换位置，然后和父节点比较，直到把这个元素
//					放在了正确的层，这种也叫做自下而上的堆化
//	  删除一个元素：	假如删除最大的元素，然后从它的字节点找到第二大元素，放到堆顶 ，然后在第二大元素的子节点寻找
//	  堆排序：	比如有n个数据，我们先把数据建堆，生成一个大顶堆，元素个数为n，获取堆顶数据(也就是最大元素)，删除堆顶，并且把最后一个元素放在堆顶，
//			 然后堆化成n-1的大顶堆，堆化的时间复杂度为LogN,底数为2。重复获取堆顶，堆化成（n-2）大顶堆。我们获取的数据就是从大到小的顺序。
//			


// 堆的应用
// 1. 优先队列
//		合并n个有序小文件，把n个有序的小文件的第一个元素取出，放入堆中，取出堆顶到大文件，然后从小文件中取出一个加入到堆，以此类推，这样就把小文件合并成大文件了
// 2. 用堆求 Top K（就是从一堆数据中找出前 k 大的数据）
// 		a. 针对静态数据，建立大小为K的小顶堆，遍历数组，数组元素与堆顶比较，比堆顶大，就把堆顶删除，并插入该元素到堆
//		b. 针对动态数据，在动态数据插入的时候就与堆顶进行比较，看是否入堆，始终维护这个堆，需要的时候直接返回，最坏O(nLogK)
// 3. 海量关键词搜索记录，求搜索次数TopK
//		a. 先用hashtable去重，并累加搜索次数
//		b. 在建立大小为K的小顶堆，遍历散列表，次数大于堆顶的，顶替堆顶入堆(就是上面2的做法)
//		假如散列表很大，超出内存要求
//			a. 建立n个空文件，对搜索词求哈希值，哈希值对n取模，得到关键词被分配的文件号(0～n-1)
//			b. 对每个文件，利用散列表和堆，分别求出topK，然后把n个topK(比如10个Top20)放在一起，出现次数最多的K个关键词就是这海量数据里搜索最频繁的