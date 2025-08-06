package filter

import (
	"context"
	"log"

	"github.com/lovoo/goka"

	"github.com/svirskey/kafka-practicum/module_2/goka_approved/article-namer"
	articlenamer "github.com/svirskey/kafka-practicum/module_2/goka_approved/article-namer"
	"github.com/svirskey/kafka-practicum/module_2/goka_approved/blocker"
	"github.com/svirskey/kafka-practicum/module_2/goka_approved/user"
)


var (
   filterGroup goka.Group = "filter"
)


func shouldDrop(ctx goka.Context) bool {
   v := ctx.Join(goka.GroupTable(blocker.Group))
   return v != nil && v.(*blocker.BlockValue).Blocked
}

func translate(ctx goka.Context, l *user.Like) *user.Like {
	for i, w := range words {
		if tw := ctx.Lookup(articlenamer.ArticleName, w); tw != nil {
			words[i] = tw.(string)
		}
	}
	return &user.Like{
		
	}
}

func RunFilter(brokers []string, inputTopic goka.Stream, outputTopic goka.Stream) {
   g := goka.DefineGroup(filterGroup,
      goka.Input(inputTopic, new(user.LikeCodec), func(ctx goka.Context, msg interface{}) {
         if shouldDrop(ctx) {
            return
         }
         m:= translate(ctx, msg)
         ctx.Emit(outputTopic, ctx.Key(), m)
      }),
      goka.Output(outputTopic, new(user.LikeCodec)),
      goka.Join(goka.GroupTable(blocker.Group), new(blocker.BlockValueCodec)),
      goka.Lookup(goka.GroupTable(articlenamer.Group),new(articlenamer.ArticleNameCodec)),
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