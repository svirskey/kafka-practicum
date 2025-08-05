package filter


import (
   "context"
   "log"


   "github.com/lovoo/goka"

   "github.com/svirskey/kafka-practicum/module_2/goka_approved/user"
   "github.com/svirskey/kafka-practicum/module_2/goka_approved/blocker"
)


var (
   filterGroup goka.Group = "filter"
)


func shouldDrop(ctx goka.Context) bool {
   v := ctx.Join(goka.GroupTable(blocker.Group))
   return v != nil && v.(*blocker.BlockValue).Blocked
}


func RunFilter(brokers []string, inputTopic goka.Stream, outputTopic goka.Stream) {
   g := goka.DefineGroup(filterGroup,
      goka.Input(inputTopic, new(user.LikeCodec), func(ctx goka.Context, msg interface{}) {
         if shouldDrop(ctx) {
            return
         }
         ctx.Emit(outputTopic, ctx.Key(), msg)
      }),
      goka.Output(outputTopic, new(user.LikeCodec)),
      goka.Join(goka.GroupTable(blocker.Group), new(blocker.BlockValueCodec)),
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