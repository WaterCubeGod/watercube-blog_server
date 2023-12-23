package flag

import (
	"context"
	"gvb_server/global"
)

func createES() {
	es := global.ES

	ctx := context.Background()

	exists, err := es.IndexExists(index).Do(ctx)
	if err != nil {
		global.LOG.Errorf("查找索引错误：%s", err)
		return
	}
	if !exists {
		_, err := es.CreateIndex(index).Body(articleTpl).Do(ctx)
		if err != nil {
			global.LOG.Errorf("创建索引失败：%s", err)
			return
		}
	}
	global.LOG.Info("创建索引成功")
}

var index = "article"

var articleTpl = `{
  "mappings": {
    "properties": {
      "title": {
        "type": "text",
        "analyzer": "ik_max_word"
      },
      "abstract": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "content": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "source": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "nick_name": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "tags": {
        "type": "keyword"
      },
      "category": {
        "type": "keyword"
      }
    }
  }
}`
