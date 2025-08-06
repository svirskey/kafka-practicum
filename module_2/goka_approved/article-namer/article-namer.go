package articlenamer


import (
   "context"
   "encoding/json"
   "log"
   "fmt"

   "github.com/lovoo/goka"
)


var (
   Group goka.Group = "article-namer"
)


type ArticleName struct {
   Name string
}

type ArticleNameCodec struct{}

// Encode переводит UserLike в []byte
func (bec *ArticleNameCodec) Encode(value any) ([]byte, error) {
   if _, ok := value.(*ArticleName); !ok {
      return nil, fmt.Errorf("тип должен быть *ArticleNameEvent, получен %T", value)
   }
   return json.Marshal(value)
}


// Decode переводит  из []byte в структуру.
func (bec *ArticleNameCodec) Decode(data []byte) (any, error) {
   var (
      articleNameEvent ArticleName
      err      error
   )
   err = json.Unmarshal(data, &articleNameEvent)
   if err != nil {
      return nil, fmt.Errorf("ошибка десериализации: %v", err)
   }
   return &articleNameEvent, nil
}


func setArticle(ctx goka.Context, msg interface{}) {
   // Получаем значение из контекста
   // Если значение отсутствует, создаем новый BlockValue
   // Если значение существует, приводим его к типу *BlockValue
   
   // Приводим сообщение к типу *BlockEvent и проверяем, нужно ли разблокировать
    // Если нужно Разблокируем
    // Иначе - нет

   var ok bool
   var articleName *ArticleName

   if articleName, ok = msg.(*ArticleName); !ok || articleName == nil {
      return
   }

   if val := ctx.Value(); val != nil {
      articleName = val.(*ArticleName)
   } else {
      articleName = &ArticleName{}
   }

   ctx.SetValue(articleName)
   log.Printf("[proc] key: %s articleName: %s \n", ctx.Key(), articleName.Name)

}




func RunArticleNamer(brokers []string, inputTopic goka.Stream) {
   g := goka.DefineGroup(Group,
      goka.Input(inputTopic, new(ArticleNameCodec), setArticle),
      goka.Persist(new(ArticleNameCodec)),
   )
   p, err := goka.NewProcessor(brokers, g)


   if err != nil {
      log.Fatal(err)
   }
   err = p.Run(context.Background())
   if err != nil {
      log.Fatal(err)
   }
}