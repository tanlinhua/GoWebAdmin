package bloom

import (
	"github.com/bits-and-blooms/bloom/v3"
)

type BloomFilter struct {
	*bloom.BloomFilter
}

func New(n uint, fp float64) *BloomFilter {
	filter := bloom.NewWithEstimates(n, fp) // n=容量, fp=误报率 (10000000, 0.01)
	return &BloomFilter{filter}
}

func (f *BloomFilter) Add(key string) {
	f.AddString(key)
}

func (f *BloomFilter) Test(key string) bool {
	return f.TestString(key)
}

func (f *BloomFilter) TestAndAdd(key string) bool {
	return f.TestAndAddString(key)
}

func (f *BloomFilter) Clear() {
	f.ClearAll()
}

/*
1. 布隆过滤器的概念
布隆过滤器（Bloom Filter） 是由 Howard Bloom在1970年提出的二进制向量数据结构，它具有很好的空间和时间效率，
被用来检测一个元素是不是集合中的一个成员，即判定 “可能已存在和绝对不存在” 两种情况。
如果检测结果为是，该元素不一定在集合中；但如果检测结果为否，该元素一定不在集合中,因此Bloom filter具有100%的召回率。

2. 布隆过滤器应用场景
垃圾邮件过滤
防止缓存击穿
比特币交易查询
爬虫的URL过滤
IP黑名单
查询加速【比如基于KV结构的数据】
集合元素重复的判断

3. 布隆过滤器工作原理
布隆过滤器的核心是一个超大的位数组和几个哈希函数。假设位数组的长度为m,哈希函数的个数为k。
下图表示有三个hash函数，比如一个集合中有x，y，z三个元素，分别用三个hash函数映射到二进制序列的某些位上，
假设我们判断w是否在集合中，同样用三个hash函数来映射，结果发现取得的结果不全为1，则表示w不在集合里面。

工作流程:
第一步：开辟空间：
开辟一个长度为m的位数组（或者称二进制向量），这个不同的语言有不同的实现方式，甚至你可以用文件来实现。
第二步：寻找hash函数
获取几个hash函数，前辈们已经发明了很多运行良好的hash函数，比如BKDRHash，JSHash，RSHash等等。这些hash函数我们直接获取就可以了。
第三步：写入数据
将所需要判断的内容经过这些hash函数计算，得到几个值，比如用3个hash函数，得到值分别是1000，2000，3000。之后设置m位数组的第1000，2000，3000位的值位二进制1。
第四步：判断
接下来就可以判断一个新的内容是不是在我们的集合中。判断的流程和写入的流程是一致的。

4. 布隆过滤器的优缺点
1、优点：
有很好的空间和时间效率
存储空间和插入/查询时间都是常数。
Hash函数相互之间没有关系，方便由硬件并行实现。
不需要存储元素本身，在某些对保密要求非常严格的场合有优势。
布隆过滤器可以表示全集，其它任何数据结构都不能。
2、缺点：
误判率会随元素的增加而增加
不能从布隆过滤器中删除元素

5. 布隆过滤器注意事项
布隆过滤器思路比较简单，但是对于布隆过滤器的随机映射函数设计，需要计算几次，向量长度设置为多少比较合适，这个才是需要认真讨论的。
如果向量长度太短，会导致误判率直线上升。
如果向量太长，会浪费大量内存。
如果计算次数过多，会占用计算资源，且很容易很快就把过滤器填满。
*/
