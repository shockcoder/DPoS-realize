# 共识机制 - 股份授权证明机制 DPOS

### 什么是DPoS

对于PoS机制的加密货币，每个节点都可以创建区块，并按照个人的持股比例获得“利息”。DPoS是由被社区选举的可信帐户（受托人，得票数排行前101位）来创建区块。为了成为正式受托人，用户要去社区拉票，获得足够多用户的信任。用户根据自己持有的加密货币数量占总量的百分比来投票。DPoS机制类似于股份制公司，普通股民进不了董事会，要投票选举代表（受托人）代他们做决策。
这101个受托人可以理解为101个矿池，而这101个矿池彼此的权利是完全相等的。那些握着加密货币的用户可以随时通过投票更换这些代表（矿池），只要他们提供的算力不稳定，计算机宕机、或者试图利用手中的权力作恶，他们将会立刻被愤怒的选民门踢出整个系统，而后备代表可以随时顶上去。

### DPoS的优缺点

```go
优点
```

- 能耗更低。DPoS机制将节点数量进一步减少到101个，在保证网络安全的前提下，整个网络的能耗进一步降低，网络运行成本最低。
- 更加去中心化。目前，对于比特币而言，个人挖矿已经不现实了，比特币的算力都集中在几个大的矿池手里，每个矿池都是中心化的，就像DPoS的一个受托人，因此DPoS机制的加密货币更加去中心化。PoS机制的加密货币（比如未来币），要求用户开着客户端，事实上用户并不会天天开着电脑，因此真正的网络节点是由几个股东保持的，去中心化程度也不能与DPoS机制的加密货币相比。
- 更快的确认速度。每个块的时间为10秒，一笔交易（在得到6-10个确认后）大概1分钟，一个完整的101个块的周期大概仅仅需要16分钟。而比特币（PoW机制）产生一个区块需要10分钟，一笔交易完成（6个区块确认后）需要1个小时。点点币（PoS机制）确认一笔交易大概也需要1小时。

```go
缺点
```

- 投票的积极性并不高。绝大多数持股人（90％+）从未参与投票。这是因为投票需要时间、精力以及技能，而这恰恰是大多数投资者所缺乏的。
- 对于坏节点的处理存在诸多困难。社区选举不能及时有效的阻止一些破坏节点的出现，给网络造成安全隐患。