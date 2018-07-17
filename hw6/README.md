# 說明
因為govendor 會抓預設的/bin 位置，所以$GOPATH 有另外調整的話，使用的govendor位置也要自己指定

我的GOPATH : ~/local_dev/gitlab/go/

# 移除相依的套件
~/local_dev/gitlab/go/bin/govendor remove +v

# 自動加入相依的套件
~/local_dev/gitlab/go/bin/govendor add +e

# 加上自己寫的套件(我放在~/local_dev/gitlab/go/src底下)
~/local_dev/gitlab/go/bin/govendor add rocket
