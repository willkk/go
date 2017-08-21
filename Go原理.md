# Go原理
## 1. 协程(原理, 优势)

并发模型演进：进程 => 线程 => 协程

特点：

|  类型 |  优势  |  劣势  |
| --------- | -------- | --------- |
| 进程   | 优点是运行空间相互隔离，保证可靠性 | 进程调度粒度大，资源消耗大（内核结构，堆栈，fd），并发量低，进程切换开销大
| 线程   | 进程空间内运行，提高共享效率，并发执行  | 线程调度粒度较小，资源消耗较多（主要是栈），大并发量受限，锁竞争，上下文切换开销较大
| 协程 | 协程调度粒度小，资源开销低（把栈放到堆上），内存有多大，并发就有多大，只需少量锁，协程切换开销小。| （go支持多核）不保证可靠性。

## 2. 调度算法(避开瓶颈，http://www.sizeofvoid.net/goroutine-under-the-hood/)

协作式调度，即需要自己主动让出cpu，响应时间方面会有劣势。

未来，我们会为了代码的简洁和可维护性而放弃一点点的性能损失。

本节目的：

#### a) goroutine也有限制，不要奔放的使用goroutine，否则系统会更奔放给你看。所以，要高效的使用和管理goroutine。

#### b) C性能更高，但逻辑更重，代码更复杂，使用go可以线性编程，降低逻辑复杂度，使代码更简单和具有可维护性，即使损失一点性能。

三个主要结构：M（负责管理内核线程），P（负责goroutine上下文切换），G（goroutine）

下面以Gopher搬砖来说明它们之间的关系：

![](https://github.com/willkk/go/blob/master/images/gopher.png)

M（Gopher），P（小车），G（砖块）
进程启动的时候生成一个主Gopher（M），并最多生成GOMAXPROCS（不大于系统逻辑核心数）个小车（P）。go命令随后创建goroutine（砖，G）。砖太多，小车空闲，就创建新的地鼠M去搬砖。
![](https://github.com/willkk/go/blob/master/images/gopher2.png)

| 操作 | 条件 |
| --- | --- |
| 生成M | 砖太多，现有的地鼠M忙不过来，并且低于低于GOMAXPROCS，生成新的M。
| sleep | 自己的小车没有砖，工厂仓库也没砖，别人的小车也没砖
| wakeup | 砖很多，唤醒地鼠M去工作。

调度点（让出CPU）：

| 操作 |调度 |状态转移 | 说明 
| --- | --- | --- | --- 
| channel读写，定时器，网络poll | runtime.park | running -> waiting | 并不放入全局队列，需要runtime.ready
| | runtime.gosched | 其他状态 -> runnable | 放入全局等待队列
| syscall | syscall | running -> syscall | entersyscall让出cpu，exitsyscall放入全局队列。

结论：相对简单，goroutine调度公平性有待提高。
