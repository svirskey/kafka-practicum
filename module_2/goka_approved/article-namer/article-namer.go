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


type ArticleNameEvent struct {
   Name string
}

type ArticleNameEventCodec struct{}

// Encode переводит UserLike в []byte
func (bec *ArticleNameEventCodec) Encode(value any) ([]byte, error) {
   if _, ok := value.(*ArticleNameEvent); !ok {
      return nil, fmt.Errorf("тип должен быть *ArticleNameEvent, получен %T", value)
   }
   return json.Marshal(value)
}


// Decode переводит  из []byte в структуру.
func (bec *ArticleNameEventCodec) Decode(data []byte) (any, error) {
   var (
      articleNameEvent ArticleNameEvent
      err      error
   )
   err = json.Unmarshal(data, &articleNameEvent)
   if err != nil {
      return nil, fmt.Errorf("ошибка десериализации: %v", err)
   }
   return &articleNameEvent, nil
}


type BlockValue struct {
   Blocked bool
}


type BlockValueCodec struct{}


// Encode переводит UserLike в []byte
func (bvc *BlockValueCodec) Encode(value any) ([]byte, error) {
   if _, ok := value.(*BlockValue); !ok {
      return nil, fmt.Errorf("тип должен быть *BlockValue, получен %T", value)
   }
   return json.Marshal(value)
}


// Decode переводит UserLike из []byte в структуру.
func (bvc *BlockValueCodec) Decode(data []byte) (any, error) {
   var (
      blockValue BlockValue
      err      error
   )
   err = json.Unmarshal(data, &blockValue)
   if err != nil {
      return nil, fmt.Errorf("ошибка десериализации: %v", err)
   }
   return &blockValue, nil
}



func block(ctx goka.Context, msg interface{}) {
   // FIXME: Выполните 
   // Получаем значение из контекста
   // Если значение отсутствует, создаем новый BlockValue
   // Если значение существует, приводим его к типу *BlockValue
   
   // Приводим сообщение к типу *BlockEvent и проверяем, нужно ли разблокировать
    // Если нужно Разблокируем
    // Иначе - нет

   var ok bool
   var blockValue *BlockValue
   var blockEvent *BlockEvent

   if blockEvent, ok = msg.(*BlockEvent); !ok || blockEvent == nil {
      return
   }


   if val := ctx.Value(); val != nil {
      blockValue = val.(*BlockValue)
   } else {
      blockValue = &BlockValue{}
   }
   log.Printf("[proc] key: %s IsBlocked: %s, Unblock: %s \n", ctx.Key(), blockValue.Blocked, blockEvent.Unblock)
   if blockValue.Blocked && blockEvent.Unblock {
      blockValue.Blocked = false
      ctx.SetValue(blockValue)
      log.Printf("[proc] unblocked user %s \n", ctx.Key())
   } else if !blockValue.Blocked && !blockEvent.Unblock {
      blockValue.Blocked = true
      ctx.SetValue(blockValue)
      log.Printf("[proc] blocked user %s \n", ctx.Key())
   }
}




func RunBlocker(brokers []string, inputTopic goka.Stream) {
   g := goka.DefineGroup(Group,
      goka.Input(inputTopic, new(BlockEventCodec), block),
      goka.Persist(new(BlockValueCodec)),
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