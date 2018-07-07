package persist

import (
	"context"
	"errors"
	"log"

	"github.com/zzayne/go-crawler/engine"
	"gopkg.in/olivere/elastic.v5"
)

//ItemSaver 。...
func ItemSaver(index string) (chan engine.Item, error) {
	//连接docker中的elastic时，需要SetSniff off
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:Saved  item #%d:%v", itemCount, item)
			itemCount++
			err = save(client, item, index)
			if err != nil {
				log.Printf("Item Saver:error saving item %v:%v", item, err)
			}

		}
	}()
	return out, err

}

//ItemSave ...
func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("Type can't be null")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.ID != "" {
		indexService.Id(item.ID)
	}
	_, err := indexService.Do(context.Background())

	return err

}
